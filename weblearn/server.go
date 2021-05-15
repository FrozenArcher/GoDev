package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	mux := httprouter.New()
	mux.GET("/", index)
	mux.GET("/hello/:name", hello)
	mux.GET("/headers", headers)
	mux.GET("/headers/:key", headersByKey)
	mux.GET("/body", body) // !
	mux.GET("/process", process)
	mux.GET("/write", writeExample)
	mux.GET("/writeheader", write501)
	mux.GET("/redirect", redirect)
	mux.GET("/json", jsonExample)
	mux.GET("/set_cookie", setCookie)
	mux.GET("/get_cookie", getCookie)
	mux.GET("/set_message", setMessage)
	mux.GET("/show_message", showMessage)
	mux.GET("/template", parseTpl)

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	_ = server.ListenAndServe()
}

func index(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	fmt.Fprintln(w, "hello, this is a web learning application.")
}
