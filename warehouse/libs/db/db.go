package db

import (
	log "github.com/ctlog"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open("mysql", "<root>:<qwe123>/<contract>?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalln(err)
		panic(err)
	}
}
