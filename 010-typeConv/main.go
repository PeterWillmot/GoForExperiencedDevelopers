package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	myInt1 := 42
	myString1 := strconv.Itoa(myInt1)
	myInt2, err := strconv.Atoi(myString1)

	if err != nil {
		log.Fatal("invalid Atoi arg")
	}

	fmt.Println(myInt1, myString1, myInt2)
}
