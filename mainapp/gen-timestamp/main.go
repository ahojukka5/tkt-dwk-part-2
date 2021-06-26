package main

import (
	"os"
	"time"
)

func main() {
	for {
		f, err := os.Create("/cache/timestamp")
		if err != nil {
			panic(err)
		}
		t := time.Now().UTC().Format(time.RFC3339)
		println("Writing timestamp t = " + t + " to /cache/timestamp")
		f.WriteString(t)
		f.Close()
		time.Sleep(5 * time.Second)
	}
}
