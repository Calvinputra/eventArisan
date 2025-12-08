package model

type BaseDoorprizeRequest struct {
	EventRecid           string `json:"eventRecid"`
	AttendanceRecid           string `json:"attendanceRecid"`
}

type BaseDoorprizeRequestError struct {
	EventRecid           []string `json:"eventRecid"`
	AttendanceRecid           []string `json:"attendanceRecid"`
}

func (BaseDoorprizeRequestError) Empty() BaseDoorprizeRequestError {
	return BaseDoorprizeRequestError{
		EventRecid:           []string{},
		AttendanceRecid:           []string{},
	}
}
