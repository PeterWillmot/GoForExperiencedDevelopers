package main

import (
	"fmt"
	"time"
)

func main() {
	doStuff(2)
	doStuff(3)
	doStuff(1)

	fmt.Println("All done .... hit enter")
	fmt.Scanln()
}

func doStuff(id int) {
	time.Sleep(time.Duration(id) * time.Second)
	fmt.Println(id, "Done")
}
