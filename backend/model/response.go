package model

type Response struct {
	HttpCode    int     `json:"code"`
	Message     string  `json:"message"`
	MessageId   string  `json:"messageId"`
	MessageCode int     `json:"messageCode"`
	Errors      any     `json:"errors,omitempty"`
	Data        any     `json:"data,omitempty"`
	NextPage    *string `json:"nextPage,omitempty"`
}

func (e *Response) Error() string {
	return e.Message
}
