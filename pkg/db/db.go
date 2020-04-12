package db

import (
	"fmt"
	"iceforg/pkg/config"
	"reflect"
	"strings"
)

const (
	mysqlType = "mysql"
	mongoType = "mongo"
)

type initer interface {
	initer()
}

var initMap = map[string]initer{}

func registerDB(dbType string, value reflect.Value) {
	switch strings.ToLower(dbType) {
	case mysqlType:
		MysqlProvider.config = value.Interface().(*config.Mysql)
		initMap[dbType] = MysqlProvider
	default:
		panic(fmt.Sprintf("register db failed,unsupport db type(%s)", dbType))
	}
}

func registerDBSet(dbs *config.DB) {
	object := reflect.ValueOf(dbs)
	myref := object.Elem()
	typeOfType := myref.Type()
	for i := 0; i < myref.NumField(); i++ {
		if !myref.Field(i).IsNil() {
			tp := typeOfType.Field(i).Tag.Get("type")
			registerDB(tp, myref.Field(i))
		}
	}
}

func InitDB(cfg *config.DB) {
	registerDBSet(cfg)
	for _, v := range initMap {
		v.initer()
	}
}
