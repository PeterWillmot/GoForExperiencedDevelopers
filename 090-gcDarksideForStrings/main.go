package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

var mem runtime.MemStats

func main() {
	initial := getMemStats()
	doStuff()
	//doStuffBetter()
	final := getMemStats()

	fmt.Println("Alloc: ", final-initial)
}

func getMemStats() uint64 {
	runtime.ReadMemStats(&mem)

	fmt.Printf("Alloc: %d Heap: %d Total : %d\n", mem.Alloc, mem.HeapAlloc, mem.TotalAlloc)

	return mem.TotalAlloc
}

func doStuff() {
	s := ""

	for i := 0; i < 100; i++ {
		t := strconv.Itoa(i) + " "
		s += t
	}

	fmt.Println("result len: ", len(s))
}

func doStuffBetter() {
	var sb strings.Builder

	for i := 0; i < 100; i++ {
		t := strconv.Itoa(i) + " "
		sb.WriteString(t)
	}

	fmt.Println("Result len: ", len(sb.String()))
}
