package convert

import (
	"net/http"
)

func sendErrorResponse(w http.ResponseWriter, status int, params []string, cType contentType) {
	w.WriteHeader(status)
	w.Write(createErrorResponse(params, cType))
}

func createErrorResponse(params []string, cType contentType) []byte {
	errResp := &errorResponse{params}

	resp, err := cType.Marshal(errResp)

	if err != nil {
		panic(err.Error())
	}

	return resp
}

type errorResponse struct {
	Error []string `xml:"error,attr" json:"error"`
}
