package dekanatEvents

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"reflect"
	"strconv"
)

type CommonEventData struct {
	ReceiptHandle *string
	Timestamp     int64
	SessionId     string `mapstructure:"sesID"`
	LessonId      string `mapstructure:"prti"`
	DisciplineId  string `mapstructure:"prt"`
	Semester      string `mapstructure:"hlf"`
}

func (data *CommonEventData) GetLessonId() uint {
	return parseUint(data.LessonId)
}

func (data *CommonEventData) GetDisciplineId() uint {
	return parseUint(data.DisciplineId)
}

func parseUint(s string) uint {
	value, _ := strconv.ParseUint(s, 10, 0)
	return uint(value)
}

func EventToMessage(event interface{}) *Message {
	tmpForm := map[string]interface{}{}

	message := &Message{
		Form: Form{},
	}

	_ = mapstructure.Decode(event, &tmpForm)
	for key, value := range tmpForm {
		if key == "CommonEventData" {
			commonEventDataMap := value.(map[string]interface{})

			message.ReceiptHandle = commonEventDataMap["ReceiptHandle"].(*string)
			message.Timestamp = commonEventDataMap["Timestamp"].(int64)

			delete(commonEventDataMap, "ReceiptHandle")
			delete(commonEventDataMap, "Timestamp")

			_ = mapstructure.Decode(commonEventDataMap, &message.Form)
		} else if reflect.TypeOf(value).Kind() != reflect.Map {
			message.Form[key] = fmt.Sprintf("%v", value)
		}
	}

	return message
}
