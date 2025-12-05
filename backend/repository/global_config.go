package repository

import (
	"event/backend/helper"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"strings"
)

// Function to get global_config data -> label for JSON_ARRAY
func (r *BaseRepository[T]) GetLabelFromGlobalConfigJsonArrayWithLanguageCode(db *gorm.DB, recid, valueToMatch, columnName, languageCode string) (string, error) {
	var jobType []string
	if err := json.Unmarshal([]byte(valueToMatch), &jobType); err != nil {
		return "", err
	}

	var resolvedLabels []string

	for _, tipe := range jobType {
		var labelValue string

		query := fmt.Sprintf(`SELECT jsonb_extract_path(data_obj -> 'label', $3) AS %s FROM general."GLOBAL_CONFIG", jsonb_array_elements(data::jsonb) AS data_obj WHERE recid = $1 AND type = 'JSON_ARRAY' AND data_obj ->> 'value' = $2;`,
			columnName)
		err := db.Raw(query, recid, tipe, languageCode).Row().Scan(&labelValue)
		if err != nil {
			labelValue = ""
		}

		resolvedLabels = append(resolvedLabels, strings.Trim(labelValue, "\""))
	}

	return helper.ConvertToJsonString(resolvedLabels), nil
}
