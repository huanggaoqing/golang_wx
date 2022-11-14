package cronConfig

import (
	"github.com/robfig/cron/v3"
	"go_http/utils"
)

var cronInstance *cron.Cron

func init() {
	cronInstance = cron.New()
}

func CronInit() *cron.Cron {
	// 定时发送《早上好》消息
	cronInstance.AddFunc("30 8 * * *", utils.StartSend)
	// 任务开始
	cronInstance.Start()
	return cronInstance
}
