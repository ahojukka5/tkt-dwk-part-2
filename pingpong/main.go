package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func main() {
	cnt := 0

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, cnt)

		f, err := os.Create("/media/shared/number_of_pingpongs")
		if err != nil {
			panic(err)
		}
		defer f.Close()
		f.WriteString(strconv.Itoa(cnt))

		cnt += 1
	})

	port := "8002"
	println("Ping/pong server started in port " + port)
	http.ListenAndServe(":"+port, nil)
}
