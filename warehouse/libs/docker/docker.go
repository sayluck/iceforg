package docker

import (
	"context"
	log "github.com/ctlog"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

var DockerClient *client.Client

func getDockerClient() *client.Client {
	cli, err := client.NewEnvClient()
	if err != nil {
		log.Fatalln("get docker client error:", err)
		panic(err)
	}
	log.Infoln("get docker client succ.")
	return cli
}

func NewClient() *client.Client {
	if DockerClient == nil {
		DockerClient = getDockerClient()
	}
	return DockerClient
}

func DockerExecCmd(containerName string, cmds []string) error {
	cfg := types.ExecConfig{
		Cmd: cmds,
	}
	ret, err := DockerClient.ContainerExecCreate(context.Background(), containerName, cfg)
	if err != nil {
		log.Errorln("ContainerExecCreate error:", err.Error())
		return err
	}
	log.Debugf("docker exec info:containerName %s,cfg:%+v\n", containerName, cfg)
	check := types.ExecStartCheck{
		Tty: true,
	}
	err = DockerClient.ContainerExecStart(context.Background(), ret.ID, check)
	if err != nil {
		log.Errorln("ContainerExecStart error:", err.Error())
		return err
	}
	log.Infof("docker exec succ:containerName %s,cfg:%+v\n", containerName, cfg)
	return nil
}
