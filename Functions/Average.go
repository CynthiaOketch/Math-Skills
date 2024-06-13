package Functions

// Average calculates the arithmetic mean of a slice of float64 values.
func Average(v []float64) float64 {
	sum := 0.0
	mean := 0.0
	for i := 0; i <= len(v)-1; i++ { // Iterate through each element of the slice adding the current element to the sum
		sum += v[i]
		mean = sum / (float64(len(v))) // Calculate the mean by dividing the sum by the number of elements
	}
	return mean // Return the calculated mean
}
