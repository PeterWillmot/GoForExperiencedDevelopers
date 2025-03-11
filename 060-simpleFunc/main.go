package main

import (
	"errors"
	"fmt"
)

func main() {
	dividend, _ := myDivider(10, 5)

	fmt.Println(dividend)
}

func myDivider(numer int, denom int) (int, error) {
	if denom == 0 {
		return 0, errors.New("Divide by zero attempted")
	}

	return (numer / denom), nil
}
