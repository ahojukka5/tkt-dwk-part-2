package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	cnt := 0

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, cnt)
		cnt += 1
	})

	port := "8000"
	println("Ping/pong server started in port " + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
