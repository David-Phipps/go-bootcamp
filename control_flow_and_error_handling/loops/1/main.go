package main

import "fmt"

func main() {
	var sum int

	//init statement; condition; counter/post statement
	for i := 1; i <= 1000; i++ {
		sum += i
	}
	fmt.Println(sum)
}
