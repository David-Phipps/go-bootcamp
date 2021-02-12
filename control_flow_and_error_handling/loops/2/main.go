package main

//for loop with a break
func main() {
	var sum, i int

	for {
		if i > 5 {
			break
		}

		sum += i
		i++
	}
}
