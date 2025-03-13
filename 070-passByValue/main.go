package main

import (
	"fmt"
)

func main() {
	myInt := 42
	fmt.Println("First:", myInt) // we start with 42
	myFunc(myInt)
	fmt.Println("Fourth:", myInt) // "back to" 42 - copy of myInt was changed in myFunc() not the original
}

func myFunc(mi int) {
	fmt.Println("Second:", mi) // parameter supplied value is 42
	mi += 7
	fmt.Println("Third:", mi) // changed to 49
}
