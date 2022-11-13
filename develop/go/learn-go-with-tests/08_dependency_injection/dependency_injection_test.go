package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Ray")

	got := buffer.String()
	want := "Hello, Ray"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
