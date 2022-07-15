package model

type Photo struct {
	Id              int         `json:"id"`
	Width           int         `json:"width"`
	Height          int         `json:"height"`
	Url             string      `json:"url"`
	Photographer    string      `json:"photographer"`
	PhotographerUrl string      `json:"photographer_url"`
	PhotographerId  int         `json:"photographer_id"`
	AvgColor        string      `json:"avg-color"`
	Src             PhotoSource `json:"src"`
	Alt             string      `json:"alt"`
}

type PhotoSource struct {
	Original  string `json:"original"`
	Large     string `json:"large"`
	Large2x   string `json:"large2x"`
	Medium    string `json:"medium"`
	Small     string `json:"small"`
	Portrait  string `json:"portrait"`
	Landscape string `json:"landscape"`
	Tiny      string `json:"tiny"`
}
