package helper

import (
	"encoding/json"
	"fmt"
)

func StringToJsonArray(jsonStr string) []string {
	var jsonArr []string
	err := json.Unmarshal([]byte(jsonStr), &jsonArr)
	if err != nil {
		panic(fmt.Sprintf("Empty when converting string to []string: %s", err.Error()))
	}
	return jsonArr
}

func DecodeJSONArrayOfJSONStrings(raw string) ([]map[string]interface{}, error) {
	// Step 1: Unmarshal the outer array into []string
	var jsonStrings []string
	if err := json.Unmarshal([]byte(raw), &jsonStrings); err != nil {
		return nil, fmt.Errorf("error unmarshalling outer array: %w", err)
	}

	// Step 2: Unmarshal each JSON string into a map
	var result []map[string]interface{}
	for _, s := range jsonStrings {
		var obj map[string]interface{}
		if err := json.Unmarshal([]byte(s), &obj); err != nil {
			return nil, fmt.Errorf("error unmarshalling inner object: %w", err)
		}
		result = append(result, obj)
	}

	return result, nil
}
