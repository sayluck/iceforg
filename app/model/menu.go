package model

import (
	"iceforg/pkg/db"
	"iceforg/pkg/utils"
)

const (
	MENU_TABLE_NAME = "t_menu"
)

func (*Menu) TableName() string {
	return MENU_TABLE_NAME
}

type Menu struct {
	Base    `json:"base"`
	PageNum int    `gorm:"column:m_page_num"`
	Name    string `gorm:"column:m_name"`
	Sort    int    `gorm:"column:m_sort"`
	SupCode string `gorm:"column:m_sup_code"`
	Route   string `gorm:"column:m_route"`
}

func (m *Menu) Save() (string, error) {
	m.Code = utils.CodeGenerate()
	dbRet := db.GetMysqlProvider().Save(&m)
	return m.Code, dbRet.Error
}
