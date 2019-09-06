package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, 256)
	fmt.Println(b)
	fmt.Println(binary.BigEndian.Uint32(b))
}
