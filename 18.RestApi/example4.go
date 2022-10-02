package main

import (
	"fmt"
	"net/http"
)

const htmlCode = `
	<h1 style="color: red;">First Html Page</h1>
`

func main() {

	http.HandleFunc("/home", home)

	http.ListenAndServe(":3000", nil)

}

func home(res http.ResponseWriter, req *http.Request) {
	url := req.FormValue("url")
	if url == "" {
		res.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(res, htmlCode)
	}
}