package main

import (
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
)

func main() {
	r := uuid.NewV4()
	for {
		t := time.Now().UTC().Format(time.RFC3339)
		fmt.Printf("%s: %s\n", t, r)
		time.Sleep(5 * time.Second)
	}
}
