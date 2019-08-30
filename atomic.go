package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var total1 uint64

func worker1(wg *sync.WaitGroup) {
	defer wg.Done()

	var i uint64

	for i = 0; i <= 100; i++ {
		atomic.AddUint64(&total1, i)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go worker1(&wg)
	go worker1(&wg)

	wg.Wait()

	fmt.Println(total1)
}
