package dekanatEvents

import (
	"github.com/mitchellh/mapstructure"
	"regexp"
	"strconv"
)

type ScoreEditEvent struct {
	CommonEventData
	// n = 4 and have keys "^st[0-9]+_(1|2)_[0-9]+$" st119906_2-2616031
	Date   string `mapstructure:"d2"`
	Scores map[int]map[uint8]string
	/* map from form:
	st119905_1-2616031:
	st119905_2-2616031: нб/нп
	st119906_1-2616031:
	st119909_1-2616031:	6
	st119909_2-2616031:
	st119910_1-2616031:
	st119910_2-2616031:
	*/

}

var StudentScoreFieldRegexp = regexp.MustCompile("^st([0-9]+)_(1|2)-[0-9]+$")

func createScoreEditEvent(form Form, eventData *CommonEventData) (event ScoreEditEvent, err error) {
	var matches []string
	var studentId int
	var lessonHalf uint64
	var hasKey bool

	err = mapstructure.Decode(form, &event)
	event.CommonEventData = *eventData
	event.Scores = make(map[int]map[uint8]string)

	if err == nil {
		for key, value := range form {
			matches = StudentScoreFieldRegexp.FindStringSubmatch(key)
			if len(matches) >= 3 {
				studentId, _ = strconv.Atoi(matches[1])
				lessonHalf, _ = strconv.ParseUint(matches[2], 10, 8)

				if _, hasKey = event.Scores[studentId]; !hasKey {
					event.Scores[studentId] = make(map[uint8]string)
				}
				event.Scores[studentId][uint8(lessonHalf)] = value
			}
		}
	}
	return
}

func (event *ScoreEditEvent) ToMessage() *Message {
	message := EventToMessage(event)
	message.Form["n"] = strconv.Itoa(ScoreEditFormActionNumber)

	for studentId, scores := range event.Scores {
		for lessonHalf, score := range scores {
			key := "st" + strconv.Itoa(studentId) + "_" + strconv.Itoa(int(lessonHalf)) + "-" + event.LessonId
			message.Form[key] = score
		}
	}
	return message
}
