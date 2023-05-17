package main

import "fmt"

type user struct {
	id   int
	name string
}

func (u user) String() string {
	return fmt.Sprintf("ID=%d, Name=%s", u.id, u.name)
}

func main() {
	u := user{}
	fmt.Printf("%+v\n", u)
	fmt.Println(u)
}
