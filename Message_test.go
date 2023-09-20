package dekanat_events

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

const EventMessageJSON = `{
	"timestamp": 1673000000,
	"ip": "127.0.0.1",
	"referer": "http://example.com",
	"form": %s
}`

const LessonCreateEventFormJSON = `{
	"hlf":"0",
	"prt":"193000",
	"prti":"0",
	"teacher":"9999",
	"action":"insert",
	"n":"10",
	"sesID":"00AB0000-0000-0000-0000-000CD0000AA0",
	"m":"-1",
	"date_z":"23.12.2022",
	"tzn":"1",
	"result":"3",
	"grade":""
}`

const LessonEditEventFormJSON = `{
	"hlf":"0",
	"prt":"193000",
	"prti":"999999",
	"teacher":"9999",
	"action":"edit",
	"n":"10",
	"sesID":"00AB0000-0000-0000-0000-000CD0000AA0",
	"m":"-1",
	"date_z":"12.12.2022",
	"tzn":"1",
	"result":"",
	"grade":"2"
}`

const LessonDeletedEventFormJSON = `{
	"sesID":"00AB0000-0000-0000-0000-000CD0000AA0",
	"n":"11",
	"action":"delete",
	"prti":"999999",
	"prt":"193000",
	"d1":"",
	"d2":"",
	"m":"-1",
	"hlf":"0",
	"course":"undefined"
}`

const ScoreEditEventFormJSON = `{
        "hlf":"0",
        "prt":"188619",
        "prti":"999999",
        "action":"",
        "n":"4",
        "sesID":"99FED80A-2E33-40CB-9CEF-01E25B5AA66B",
        "d1":"09.09.2022",
        "course":"3",
        "m":"-1",
        "d2":"18.12.2022",
        "st110030_1-999999":"",
        "st110043_2-999999":"",
        "st110044_1-999999":"-11",
        "st110044_2-999999":"",
        "st110054_1-999999":"нб/нп",
        "st110054_2-999999":"",
        "st110055_1-999999":"",
        "st110055_2-999999":"",
        "AddEstim":"0"
    }`

var testReceiptHandle = "ReceiptHandle"

func TestMakeEventFromMessageJson(t *testing.T) {
	t.Run("Fetch LessonCreateEvent", func(t *testing.T) {
		expectedEvent := LessonCreateEvent{
			CommonEventData: CommonEventData{
				ReceiptHandle: &testReceiptHandle,
				Timestamp:     1673000000,
				SessionId:     "00AB0000-0000-0000-0000-000CD0000AA0",
				LessonId:      "0",
				DisciplineId:  "193000",
				Semester:      "0",
			},
			TypeId:    "1",
			Date:      "23.12.2022",
			TeacherId: "9999",
		}

		actualEvent := transformMessageJsonToEvent(t, LessonCreateEventFormJSON)

		assert.IsType(t, LessonCreateEvent{}, actualEvent)
		assert.Equal(t, expectedEvent, actualEvent)

		event := actualEvent.(LessonCreateEvent)
		assert.Equal(t, uint(0), event.GetLessonId())
		assert.Equal(t, uint(193000), event.GetDisciplineId())
		assert.Equal(t, "0", event.Semester)
	})

	t.Run("Fetch LessonEditEvent", func(t *testing.T) {
		expectedEvent := LessonEditEvent{
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

		actualEvent := transformMessageJsonToEvent(t, LessonEditEventFormJSON)

		assert.IsType(t, LessonEditEvent{}, actualEvent)
		assert.Equal(t, expectedEvent, actualEvent)

		event := actualEvent.(LessonEditEvent)
		assert.Equal(t, uint(999999), event.GetLessonId())
		assert.Equal(t, uint(193000), event.GetDisciplineId())
		assert.Equal(t, "0", event.Semester)
	})

	t.Run("Fetch LessonDeletedEvent", func(t *testing.T) {
		expectedEvent := LessonDeletedEvent{
			CommonEventData: CommonEventData{
				ReceiptHandle: &testReceiptHandle,
				Timestamp:     1673000000,
				SessionId:     "00AB0000-0000-0000-0000-000CD0000AA0",
				LessonId:      "999999",
				DisciplineId:  "193000",
				Semester:      "0",
			},
		}

		actualEvent := transformMessageJsonToEvent(t, LessonDeletedEventFormJSON)

		assert.IsType(t, LessonDeletedEvent{}, actualEvent)
		assert.Equal(t, expectedEvent, actualEvent)
		event := actualEvent.(LessonDeletedEvent)

		assert.Equal(t, uint(999999), event.GetLessonId())
		assert.Equal(t, uint(193000), event.GetDisciplineId())
		assert.Equal(t, "0", event.Semester)

		fmt.Printf("%#v\n", event)
	})

	t.Run("Fetch ScoreEditEvent", func(t *testing.T) {
		expectedEvent := ScoreEditEvent{
			CommonEventData: CommonEventData{
				ReceiptHandle: &testReceiptHandle,
				Timestamp:     1673000000,
				SessionId:     "99FED80A-2E33-40CB-9CEF-01E25B5AA66B",
				LessonId:      "999999",
				DisciplineId:  "188619",
				Semester:      "0",
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

		actualEvent := transformMessageJsonToEvent(t, ScoreEditEventFormJSON)

		assert.IsType(t, ScoreEditEvent{}, actualEvent)
		assert.Equal(t, expectedEvent, actualEvent)
		event := actualEvent.(ScoreEditEvent)

		assert.Equal(t, uint(999999), event.GetLessonId())
		assert.Equal(t, uint(188619), event.GetDisciplineId())
		assert.Equal(t, "0", event.Semester)
		assert.Equal(t, "-11", event.Scores[110044][1])
		assert.Equal(t, "", event.Scores[110044][2])
		assert.Equal(t, "нб/нп", event.Scores[110054][1])
		assert.Equal(t, "", event.Scores[110054][2])
	})

	t.Run("Empty Form", func(t *testing.T) {
		messageJson := fmt.Sprintf(EventMessageJSON, "{}")
		receiptHandle := "ReceiptHandle"

		message, err := CreateMessage(&messageJson, &receiptHandle)
		assert.False(t, message.HasValidForm())
		assert.NoError(t, err)

		event, err := message.ToEvent()

		assert.Error(t, err)
		assert.Nil(t, event)
	})

	t.Run("Wrong form", func(t *testing.T) {
		eventFormJSON := strings.Replace(ScoreEditEventFormJSON, `"st`, `"__`, -1)

		messageJson := fmt.Sprintf(EventMessageJSON, eventFormJSON)
		receiptHandle := "ReceiptHandle"

		message, err := CreateMessage(&messageJson, &receiptHandle)
		assert.True(t, message.HasValidForm())
		assert.NoError(t, err)

		event, err := message.ToEvent()

		assert.Error(t, err)
		assert.Nil(t, event)
	})

	t.Run("Fetch wrong message", func(t *testing.T) {
		messageJson := fmt.Sprintf(EventMessageJSON, "test")
		receiptHandle := "ReceiptHandle"

		message, err := CreateMessage(&messageJson, &receiptHandle)
		assert.Error(t, err)
		assert.False(t, message.HasValidForm())
		assert.NotNil(t, message)
		assert.Equal(t, "ReceiptHandle", *message.ReceiptHandle)

		event, _ := message.ToEvent()
		assert.Nil(t, event)
	})
}

func transformMessageJsonToEvent(t *testing.T, formJSON string) interface{} {
	messageJson := fmt.Sprintf(EventMessageJSON, formJSON)

	message, err := CreateMessage(&messageJson, &testReceiptHandle)
	assert.NoError(t, err)

	assert.True(t, message.HasValidForm())
	event, err := message.ToEvent()
	assert.NoError(t, err)

	return event
}

func ExecuteTestMessageToEvent(t *testing.T, originalEvent interface{}, message *Message) {
	json := message.ToJson()

	message, err := CreateMessage(json, nil)
	assert.NoError(t, err)

	parsedEvent, err := message.ToEvent()
	assert.NoError(t, err)
	assert.IsType(t, originalEvent, parsedEvent)
	assert.Equal(t, originalEvent, parsedEvent)
}
