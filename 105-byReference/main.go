package main

import (
	"fmt"
)

func main() {
	var myInt int16 = 42
	myIntPtr := &myInt           // we don't really need this
	fmt.Println("First:", myInt) // we start with 42
	myFunc(myIntPtr)             // we are passing myInt ADDRESS (as a pointer)
	//myFunc(&myInt)			  // alternative - instead of pointer variable
	fmt.Println("Fourth:", myInt) // value of myInt was CHANGED by myFunc()
}

func myFunc(myPtr *int16) {
	fmt.Println("Second:", *myPtr) // contents of address in parameter
	*myPtr += 7                    // change the value of variable at address in myPtr
	fmt.Println("Third:", *myPtr)  // changed to 49
}
