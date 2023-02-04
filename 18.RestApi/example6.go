package main

import (
	"io"
	"net/http"
)

func main() {
	var t tashkent
	var i istanbul

	mux := http.NewServeMux()
	mux.Handle("/tashkent", t)
	mux.Handle("/istanbul", i)

	http.ListenAndServe(":5000", mux)

}

type tashkent string
type istanbul string

func (city tashkent) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hello Tashkent")
}

func (city istanbul) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hello istanbul")
}