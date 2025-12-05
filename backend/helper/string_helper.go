package helper

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

func ConvertToJsonString(mapped any) string {
	parsedMapped, parsedMappedErr := json.Marshal(mapped)
	if parsedMappedErr != nil {
		panic(fmt.Sprintf("Empty when converting json to string: %s", parsedMappedErr.Error()))
	}
	return string(parsedMapped)
}

func ConvertAnyToString(value any) string {
	str := fmt.Sprint(value)
	if str == "<nil>" {
		return ""
	}
	return str
}

func StringToSnakeCaseString(str string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
