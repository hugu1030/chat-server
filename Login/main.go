package main

import (
	"log"
	"net/http"
	"./login"
)

func handleLoginRequests(){
	 http.HandleFunc("/login",login.ReturnCanLogin)
	 http.HandleFunc("/newLogin",login.ReturnMadeAccount)
	 http.HandleFunc("/resetLogin",login.ReturnResetAccount)
     log.Fatal(http.ListenAndServe(":8081",nil))
}


func main() {
	handleLoginRequests()
}