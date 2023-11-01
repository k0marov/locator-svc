package service

import "time"

// MissingPerson is the main domain entity
type MissingPerson struct {
	PhotoURL   string
	DateOfLoss time.Time
	Locations  []GeoPoint
	// Relevance is an indicator of how this person is relevant to the defined location
	Relevance float64
}

type GeoPoint struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
