package convert

import (
	"encoding/json"
	"net/url"
	"reflect"
	"testing"
)

func Test_unpackQuery(t *testing.T) {
	type args struct {
		params url.Values
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 []string
	}{
		{"Return unpacked params from valid http request",
			args{url.Values{"amount": []string{"1"}, "currency": []string{"EUR"}}},
			true,
			[]string{"1", "EUR"},
		},
		{"Return errors for invalid http request -- amount",
			args{url.Values{"invalid_param": []string{"1"}, "currency": []string{"EUR"}}},
			false,
			[]string{"Is amount param correct: false",
				"Is currency param present: true"},
		},
		{"Return errors for invalid http request -- currency",
			args{url.Values{"amount": []string{"1"}, "invalid_param": []string{"EUR"}}},
			false,
			[]string{"Is amount param correct: true",
				"Is currency param present: false"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := unpackQuery(tt.args.params)
			if got != tt.want {
				t.Errorf("unpackQuery() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("unpackQuery() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_createErrorResponse(t *testing.T) {
	properErrorMessage, _ := json.Marshal(&errorResponse{[]string{"Some error message."}})
	type args struct {
		params []string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"Should return proper json error message",
			args{[]string{"Some error message."}},
			properErrorMessage},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createErrorResponse(tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createErrorResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
