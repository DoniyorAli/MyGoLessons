package main

import (
	"fmt"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	x1 := &messageHandler{"This is a x1 message"}
	x2 := &messageHandler{"This is a x2 message"}

	mux.Handle("/one", x1)
	mux.Handle("/two", x2)

	fmt.Println("Port is available in 3000 port!")

	http.ListenAndServe(":8000", mux)

}

type messageHandler struct {
	message string
}

func (x *messageHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(res, x.message)

}