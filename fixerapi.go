package fixerapi

import (
	"fixerapi/convert"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/convert", convert.Convert)
	log.Fatal(http.ListenAndServe(":8080", router))
}
