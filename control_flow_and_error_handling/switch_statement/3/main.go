package main

import "fmt"

//short switch example

func main() {
	switch i := 10; {
	case i > 0:
		fmt.Println("positive")
	case i < 0:
		fmt.Println("negative")
	default:
		fmt.Println("zero")

	}
}
