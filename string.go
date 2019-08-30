package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	s := "hello world"
	hello := s[5:]
	world := s[6:]

	s1 := "hello world"[:5]
	s2 := "hello world"[6:]
	fmt.Println("hello:", hello, " world:", world)
	fmt.Println("hello:", s1, " world:", s2)
	fmt.Println("-------------------")

	fmt.Println("len(s):", (*reflect.StringHeader)(unsafe.Pointer(&s)).Len)
	fmt.Println(len(s), " ", len(s1), " ", len(s2))

	fmt.Println("-----------------------------")
	fmt.Printf("%#v\n", []byte("hello,世界"))

	fmt.Println("--------------------------------")
	for i, c := range "\xe4\x00\x00\xe7\x95\x8cabc" {
		fmt.Println(i, c)
	}
}
