package dekanat_events

import "encoding/json"

type Message struct {
	ReceiptHandle *string
	Timestamp     int64  `json:"timestamp"`
	Ip            string `json:"ip"`
	Referer       string `json:"referer"`
	Form          Form   `json:"form"`
}

func CreateMessage(jsonString *string, ReceiptHandle *string) (*Message, error) {
	eventMessage := Message{}
	eventMessage.ReceiptHandle = ReceiptHandle
	err := json.Unmarshal([]byte(*jsonString), &eventMessage)
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
