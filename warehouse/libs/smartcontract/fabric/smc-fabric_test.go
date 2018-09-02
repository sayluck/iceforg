package fabric

import "testing"
import (
	"github.com/iceforg/warehouse/libs/docker"
	"github.com/iceforg/warehouse/libs/storage"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSmcExec(t *testing.T) {
	Convey("test smc exec", t, func() {
		docker.NewClient()
		Convey("test smc exec,case one", func() {
			retS := storage.GenerateStoragePath("fabric", "test", "3.6.9") + "example.go"
			e := smcExec(retS)
			t.Logf("retS:%+v\n", retS)
			So(e, ShouldBeNil)
		})
	})
}
