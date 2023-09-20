package dekanat_events

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var testLessonEditEvent = LessonEditEvent{
	CommonEventData: CommonEventData{
		ReceiptHandle: &testReceiptHandle,
		Timestamp:     1673000000,
		SessionId:     "00AB0000-0000-0000-0000-000CD0000AA0",
		LessonId:      "999999",
		DisciplineId:  "193000",
		Semester:      "0",
	},
	TypeId:    "12",
	Date:      "12.12.2022",
	TeacherId: "9999",
	IsDeleted: false,
}

func TestLessonEditEvent_ToMessage(t *testing.T) {
	t.Run("TestLessonEditEvent_ToMessage", func(t *testing.T) {
		originalEvent := LessonEditEvent{
			CommonEventData: CommonEventData{
				ReceiptHandle: &testReceiptHandle,
				Timestamp:     1673000000,
				SessionId:     "00AB0000-0000-0000-0000-000CD0000AA0",
				LessonId:      "999999",
				DisciplineId:  "193000",
				Semester:      "0",
			},
			TypeId:    "1",
			Date:      "12.12.2022",
			TeacherId: "9999",
			IsDeleted: false,
		}

		ExecuteTestMessageToEvent(t, originalEvent, originalEvent.ToMessage())
	})
}

func TestLessonEditEvent_GetTypeId(t *testing.T) {
	assert.Equal(t, uint8(12), testLessonEditEvent.GetTypeId())
}

func TestLessonEditEvent_GetDate(t *testing.T) {
	t.Run("not empty date", func(t *testing.T) {
		expectedDate := time.Date(2022, 12, 12, 0, 0, 0, 0, time.Local)
		assert.Equal(t, expectedDate, testLessonEditEvent.GetDate())
	})

	t.Run("empty date", func(t *testing.T) {
		assert.True(t, (&LessonEditEvent{}).GetDate().IsZero())
	})
}
