package main

import (
	"WowjoyProject/DataAcceptanceHospital/global"
	initialize "WowjoyProject/DataAcceptanceHospital/internal/init"
	"WowjoyProject/DataAcceptanceHospital/pkg/object"

	"github.com/robfig/cron"
)

// @title 模拟第三方PACS上传申请单数据
// @version 1.0.0.0
// @description 模拟第三方PACS上传申请单数据
// @termsOfService https://github.com/jianghuxiaoloulou/DataAcceptanceSystem.git
func main() {
	initialize.InitSetup()
	global.Logger.Info("***开始执行定时任务上传任城区妇保院申请单数据***")
	TimerTask()
}

func TimerTask() {
	MyCron := cron.New()
	MyCron.AddFunc(global.ObjectSetting.CronSpec, func() {
		global.Logger.Info("开始执行定时任务")
		object.UploadApply()

	})
	MyCron.Start()
	defer MyCron.Stop()
	select {}
}
