package main

import "time"

func main() {
	go hello()
	println("World")
	time.Sleep(1 * time.Second)
}
func hello() {
	println("Hello")
}
