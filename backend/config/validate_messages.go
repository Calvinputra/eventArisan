package config

import "event/backend/constants"

func getValidateMessages() map[string]string {
	return map[string]string{
		"required":         constants.ValidationEmpty,
		"required_without": constants.ValidationEmpty,
		"min":              constants.ValidationMinNumber,
		"max":              constants.ValidationMaxNumber,
		"oneof":            constants.ValidationChoices,
		"dive":             constants.ValidationChoices,
	}
}
