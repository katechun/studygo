package main

import (
	"fmt"
)

func main() {
	var a = []int{1, 2, 3}

	a = append([]int{-3, -2, -1}, a...)
	fmt.Println(a, " len:", len(a), " sap:", cap(a))
}
