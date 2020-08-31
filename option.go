package jpushclient

type Option struct {
	SendNo          int    `json:"sendno,omitempty"`
	TimeLive        int    `json:"time_to_live,omitempty"`
	ApnsProduction  bool   `json:"apns_production"`
	ApnsCollapseID  string `json:"apns_collapse_id,omitempty"`
	OverrideMsgId   int64  `json:"override_msg_id,omitempty"`
	BigPushDuration int    `json:"big_push_duration,omitempty"`
}

func (option *Option) SetSendNo(no int) {
	option.SendNo = no
}

func (option *Option) SetTimeLive(timeLive int) {
	option.TimeLive = timeLive
}

func (option *Option) SetOverrideMsgId(id int64) {
	option.OverrideMsgId = id
}

func (option *Option) SetApns(apns bool) {
	option.ApnsProduction = apns
}

func (option *Option) SetBigPushDuration(bigPushDuration int) {
	option.BigPushDuration = bigPushDuration
}
