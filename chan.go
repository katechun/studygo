package main

import (
	"fmt"
)

func main() {
	done := make(chan int)

	go func() {
		fmt.Println("你好")
		done <- 1
	}()
	m := <-done

	fmt.Println(m)

	fmt.Println("结束")
}
