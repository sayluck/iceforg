package docker

import (
	"github.com/iceforg/warehouse/libs/utils"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDockerExecCmd(t *testing.T) {
	Convey("docker client exec a test cmd", t, func() {
		New()
		// container must be running
		Convey("docker client exec a test cmd,case one.", func() {
			testCmds := []string{"sleep", "5"}
			e := DockerExecCmd(utils.CLICONTAINER, testCmds)
			So(e, ShouldBeNil)
		})

		// exec a bin file
		Convey("docker client exec a test cmd,case two.", func() {
			testCmds := []string{"/home/main"}
			e := DockerExecCmd(utils.CLICONTAINER, testCmds)
			So(e, ShouldBeNil)
		})
	})
}
