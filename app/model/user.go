package model

import (
	"iceforg/pkg/db"
	"iceforg/pkg/multilingual"
	"iceforg/pkg/utils"

	"github.com/jinzhu/gorm"
)

const (
	UserKeyProperty = "name"
	tableName       = "t_user"
)

func (*User) TableName() string {
	return tableName
}

type User struct {
	Base
	UserID   string `gorm:"column:user_id"`
	UserName string `gorm:"column:name"`
	Password string `gorm:"column:password"`
	NickName string `gorm:"column:nick_name"`
}

func (u *User) Save() (string, error) {
	u.UserID = utils.HashCodeUUID(20)
	db := db.GetMysqlProvider().Save(&u)
	return u.UserID, db.Error
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
	db := db.GetMysqlProvider().Table(tableName).Where("name = ?", u.UserName).Count(&cnt)
	return cnt > 0, db.Error
}

func (u *User) IsExisted() (bool, error) {
	var cnt int64
	db := db.GetMysqlProvider().Table(tableName).
		Where("name = ? and password = ?",
			u.UserName, u.Password).Count(&cnt)
	return cnt > 0, db.Error
}

func (u *User) DetailByNameAndPw() error {
	db := db.GetMysqlProvider().Table(tableName).
		Where("name = ? and password = ?",
			u.UserName, u.Password).Find(&u)
	return db.Error
}
