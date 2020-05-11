package user

import (
	"errors"
	"iceforg/app/common"
	"iceforg/app/model"
	"iceforg/pkg/multilingual"
	"iceforg/pkg/utils"
	"strings"
	"time"

	"github.com/jinzhu/gorm"

	. "iceforg/app/log"

	jwt "github.com/dgrijalva/jwt-go"
)

func Register(user *UserRegister) error {
	var err error

	if user.NickName == "" {
		user.NickName = user.UserName
	}

	userM := model.User{}

	err = utils.TramsStruct(&user, &userM)
	if err != nil {
		return err
	}
	_, err = userM.Save()
	return err
}

func Detail(name string) (interface{}, error) {
	user := model.User{
		UserName: name,
	}
	return user.DetailByKeyProperty()
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
		Log.Errorf("generate token error:%v", err)
		return "", multilingual.SystemOperationError
	}
	return token, nil
}

func generateToken(u *model.User) (string, error) {
	claim := jwt.MapClaims{
		"id":       u.ID,
		"username": u.UserName,
		"nbf":      time.Now().Unix(),
		"iat":      time.Now().Unix(),
		"iss":      "iceforg",
		"exp":      int64(time.Now().Unix() + 10), //  60*60*24*1 one day, 10s
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(common.TokenSecret))
}

func ParseToken(t string) (*UserLogin, error) {
	if t == "" {
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

	user.ID = int64(claim["id"].(float64))
	user.UserName = claim["username"].(string)
	return user, nil
}
