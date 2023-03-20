package main

func Sum(numbers []int) int {
	var sum int
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	return sum
}
