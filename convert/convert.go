package convert

import (
	"net/http"
	"net/url"
	"strconv"
)

// Convert handles http request
func Convert(w http.ResponseWriter, r *http.Request) {
	status, _ := unpackQuery(r.URL.Query())
	r.Header

	if !status {
		w.WriteHeader(http.StatusBadRequest)
	} else {

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
