// Package space implements utility to convert an age from a specified Planet in Earth years.
package space

// Planet represents the required Planet name.
type Planet string

// yearSeconds holds the number of seconds in 1 year on Earth.
const yearSeconds float64 = 31557600

// Age calculates age in Earth years when given a Planet and age in seconds.
func Age(seconds float64, planet Planet) float64 {
	earthyears := map[Planet]float64{
		"Earth":   1,
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}
	return seconds / (yearSeconds * earthyears[planet])
}
