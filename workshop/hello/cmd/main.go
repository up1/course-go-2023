package main

import (
	"hello"
	xx "hello/utils"
)

func main() {
	d := hello.NewDemo("somkiat")
	r := d.SayHello()
	println(r)

	xx.FormatXXX()

}

//go fmt ./...
//go run hello.go
//go build -o xxx
