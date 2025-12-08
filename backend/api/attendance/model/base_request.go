package model

type BaseAttendanceRequest struct {
	EventRecid           string `json:"eventRecid"`
	Name    string `json:"name"`
	Code  string  `json:"code"`
	NoTable    int64  `json:"noTable"`
	StatusCheckin         int8 `json:"statusCheckin"`
	StatusSouvenir         int8 `json:"StatusSouvenir"`
	CheckinTime         int64 `json:"CheckinTime"`
	SouvenirTime         int64 `json:"SouvenirTime"`
}

type BaseAttendanceRequestError struct {
	EventRecid           []string `json:"eventRecid"`
	Name    []string `json:"name"`
	Code  []string  `json:"code"`
	NoTable    []int64  `json:"noTable"`
	StatusCheckin         []int8 `json:"statusCheckin"`
	StatusSouvenir         []int8 `json:"StatusSouvenir"`
	CheckinTime         []int64 `json:"CheckinTime"`
	SouvenirTime         []int64 `json:"SouvenirTime"`
}

func (BaseAttendanceRequestError) Empty() BaseAttendanceRequestError {
	return BaseAttendanceRequestError{
		EventRecid:           []string{},
		Name:        []string{},
        Code:      []string{},
        NoTable:           []int64{},
        StatusCheckin:     []int8{},
        StatusSouvenir:    []int8{},
        CheckinTime:       []int64{},
        SouvenirTime:       []int64{},
	}
}
