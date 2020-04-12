package db

import (
	"fmt"
	"iceforg/pkg/config"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type mysql struct {
	config *config.Mysql
	db     *gorm.DB
	locker sync.Mutex
}

var MysqlProvider = new(mysql)

func GetMysqlProvider() *gorm.DB {
	if MysqlProvider.db == nil {
		MysqlProvider.initer()
	}
	return MysqlProvider.db
}

func (m *mysql) getConnectStr() string {
	if m == nil {
		return ""
	}
	return fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		m.config.UserName,
		m.config.Password,
		m.config.URL,
		m.config.DBName)
}

func (m *mysql) initer() {
	var (
		err error
	)
	if m.db == nil {
		m.locker.Lock()
		defer m.locker.Unlock()

		if m.db, err = gorm.Open(mysqlType, m.getConnectStr()); err != nil {
			// TODO add log
			panic(fmt.Sprintf("mysql start failed,%s", err.Error()))
		}

		m.db.DB().SetMaxIdleConns(m.config.MaxIdleConns)
		m.db.DB().SetMaxOpenConns(m.config.MaxOpenConns)
		m.db.LogMode(m.config.LogMode)
	}
}
