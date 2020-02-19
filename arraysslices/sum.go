package arraysslices

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var result []int

	for _, numbers := range numbersToSum {
		result = append(result, Sum(numbers))
	}

	return result
}

func SumAllTails(tails ...[]int) []int {
	var result []int

	for _, numbers := range tails {
		sum := 0
		if len(numbers) != 0 {
			sum = Sum(numbers[1:])
		}
		result = append(result, sum)
	}

	return result
}
