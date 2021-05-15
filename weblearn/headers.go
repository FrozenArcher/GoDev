package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func headers(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	h := r.Header
	fmt.Fprintln(w, h)
}
func headersByKey(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	h := r.Header.Get(p.ByName("key"))
	fmt.Fprintln(w, "You are asking the header: ", p.ByName("key"))
	fmt.Fprintln(w, h)
}
