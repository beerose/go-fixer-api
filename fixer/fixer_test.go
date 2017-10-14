package fixer

import (
	"testing"
)

func Test_getLatestRates(t *testing.T) {
	type args struct {
		currency string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Should return proper fixerResponse", args{"EUR"}, true},
		{"Should return error for invalid currency", args{"invalid_curr"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := getLatestRates(tt.args.currency)
			if got != tt.want {
				t.Errorf("getLatestRates() got = %v, want %v", got, tt.want)
			}
		})
	}
}
