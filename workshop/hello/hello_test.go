package hello_test

import (
	"hello"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSayHi(t *testing.T) {
	d := hello.NewDemo("somkiat")
	r := d.SayHello()
	if r != "Hello World 1" {
		t.Errorf("Test fail ... got %s", r)
	}
	assert.Equal(t, "Hello World 1", r)
}

//go test
//go test ./...
//go test -v
//go test -v -cover
//go test ./... -coverprofile=coverage.out
//go tool cover -html=coverage.out
