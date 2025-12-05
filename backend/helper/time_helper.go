package helper

import (
	"fmt"
	"strconv"
	"time"
)

const LAYOUT_TIME = "2006-01-02 15:04:05"

func SetCurrentTimezone(timezone string) {
	location, errLocation := time.LoadLocation(timezone)
	fmt.Printf("timezone: %s\n", timezone)
	if errLocation != nil {
		panic(errLocation)
	}
	time.Local = location
}

func GetDefaultTimeString() string {
	t := time.Now()
	format := t.Format(LAYOUT_TIME)
	return format
}

func GetNowDateTimeInMs() int64 {
	return time.Now().UnixMilli()
}

func GetDefaultTime(timezone string) time.Time {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		panic(err)
	}

	parsedTime, err := time.ParseInLocation(LAYOUT_TIME, GetDefaultTimeString(), loc)
	if err != nil {
		panic(err)
	}
	return parsedTime
}

func GetDefaultTimeStrWithFormat(timezone, format string) string {
	_, err := time.LoadLocation(timezone)
	if err != nil {
		panic(err)
	}

	currentTime := time.Now()
	return currentTime.Format(format)
}

func FormatTimeStringToAnotherLayout(timezone, currentLayout, newLayout, timeStr string) string {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		panic(err)
	}

	parsedTime, err := time.ParseInLocation(currentLayout, timeStr, loc)
	if err != nil {
		panic(err)
	}

	return parsedTime.Format(newLayout)
}

func FormatTimeStringToAnotherLayoutWithBackupCurrentLayout(timezone, currentLayout, currentLayout2, newLayout, timeStr string) string {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		panic(err)
	}

	parsedTime, err := time.ParseInLocation(currentLayout, timeStr, loc)
	if err != nil {
		parsedTime2, err2 := time.ParseInLocation(currentLayout2, timeStr, loc)
		if err2 != nil {
			panic(err)
		}
		return parsedTime2.Format(newLayout)
	}

	return parsedTime.Format(newLayout)
}

func FormatTimeToString(time time.Time) string {
	return time.Format(LAYOUT_TIME)
}

func FormatStringToTime(timeStr, timezone string) time.Time {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		panic(err)
	}

	parsedTime, err := time.ParseInLocation(LAYOUT_TIME, timeStr, loc)
	if err != nil {
		fmt.Println(err)
	}
	return parsedTime
}

func GetTimeInMillis() int64 {
	return time.Now().UnixMilli()
}

func ConvertMsToDateTime(timezone string, epochMs int64) string {
	seconds := epochMs / 1000
	nanoseconds := (epochMs % 1000) * int64(time.Millisecond)

	t := time.Unix(seconds, nanoseconds)

	location, err := time.LoadLocation(timezone)
	if err != nil {
		panic(err)
	}

	t = t.In(location)
	formattedTime := t.Format(LAYOUT_TIME)

	return formattedTime
}

func GetPresentYear() string {
	currentTime := time.Now()
	year := currentTime.Year()
	return strconv.Itoa(year)
}

func ConvertStringToUnixMili(request string, timezone string) int64 {
	// Load the specified timezone
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		fmt.Println("Error loading timezone:", err)
		return 0
	}

	// Parse the input time string in UTC first
	parsedTime, err := time.ParseInLocation(LAYOUT_TIME, request, loc)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return 0
	}

	// Return Unix time in milliseconds
	return parsedTime.UnixMilli()
}
