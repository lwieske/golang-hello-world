package main

import "testing"

func TestHelloWorld(t *testing.T) {
	if helloworld() != "Hello, world!" {
		t.Fatal("Test failed.")
	}
}
