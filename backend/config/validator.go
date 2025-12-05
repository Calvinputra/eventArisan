package config

import (
	"event/backend/constants"
	"event/backend/helper"
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

type ErrorStruct[T any] struct {
	Struct T
}

type ErrorField struct {
	ActualTag string
	Param     string
}

func NewValidator() *validator.Validate {
	v := validator.New()
	v.RegisterValidation("is_date_format", helper.IsDateFormat)
	err := v.RegisterValidation("is_date_time_format", helper.IsDateTimeFormat)
	if err != nil {
		panic(err)
	}
	return v
}

func (e *ErrorStruct[T]) GenerateStructErrorV2(err error, rp *ResponseParameter) T {
	listOfFailedField := make(map[string][]ErrorField)
	fmt.Println("err")
	fmt.Println(err)
	for _, er := range err.(validator.ValidationErrors) {
		fmt.Println("er")
		fmt.Println(er)
		listOfFailedField[er.StructField()] = []ErrorField{
			{
				ActualTag: er.ActualTag(),
				Param:     er.Param(),
			},
		}
	}
	structValue := reflect.ValueOf(&e.Struct).Elem()

	// Iterate over the fields of the struct
	for fieldName, errorList := range listOfFailedField {
		field := structValue.FieldByName(fieldName)
		if !field.IsValid() {
			// Handle the case where the field doesn't exist in the struct
			continue
		}

		// Assuming you have some logic to determine the values to assign
		valuesToAssign := rp.generateFieldError(errorList)

		// Assign the slice of strings to the corresponding field
		field.Set(reflect.ValueOf(valuesToAssign))
	}
	return e.Struct
}

func (c *ResponseParameter) generateFieldError(errorList []ErrorField) []string {
	errorMessages := []string{}
	for _, err := range errorList {
		errValidationRp := c.GetResponse(c.ValidatorErrorMapping[err.ActualTag])
		errorMessages = append(errorMessages, formatMessage(errValidationRp, err))
	}

	return errorMessages
}

func formatMessage(errValidationRp ResponseData, errField ErrorField) string {
	switch errValidationRp.Recid {
	case constants.ValidationChoices:
		return fmt.Sprintf(errValidationRp.Message, strings.ReplaceAll(errField.Param, " ", " / "))
	case constants.ValidationMinNumber:
		return fmt.Sprintf(errValidationRp.Message, errField.Param)
	}

	return errValidationRp.Message
}
