package formulas

import (
	"math"
)

// adapted from: https://gist.github.com/cdipaolo/d3f8db3848278b49db68
func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}

const earthRadiusKM = float64(6378.1)
const piRad = math.Pi / 180

// DistanceKM returns distance from p1 to p2 in meters
func DistanceKM(lat1, long1, lat2, long2 float64) float64 {
	la1 := lat1 * piRad
	lo1 := long1 * piRad
	la2 := lat2 * piRad
	lo2 := long2 * piRad

	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)

	meters := 2 * earthRadiusKM * math.Asin(math.Sqrt(h))
	return meters
}
