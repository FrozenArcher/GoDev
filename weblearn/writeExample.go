package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Post struct {
	User    string
	Threads []string
}

func writeExample(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	str := `<html>
<head><title>WebLearn</title></head>
<body><h1>Hello World!</h1></body>
</html>`
	w.Write([]byte(str))
}

func write501(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "No such service, try next door.")
}

func redirect(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	w.Header().Set("Location", "https://www.baidu.com")
}

func jsonExample(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		User:    "yhy",
		Threads: []string{"First", "Second"},
	}
	json, _ := json.Marshal(post)
	w.Write(json)
}
