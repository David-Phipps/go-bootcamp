package main

import "fmt"

func main() {
	prev := [3]string{"Kafka's Revenge",
		"Stay Golden", "Everythingship"}

	//go will create a new array
	var books [4]string

	for i := range prev {
		books[i] = prev[i] + " 2nd Ed."
	}

	fmt.Printf("last year:\n%#v\n", prev)
	fmt.Printf("last year:\n%#v\n", books)
}
