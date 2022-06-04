package main

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	res := foo(100)
	fmt.Printf("Result: %d", res)
}
