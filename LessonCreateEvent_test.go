package dekanat_events

import (
	"testing"
)

func TestLessonCreateEvent_ToMessage(t *testing.T) {
	receipt := "asdads"

	createEvent := LessonCreateEvent{
		TypeId: "211",
		CommonEventData: CommonEventData{
			ReceiptHandle: &receipt,
			LessonId:      "4989",
		},
	}

	createEvent.ToMessage()

	//	fmt.Printf("%+v", message)

}
