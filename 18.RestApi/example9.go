package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {

	http.HandleFunc("/add", add)

	http.ListenAndServe(":7000", nil)

}

func add(res http.ResponseWriter, req *http.Request) {

	n1 := req.FormValue("n1")
	intN1, _ := strconv.Atoi(n1)
	n2 := req.FormValue("n2")
	intN2, _ := strconv.Atoi(n2)

	data := intN1 + intN2

	res.Write([]byte (fmt.Sprint(data)))

}
