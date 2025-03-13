package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 5)

	ch <- 1
	ch <- 2
	ch <- 3

	fmt.Println("Capacity:", cap(ch))
	fmt.Println("Items:", len(ch))

	ch <- 4
	ch <- 5

	close(ch)

	for i := range ch {
		fmt.Println(i)
	}
}
