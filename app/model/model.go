package model

import "time"

type Base struct {
	ID          int       `gorm:"column:id"`
	CreateTime  time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP"`
	LastModTime time.Time `gorm:"column:last_mod_time;default:CURRENT_TIMESTAMP"`
}
