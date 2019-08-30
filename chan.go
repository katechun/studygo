package main

import (
	"fmt"
)

func main() {
	done := make(chan int)

	go func() {
		println("你好")
		done <- 1
	}()
	m := <-done

	fmt.Println(m)
}
