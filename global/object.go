package global

type ApplyAndDicomInfo struct {
	Bizno      string  `json:"bizno"`
	Time       string  `json:"time`
	HospitalId string  `json:"hospital_id"`
	Param      ReqInfo `json:"req_info"`
}

type ReqInfo struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}
