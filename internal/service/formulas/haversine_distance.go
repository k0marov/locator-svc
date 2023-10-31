package formulas

import (
	"gitlab.com/samkomarov/locator-svc.git/internal/service"
	"math"
)

// adapted from: https://gist.github.com/cdipaolo/d3f8db3848278b49db68
// haversin(θ) function
func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}

const earthRadiusMeters = float64(6378100)
const piRad = math.Pi / 180

// MetersDistance returns distance from p1 to p2 in meters
func MetersDistance(p1, p2 service.GeoPoint) float64 {
	la1 := p1.Latitude * piRad
	lo1 := p1.Longitude * piRad
	la2 := p2.Latitude * piRad
	lo2 := p2.Longitude * piRad

	// calculate
	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)

	meters := 2 * earthRadiusMeters * math.Asin(math.Sqrt(h))
	return meters
}
