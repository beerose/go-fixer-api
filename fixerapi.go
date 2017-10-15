package fixerapi

import (
	"fixerapi/convert"
	"fmt"
	"net/http"
)

func init() {

	fmt.Println("Started application.")

	http.HandleFunc("/convert", convert.Convert)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Try use convert?amount={some_amount}&currency={some_currency} in url.")
	})
}
