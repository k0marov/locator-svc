package service

import "gitlab.com/samkomarov/locator-svc.git/internal/core"

type MissingPerson struct {
	VerticalURL string     `json:"vertical_url"`
	DateOfLoss  core.Date  `json:"date_of_loss"`
	GeoPoints   []GeoPoint `json:"geo_points"`
}

type GeoPoint struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
