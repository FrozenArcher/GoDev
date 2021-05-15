package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	_, _ = fmt.Fprintf(w, "hello, %s!\n", p.ByName("name"))
}
