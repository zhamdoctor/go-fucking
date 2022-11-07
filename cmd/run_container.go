package main

import (
	"awesomeProject/container"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
)

const usage = "container runtime implementation"

func main() {
	app := cli.NewApp()
	app.Name = "mydocker"
	app.Usage = usage
	app.Commands = []cli.Command{}
	app.Before = func(context *cli.Context) error {
		log.SetFormatter(&log.JSONFormatter{})
		log.SetOutput(os.Stdout)
		return nil
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}

var runCommand = cli.Command{
	Name:  "run",
	Usage: "创建容器并通过ns和cgroup限制资源,docker run -ti [command]",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "ti",
			Usage: "enable tty",
		},
	},
	//判断参数是否含command,获取用户指定command,调用run function准备启动容器
	Action: func(context *cli.Context) error {
		if len(context.Args()) < 1 {
			return fmt.Errorf("missing container command")
		}
		cmd := context.Args().Get(0)
		tty := context.Bool("ti")
		Run(tty, cmd)
		return nil
	},
}

var initCommand = cli.Command{
	Name:  "init",
	Usage: "init container process run users process in container,inner function only for inner",
	//获取传递过来的command参数 执行容器初始化操作
	Action: func(context *cli.Context) error {
		log.Info("init come on")
		cmd := context.Args().Get(0)
		log.Infof("command %s", cmd)
		container.RunContainerInitProcess(cmd, nil)
		return nil
	},
}

func Run(tty bool, command string) {
	parent := container.NewParentProcess(tty, command)
	if err := parent.Start(); err != nil {
		log.Error(err)
	}
	parent.Wait()
	os.Exit(-1)
}
