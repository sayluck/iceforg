package docker

import (
	log "github.com/ctlog"
	"github.com/docker/docker/client"
)

func GetDockerClient() *client.Client {
	cli, err := client.NewEnvClient()
	if err != nil {
		log.Fatalln("get docker client error:", err)
		panic(err)
	}
	log.Infoln("get docker client succ.")
	return cli
}

//func dockerExec(cmd []string) {
//
//}

//
//func dockerClient() {
//
//	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
//	if err != nil {
//		panic(err)
//	}
//
//	for _, container := range containers {
//		fmt.Printf("%s %s %s\n", container.Names, container.ID[:10], container.Image)
//	}
//
//	cfg := types.ExecConfig{
//		Cmd: []string{"sleep", "10"},
//	}
//	ret, err := cli.ContainerExecCreate(context.Background(), "cli", cfg)
//	if err != nil {
//		fmt.Println("exec error:", err.Error())
//		return
//	}
//	fmt.Printf("msg:%+v\n", ret)
//
//	check := types.ExecStartCheck{
//		Tty: true,
//	}
//	err = cli.ContainerExecStart(context.Background(), ret.ID, check)
//	if err != nil {
//		fmt.Println("start error:", err.Error())
//		return
//	}
//}
