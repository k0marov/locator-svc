package service

import (
	"gitlab.com/samkomarov/locator-svc.git/internal/service/formulas"
	"math"
	"time"
)

func personRelevance(point GeoPoint, person MissingPerson) float64 {
	minDistance := math.Inf(1)
	for _, personPosition := range person.GeoPoints {
		d := formulas.MetersDistance(point, personPosition)
		if d < minDistance {
			minDistance = d
		}
	}
	return computeRelevance(minDistance, time.Time(person.DateOfLoss))
}

// TODO: fine-tune formula for computing relevance
func computeRelevance(distanceMeters float64, dateOfLoss time.Time) float64 {
	daysDiff := time.Now().Sub(dateOfLoss).Hours() * 24
	distanceMeters += 100000 * daysDiff // 10 km for each day
	return 1 / distanceMeters
}
