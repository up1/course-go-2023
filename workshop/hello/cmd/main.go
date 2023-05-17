package main

import "hello"

func main() {
	d := hello.Demo{
		Name: "somkiat",
	}
	d.SayHello()
}

//go fmt ./...
//go run hello.go
//go build -o xxx
