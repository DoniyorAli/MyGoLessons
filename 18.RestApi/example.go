package main

import (
	"net/http"	//* http paketidan foydalanamiz
)

func main() {

//! Portni ishga tushurip va portni bosh sahifasiga "Hello New York" shu yozuvni chiqarish

	//* ishga tushgan portimda "/" bosh sohifasiga --->
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte ("Hello NewYork")) //* ---> ekranga []byte sifatida shuni yozdirdim
		//* bunda Write orqali stringni byte ga ko'chirgan holatda yozdirdim, ekranga string ko'rinishida chiqishi uchun
	})

	http.ListenAndServe(":3000", nil)	//* port yani serverim ishga tushdi

}

//! ESLATMA
//* Google, Yandex yoki barcha shunga oxshash saytlaring barcha sahifalarida baytlar almashadi, turli xil sahifalar
//* baytlar asosida bo'ladi.
//* Aslini olganda ko'rinib turgan barcha narsalar bayt korinishida amalga oshiriladi
//* Dasturlashda hamma ishlash jarayonlari byte orqali amalga oshiriladi!