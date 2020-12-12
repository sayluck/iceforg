package userCenter

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

func (u *User) Save() (string, error) {
	u.Code = utils.CodeGenerate()
	dbRet := db.GetMysqlProvider().Save(&u)
	return u.Code, dbRet.Error
}

func (u *User) DetailByKeyProperty() (interface{}, error) {
	var (
		err error
	)
	err = db.GetMysqlProvider().Where("m_code = ?", u.Code).First(&u).Error
	if gorm.IsRecordNotFoundError(err) {
		return u, multilingual.UserNotExisted
	}
	return u, err
}

func (u *User) IsExistedByKeyProperty() (bool, error) {
	var cnt int64
	dbRet := db.GetMysqlProvider().Table(USER_TABLE_NAME).Where("name = ?", u.UserName).Count(&cnt)
	return cnt > 0, dbRet.Error
}

func (u *User) IsExisted() (bool, error) {
	var cnt int64
	dbRet := db.GetMysqlProvider().Table(USER_TABLE_NAME).
		Where("name = ? and password = ?",
			u.UserName, u.Password).Count(&cnt)
	return cnt > 0, dbRet.Error
}

func (u *User) DetailByNameAndPw() error {
	dbRet := db.GetMysqlProvider().Table(USER_TABLE_NAME).
		Where("name = ? and password = ?",
			u.UserName, u.Password).Find(&u)
	return dbRet.Error
}
