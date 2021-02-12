package main

import "fmt"

func main() {

	names := []string{
		"bob", "mike", "randy", "david", "leroy",
	}
	fmt.Println(names[0:2])
	fmt.Println(names[:2])
	fmt.Println(names[2:])

}
