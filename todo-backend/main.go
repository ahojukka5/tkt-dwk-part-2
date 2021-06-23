package main

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func Health(w http.ResponseWriter, r *http.Request) {
	println("API Health is OK")
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"status": "ok"}`)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/health", Health).Methods("GET")
	http.ListenAndServe(":8000", router)
}
