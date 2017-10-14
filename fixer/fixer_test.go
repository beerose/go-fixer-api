package fixer

import (
	"reflect"
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

func Test_createFixerResponse(t *testing.T) {
	type args struct {
		body []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *fixerResponse
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := createFixerResponse(tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("createFixerResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createFixerResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
