package main

import (
	"fmt"
	"unsafe"
)

func main() {
	type first struct {
		f float64
		i int32
		b bool
	}
	a := first{}
	fmt.Println(unsafe.Sizeof(a)) // 24 bytes
}
