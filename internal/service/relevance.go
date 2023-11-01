package service

import (
	"gitlab.com/samkomarov/locator-svc.git/internal/service/formulas"
	"math"
	"time"
)

func personRelevance(point GeoPoint, person MissingPerson) float64 {
	minDistance := math.Inf(1)
	for _, personPosition := range person.GeoPoints {
		d := formulas.DistanceKM(point.Latitude, point.Longitude, personPosition.Latitude, personPosition.Longitude)
		if d < minDistance {
			minDistance = d
		}
	}
	return computeRelevance(minDistance, time.Time(person.DateOfLoss))
}

// TODO: fine-tune formula for computing relevance
func computeRelevance(distanceKM float64, dateOfLoss time.Time) float64 {
	relevance := 1 / distanceKM
	daysDiff := time.Now().Sub(dateOfLoss).Hours() / 24
	if daysDiff > 30 {
		relevance *= 0.5
	}
	return relevance
}
