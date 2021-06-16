package main

import (
	"os"
	"time"
)

func main() {
	for {
		f, err := os.Create("timestamp")
		if err != nil {
			panic(err)
		}
		defer f.Close()
		t := time.Now().UTC().Format(time.RFC3339)
		f.WriteString(t)
		time.Sleep(5 * time.Second)
	}
}
