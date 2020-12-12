package userCenter

import (
	"context"
	"errors"
	"iceforg/app/common"
	"iceforg/pkg/config"
	"iceforg/pkg/multilingual"
	"iceforg/pkg/utils"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

const (
	UserID   = "userID"
	UserName = "userName"
)

type UserCenter struct {
}

func (uc *UserCenter) Register(ctx context.Context, u *UserRegister, token *string) error {
	var (
		err   error
		userM User
	)

	if u.NickName == "" {
		u.NickName = u.UserName
	}

	err = utils.TramsStruct(&u, &userM)
	if err != nil {
		return err
	}

	userM.Code, err = userM.Save()
	if err != nil {
		return err
	}
	*token, err = generateToken(&userM)
	return err
}

func (uc *UserCenter) Detail(ctx context.Context, name *string, detail *UserDetail) error {
	user := &User{
		UserName: *name,
	}

	var (
		err error
	)

	_, err = user.DetailByKeyProperty()
	if err != nil {
		return err
	}
	return utils.TramsStruct(user, detail)
}

func (uc *UserCenter) SaveUser(ctx context.Context, u *User, code *string) error {
	var (
		err error
	)
	*code, err = u.Save()
	return err
}

func (uc *UserCenter) Login(ctx context.Context, user *UserLogin, token *string) error {
	var (
		err   error
		userM User
	)

	err = utils.TramsStruct(&user, &userM)
	if err != nil {
		return err
	}
	err = userM.DetailByNameAndPw()
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return multilingual.UserLoginErr
		}
		return multilingual.SystemOperationError
	}

	*token, err = generateToken(&userM)
	if err != nil {
		return multilingual.SystemOperationError
	}
	return nil
}

func generateToken(u *User) (string, error) {
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

func (uc *UserCenter) ParseToken(t string) (*UserLogin, error) {
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
