package convert

import (
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
		fixerStatus, latestRates := fixer.GetLatestRates(params[1])
		if !fixerStatus {
			sendErrorResponse(w, http.StatusFailedDependency,
				[]string{"Cannot GET to fixer.io."})
		} else {

		}

	}

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
	errResponse := &errorResponse{params}
	jsonResp, err := json.Marshal(errResponse)

	if err != nil {
		panic(err.Error())
	}

	return jsonResp
}

type validResponse struct {
	Amount    int                `json:"amount"`
	Currency  string             `json:"currency"`
	Converted map[string]float64 `json:"converted"`
}

type errorResponse struct {
	Error []string `json:"error"`
}
