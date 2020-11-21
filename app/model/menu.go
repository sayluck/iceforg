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
	Level   int    `gorm:"column:m_level"`
	SupCode string `gorm:"column:m_sup_code"`
	Route   string `gorm:"column:m_route"`
}

func (m *Menu) List() (interface{}, error) {
	var data = []*Menu{}
	dbRet := db.GetMysqlProvider().Table(MENU_TABLE_NAME).
		Find(&data, "m_page_num = ?", m.PageNum)
	return data, dbRet.Error
}

func (m *Menu) DetailByKeyProperty() (interface{}, error) {
	return nil, nil
}

func (m *Menu) IsExistedByKeyProperty() (bool, error) {
	return false, nil
}

func (m *Menu) Save() (string, error) {
	m.Code = utils.CodeGenerate()
	dbRet := db.GetMysqlProvider().Save(&m)
	return m.Code, dbRet.Error
}
