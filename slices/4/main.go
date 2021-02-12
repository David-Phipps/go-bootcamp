package main

import "fmt"

func main() {
	names := []string{
		"bob", "david", "mark", "jimmuy",
		"larry", "bobby", "fran", "joey",
		"mark", "MARY", "elvie", "gary",
		"ed",
	}

	const pageSize = 4

	l := len(names)

	for from := 0; from < l; from += pageSize {
		to := from + pageSize
		if to > l {
			to = l
		}

		currentPage := names[from:to]

		fmt.Println("", currentPage)
	}
}
