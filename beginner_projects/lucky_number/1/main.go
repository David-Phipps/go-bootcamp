package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// gets a time value to use as random nubmer seed
	//t := time.Now()
	//seed with current unix time value t
	//rand.Seed(t.UnixNano())

	//one line way to do it
	rand.Seed(time.Now().UnixNano())

	guess := 10

	for n := 0; n != guess; {
		n = rand.Intn(guess + 1)
		fmt.Printf("%d ", n)
	}
	fmt.Println()
}
