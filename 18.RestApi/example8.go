package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/search", search)

	http.ListenAndServe(":7000", nil)

}

func search(res http.ResponseWriter, req *http.Request) {

	source := req.FormValue("source")
	q := req.FormValue("q")

	data := "/search?q=" + q + "&source=" + source

	res.Write([]byte (data))

}

// * https://www.google.com/search?q=golang&source=hp

// * https://www.google.com/search?q=golang&oq=golang&aqs=chrome
