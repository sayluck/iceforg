package model

import (
	"iceforg/pkg/db"
	"iceforg/pkg/multilingual"
	"iceforg/pkg/utils"

	"github.com/jinzhu/gorm"
)

const (
	USER_TABLE_NAME = "t_user"
)

func (*User) TableName() string {
	return USER_TABLE_NAME
}

type User struct {
	Base
	UserName string `gorm:"column:name"`
	Password string `gorm:"column:password"`
	NickName string `gorm:"column:nick_name"`
}

func (u *User) Save() (string, error) {
	u.Code = utils.CodeGenerate()
	dbRet := db.GetMysqlProvider().Save(&u)
	return u.Code, dbRet.Error
}

func (u *User) DetailByKeyProperty() (interface{}, error) {
	var (
		err error
	)
	err = db.GetMysqlProvider().Where("name = ?", u.UserName).First(&u).Error
	if gorm.IsRecordNotFoundError(err) {
		return u, multilingual.UserNotExisted
	}
	return u, err
}

func (u *User) IsExistedByKeyProperty() (bool, error) {
	var cnt int64
	db := db.GetMysqlProvider().Table(USER_TABLE_NAME).Where("name = ?", u.UserName).Count(&cnt)
	return cnt > 0, db.Error
}

func (u *User) IsExisted() (bool, error) {
	var cnt int64
	db := db.GetMysqlProvider().Table(USER_TABLE_NAME).
		Where("name = ? and password = ?",
			u.UserName, u.Password).Count(&cnt)
	return cnt > 0, db.Error
}

func (u *User) DetailByNameAndPw() error {
	db := db.GetMysqlProvider().Table(USER_TABLE_NAME).
		Where("name = ? and password = ?",
			u.UserName, u.Password).Find(&u)
	return db.Error
}
