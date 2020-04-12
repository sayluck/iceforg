package db

import (
	"iceforg/pkg/config"
	"testing"
)

func Test_mysql_getConnectStr(t *testing.T) {
	my := new(mysql)
	my.config = &config.Mysql{
		URL:          "url",
		UserName:     "username",
		Password:     "password",
		DBName:       "dbname",
		MaxIdleConns: 0,
		MaxOpenConns: 0,
		LogMode:      false,
	}
	str := my.getConnectStr()
	if str != "username:password@url/dbname?charset=utf8&parseTime=True&loc=Local" {
		t.Fatal(str)
	}
}
