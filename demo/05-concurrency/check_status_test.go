package demo_test

import (
	"demo"
	"reflect"
	"testing"
)

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://gmail.com",
		"http://nodata.xxx",
	}

	want := map[string]bool{
		"http://google.com": true,
		"http://gmail.com":  true,
		"http://nodata.xxx": false,
	}

	got := demo.CheckStatusWithRoutineAndChannel(websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}
