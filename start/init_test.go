package start

import (
	"iceforg/pkg/db"
	"testing"
)

func TestHasDBSet(t *testing.T) {
	db.InitDB(cfg.DB)
	mysql := db.GetMysqlProvider()
	t.Log(mysql.HasTable("t_user"))
}
