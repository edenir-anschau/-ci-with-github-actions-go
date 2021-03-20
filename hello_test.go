package main

import "testing"

func Test_Name(t *testing.T) {
	r := Hello()

	if r != "Hello World!" {
		t.Error()
	}
}
