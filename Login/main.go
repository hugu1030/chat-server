package main

import (
	"./login"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func forCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		// プリフライトリクエストの対応
		if r.Method == "OPTIONS" {
			log.Println("OPTIONS")
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
		return
	})
}

func main() {
	router := mux.NewRouter()
	router.Use(forCORS)

	router.HandleFunc("/login", login.ReturnCanLogin)
	router.HandleFunc("/newLogin", login.ReturnMadeAccount)
	router.HandleFunc("/resetLogin", login.ReturnResetAccount)
	log.Fatal(http.ListenAndServe(":8081", router))

	println("start server.")
}
