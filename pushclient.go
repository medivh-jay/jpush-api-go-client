package jpushclient

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

const (
	SuccessFlag  = "msg_id"
	HostNameSsl  = "https://api.jpush.cn/v3/push"
	HostSchedule = "https://api.jpush.cn/v3/schedules"
	HostReport   = "https://report.jpush.cn/v3/received"
	Base64Table  = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
)

var base64Coder = base64.NewEncoding(Base64Table)

type PushClient struct {
	MasterSecret string
	AppKey       string
	AuthCode     string
	BaseUrl      string
}

func NewPushClient(secret, appKey string) *PushClient {
	//base64
	auth := "Basic " + base64Coder.EncodeToString([]byte(appKey+":"+secret))
	pusher := &PushClient{secret, appKey, auth, HostNameSsl}
	return pusher
}

func (client *PushClient) Send(data []byte) (string, error) {
	return client.SendPushBytes(data)
}
func (client *PushClient) CreateSchedule(data []byte) (string, error) {
	// this.BaseUrl = HOST_SCHEDULE
	return client.SendScheduleBytes(data, HostSchedule)
}
func (client *PushClient) DeleteSchedule(id string) (string, error) {
	// this.BaseUrl = HOST_SCHEDULE
	return client.SendDeleteScheduleRequest(id, HostSchedule)
}
func (client *PushClient) GetSchedule(id string) (string, error) {
	// GET https://api.jpush.cn/v3/schedules/{schedule_id}
	// this.BaseUrl = HOST_SCHEDULE
	return client.SendGetScheduleRequest(id, HostSchedule)

}
func (client *PushClient) GetReport(msg_ids string) (string, error) {
	// this.BaseUrl = HOST_REPORT
	return client.SendGetReportRequest(msg_ids, HostReport)
}
func (client *PushClient) SendPushString(content string) (string, error) {
	ret, err := SendPostString(client.BaseUrl, content, client.AuthCode)
	if err != nil {
		return ret, err
	}
	if strings.Contains(ret, "msg_id") {
		return ret, nil
	} else {
		return "", errors.New(ret)
	}
}

func (client *PushClient) SendPushBytes(content []byte) (string, error) {
	//ret, err := SendPostBytes(this.BaseUrl, content, this.AuthCode)
	ret, err := SendPostBytes2(client.BaseUrl, content, client.AuthCode)
	if err != nil {
		return ret, err
	}
	if strings.Contains(ret, "msg_id") {
		return ret, nil
	} else {
		return "", errors.New(ret)
	}
}

func (client *PushClient) SendScheduleBytes(content []byte, url string) (string, error) {
	ret, err := SendPostBytes2(url, content, client.AuthCode)
	if err != nil {
		return ret, err
	}
	if strings.Contains(ret, "schedule_id") {
		return ret, nil
	} else {
		return "", errors.New(ret)
	}

}

func (client *PushClient) SendGetReportRequest(msg_ids string, url string) (string, error) {
	return Get(url).SetBasicAuth(client.AppKey, client.MasterSecret).Param("msg_ids", msg_ids).String()
}

func UnmarshalResponse(rsp string) (map[string]interface{}, error) {
	mapRs := map[string]interface{}{}
	if len(strings.TrimSpace(rsp)) == 0 {
		return mapRs, nil
	}
	err := json.Unmarshal([]byte(rsp), &mapRs)
	if err != nil {
		return nil, err
	}
	if _, ok := mapRs["error"]; ok {
		return nil, fmt.Errorf(rsp)
	}
	return mapRs, nil
}

func (client *PushClient) SendDeleteScheduleRequest(schedule_id string, url string) (string, error) {
	rsp, err := Delete(strings.Join([]string{url, schedule_id}, "/")).Header("Authorization", client.AuthCode).String()
	if err != nil {
		return "", err
	}
	_, err = UnmarshalResponse(rsp)
	if err != nil {
		return "", err
	}
	return rsp, nil
}
func (client *PushClient) SendGetScheduleRequest(schedule_id string, url string) (string, error) {
	rsp, err := Get(strings.Join([]string{url, schedule_id}, "/")).Header("Authorization", client.AuthCode).String()
	if err != nil {
		return "", err
	}
	_, err = UnmarshalResponse(rsp)
	if err != nil {
		return "", err
	}
	return rsp, nil
}
