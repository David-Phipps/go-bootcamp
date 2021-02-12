package main

import "fmt"

func main() {
	var (
		blue   = [3]int{6, 9, 3}
		red    = [3]int{6, 9, 3}
		yellow = [2]int{6, 9}
	)

	fmt.Printf("blue bookcase : %v\n", blue)
	fmt.Printf("red bookcase : %v\n", red)

	fmt.Println("Are they equal?", blue == red)
	fmt.Println("Are they equal?", blue == yellow) //length of an array is part of its type!
}
