package model

type BaseEventRequest struct {
	Name           string `json:"name"`
	Description    string `json:"description"`
	Location    string `json:"location"`
	StartDateTime  int64  `json:"startDateTime"`
	Status         string `json:"status"`
}

type BaseEventRequestError struct {
	Name        []string `json:"name"`
	Description []string `json:"description"`
	Location    []string `json:"location"`
	StartDateTime  []int64  `json:"startDateTime"`
	Status         []string `json:"status"`
}

func (BaseEventRequestError) Empty() BaseEventRequestError {
	return BaseEventRequestError{
     	Name:          []string{},
        Description:   []string{},
        Location:   []string{},
        StartDateTime: []int64{},
        Status:        []string{},
	}
}
