package main

import (
	"fmt"
	"net/http"
)

func main() {
	response, err := http.Get("https://go.junk")

	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Println("RESP:", response)
	}

}
