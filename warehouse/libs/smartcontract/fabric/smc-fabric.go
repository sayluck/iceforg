package fabric

import (
	"github.com/iceforg/warehouse/libs/docker"
	"github.com/iceforg/warehouse/libs/utils"
)

type SmcFabric struct {
	Id          int    `json:"id",gorm:"AUTO_INCREMENT"`
	BcType      string `json:"bctype" gorm:"column:bc_type"`
	Name        string `json:"name" gorm:"column:smc_name"`
	Version     string `json:"version" gorm:"column:version"`
	ExecName    string `json:"execName" gorm:"-"`
	Lang        string `json:"lang" gorm:"column:lang"`
	Path        string `json:"path" gorm:"column:path"`
	State       string `json:"state" gorm:"column:state"`
	Descripe    string `json:"descripe" gorm:"column:descripe"`
	OperateInfo string `json:"OperateInfo" gorm:"column:operate_info"`
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
