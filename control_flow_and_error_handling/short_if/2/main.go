package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if a := os.Args; len(a) != 2 {
		fmt.Println("Give me 1 number only")
	} else if _, err := strconv.Atoi(a[1]); err != nil {
		fmt.Println("There was an error: ", err)
	} else {
		fmt.Println("Success, your thing is:", a[1])
	}

}
