package service

import (
	"gitlab.com/samkomarov/locator-svc.git/internal/service/formulas"
	"math"
	"time"
)

func personRelevance(target GeoPoint, locations []GeoPoint, dateOfLoss time.Time) float64 {
	minDistance := math.Inf(1)
	for _, personPosition := range locations {
		d := formulas.DistanceKM(target.Latitude, target.Longitude, personPosition.Latitude, personPosition.Longitude)
		if d < minDistance {
			minDistance = d
		}
	}
	return computeRelevance(minDistance, dateOfLoss)
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
