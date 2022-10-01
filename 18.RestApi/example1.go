package main

import (
	"net/http"
)


func aboutHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte ("About Page"))
}

func main() {
	//* Handler
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte ("Hi first Page")) // * bunda localhostga "/" deb yozsak "Hi first page" ni sahifaga chiqaradi
	})
//* Yangi endpoint kiritish yani "/" bunga parametr sifatida "/about" endpointni nomini yozamiz
//* localhostga "/about" deb yozsak tepadegi (func aboutHandler) ni ishga tushurib beradi va sahifani ochadi
	http.HandleFunc("/about", aboutHandler) 

	http.ListenAndServe(":3000", nil)

}