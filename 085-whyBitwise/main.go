package main

import (
	"fmt"
)

func main() {
	var myBools [8]bool = [8]bool{true, false, true, false, true, true, false, true}
	var myByte byte = 173 // in binary 10101101

	fmt.Println("Let's start as beginners:\n")

	for i := 0; i < 8; i++ {
		fmt.Printf("(%d) %v ", i, checkIsSetForBabies(myBools, i))
	}

	fmt.Println("\n\nand now for the data size conscious:\n")

	for i := 0; i < 8; i++ {
		fmt.Printf("(%d) %v ", i, checkIsSetForSemiPros(myByte, i))
	}

	fmt.Println("\n\nAn 87.5% saving in data overhead!\n")
	fmt.Println("Now lets use less code:\n")

	for i := 0; i < 8; i++ {
		fmt.Printf("(%d) %v ", i, checkIsSetForPros(myByte, i))
	}

	fmt.Println("\n\nHow low can you go? :-)")
}

func checkIsSetForBabies(data [8]bool, switchNo int) bool {
	if data[switchNo] {
		return true
	} else {
		return false
	}
}

func checkIsSetForSemiPros(data byte, switchNo int) bool {
	switch switchNo {
	case 0:
		return (data & 128) != 0
	case 1:
		return (data & 64) != 0
	case 2:
		return (data & 32) != 0
	case 3:
		return (data & 16) != 0
	case 4:
		return (data & 8) != 0
	case 5:
		return (data & 4) != 0
	case 6:
		return (data & 2) != 0
	default:
		return (data & 1) != 0
	}
}

func checkIsSetForPros(data byte, switchNo int) bool {
	if (data & (1 << (7 - switchNo))) != 0 {
		return true
	} else {
		return false
	}
}
