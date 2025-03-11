package main

import (
	"fmt"
)

func main() {
	var myInt int16 = 42
	var myIntPtr *int16 = &myInt // myIntPtr := &myInt
	fmt.Println("Value contained in variable pointed to by the address in myIntPtr is:", *myIntPtr)
}
