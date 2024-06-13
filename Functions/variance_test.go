package Functions

import "testing"

func TestVariance(t *testing.T) {
	type args struct {
		v    []float64
		mean float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Test case 1", args{[]float64{1, 2, 3, 4, 5}, 3.0}, 2.0},         // Test case with a simple sequence
		{"Test case 2", args{[]float64{10, 20, 30, 40, 50}, 30.0}, 200.0}, // Test case with larger numbers
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Variance(tt.args.v, tt.args.mean); got != tt.want {
				t.Errorf("Variance() = %v, want %v", got, tt.want)
			}
		})
	}
}
