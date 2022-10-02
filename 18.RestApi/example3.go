package main

import (
	"net/http"
)

func indexHandler(res http.ResponseWriter, req *http.Request) {

	data := req.URL.Path[1:] //* yani birinchi / dan keyingi qismini pathni oladi
	data = "Hello " + data	 //* va sahifaga (Hello Alixan) dep chiqaradi
	
	if len(req.URL.Path[1:]) > 0 {
		res.Write([]byte (data))
	}else {
		res.Write([]byte ("Paramas no entered!"))
	}
}

func main() {

	http.HandleFunc("/", indexHandler) //* yani localhost:3000/Alixan, agar "/" dan keyin yozmasak (Hello) ni ozi chiqadi

	http.ListenAndServe(":3000", nil)

}