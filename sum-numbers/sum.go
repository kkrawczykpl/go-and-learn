package main

func Sum(numbers []int) int {
	sum := 0

	// blank identifier - https://golang.org/doc/effective_go#blank
	for _, number := range numbers {
		sum += number
	}
	
	return sum
}