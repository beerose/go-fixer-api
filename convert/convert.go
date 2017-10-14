package convert

import (
	"api/currency"
	"api/fixer"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

// Convert handles http request
func Convert(w http.ResponseWriter, r *http.Request) {
	status, params := unpackQuery(r.URL.Query())

	if !status {
		sendErrorResponse(w, http.StatusBadRequest, params)
	} else {
		amountString, currency := params[0], params[1]
		fixerStatus, latestRates := fixer.GetLatestRates(currency)
		if !fixerStatus {
			sendErrorResponse(w, http.StatusFailedDependency,
				[]string{"Cannot GET to fixer.io."})
		} else {
			amount, _ := strconv.ParseFloat(amountString, 64)
			converted := latestRates.Rates.Multiply(amount)
			sendValidResponse(w, amount, currency, &converted)
		}
	}
}

func sendValidResponse(w http.ResponseWriter, amount float64,
	currency string, converted *currency.Rates) {
	w.WriteHeader(http.StatusOK)
	w.Write(createValidRespone(amount, currency, converted))
}

func sendErrorResponse(w http.ResponseWriter, status int, params []string) {
	w.WriteHeader(status)
	w.Write(createErrorResponse(params))
}

func unpackQuery(params url.Values) (bool, []string) {
	amount, amountPresent := params["amount"]
	currency, currencyPresent := params["currency"]

	if !(amountPresent && currencyPresent) {
		return false, []string{
			"Is amount param correct: " +
				strconv.FormatBool(amountPresent),
			"Is currency param present: " +
				strconv.FormatBool(currencyPresent)}
	}

	return true, []string{amount[0], currency[0]}
}

func createErrorResponse(params []string) []byte {
	errResp := &errorResponse{params}
	jsonResp, err := json.Marshal(errResp)

	if err != nil {
		panic(err.Error())
	}

	return jsonResp
}

func createValidRespone(amount float64, currency string, converted *currency.Rates) []byte {
	validResp := &validResponse{amount, currency, *converted}
	jsonResp, err := json.Marshal(validResp)

	if err != nil {
		panic(err.Error())
	}

	return jsonResp
}

type validResponse struct {
	Amount    float64        `json:"amount"`
	Currency  string         `json:"currency"`
	Converted currency.Rates `json:"converted"`
}

type errorResponse struct {
	Error []string `json:"error"`
}
