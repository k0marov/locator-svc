package service

import (
	"gitlab.com/samkomarov/locator-svc.git/internal/core"
)

// LizaAlertPerson is the model that is returned from LizaAlert API
type LizaAlertPerson struct {
	VerticalURL string     `json:"vertical_url"`
	DateOfLoss  core.Date  `json:"date_of_loss"`
	GeoPoints   []GeoPoint `json:"geo_points"`
}
