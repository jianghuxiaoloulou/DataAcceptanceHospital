package model

import (
	"WowjoyProject/DataAcceptanceHospital/global"
	"WowjoyProject/DataAcceptanceHospital/pkg/setting"
	"database/sql"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// PACS集成平台数据库操作
func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*sql.DB, error) {
	db, err := sql.Open(databaseSetting.DBType, databaseSetting.DBConn)
	if err != nil {
		return nil, err
	}
	// 数据库最大连接数
	db.SetConnMaxLifetime(time.Duration(databaseSetting.DBMaxLifetime) * time.Minute)
	db.SetMaxOpenConns(databaseSetting.DBMaxOpenConns)
	db.SetMaxIdleConns(databaseSetting.DBMaxIdleConns)

	return db, nil
}

// 获取该医院申请单上传成功的时间
func GetUploadApplyTime(hospitalid string) (data string) {
	global.Logger.Info("获取该医院申请单上传成功的时间: ", hospitalid)
	sqlstr := `SELECT upload_time FROM sys_dict_hospital_config WHERE hospital_id = ?`
	var err error
	err = global.DBEngine.Ping()
	if err != nil {
		global.Logger.Error(err.Error())
		global.DBEngine, _ = NewDBEngine(global.DatabaseSetting)
	}

	row := global.DBEngine.QueryRow(sqlstr, hospitalid)
	var uploadtime sql.NullString
	err = row.Scan(&uploadtime)
	if err != nil {
		global.Logger.Error(err)
		return
	}
	data = uploadtime.String
	return
}
