package Functions

import "testing"

func TestStandardDev(t *testing.T) {
	type args struct {
		variance float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Test case 1", args{9.0}, 3.0},  // Test case with variance 9.0
		{"Test case 2", args{16.0}, 4.0}, // Test case with variance 16.0es
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StandardDev(tt.args.variance); got != tt.want {
				t.Errorf("StandardDev() = %v, want %v", got, tt.want)
			}
		})
	}
}
