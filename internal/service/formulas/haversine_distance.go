package formulas

import (
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
func MetersDistance(lat1, long1, lat2, long2 float64) float64 {
	la1 := lat1 * piRad
	lo1 := long2 * piRad
	la2 := lat2 * piRad
	lo2 := long2 * piRad

	// calculate
	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)

	meters := 2 * earthRadiusMeters * math.Asin(math.Sqrt(h))
	return meters
}
