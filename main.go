package main

import (
	"fmt"
	logic "gocircleci/app/src"
)

func main() {
	fmt.Println("Hello CircleIC")
	fmt.Printf("Foo parameter is: %d", 100)
	logic.WhichInput(100)
}
