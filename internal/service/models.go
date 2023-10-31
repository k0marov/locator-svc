package service

import "time"

type MissingPerson struct {
	VerticalURL string     `json:"vertical_url"`
	DateOfLoss  time.Time  `json:"date_of_loss"`
	GeoPoints   []GeoPoint `json:"geo_points"`
}

type GeoPoint struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
