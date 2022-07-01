package request

import (
	loggers "helle/log"
)

func init() {
	loggers.Init()
}

type Acc struct {
	Client        string `json:"client" validate:"required"`
	AccountNumber string `json:"account_number" validate:"required"`
	RequestRefnum string `json:"request_refnum" validate:"required"`
	RequestId     string `json:"request_id" validate:"required"`
}

type Name struct {
	Client        string `json:"client" validate:"required"`
	Username      string `json:"username" validate:"required"`
	RequestRefnum string `json:"request_refnum" validate:"required"`
	RequestId     string `json:"request_id" validate:"required"`
}

type InquiryHpByAccount struct{}

func SendRequest(result interface{}) {
	loggers.LogRequest("ideapad e7", "Dion", "inquiry_hp_byaccount", result, nil)

}

// func SendRequest(result interface{}) {

// 	log := logrus.New()
// 	log.SetFormatter(&logrus.JSONFormatter{
// 		FieldMap: logrus.FieldMap{
// 			logrus.FieldKeyTime:  "timestamp",
// 			logrus.FieldKeyMsg:   "message",
// 			logrus.FieldKeyLevel: "level",
// 		},
// 	})
// 	js, _ := json.Marshal(result)
// 	log.Info("request: ", string(js))
// }
