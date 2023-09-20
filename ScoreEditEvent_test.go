package dekanat_events

import (
	"testing"
)

func TestScoreEditEvent_ToMessage(t *testing.T) {
	t.Run("TestLessonEditEvent_ToMessage", func(t *testing.T) {
		originalEvent := ScoreEditEvent{
			CommonEventData: CommonEventData{
				Timestamp:    1673000000,
				SessionId:    "99FED80A-2E33-40CB-9CEF-01E25B5AA66B",
				LessonId:     "999999",
				DisciplineId: "188619",
				Semester:     "0",
			},
			Date: "18.12.2022",
			Scores: map[int]map[uint8]string{
				110030: {
					1: "",
				},
				110043: {
					2: "",
				},
				110044: {
					1: "-11",
					2: "",
				},
				110054: {
					1: "нб/нп",
					2: "",
				},
				110055: {
					1: "",
					2: "",
				},
			},
		}

		ExecuteTestMessageToEvent(t, originalEvent, originalEvent.ToMessage())
	})
}
