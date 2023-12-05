package dekanatEvents

import (
	"github.com/mitchellh/mapstructure"
	"strconv"
	"strings"
	"time"
)

const LessonEditEventActionNumber = 10

type LessonEditEvent struct {
	CommonEventData
	// n=4 and 	action: edit
	TypeId    string `mapstructure:"tzn"`
	Date      string `mapstructure:"date_z"`
	TeacherId string `mapstructure:"teacher"`
	IsDeleted bool   `mapstructure:"-"`
}

func createLessonEditEvent(form Form, eventData *CommonEventData) (event LessonEditEvent, err error) {
	err = mapstructure.Decode(form, &event)
	event.CommonEventData = *eventData
	return
}

func (event *LessonEditEvent) ToMessage() *Message {
	message := EventToMessage(event)
	message.Form["n"] = strconv.Itoa(LessonEditEventActionNumber)
	message.Form["action"] = LessonEditFormAction

	return message
}

func (event *LessonEditEvent) GetTypeId() uint8 {
	return uint8(parseUint(event.TypeId))
}

func (event *LessonEditEvent) GetDate() (date time.Time) {
	dateParts := strings.Split(event.Date, ".")
	if len(dateParts) == 3 {
		year, _ := strconv.Atoi(dateParts[2])
		month, _ := strconv.Atoi(dateParts[1])
		day, _ := strconv.Atoi(dateParts[0])
		date = time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	}
	return
}
