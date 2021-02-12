package main

import "fmt"

// fallthrough jumps to the end
func main() {

	i := 142
	switch {
	case i > 100:
		fmt.Println("big")
	case i > 0:
		fmt.Println("positive")
		fallthrough
	default:
		fmt.Println("number")
	}
}
