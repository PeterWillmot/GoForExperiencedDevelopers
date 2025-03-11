package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wgrp sync.WaitGroup

var totalDelay uint64 = 0

//var safer sync.Mutex

func main() {
	rng := rand.New(rand.NewSource(42))

	for w := 1; w < 20; w++ {
		wgrp.Add(1)
		go worker(w, (rng.Intn(5) + 1))
	}

	wgrp.Wait()

	fmt.Print("total delay:", totalDelay)
}

func worker(id int, delay int) {
	fmt.Println(id, " starting for: ", delay)
	time.Sleep(time.Duration(delay) * time.Second)

	//safer.Lock()
	totalDelay += uint64(delay)
	//safer.Unlock()

	fmt.Println(id, " completed after: ", delay)
	wgrp.Done()
}
