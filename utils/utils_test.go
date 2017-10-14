package utils

import "testing"

func TestRound(t *testing.T) {
	type args struct {
		some float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Rounds up", args{3.66666}, 3.67},
		{"Rounds down", args{4.22222}, 4.22},
		{"Does not truncate when not needed", args{1.0}, 1.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Round(tt.args.some); got != tt.want {
				t.Errorf("Round() = %v, want %v", got, tt.want)
			}
		})
	}
}
