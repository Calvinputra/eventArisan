package model

type BaseEventRequest struct {
	Name           string `json:"name"`
	Description    string `json:"description"`
	StartDateTime  int64  `json:"startDateTime"`
	EndDateTime    int64  `json:"endDateTime"`
	Status         string `json:"status"`
}

type BaseEventRequestError struct {
	Name        []string `json:"name"`
	Description []string `json:"description"`
	StartDateTime  []int64  `json:"startDateTime"`
	EndDateTime    []int64  `json:"endDateTime"`
	Status         []string `json:"status"`
}

func (BaseEventRequestError) Empty() BaseEventRequestError {
	return BaseEventRequestError{
     	Name:          []string{},
        Description:   []string{},
        StartDateTime: []int64{},
        EndDateTime:   []int64{},
        Status:        []string{},
	}
}
