package db

import "strings"

type connect interface {
	Connect() (interface{}, error)
}

const (
	mysqlType = "mysql"
)

type DBer interface {
	GetType() string
	GetOptions()
}

func ConnectFactory(db DBer) connect {
	switch strings.ToLower(db.GetType()) {
	case mysqlType:
		return &mysql{}
	}
	return nil
}
