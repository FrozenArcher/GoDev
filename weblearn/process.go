package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.ParseForm()
	fmt.Fprintln(w, r.Form)
}
