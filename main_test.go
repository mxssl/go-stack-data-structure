package main

import (
	"testing"
)

func TestPush(t *testing.T) {

	push(100)

	if !traverse(stack) {
		t.Error("Push is incorrect")
	}
}

func TestPop(t *testing.T) {

	pop(stack)

	if traverse(stack) {
		t.Error("Pop is incorrect")
	}
}
