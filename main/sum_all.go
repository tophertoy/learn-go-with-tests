package main

// func SumAll(numbersToSum ...[]int) []int {
// 	lengthOfNumbers := len(numbersToSum)
// 	sums := make([]int, lengthOfNumbers)

// 	for i, numbers := range numbersToSum {
// 		sums[i] = Sum(numbers)
// 	}

// 	return sums
// }

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}

func SumAllHeads(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	headsTotals := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		headNumber := []int{numbers[0]}
		headsTotals[i] = Sum(headNumber)
	}
	return headsTotals
}

// len(numbers) == 0
// empty?(numbers)
