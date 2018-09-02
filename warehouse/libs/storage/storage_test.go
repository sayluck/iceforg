package storage

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGenerateStoragePath(t *testing.T) {
	Convey("test GenerateStoragePath", t, func() {
		Convey("test GenerateStoragePath,case one", func() {
			retS := GenerateStoragePath("fabric", "test", "3.6.9")
			t.Logf("retS:%+v\n", retS)
			So(retS, ShouldNotBeEmpty)
		})
		Convey("test GenerateStoragePath,case two", func() {
			retS := GenerateStoragePath("fabric", "", "")
			t.Logf("retS:%+v\n", retS)
			So(retS, ShouldBeEmpty)
		})
	})
}

func TestCreateContractDirector(t *testing.T) {
	Convey("test CreateContractDirector", t, func() {
		Convey("test CreateContractDirector,case one", func() {
			p := GenerateStoragePath("fabric", "test", "3.6.9")
			e := CreateContractDirector(p)
			// os.RemoveAll(p)
			So(e, ShouldBeNil)

		})
	})
}
