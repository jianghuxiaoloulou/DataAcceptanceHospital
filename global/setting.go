package global

import (
	"WowjoyProject/DataAcceptanceHospital/pkg/logger"
	"WowjoyProject/DataAcceptanceHospital/pkg/setting"
)

var (
	GeneralSetting  *setting.GeneralSettingS
	DatabaseSetting *setting.DatabaseSettingS
	ObjectSetting   *setting.ObjectSettingS
	Logger          *logger.Logger
)
