package main

import "testing"

func TestReverse(t *testing.T) {
	in := "jacob"
	want := "nikhil"
	got := Reverse(in)

	if got != want {
		t.Errorf("Reverse(%q) = %q, want %q", in, got, want)
	}
}
