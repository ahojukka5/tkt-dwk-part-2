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
		t2, err2 := ioutil.ReadFile("/media/shared/number_of_pingpongs")
		if err2 != nil {
			panic(err2)
		}
		fmt.Fprintf(w, "Ping Pongs: %s\n", t2)
	})

	port := "8001"
	println("Server started in port " + port)
	http.ListenAndServe(":"+port, nil)

}
