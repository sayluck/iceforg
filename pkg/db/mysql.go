package db

import (
	"github.com/jinzhu/gorm"
	"iceforg/pkg/config"
)

var db *gorm.DB

type mysql struct {
	config.Mysql
}

func (m *mysql) Connect() (interface{}, error) {
	var (
		err error
	)
	if db == nil {
		db, err = gorm.Open(m.URL)
		if err != nil {
			return nil, err
		}
	}
	return db, nil
}
