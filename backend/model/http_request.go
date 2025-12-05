package model

type Header struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

type ActivityLogHeader struct {
	UserAgent string `json:"userAgent"`
	Country   string `json:"country"`
	IPAddress string `json:"ipAddress"`
	Host      string `json:"host"`
}
