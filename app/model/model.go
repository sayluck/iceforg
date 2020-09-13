package model

import "time"

type Base struct {
	Code        string    `gorm:"column:m_code"`
	CreateTime  time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP"`
	LastModTime time.Time `gorm:"column:last_mod_time;default:CURRENT_TIMESTAMP"`
	Creator     string    `gorm:"column:creator"`
	Modifier    string    `gorm:"column:modifier"`
}
