package service

import "gitlab.com/samkomarov/locator-svc.git/internal/core"

// MissingPerson is the main domain entity
type MissingPerson struct {
	VerticalURL string
	DateOfLoss  core.Date
	GeoPoints   []GeoPoint
}

type GeoPoint struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
