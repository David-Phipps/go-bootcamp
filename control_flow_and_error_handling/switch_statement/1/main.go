package main

import (
	"fmt"
	"os"
)

func main() {
	// Switch statement looks for an equal value
	city := os.Args[1]

	switch city {
	case "Paris":
		fmt.Println("France")
	case "Winfield":
		fmt.Println("FrAAnce")
	default:
		fmt.Print("FFFFrance \n")
	}
}
