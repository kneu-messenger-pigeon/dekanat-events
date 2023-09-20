package dekanatEvents

import "testing"

func TestLessonDeletedEvent_ToMessage(t *testing.T) {
	t.Run("TestLessonEditEvent_ToMessage", func(t *testing.T) {
		originalEvent := LessonDeletedEvent{
			CommonEventData: CommonEventData{
				Timestamp:    1673000000,
				SessionId:    "00AB0000-0000-0000-0000-000CD0000AA0",
				LessonId:     "999999",
				DisciplineId: "193000",
				Semester:     "0",
			},
		}

		ExecuteTestMessageToEvent(t, originalEvent, originalEvent.ToMessage())
	})
}
