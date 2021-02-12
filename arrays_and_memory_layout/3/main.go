package main

import "fmt"

const (
	winter = 1
	summer = 3
	yearly = winter + summer
)

func main() {
	books := [...]string{
		"Kafka's Revenge",
		"Kafka's Revenge 2nd Edition",
		"Stay Golden",
		"Everythingship",
	}

	fmt.Printf("books : %#v\n", books)

	var (
		wBooks [winter]string
		sBooks [summer]string
	)
	wBooks[0] = books[0]

	// // USING RANGE:
	for i := range sBooks {
		sBooks[i] = books[i+1]
	}

	fmt.Printf("\nwinter : %#v\n", wBooks)
	fmt.Printf("\nsummer : %#v\n", sBooks)

	var published [len(books)]bool

	published[0] = true
	published[len(books)-1] = true

	fmt.Println("\nPublished Books:")
	for i, ok := range published {
		if ok {
			fmt.Printf(" + %s\n", books[i])
		}
	}
}
