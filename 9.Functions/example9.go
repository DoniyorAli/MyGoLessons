package main

import "fmt"

//* DEFER

var isConnected = false

func main() {
	fmt.Println("Connection status:", isConnected)	//! 1 ===> Connection status: false
	dbProcessing()
	fmt.Println("Connection status:", isConnected)	//! 7 ===> Connection status: false  // ???????
}

func dbProcessing() {
	defer disconnected()
	connect()
	fmt.Println("Defering disconnected!")			//! 3 ===> Defering disconnected!
	fmt.Println("Connection status", isConnected)	//! 4 ===> Connection status true  // true ??????
	fmt.Println("DataBase is working !")		 	//! 5 ===> DataBase is working !
}

func connect() {
	isConnected = true
	fmt.Println("Connected To DataBase!!!")			//! 2 ===> Connected To DataBase!!!
}

func disconnected() {
	isConnected = false
	fmt.Println("Disconnected!!!")					//! 6 ===> Disconnected!!!
}
//? QUESTON