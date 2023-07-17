package init

import (
	"WowjoyProject/DataAcceptanceHospital/global"
	"WowjoyProject/DataAcceptanceHospital/internal/model"
	"WowjoyProject/DataAcceptanceHospital/pkg/logger"
	"WowjoyProject/DataAcceptanceHospital/pkg/setting"
	"io"
	"log"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
)

func InitSetup() {
	ReadSetup()
}

func SetupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("General", &global.GeneralSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Object", &global.ObjectSetting)
	if err != nil {
		return err
	}
	return nil
}

func SetupLogger() error {
	lunberLogger := &lumberjack.Logger{
		Filename:  global.GeneralSetting.LogSavePath + "/" + global.GeneralSetting.LogFileName + global.GeneralSetting.LogFileExt,
		MaxSize:   global.GeneralSetting.LogMaxSize,
		MaxAge:    global.GeneralSetting.LogMaxAge,
		LocalTime: true,
	}
	global.Logger = logger.NewLogger(io.MultiWriter(lunberLogger, os.Stdout), "", log.LstdFlags).WithCaller(2)
	return nil
}

func SetupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func ReadSetup() {
	err := SetupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	err = SetupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
	err = SetupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
}
