package hello_test

import (
	"hello"
	"testing"
)

func TestSayHi(t *testing.T) {
	d := hello.NewDemo("somkiat")
	r := d.SayHello()
	if r != "Hello World 1" {
		t.Errorf("Test fail ... got %s", r)
	}
}

//go test
//go test ./...
//go test -v
//go test -v -cover
//go test -coverprofile=coverage.out
//go tool cover -html=coverage.out
