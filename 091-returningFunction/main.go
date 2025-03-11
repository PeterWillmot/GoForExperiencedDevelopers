package main

import (
	"fmt"
)

func main() {
	callMe := incrementMe()

	fmt.Println(callMe())
	fmt.Println(callMe())
	fmt.Println(callMe())
}

func incrementMe() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
