package main

import "hello"

func main() {
	d := hello.NewDemo("somkiat")
	r := d.SayHello()
	println(r)

}

//go fmt ./...
//go run hello.go
//go build -o xxx
