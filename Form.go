package dekanatEvents

import (
	"errors"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"strconv"
)

const DekanatFormDateFormat = "02.01.2006"

const LessonFormActionNumber = 10
const LessonDeleteFormActionNumber = 11
const ScoreEditFormActionNumber = 4

const LessonEditFormAction = "edit"
const LessonInsertFormAction = "insert"
const LessonDeleteFormAction = "delete"

type Form map[string]string

func (form Form) toEvent(timestamp int64, receiptHandle *string) (event interface{}, err error) {
	if !form.isValidEventForm() {
		return nil, errors.New("invalid form")
	}

	commonEventData := &CommonEventData{}
	_ = mapstructure.Decode(form, commonEventData)
	commonEventData.Timestamp = timestamp
	commonEventData.ReceiptHandle = receiptHandle

	actionNumber, _ := strconv.Atoi(form["n"])
	action, _ := form["action"]

	if actionNumber == LessonFormActionNumber && action == LessonEditFormAction && form.hasLessonFormFields() {
		event, err = createLessonEditEvent(form, commonEventData)
	} else if actionNumber == LessonFormActionNumber && action == LessonInsertFormAction && form.hasLessonFormFields() {
		event, err = createLessonCreateEvent(form, commonEventData)
	} else if actionNumber == LessonDeleteFormActionNumber && action == LessonDeleteFormAction {
		event, err = createLessonDeleteEvent(form, commonEventData)
	} else if actionNumber == ScoreEditFormActionNumber && form.hasStudentScoreField() {
		event, err = createScoreEditEvent(form, commonEventData)
	} else {
		err = errors.New(fmt.Sprintf("Unknown form: n=%d, action=%s", actionNumber, action))
	}

	return
}

func (form Form) isValidEventForm() (valid bool) {
	return form.isFieldExist("n") && form.isFieldExist("prti") && form.isFieldExist("prt")
}

func (form Form) hasLessonFormFields() (valid bool) {
	return form.isFieldExist("tzn") && form.isFieldExist("date_z")
}

func (form Form) isFieldExist(field string) (exist bool) {
	_, exist = form[field]
	return
}

func (form Form) hasStudentScoreField() bool {
	for key := range form {
		if StudentScoreFieldRegexp.MatchString(key) {
			return true
		}
	}
	return false
}
