package main

import (
	"fmt"
	"math"
)

func main() {
	// find the sub-array [a:b], where 0 <= a <= b <= array_length, whose numbers sum to the greatest value
	// values can be negative or positive

	// creating example slices to use
	intSlice := []int{1, 5, 7, -10, 1, 0, -5, 9, -2, 10}
	// intSlice := []int{1, 1, -5, 1, 1}
	// intSlice := []int{4, 6, 5, 5, 10}
	// intSlice := []int{0, 0, 0, 0, 0}
	// intSlice := []int{-5, -100, -2, -3, -4}

	// brute force method commented out below

	// Kadane's Algorithm - O(n) for all cases
	// Kadane's Algorithm focuses on the sub-arrays that end at the given index a[i]
	// the maximum value will either be 1) the maximal sub-array ending at a[i-1] + a[i] OR 2) a[i]

	// maxEnding will keep track of the maximum value from the maximal sub-array ending at a[i-1]
	maxEnding := 0
	// currentMax will keep track of the current maximum value across all sub-arrays
	currentMax := 0
	// indices will map the indices of the maximal sub-array ending at a[i]
	indices := make(map[int][]string)
	aIndex := 0
	bIndex := 0

	// allZeros will be used to check if all values in slice are zero
	allZeros := true

	for i, v := range intSlice {
		// we want to treat an array of all zeros as if it were an array of all positive numbers greater than zero
		// in other words, the maximal sub-array would be [0:n-1] where n is the length of the slice/array
		// if allZeros remains true, we will adjust the indices manually.
		if v != 0 {
			allZeros = false
		}
		if i == 0 {
			// at intSlice[0], there is only one element to consider
			maxEnding = v
			currentMax = maxEnding
			// at intSlice[0], the maximal sub-array will always be itself [0:0]
			indices[maxEnding] = append(indices[maxEnding], fmt.Sprintf("[%d:%d]", aIndex, bIndex))
		} else {
			// indices are adjusted based on which is greater
			// if maxEnding + v >= v, then the bIndex is set to i to include a[i] in the indices of the sub-array
			// if v > maxEnding + v, then both aIndex and bIndex are set to i as a[i] is a greater sub-array than a[i-1]+a[i]
			if maxEnding+v >= v {
				bIndex = i
			} else {
				aIndex = i
				bIndex = i
			}
			// maxEnding is updated to the max of either the previous maxEnding + v OR v itself
			maxEnding = int(math.Max(float64(maxEnding+v), float64(v)))
			// currentMax is updated to the max of either the new maxEnding OR to currentMax itself
			currentMax = int(math.Max(float64(currentMax), float64(maxEnding)))
			indices[maxEnding] = append(indices[maxEnding], fmt.Sprintf("[%d:%d]", aIndex, bIndex))
		}
	}

	if allZeros {
		// manually set the indices to the entire slice/array
		indices[currentMax] = []string{fmt.Sprintf("[%d:%d]", 0, len(intSlice)-1)}
	}

	fmt.Printf("The sub-array(s) with the greatest sum is/are %v and the sum is %d.", indices[currentMax], currentMax)

	/*
		// brute force method - O(n) under the conditions that all numbers are positive or zero, O(n^2) otherwise
		// maxSum will keep track of the greatest value summed thus far - it is set to the minimum integer value to accommodate
		// negative values at intSlice[0]
		maxSum := math.MinInt
		// in the event that more than one sub-array sums up to the greatest value summed, m will map all indices at which the
		// intermediate sum is greater than or equal to the current maxSum
		m := make(map[int][]string)

		// testing for all positive numbers
		allPositive := true
		// testing if all values are zero
		allZero := true

		for indexA, valueA := range intSlice {
			sum := 0
			// when indexA is zero, we iterate through the entire slice
			// we check for all numbers being positive or zero during this iteration
			// we can terminate the for loops in either case as the sub-array that sums
			// to the greatest value will always be the entire "array" or slice itself
			if indexA > 0 && allPositive {
				// if the slice contained all positive numbers but not all zeros
				if !allZero {
					// the maximum is the sum of all elements in the array which
					// is achieved by the first pass of the nested for loop
					break
				}
				// if the slice contained all zeros, set the indices to be the entire slice
				// the maximum is zero which is the result of completing the first pass of
				// the nested for loop
				m[0] = []string{fmt.Sprintf("[0:%d]", len(intSlice)-1)}
				break
			}
			for indexB := indexA; indexB < len(intSlice); indexB++ {
				// check for negative values while iterating through the entire slice
				if indexA == 0 && intSlice[indexB] < 0 {
					allPositive = false
				}
				// check if the value is not zero
				if indexA == 0 && intSlice[indexB] != 0 {
					allZero = false
				}
				if indexA == indexB {
					// when indexA and indexB are equal, there is only one element to consider
					// the sum is set to valueA since there isn't anything to sum together yet
					sum = valueA
					if sum >= maxSum {
						maxSum = sum
						m[maxSum] = append(m[maxSum], fmt.Sprintf("[%d:%d]", indexA, indexB))
					}
				} else {
					sum += intSlice[indexB]
					if sum >= maxSum {
						// update maxSum
						maxSum = sum
						// update map
						m[maxSum] = append(m[maxSum], fmt.Sprintf("[%d:%d]", indexA, indexB))
					}
				}
			}
		}

		fmt.Printf("The sub-array(s) with the greatest sum is/are %v and the sum is %d.", m[maxSum], maxSum)
	*/
}
