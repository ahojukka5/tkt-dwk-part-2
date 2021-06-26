package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = ":8000"

func main() {
	cnt := 0

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		println("pingpong called, cnt =", cnt)
		fmt.Fprint(w, cnt)
		cnt += 1
	})

	println("Ping/pong server listening in address http://localhost" + port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
