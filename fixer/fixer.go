package fixer

import (
	"encoding/json"
	"fixerapi/currency"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetLatestRates gets latest exchange rates from fixer.io
func GetLatestRates(currency string) (bool, *Response) {
	res, err := http.Get("http://api.fixer.io/latest?base=" + currency)
	if res.StatusCode != 200 || err != nil {
		return false, nil
	}
	body, err := ioutil.ReadAll(res.Body)
	s, err := createResponse([]byte(body))

	return true, s
}

func createResponse(body []byte) (*Response, error) {
	s := new(Response)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("createFixerResponed failed. Coudn't unmarshal body.", err)
	}
	return s, err
}

// Response represents response delivered by fixer.io
type Response struct {
	Base  string         `json:"base"`
	Date  string         `json:"date"`
	Rates currency.Rates `json:"rates"`
}
