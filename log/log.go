package log

import (
	"encoding/json"
	"fmt"
	"helle/entity/response"
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

var (
	serviceName string
	logText     *log.Logger
	logJSON     *log.Logger
)

const (
	httpRequest  = "REQUEST"
	httpResponse = "RESPONSE"
	timeformat   = "2006-01-02T15:04:05-0700"
	nameformat   = "log-2006-01-02.log"
)

func LogDebug(msg string) {
	timestamp := setLogFile()
	logText.Debug(fmt.Sprintf("%s [%s] %s", timestamp, "", msg))
}
func LogWarn(err, location string) {
	timestamp := setLogFile()
	logText.WithFields(log.Fields{
		"service":        serviceName,
		"error":          err,
		"error_location": location,
		"timestamp":      timestamp,
	}).Warn("WARNING")
}

func LogResponse(req *http.Request, deviceID, username, trx_type, code string, response *response.Response, header http.Header) {
	timestamp := setLogFile()
	// trace, _ := context.Get(req, "trace").([]interface{})
	logJSON.WithFields(log.Fields{
		"service":         serviceName,
		"http_type":       httpResponse,
		"response_header": header,
		"response_body":   response,
		"trx_type":        trx_type,
		"username":        username,
		"device_id":       deviceID,
		// "trace":           trace,
		"response_code": code,
		"timestamp":     timestamp,
	}).Info("RESPONSE")
}

func LogRequest(deviceID, username, trx_type string, request interface{}, header http.Header) {
	timestamp := setLogFile()
	mapRequest := minifyRequest(request)
	delete(mapRequest, "password")
	delete(mapRequest, "otp")
	delete(mapRequest, "key_email")
	logJSON.WithFields(log.Fields{
		"service":        serviceName,
		"http_type":      httpRequest,
		"request_header": header,
		"request_body":   mapRequest,
		"trx_type":       trx_type,
		"username":       username,
		"device_id":      deviceID,
		"timestamp":      timestamp,
	}).Info("REQUEST")

}

func setLogFile() string {
	currentTime := time.Now()
	timestamp := currentTime.Format(timeformat)
	filename := "logs/" + currentTime.Format(nameformat)
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		LogWarn(err.Error(), "set logfile")
		//fmt.Println(err)
	} else {
		logText.SetOutput(file)
		logJSON.SetOutput(file)
	}
	return timestamp
}
func setJSON() {
	logJSON = log.New()
	formatter := new(log.JSONFormatter)
	formatter.DisableTimestamp = true
	logJSON.SetFormatter(formatter)
}

func setText() {
	logText = log.New()
	formatter := new(log.TextFormatter)
	formatter.DisableTimestamp = true
	formatter.DisableQuote = true
	logText.SetFormatter(formatter)
}

func Init() {
	setText()
	setJSON()
	if os.Getenv("ENV") == "development" {
		logText.SetLevel(log.DebugLevel)
		logJSON.SetLevel(log.DebugLevel)
	} else {
		logText.SetLevel(log.InfoLevel)
		logJSON.SetLevel(log.InfoLevel)
	}

	// serviceName = constant.ServerName
}

func minifyRequest(r interface{}) map[string]interface{} {
	js, _ := json.Marshal(r)
	var m map[string]interface{}
	_ = json.Unmarshal(js, &m)
	for k, v := range m {
		s := fmt.Sprintf("%v", v)
		if len(s) > 100 {
			m[k] = "panjang"
		}
	}

	// for _, mKey := range constant.MaskedKey {
	// 	if m[mKey] != nil {
	// 		//exist, must be masked
	// 		m[mKey] = constant.MaskedStr
	// 	}
	// }

	return m
}
