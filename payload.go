package jpushclient

import (
	"encoding/json"
)

type PayLoad struct {
	Platform     interface{} `json:"platform"`
	Audience     interface{} `json:"audience"`
	Notification interface{} `json:"notification,omitempty"`
	Message      interface{} `json:"message,omitempty"`
	Options      *Option     `json:"options,omitempty"`
}

func NewPushPayLoad() *PayLoad {
	pl := &PayLoad{}
	o := &Option{}
	o.ApnsProduction = false
	pl.Options = o
	return pl
}

func (payload *PayLoad) SetPlatform(pf *Platform) {
	payload.Platform = pf.Os
}

func (payload *PayLoad) SetAudience(ad *Audience) {
	payload.Audience = ad.Object
}

func (payload *PayLoad) SetOptions(o *Option) {
	payload.Options = o
}

func (payload *PayLoad) SetMessage(m *Message) {
	payload.Message = m
}

func (payload *PayLoad) SetNotice(notice *Notice) {
	payload.Notification = notice
}

func (payload *PayLoad) ToBytes() ([]byte, error) {
	content, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return content, nil
}
