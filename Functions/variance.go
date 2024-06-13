package Functions

func Variance(v []float64, mean float64) float64 {
	variance := 0.0
	for _, val := range v { // Iterate through each element of the slice calculating the difference between the current element and the mean,
		diff := val - mean
		variance += diff * diff // Square the difference and add it to the variance
	}
	variance /= float64(len(v)) // Divide the sum of squared differences by the number of elements
	return variance
}
