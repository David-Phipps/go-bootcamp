package main

import "fmt"

func main() {
	// underlying type is still [5]int
	type bookcase [5]int

	blue := bookcase{6, 9, 3, 2, 1}
	red := [5]int{6, 9, 3, 2, 1}

	fmt.Print("are they equal ")

	if blue == red {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

	fmt.Printf("blue: %#v\n", blue)
	fmt.Printf("red : %#v\n", red)

}
