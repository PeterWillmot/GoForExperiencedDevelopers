package main

import (
	"fmt"
	"time"
	//	"time"
)

func main() {
	ch := make(chan int)

	defer close(ch)

	for i := 3; i > 0; i-- {
		fmt.Println("starting", i)
		go doStuff(i, ch)
	}

	id := <-ch

	/*
		id, ok := <-ch

		if !ok {
			fmt.Println("channel was closed")
		}
	*/

	//close(ch)    // this will cause a panic - channel closed too early

	fmt.Println("First response from:", id)
	fmt.Println("All done .... hit enter")
	fmt.Scanln()
}

func doStuff(id int, ch chan int) {
	time.Sleep(time.Duration(id) * time.Second)
	fmt.Println(id, "Done")
	ch <- id
}
