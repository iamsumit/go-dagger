package main

import "testing"

func Test_Main(t *testing.T) {
	exp := "something"
	if exp != "something" {
		t.Errorf("expected %s, got %s", exp, "something")
	}
}
