package main

func Sum(numbers []int) int {
	sum := 0

	// blank identifier - https://golang.org/doc/effective_go#blank
	for _, number := range numbers {
		sum += number
	}
	
	return sum
}

func SumAll(numbersToSum ...[]int) (sums []int) {

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}