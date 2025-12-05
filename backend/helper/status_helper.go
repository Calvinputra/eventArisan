package helper

import (
	"event/backend/constants"
	"strings"
)

func SetStatusBasedOnApiKey(apiKey string) string {
	switch strings.ToUpper(apiKey) {
	case constants.INSERT:
		return constants.PUBLISHED

	case constants.HOLD_CREATE_AND_UPDATE:
		return constants.DRAFT

	case constants.HOLD_TO_NAU:
		return constants.IN_REVIEW

	case constants.REJECT:
		return constants.REJECTED

	case constants.UPDATE:
		return constants.PUBLISHED

	case constants.EDIT_AUTH:
		return constants.PUBLISHED

	case constants.LIST:
		return constants.PUBLISHED

	default:
		return ""
	}
}
