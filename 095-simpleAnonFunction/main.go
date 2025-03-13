package main

import (
	"fmt"
)

func main() {
	pi := 3.141
	i := 42

	func(j int) {
		fmt.Println(j, pi) // Note: pi is not declared in the func - is a closure - Here be dragons!
		pi *= 2
	}(i)

	fmt.Println(pi)
}
