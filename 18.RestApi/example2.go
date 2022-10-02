package main

import (
	"net/http"
)
//* Endi FrontEnd dan BackEnd ga data opkelamiz, yani parametrlarni ko'ramiz
//* misol uchun localhost:3000(/about/hello/)  <--- bular parametrlar 

func aboutHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte ("About Page !!!"))

	data := req.URL.Path[1:] //* bunda request jonatilgan URL ni bor Path ni ol va data ga tenglashtir
	res.Write([]byte (data)) //* yani sahifaga (About Page !!! about) parametr bilan birga chiqaradi
}

func main() {

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte ("Hi first Page"))
	})

	http.HandleFunc("/about", aboutHandler)

	http.ListenAndServe(":3000", nil)

}