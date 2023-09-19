package dekanat_events

import "github.com/mitchellh/mapstructure"

type LessonDeletedEvent struct {
	CommonEventData
	// n = 11
}

func createLessonDeleteEvent(form Form, eventData *CommonEventData) (event LessonDeletedEvent, err error) {
	err = mapstructure.Decode(form, &event)
	event.CommonEventData = *eventData
	return
}

func (event *LessonDeletedEvent) ToMessage() (message Message) {
	return EventToMessage(event)
}
