package main

import (
	"fmt"
	"time"
	//	"time"
)

func main() {
	scheduler()
	fmt.Println("All done ... back in main.... hit enter")
	fmt.Scanln()
}

func scheduler() {
	ch := make(chan int)

	//defer close(ch)

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

	fmt.Println("Leaving scheduler")
}

func doStuff(id int, ch chan int) {
	fmt.Println(id, "Started")
	time.Sleep(time.Duration(id) * time.Second)
	ch <- id
	fmt.Println(id, "Done")
}
