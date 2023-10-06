package dekanatEvents

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

func TestCommonEventData_GetSemester(t *testing.T) {
	semester1 := CommonEventData{Semester: "0"}
	semester2 := CommonEventData{Semester: "1"}

	assert.Equal(t, uint(1), semester1.GetSemester())
	assert.Equal(t, uint(2), semester2.GetSemester())
}
