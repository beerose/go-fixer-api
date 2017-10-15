package convert

import (
	"api/fixer"
	"net/http"
	"net/url"
	"strconv"
)

// Convert handles http request
func Convert(w http.ResponseWriter, r *http.Request) {
	status, params := unpackQuery(r.URL.Query())
	cType := contentType(r.Header.Get("Accept"))
	if cType != ctXML {
		cType = ctJSON
	}

	if !status {
		sendErrorResponse(w, http.StatusBadRequest, params, cType)
	} else {
		amountString, currency := params[0], params[1]
		fixerStatus, latestRates := fixer.GetLatestRates(currency)
		if !fixerStatus {
			sendErrorResponse(w, http.StatusFailedDependency,
				[]string{"Cannot GET to fixer.io."}, cType)
		} else {
			amount, _ := strconv.ParseFloat(amountString, 64)
			converted := latestRates.Rates.Multiply(amount)
			sendValidResponse(w, amount, currency, &converted, cType)
		}
	}
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
