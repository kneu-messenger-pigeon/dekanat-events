package dekanat_events

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testCommonEventDataEvent = CommonEventData{
	ReceiptHandle: &testReceiptHandle,
	Timestamp:     1673000000,
	SessionId:     "00AB0000-0000-0000-0000-000CD0000AA0",
	LessonId:      "999999",
	DisciplineId:  "193000",
	Semester:      "0",
}

func TestCommonEventData_GetDisciplineId(t *testing.T) {
	assert.Equal(t, uint(193000), testCommonEventDataEvent.GetDisciplineId())
}

func TestCommonEventData_GetLessonId(t *testing.T) {
	assert.Equal(t, uint(999999), testCommonEventDataEvent.GetLessonId())
}
