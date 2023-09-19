package dekanat_events

import (
	"github.com/mitchellh/mapstructure"
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

func (event *LessonCreateEvent) ToMessage() (message Message) {
	return EventToMessage(event)
}
