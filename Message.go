package dekanatEvents

import "encoding/json"

type Message struct {
	ReceiptHandle *string `json:"-"`
	Timestamp     int64   `json:"timestamp"`
	Ip            string  `json:"ip"`
	Referer       string  `json:"referer"`
	Form          Form    `json:"form"`
}

func CreateMessage(jsonString *string, ReceiptHandle *string) (*Message, error) {
	eventMessage := Message{}
	err := json.Unmarshal([]byte(*jsonString), &eventMessage)
	eventMessage.ReceiptHandle = ReceiptHandle
	return &eventMessage, err
}

func (message *Message) ToEvent() (event interface{}, err error) {
	return message.Form.toEvent(message.Timestamp, message.ReceiptHandle)
}

func (message *Message) ToJson() *string {
	jsonBytes, _ := json.Marshal(message)
	jsonString := string(jsonBytes)
	return &jsonString
}

func (message *Message) HasValidForm() (valid bool) {
	return message.Form.isValidEventForm()
}
