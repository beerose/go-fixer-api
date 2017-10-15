package convert

import (
	"api/currency"
	"net/http"
)

func sendValidResponse(w http.ResponseWriter, amount float64,
	currency string, converted *currency.Rates, cType contentType) {
	w.WriteHeader(http.StatusOK)
	w.Write(createValidRespone(amount, currency, converted, cType))
}

func createValidRespone(amount float64, currency string, converted *currency.Rates, cType contentType) []byte {
	validResp := &validResponse{amount, currency, *converted}

	resp, err := cType.Marshal(validResp)

	if err != nil {
		panic(err.Error())
	}

	return resp
}

type validResponse struct {
	Amount    float64        `xml:"amount,attr" json:"amount"`
	Currency  string         `xml:"currency,attr" json:"currency"`
	Converted currency.Rates `xml:"converted" json:"converted"`
}
