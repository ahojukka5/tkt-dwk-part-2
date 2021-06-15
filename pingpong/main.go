package main

import (
	"fmt"
	"net/http"
)

func main() {
	cnt := 0

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, cnt)
		cnt += 1
	})

	port := "8002"
	println("Ping/pong server started in port " + port)
	http.ListenAndServe(":"+port, nil)
}
