package models

type HttpGetMock struct {
	Port     int    `json:"port"`
	Url      string `json:"url"`
	Response string `json:"response"`
}
