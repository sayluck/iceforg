package service

import (
	"iceforg/app/model"
	"iceforg/pkg/db"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	// TODO
}

func Register(user *model.User) error {
	var err error

	if user.NickName == "" {
		user.NickName = user.UserName
	}

	err = db.GetMysqlProvider().Save(user).Error
	if err != nil {
		return err
	}
	return nil
}
