package Functions

import "testing"

func TestMedian(t *testing.T) {
	type args struct {
		v []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Test case 1", args{[]float64{1, 2, 3, 4, 5}}, 3.0},           // Test case with odd number of elements
		{"Test case 2", args{[]float64{10, 20, 30, 40, 50, 60}}, 35.0}, // Test case with even number of elements
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Median(tt.args.v); got != tt.want {
				t.Errorf("Median() = %v, want %v", got, tt.want)
			}
		})
	}
}
