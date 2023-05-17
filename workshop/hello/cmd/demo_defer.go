package main

func main() {
	call()
}

func call() {
	for i := 0; i < 10; i++ {
		defer print(i)
	}

	print("Hello ")
}
