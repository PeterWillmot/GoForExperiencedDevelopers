package main

import (
	"fmt"
	"time"
)

func main() {
	scheduler()
	fmt.Println("All done ... back in main")

	// at this stage we could be doing other things
	// all heap allocations made by "leaked goroutines" will never be collected

	fmt.Println("Ready to exit main")
	fmt.Scanln()
}

func scheduler() {
	ch := make(chan int)

	//defer close(ch)	// if this is used then we will get panics for later attempts to write to ch

	for i := 3; i > 0; i-- {
		go doStuff(i, ch)
	}

	id := <-ch

	fmt.Println("First response from:", id)

	/*

		//doStuff(2) will be blocked until we read its content from the ch

		fmt.Println("Hit enter to get next")

		id = <-ch

		fmt.Println("Next response from:", id)
	*/

	fmt.Println("Leaving scheduler")
}

func doStuff(id int, ch chan int) {
	fmt.Println(id, "Started") // display order for this output not deterministic!

	//var myHeapData []float64 = make([]float64, 1000000)

	//myHeapData[0] = float64(id)

	time.Sleep(time.Duration(id) * time.Second)

	ch <- id // potential "goroutine leak"

	fmt.Println(id, "Done")
}
