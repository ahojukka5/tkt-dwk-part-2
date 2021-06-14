package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, world!")
	})

	port := "8000"
	println("Server started in port " + port)
	http.ListenAndServe(":"+port, nil)
}
