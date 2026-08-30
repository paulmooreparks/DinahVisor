package main

import "testing"

func TestGreetingNamesTheProgram(t *testing.T) {
	got := greeting()
	if got != "DinahVisor is not built yet." {
		t.Fatalf("greeting() = %q", got)
	}
}
