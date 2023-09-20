package dekanat_events

import (
	"testing"
)

func TestLessonCreateEvent_ToMessage(t *testing.T) {
	originalEvent := LessonCreateEvent{
		CommonEventData: CommonEventData{
			Timestamp:    1673000000,
			SessionId:    "00AB0000-0000-0000-0000-000CD0000AA0",
			LessonId:     "0",
			DisciplineId: "193000",
			Semester:     "0",
		},
		TypeId:    "1",
		Date:      "23.12.2022",
		TeacherId: "9999",
	}

	ExecuteTestMessageToEvent(t, originalEvent, originalEvent.ToMessage())
}
