package geo

import "math"

// DegreesToRadians converts degrees to radians.
func DegreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

// RadiansToDegrees converts radians to degress.
func RadiansToDegrees(radians float64) float64 {
	return radians * 180 / math.Pi
}
