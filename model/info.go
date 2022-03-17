package model

type Info struct {
	Date string `json:"date"`
	Weather string `json:"weather"`
	HighTemp string `json:"highTemp"`
	LowTemp string `json:"lowTemp"`
	Times []string `json:"times"`
	Precipitations []string `json:"precipitations"`
	Wind string `json:"wind"`
}