package model

type Route struct {
	Start      string  `json:"start"`
	End        string  `json:"end"`
	StopsCount int     `json:"stopsCount"`
	Distance   float32 `json:"distance"`
}
