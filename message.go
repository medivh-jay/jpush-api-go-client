package jpushclient

type Message struct {
	Content     string                 `json:"msg_content"`
	Title       string                 `json:"title,omitempty"`
	ContentType string                 `json:"content_type,omitempty"`
	Extras      map[string]interface{} `json:"extras,omitempty"`
}

func (message *Message) SetContent(c string) {
	message.Content = c

}

func (message *Message) SetTitle(title string) {
	message.Title = title
}

func (message *Message) SetContentType(t string) {
	message.ContentType = t
}

func (message *Message) AddExtras(key string, value interface{}) {
	if message.Extras == nil {
		message.Extras = make(map[string]interface{})
	}
	message.Extras[key] = value
}
