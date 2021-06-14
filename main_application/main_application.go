package main

import (
	"fmt"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

func main() {
	random_string := uuid.NewV4()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now().UTC().Format(time.RFC3339)
		fmt.Fprintf(w, "%s: %s\n", t, random_string)
	})

	port := "8001"
	println("Server started in port " + port)
	http.ListenAndServe(":"+port, nil)

}
