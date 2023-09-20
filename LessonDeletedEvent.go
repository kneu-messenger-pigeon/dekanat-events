package dekanat_events

import (
	"github.com/mitchellh/mapstructure"
	"strconv"
)

type LessonDeletedEvent struct {
	CommonEventData
	// n = 11
}

func createLessonDeleteEvent(form Form, eventData *CommonEventData) (event LessonDeletedEvent, err error) {
	err = mapstructure.Decode(form, &event)
	event.CommonEventData = *eventData
	return
}

func (event *LessonDeletedEvent) ToMessage() *Message {
	message := EventToMessage(event)
	message.Form["n"] = strconv.Itoa(LessonDeleteFormActionNumber)
	message.Form["action"] = LessonDeleteFormAction
	return message
}
