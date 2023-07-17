package setting

type GeneralSettingS struct {
	LogSavePath string
	LogFileName string
	LogFileExt  string
	LogMaxSize  int
	LogMaxAge   int
}

type DatabaseSettingS struct {
	// 数据库模块修改后的参数配置
	DBMaxIdleConns int
	DBMaxOpenConns int
	DBMaxLifetime  int
	DBType         string
	DBConn         string
}

type ObjectSettingS struct {
	CronSpec     string
	HospitalID   string
	InterfaceURL string
	TimeInterval int
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}
