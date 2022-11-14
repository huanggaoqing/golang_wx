package main

import (
	"go_http/cronConfig"
	"go_http/router"
)

func main() {
	// 初始化定时任务
	c := cronConfig.CronInit()
	defer c.Stop()
	r := router.RegRouter()
	r.Run(":8066")
}
