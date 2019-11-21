package main

import "testing"

func TestMapAssign(t *testing.T) {
	m := make(map[string]string)
	assign(m)
	t.Log(m)
}

func assign(m map[string]string) {
	m["hello"] = "world"
}
