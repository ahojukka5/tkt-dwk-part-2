package main

import (
	"net/http"
)

func main() {
	port := "8000"
	println("Server started in port " + port)
	http.ListenAndServe(":"+port, nil)
}
