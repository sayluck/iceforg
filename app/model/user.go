package model

import (
	"iceforg/pkg/db"
	"iceforg/pkg/multilingual"

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
	UserName string `gorm:"column:name"`
	Password string `gorm:"column:password"`
	NickName string `gorm:"column:nick_name"`
}

func (u *User) Save() (int64, error) {
	db := db.GetMysqlProvider().Save(u)
	return db.RowsAffected, db.Error
}

func (u *User) DetailByKeyProperty() (interface{}, error) {
	var (
		user User
		err  error
	)
	err = db.GetMysqlProvider().Where("name = ?", u.UserName).First(&user).Error
	if gorm.IsRecordNotFoundError(err) {
		return user, multilingual.UserNotExisted
	}
	return user, err
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
