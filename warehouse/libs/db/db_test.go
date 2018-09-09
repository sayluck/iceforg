package db

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

const TEST_DB_TABLE_NAME = "t_test"

type testT struct {
	Id   int    `json:"id",gorm:"AUTO_INCREMENT"`
	Name string `json:"name",gorm:"column:name"`
	Pswd string `json:"pswd",gorm:"column:pswd"`
}

func TestGetDBClient(t *testing.T) {
	Convey("test db connect", t, func() {
		GetDBClient()
	})
}

func TestDB(t *testing.T) {
	Convey("test db is working", t, func() {
		GetDBClient()
		testData := testT{
			Name: "testUser",
			Pswd: "qwe123",
		}
		Convey("test create data", func() {
			ret := Db.Table(TEST_DB_TABLE_NAME).Create(&testData)
			if ret.Error != nil {
				t.Fatal("create db error.")
			}
		})
		Convey("test update data", func() {
			testData1 := testT{
				Pswd: "qwe5223",
			}
			ret := Db.Table(TEST_DB_TABLE_NAME).Model(&testT{}).Where("id=?", 5).Update(&testData1)
			if ret.Error != nil {
				t.Fatal("update db error.", ret.Error)
			}
		})
		//Convey("test delete data", func() {
		//	testData.Id = 4
		//	ret := Db.Table(TEST_DB_TABLE_NAME).Delete(&testData)
		//	if ret.Error != nil {
		//		t.Fatal("delete db error.")
		//	}
		//})
	})
}
