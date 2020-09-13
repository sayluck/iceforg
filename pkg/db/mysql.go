package db

import (
	"fmt"
	"iceforg/pkg/config"
	"sync"

	"github.com/prometheus/common/log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type mysql struct {
	config *config.Mysql
	db     *gorm.DB
	locker sync.Mutex
}

var mysqlProvider = new(mysql)

func GetMysqlProvider() *gorm.DB {
	if mysqlProvider.db == nil {
		mysqlProvider.config = config.GetConfig().DB.Mysql
		mysqlProvider.initer()
	}
	return mysqlProvider.db
}

func (m *mysql) getConnectStr() string {
	if m == nil || m.config == nil {
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
			log.Error("mysql start failed,%s", err.Error())
			panic(fmt.Sprintf("mysql start failed,%s", err.Error()))
		}

		m.db.DB().SetMaxIdleConns(m.config.MaxIdleConns)
		m.db.DB().SetMaxOpenConns(m.config.MaxOpenConns)
		m.db.LogMode(m.config.LogMode)
	}
}
