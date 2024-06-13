package Functions

import "math"

// StandardDev calculates the standard deviation from the given variance.
func StandardDev(variance float64) float64 {
	stdDev := math.Sqrt(variance)
	return stdDev
}
