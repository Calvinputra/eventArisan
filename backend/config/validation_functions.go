package config

import (
	"event/backend/constants"
	"fmt"
	"golang.org/x/exp/slices"
	"strings"
	"time"
)

type IsValidated struct {
	Status bool
}

type CustomValidation struct {
	Rp *ResponseParameter
}

func NewCustomValidation(rp *ResponseParameter) *CustomValidation {
	return &CustomValidation{Rp: rp}
}

func (c *CustomValidation) SingleValidationString(mandatory bool, max int, choices []string, value string, isValidated *IsValidated) []string {
	var errors []string = []string{}

	if mandatory && (value == "" || value == "<nil>") {
		errors = append(errors, c.Rp.GetResponse(constants.ValidationEmpty).Message)
		isValidated.Status = false
	}

	if choices != nil {
		if !slices.Contains(choices, value) {
			errors = append(errors, fmt.Sprintf(c.Rp.GetResponse(constants.ValidationChoices).Message, strings.Join(choices, ",")))
			isValidated.Status = false
		}
	}

	if max != -1 && len(value) > max {
		errors = append(errors, fmt.Sprintf(c.Rp.GetResponse(constants.ValidationMaxChar).Message, max))
		isValidated.Status = false
	}

	return errors
}

func (c *CustomValidation) SingleValidationInteger(mandatory bool, max, min, value int, isValid *IsValidated) []string {
	var errors []string = []string{}

	if mandatory && value < 1 {
		errors = append(errors, c.Rp.GetResponseMessageOnly(constants.ValidationNol))
		isValid.Status = false
	}

	if max != -1 && value > max {
		errors = append(errors, fmt.Sprintf(c.Rp.GetResponseMessageOnly(constants.ValidationMaxNumber), float32(max)))
		isValid.Status = false
	}

	if min != -2 && value < min {
		errors = append(errors, fmt.Sprintf(c.Rp.GetResponseMessageOnly(constants.ValidationMinNumber), float32(min)))
		isValid.Status = false
	}

	return errors
}

func (c *CustomValidation) EmptyValidation(mandatory bool, emptyData any, value any, isValidated *IsValidated) []string {
	var errors []string = []string{}

	if mandatory && emptyData == value {
		errors = append(errors, c.Rp.GetResponse(constants.ValidationEmpty).Message)
		isValidated.Status = false
	}

	return errors
}

func (c *CustomValidation) MinMaxNumberValidation(min, max, value int, isValidated *IsValidated) []string {
	var errors []string = []string{}

	if min != -1 && value < min {
		errors = append(errors, fmt.Sprintf(c.Rp.GetResponse(constants.ValidationMinNumber).Message, min))
		isValidated.Status = false
	}

	if max != -1 && value > max {
		errors = append(errors, fmt.Sprintf(c.Rp.GetResponse(constants.ValidationMaxNumber).Message, max))
		isValidated.Status = false
	}

	return errors
}

func (c *CustomValidation) MinMaxInt64Validation(min, max, value int64, isValidated *IsValidated) []string {
	var errors []string = []string{}

	if value == 0 {
		errors = append(errors, fmt.Sprintf(c.Rp.GetResponse(constants.ValidationEmpty).Message))
		isValidated.Status = false
	}

	if min != -1 && value < min {
		errors = append(errors, fmt.Sprintf(c.Rp.GetResponse(constants.ValidationMinNumber).Message, min))
		isValidated.Status = false
	}

	if max != -1 && value > max {
		errors = append(errors, fmt.Sprintf(c.Rp.GetResponse(constants.ValidationMaxNumber).Message, max))
		isValidated.Status = false
	}

	return errors
}

func (c *CustomValidation) MinMaxFloat64Validation(min, max, value float64, isValidated *IsValidated) []string {
	var errors []string = []string{}

	if min != -1 && value < min {
		errors = append(errors, fmt.Sprintf(c.Rp.GetResponse(constants.ValidationMinNumber).Message, min))
		isValidated.Status = false

	}

	if max != -1 && value > max {
		errors = append(errors, fmt.Sprintf(c.Rp.GetResponse(constants.ValidationMaxNumber).Message, max))
		isValidated.Status = false
	}

	return errors
}

func (c *CustomValidation) ValidationTimeNotBeforeStart(startDatetime, endDatetime time.Time, isValid *IsValidated) []string {
	var errors []string = []string{}

	if endDatetime.Before(startDatetime) {
		isValid.Status = false
		errors = append(errors, fmt.Sprintf(c.Rp.GetResponseMessageOnly(constants.ValidationTimeCannotBeBefore), startDatetime.AppendFormat([]byte{}, constants.LAYOUTTIME)))
		return errors
	}

	return errors
}

func (c *CustomValidation) ValidationFormatDatetime(datetime string, isValid *IsValidated) (time.Time, []string) {
	var errors []string = []string{}

	parseDatetime, err := time.Parse(constants.LAYOUTTIME, datetime)
	if err != nil {
		isValid.Status = false
		errors = append(errors, c.Rp.GetResponseMessageOnly(constants.ValidationDateTimeFormat))

		return time.Time{}, errors
	}

	return parseDatetime, nil
}
