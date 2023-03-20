package main

func Sum(numbers [5]int) int {
	var sum int
	for i := 0; i < 5; i++ {
		sum += numbers[i]
	}
	return sum
}
