package main

import (
	"fmt"
	"hook/app"
	"hook/app/cli"
	"hook/bootstrap"
	app_router "hook/router"
	"log"
)

func init() {
	app.Init()
}

func main() {
	switch app.OPTIONS.SubCommand {
	case bootstrap.CommandServe:
		app.Boot()
		engine := app_router.SetupRouter()

		// 启动 HTTP 服务器
		log.Fatal(
			engine.Run(fmt.Sprintf(":%s", app.CFG.App.Port)).Error(),
		)
	case bootstrap.CommandDump:
		cli.Dump()
	case bootstrap.CommandSystemd:
		cli.DumpSystemd()
	default:
		panic("unknown command")
	}
}
