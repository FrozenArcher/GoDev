package main

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

func parseTpl(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	const simpleDir = "templates/simplest.html"
	t := template.Must(template.ParseFiles(simpleDir))
	t.Execute(w, "Hello, World!")
}
