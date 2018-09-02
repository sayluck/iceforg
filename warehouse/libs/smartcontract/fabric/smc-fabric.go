package fabric

import (
	"github.com/iceforg/warehouse/libs/docker"
	"github.com/iceforg/warehouse/libs/utils"
)

type SmcFabric struct {
	Name     string `json:"name"`
	Version  string `json:"version"`
	ExecName string `json:"execName"`
	Lang     string `json:"lang"`
	Path     string `json:"path"`
	Descripe string `json:"descripe"`
}

func (smc *SmcFabric) Query() error {
	return nil
}

func (smc *SmcFabric) Invoke() error {
	return nil
}

func (smc *SmcFabric) Install() error {
	// exec contract
	return nil
}

// smcExec: for test smartcontract,exec smc (golang)
func smcExec(path string) error {
	cmds := []string{
		path,
	}
	return docker.DockerExecCmd(utils.CLICONTAINER, cmds)
}
