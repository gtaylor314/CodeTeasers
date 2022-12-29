package main

import (
	"fmt"
	"math"
)

func main() {
	// Goal: Implement a power function (i.e. a to the b power) without using the multiplication operator
	// It is assumed that all values of b (the exponent) are positive but values of a can be positive or negative
	// If a is negative, -a to the b power is assumed to be "(-a) to the b power" and not "-(a to the b power)"

	// test values
	// a, b := 0, 0
	// a, b := 10, 0
	// a, b := 0, 100
	// a, b := 5, -2
	// a, b := 100, 3
	// a, b := 5, 3
	// a, b := -2, 4
	a, b := -3, 3

	switch {
	// if the exponent is zero, we have two cases to be concerned about: 1) a is zero which is undefined, 2) a is not zero -
	// if a is not zero, any value of "a to the 0 power," except for zero, is 1
	case a == 0 && b == 0:
		fmt.Println("Undefined: 0 to the 0 power is undefined. Please choose two new values.")
		return
	case a != 0 && b == 0:
		fmt.Printf("%d to the %d power = %d", a, b, 1)
		return
	// if a is zero but b is an exponent greater than zero, we can immediately return 0 - this simply saves time
	case a == 0 && b > 0:
		fmt.Printf("%d to the %d power = %d", a, b, 0)
		return
	// if b is negative, we will print an error message stating the exponent must be positive
	case b < 0:
		fmt.Println("Exponent must be positve. Please choose a positive exponent.")
		return
	}

	// we will treat a as if it were positive and manually take care of the sign when the calculation is complete
	negSign := false
	copyA := a
	if a < 0 {
		// since a is negative, make its copy (copyA) positive
		copyA = -a
		// if b is an even value, the result of (-a) to the b power will be positive and negSign remains false
		// if b is an odd value, the result of (-a) to the b power will be negative and negSign becomes true
		if b%2 != 0 {
			negSign = true
		}
	}

	resultAddForLoops := powByAdditionForLoops(copyA, b)
	resultAddRecursion := powByAdditionRecursion(copyA, b)
	resultLogarithms := powByLogarithms(copyA, b)

	// if b was odd and a was negative, negSign will be true
	if negSign {
		// invert sign of result
		resultAddForLoops = -resultAddForLoops
		resultAddRecursion = -resultAddRecursion
		resultLogarithms = -resultLogarithms
	}
	fmt.Printf("Power by addition over for loops: %d to the %d power = %d\n", a, b, resultAddForLoops)
	fmt.Printf("Power by addition via recursion: %d to the %d power = %d\n", a, b, resultAddRecursion)
	fmt.Printf("Power by logarithm: %d to the %d power = %.0f\n", a, b, resultLogarithms)

}

// powByAdditionForLoops() will take a value a and an exponent b and use addition to calculate a to the b power using
// nested for loops
func powByAdditionForLoops(a int, b int) int {
	// result will be updated throughout the for loop iterations and returned when the calculation is complete
	result := a

	for b-1 > 0 {
		interValue := result
		for i := 0; i < a-1; i++ {
			result += interValue
		}
		b--
	}

	return result
}

// powByAdditionRecursion() will take a value a and an exponent b and use addition to calculate a to the b power using
// recursion
func powByAdditionRecursion(a int, b int) int {
	// the recursive calls will terminate once the exponent b is zero
	// any value a, other than zero, to the zero power is one
	if b == 0 {
		return 1
	}

	result := 0
	interValue := powByAdditionRecursion(a, b-1)
	for i := 0; i < a; i++ {
		result += interValue
	}

	return result
}

// powByLogarithms() will take a value a and an exponent b and use the properties of logarithms to calculate a to the b power
func powByLogarithms(a int, b int) float64 {
	// fundamental logarithmic property in use:
	// ln(a^b) = b*ln(a)
	// e^ln(a^b) = e^(b*ln(a))
	// a^b = e^b*ln(a)
	// b*ln(a) is equivalent to adding ln(a), b times - meaning we can use addition to calculate b*ln(a)
	sum := 0.0
	for i := 0; i < b; i++ {
		sum += math.Log(float64(a))
	}
	return math.Exp(sum)
}
