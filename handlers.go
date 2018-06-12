package main

import (
	"fmt"
	"net/http"
)

func apihandlefunc(w http.ResponseWriter, r *http.Request) {

}

func viewhandlefunc(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid method", http.StatusBadRequest)
	}
	page, _ := viewsAppHtmlBytes()
	fmt.Fprint(w, string(page))
}
