package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// !
func body(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	fmt.Fprintln(w, string(body))
}
