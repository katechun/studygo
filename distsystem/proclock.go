package main

import "sync"

func main() {
	var wg sync.WaitGroup
	var l sync.Mutex
	var counter int
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			l.Lock()
			counter++
			l.Unlock()
		}()
	}
	wg.Wait()
	println(counter)
}
