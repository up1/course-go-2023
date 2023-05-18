package main

import (
	"demo"
	"fmt"
)

func main() {
	websites := []string{
		"http://google.com",
		"http://gmail.com",
		"http://nodata.xxx",
	}

	results := demo.CheckStatus(websites)

	for web, status := range results {
		fmt.Println(web, status)
	}
}