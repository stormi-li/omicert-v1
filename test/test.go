package main

import (
	"fmt"
	"net/http"

	"github.com/stormi-li/omicert-v1"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello")
	})
	omicert.ListenAndServeTLS(":9999", nil)
}
