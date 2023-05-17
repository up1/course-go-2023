package main

type ABC func(int) int

func main() {
	println(do(incBy5, 0))
}

func do(f ABC, x int) int {
	return f(x)
}

func incBy5(i int) int {
	return i + 5
}
