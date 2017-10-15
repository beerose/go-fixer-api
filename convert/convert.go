package convert

import (
	"fixer-api/fixer"
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
	amountString, amountPresent := params["amount"]
	currency, currencyPresent := params["currency"]

	if !(amountPresent && currencyPresent) {
		return false, []string{
			"Is amount param correct: " +
				strconv.FormatBool(amountPresent),
			"Is currency param correct: " +
				strconv.FormatBool(currencyPresent)}
	}

	_, err := strconv.ParseFloat(amountString[0], 64)
	isProperCurr := isProperCurrency(currency[0])

	if err != nil || !isProperCurr {
		return false, []string{
			"Is amount param correct: " +
				strconv.FormatBool(err != nil),
			"Is currency param correct: " +
				strconv.FormatBool(isProperCurr)}
	}
	return true, []string{amountString[0], currency[0]}
}

func isProperCurrency(currency string) bool {
	currencies := []string{"AUD", "BGN", "BRL", "CAD", "CHF", "CNY", "CZK", "DKK", "GBP", "HKD", "HRK",
		"HUF", "IDR", "ILS", "INR", "JPY", "KRW", "MXN", "MYR", "NOK", "NZD", "PHP",
		"PLN", "RON", "RUB", "SEK", "SGD", "THB", "TRY", "USD", "ZAR", "EUR"}

	for i := range currencies {
		if currency == currencies[i] {
			return true
		}
	}
	return false
}
