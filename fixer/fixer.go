package fixer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getLatestRates(currency string) (bool, *fixerResponse) {
	res, err := http.Get("http://api.fixer.io/latest?base=" + currency)
	if res.StatusCode != 200 || err != nil {
		return false, nil
	}
	body, err := ioutil.ReadAll(res.Body)
	s, err := createFixerResponse([]byte(body))

	return true, s
}

func createFixerResponse(body []byte) (*fixerResponse, error) {
	s := new(fixerResponse)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("createFixerResponed failed. Coudn't unmarshal body.", err)
	}
	return s, err
}

type fixerResponse struct {
	Base  string             `json:"base"`
	Date  string             `json:"date"`
	Rates map[string]float64 `json:"rates"`
}
