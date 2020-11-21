package user

import (
	"context"
	"errors"
	"iceforg/app/common"
	"iceforg/app/model"
	"iceforg/pkg/config"
	"iceforg/pkg/multilingual"
	"iceforg/pkg/utils"
	"strings"
	"time"

	"github.com/jinzhu/gorm"

	. "iceforg/app/log"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	UserID   = "userID"
	UserName = "userName"
	TeamCode = "TeamCode"
)

func Register(user *UserRegister) (string, error) {
	var err error

	if user.NickName == "" {
		user.NickName = user.UserName
	}

	userM := model.User{}

	err = utils.TramsStruct(&user, &userM)
	if err != nil {
		return "", err
	}

	userID, err := userM.Save()
	if err != nil {
		return "", err
	}
	userM.Code = userID
	return generateToken(&userM)
}

func Detail(name string) (*UserDetail, error) {
	user := &model.User{
		UserName: name,
	}

	var (
		err        error
		userCenter = &UserDetail{}
	)
	_, err = user.DetailByKeyProperty()
	if err != nil {
		return userCenter, err
	}
	err = utils.TramsStruct(user, userCenter)
	if err != nil {
		IceLog.Errorf(context.Background(), "currentUser trams struct failed,%v", err.Error())
	}
	return userCenter, err
}

func Login(user *UserLogin) (string, error) {
	var (
		u   model.User
		err error
	)
	err = utils.TramsStruct(&user, &u)
	if err != nil {
		return "", err
	}
	err = u.DetailByNameAndPw()
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return "", multilingual.UserLoginErr
		}
		return "", multilingual.SystemOperationError
	}

	token, err := generateToken(&u)
	if err != nil {
		IceLog.Errorf(context.Background(), "generate token error:%v", err)
		return "", multilingual.SystemOperationError
	}
	return token, nil
}

func generateToken(u *model.User) (string, error) {
	claim := jwt.MapClaims{
		UserID:   u.Code,
		UserName: u.UserName,
		"nbf":    time.Now().Unix(),
		"iat":    time.Now().Unix(),
		"iss":    "iceforg",
		"exp":    int64(time.Now().Unix() + config.GetConfig().App.Token.ExpiredAfter), //  60*60*24*1 one day, 10s
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(common.TokenSecret))
}

func ParseToken(t string) (*UserLogin, error) {
	if t == "" || t == "null" {
		return nil, multilingual.UserInvaildToken
	}
	user := &UserLogin{}
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		return []byte(common.TokenSecret), nil
	})
	if err != nil {
		if strings.Contains(err.Error(), common.TokenIsExpired) {
			return user, multilingual.UserTokenIsExpired
		}
		return user, err
	}
	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return user, errors.New("parse token error,claims change to mapclaims")
	}

	if !token.Valid {
		return user, multilingual.UserInvaildToken
	}

	user.UserID = claim[UserID].(string)
	user.UserName = claim[UserName].(string)
	return user, nil
}
