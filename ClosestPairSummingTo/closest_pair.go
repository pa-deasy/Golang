package closestpair

import "math"

func ClosestPairSummingTo(numbers []int32, target int32) (int32, int32) {
	leftNumber := numbers[0]
	rightNumber := numbers[len(numbers)-1]
	var closestSum int32 = math.MaxInt32
	sortedNumbers := quickSort(numbers)

	for leftIndex, rightIndex := 0, len(sortedNumbers)-1; leftIndex < rightIndex; {
		sum := sortedNumbers[leftIndex] + sortedNumbers[rightIndex]

		if math.Abs(float64(target-sum)) < math.Abs(float64(target-closestSum)) {
			closestSum = sum
			leftNumber = sortedNumbers[leftIndex]
			rightNumber = sortedNumbers[rightIndex]
		}

		if sum > target {
			rightIndex--
		} else {
			leftIndex++
		}
	}

	return leftNumber, rightNumber
}

func quickSort(numbers []int32) []int32 {
	if len(numbers) < 2 {
		return numbers
	}

	pivot := numbers[0]
	var lessThanOrEqualTo []int32
	var greaterThan []int32

	for i, number := range numbers {
		if i == 0 {
			continue
		}

		if number > pivot {
			greaterThan = append(greaterThan, number)
		} else {
			lessThanOrEqualTo = append(lessThanOrEqualTo, number)
		}
	}

	var sorted []int32
	sorted = append(sorted, quickSort(lessThanOrEqualTo)...)
	sorted = append(sorted, pivot)
	sorted = append(sorted, quickSort(greaterThan)...)

	return sorted
}
