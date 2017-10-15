package convert

import (
	"encoding/json"
	"net/http"
)

func sendErrorResponse(w http.ResponseWriter, status int, params []string, cType contentType) {
	w.WriteHeader(status)
	w.Write(createErrorResponse(params, cType))
}

func createErrorResponse(params []string, cType contentType) []byte {
	errResp := &errorResponse{params}

	jsonResp, err := json.Marshal(errResp)

	if err != nil {
		panic(err.Error())
	}

	return jsonResp
}

type errorResponse struct {
	Error []string `xml:"error,attr" json:"error"`
}
