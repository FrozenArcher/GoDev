package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func setCookie(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	c1 := http.Cookie{
		Name:     "first_cookie",
		Value:    "Web Learning",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name:     "second_cookie",
		Value:    "YHY",
		HttpOnly: true,
	}
	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
}

func getCookie(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	c1, err := r.Cookie("first_cookie")
	if err != nil {
		fmt.Fprintln(w, "Cannot get the first cookie.")
	}
	cs := r.Cookies()
	fmt.Fprintln(w, c1)
	fmt.Fprintln(w, cs)
}
