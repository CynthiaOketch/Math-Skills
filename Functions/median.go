package Functions

//import ("sort")

func Median(v []float64) float64 { //sorting the slice of floats using the bubble sort technique
	for i := 0; i < len(v)-1; i++ {
		for j := 0; j < len(v)-i-1; j++ {
			if v[j] > v[j+1] {
				v[j], v[j+1] = v[j+1], v[j]
			}
		}
	}

	//sort.Float64s(v)
	if len(v)%2 == 0 { //handling slice of floats with even number of elements

		mid := len(v) / 2
		return (v[mid-1] + v[mid]) / 2
	} else { //handling slice of floats with odd number of elements
		return v[len(v)/2]
	}
}
