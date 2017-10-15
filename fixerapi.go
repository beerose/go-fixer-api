package fixerapi

import (
	"fixerapi/convert"
	"fmt"
	"net/http"
)

func init() {

	http.HandleFunc("/convert", convert.Convert)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w,
			"Try with http://your-service/convert?currency={currency}&amount={amount}.")
	})
}
