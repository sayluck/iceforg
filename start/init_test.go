package start

import (
	"iceforg/pkg/config"
	"iceforg/pkg/db"
	"testing"
)

func TestHasDBSet(t *testing.T) {
	db.InitDB(config.GetConfig().DB)
	mysql := db.GetMysqlProvider()
	t.Log(mysql.HasTable("t_user"))
}
