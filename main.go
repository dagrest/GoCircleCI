package main

import "fmt"

func main() {
	fmt.Println("Hello CircleIC")
	fmt.Printf("Foo parameter is: %d", foo(100))
}

func foo(n int) int {
	return n
}
