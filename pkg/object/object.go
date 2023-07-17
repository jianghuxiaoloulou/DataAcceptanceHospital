package object

import (
	"WowjoyProject/DataAcceptanceHospital/global"
	"WowjoyProject/DataAcceptanceHospital/internal/model"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

// 调用PACS集成平台接口上传数据
func UploadApply() {
	// 获取数据库中申请单上传成功的时间
	data := model.GetUploadApplyTime(global.ObjectSetting.HospitalID)
	global.Logger.Debug("获取到上传成功的时间是：", data)
	// 请求接口上传数据
	PostPacsApplyAndDicomInfo(data)

}

func PostPacsApplyAndDicomInfo(data string) {
	global.Logger.Debug("开始调用第三方PACS申请单和影像数据上传接口")

	startTime, _ := time.Parse("2006-01-02 15:04:05", data)
	endTime := startTime.Add(time.Duration(global.ObjectSetting.TimeInterval) * time.Minute).Format("2006-01-02 15:04:05")
	global.Logger.Debug("获取申请单数据的时间范围：", data, " -- ", endTime)

	object := global.ApplyAndDicomInfo{
		Bizno:      "1002",
		Time:       time.Now().Format("20060102150405"),
		HospitalId: global.ObjectSetting.HospitalID,
		Param: global.ReqInfo{
			StartDate: data,
			EndDate:   endTime,
		},
	}
	reqdata, err := json.Marshal(object)
	if err != nil {
		global.Logger.Error(err)
		return
	}
	url := global.ObjectSetting.InterfaceURL
	global.Logger.Debug("操作的URL: ", url)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqdata))
	if err != nil {
		global.Logger.Error(err)
		return
	}
	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		global.Logger.Error(err)
		return
	}
	global.Logger.Info("resp.Body: ", string(resp_body))
}
