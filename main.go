package main

import (
	"crm/commands"
	"crm/commands/server"
	"crm/gopkg/log"
	"crm/internal/worker"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	// 创建后台工作进程
	backgroundWorker := worker.NewBackgroundWorker(10, 100)

	// 在程序退出时停止后台工作进程
	defer backgroundWorker.Stop()

	// 启动后台工作进程
	backgroundWorker.Start()

	// 提交一些示例任务
	//backgroundWorker.SubmitSampleTasks()

	app := cli.NewApp()
	app.Action = func(c *cli.Context) error {
		// 启动 Web 服务
		return server.Run(c)
	}
	app.Before = server.InitConfig
	app.After = server.Flush
	app.Commands = commands.All()
	app.Flags = server.Flags()

	if err := app.Run(os.Args); err != nil {
		log.Sugar().Fatal(err)
	}
}
