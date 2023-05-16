package main

import (
	"fmt"
	"unsafe"
)

func main() {
	type first struct {
		b bool    // 1 byte
		f float64 // 8 bytes
		i int32   // 4 bytes
	}
	a := first{}
	fmt.Println(unsafe.Sizeof(a)) // 24 bytes
}
