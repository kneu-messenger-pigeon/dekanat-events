package dekanatEvents

import (
	"github.com/mitchellh/mapstructure"
	"strconv"
)

type LessonCreateEvent struct {
	CommonEventData
	// n=4 and action:insert
	TypeId    string `mapstructure:"tzn"`
	Date      string `mapstructure:"date_z"`
	TeacherId string `mapstructure:"teacher"`
}

func createLessonCreateEvent(form Form, eventData *CommonEventData) (event LessonCreateEvent, err error) {
	err = mapstructure.Decode(form, &event)
	event.CommonEventData = *eventData
	return
}

func (event *LessonCreateEvent) ToMessage() *Message {
	message := EventToMessage(event)
	message.Form["n"] = strconv.Itoa(LessonFormActionNumber)
	message.Form["action"] = LessonInsertFormAction

	return message
}
