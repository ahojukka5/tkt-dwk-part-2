package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func main() {
	random_string := uuid.NewV4()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t, err := ioutil.ReadFile("/cache/timestamp")
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, "%s: %s.\n", t, random_string)

		// for local testing
		// resp, err := http.Get("http://localhost:8081/pingpong")

		resp, err := http.Get("http://pingpong-svc:8000")
		if err != nil {
			panic(err)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		cnt := string(body)
		fmt.Fprintf(w, "Ping Pongs: %s\n", cnt)
	})

	port := "3000"
	println("Server started in port " + port)
	http.ListenAndServe(":"+port, nil)

}
