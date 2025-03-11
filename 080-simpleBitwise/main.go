package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%d\n", 2&6)                   // 2
	fmt.Printf("%d\n", 0x02&0x06)             // 2
	fmt.Printf("%d\n", 0b00000010&0b00000110) // 2
	fmt.Printf("%d %d\n", 5, (5 ^ 5))         // 0
}
