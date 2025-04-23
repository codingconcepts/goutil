package geo

import (
	"math"
	"math/rand/v2"
)

// RandomPointWithinRadius generates a random lat and long position
// within a given radius (in miles) of a given lat and long position.
func RandomPointWithinRadius(lat, lon, radiusMiles float64) (float64, float64) {
	randomDistance := (rand.Float64() * radiusMiles) / EarthRadiusMiles
	randomBearing := rand.Float64() * 2 * math.Pi

	latRad := DegreesToRadians(lat)
	lonRad := DegreesToRadians(lon)

	sinLatRad := math.Sin(latRad)
	cosLatRad := math.Cos(latRad)
	sinRandomDistance := math.Sin(randomDistance)
	cosRandomDistance := math.Cos(randomDistance)
	cosRandomBearing := math.Cos(randomBearing)
	sinRandomBearing := math.Sin(randomBearing)

	newLatRad := math.Asin(sinLatRad*cosRandomDistance + cosLatRad*sinRandomDistance*cosRandomBearing)

	newLonRad := lonRad + math.Atan2(
		sinRandomBearing*sinRandomDistance*cosLatRad,
		cosRandomDistance-sinLatRad*math.Sin(newLatRad),
	)

	return RadiansToDegrees(newLatRad), RadiansToDegrees(newLonRad)
}
