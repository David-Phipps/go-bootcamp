package main

import (
	"fmt"
	"strconv"
)

func main() {
	s, err := strconv.Atoi("42o")
	if err != nil {
		fmt.Println("Returned error:", err)
	}
	fmt.Println(s)
}
