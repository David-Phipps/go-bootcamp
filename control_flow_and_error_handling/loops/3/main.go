package main

import "fmt"

//for loop with a continue
func main() {
	var sum, i int

	for {
		if i > 5 {
			break
		}

		sum += i
		fmt.Println((sum))
		i++

		if i%2 == 0 {
			continue
		}
	}
}
