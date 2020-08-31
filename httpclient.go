package jpushclient

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	Charset                  = "UTF-8"
	ContentTypeJson          = "application/json"
	DefaultConnectionTimeout = 20 //seconds
	DefaultSocketTimeout     = 30 // seconds
)

func SendPostString(url, content, authCode string) (string, error) {
	req := Post(url)
	req.SetTimeout(DefaultConnectionTimeout*time.Second, DefaultSocketTimeout*time.Second)
	req.Header("Connection", "Keep-Alive")
	req.Header("Charset", Charset)
	req.Header("Authorization", authCode)
	req.Header("Content-Type", ContentTypeJson)
	req.SetProtocolVersion("HTTP/1.1")
	req.Body(content)

	return req.String()
}

func SendPostBytes(url string, content []byte, authCode string) (string, error) {
	req := Post(url)
	req.SetTimeout(DefaultConnectionTimeout*time.Second, DefaultSocketTimeout*time.Second)
	req.Header("Connection", "Keep-Alive")
	req.Header("Charset", Charset)
	req.Header("Authorization", authCode)
	req.Header("Content-Type", ContentTypeJson)
	req.SetProtocolVersion("HTTP/1.1")
	req.Body(content)

	return req.String()
}

func SendPostBytes2(url string, data []byte, authCode string) (string, error) {

	client := &http.Client{}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.Header.Add("Charset", Charset)
	req.Header.Add("Authorization", authCode)
	req.Header.Add("Content-Type", ContentTypeJson)
	resp, err := client.Do(req)

	if err != nil {
		if resp != nil {
			_ = resp.Body.Close()
		}
		return "", err
	}
	if resp == nil {
		return "", nil
	}

	defer resp.Body.Close()
	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(r), nil
}
