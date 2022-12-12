package main

import (
	"fmt"
	"math"
)

func main() {
	// Goal: to add two integers without using the addition operator (+)
	// the numbers can be positive or negative
	// the increment (++) operator is allowed

	// num1 and num2 are test numbers
	num1, num2 := 20, 100

	// Method 1: addition by subtracting the opposite of the second number
	fmt.Printf("Addition via subtracting the opposite of the second number: %d\n", addBySub(num1, num2))

	// Method 2: addition via a loop to increment num1 by one exactly num2 times if num2 is positive, otherwise decrement num1
	// by one exactly num2 times
	fmt.Printf("Addition via a loop using increment and decrement operators: %d\n", addByLoop(num1, num2))

	// Method 3: addition using half-adder logic - using a "combination circuit" to add two numbers a bit at a time
	fmt.Printf("Addition via half-adder logic: %d\n", addByHalfAdder(num1, num2))

	// THIS FAILS AT LARGE VALUES FOR NUM1 AND NUM2 (POSITIVE OR NEGATIVE) DUE TO THE SIZE OF THE RESULT
	// Method 4: addition using the logarithm operation and exponentiation
	fmt.Printf("Addition using the logarithm operation and exponentiation: %d", addByLogExp(num1, num2))

	// ONLY WORKS IF BOTH NUMBERS ARE POSITIVE - included for its limited application
	// Method 5: addition using fmt.Printf() and format specifiers - fmt.Printf() returns the number of bytes written and
	// using format specifiers, we can use num1 and num2 as dynamic width values to print a one byte character num1+num2 times
	if num1 >= 0 && num2 >= 0 {
		sum, err := addByPrintf(num1, num2)
		if err != nil {
			fmt.Printf("Error while adding via fmt.Printf(): %s", err.Error())
			return
		}
		fmt.Printf("\nAddition via fmt.Printf() and format specifiers: %d\n", sum)
		return
	}
	fmt.Println("\nBoth numbers must be positive for method 5 to work. Skipping as at least one number is negative.")
}

// addBySub() adds two integers by subtracting the opposite of num2 - i.e. num1 - (-num2) = num1 + num2
func addBySub(num1 int, num2 int) int {
	return num1 - (-num2)
}

// addByLoop() adds two integers by incrementing num1 by one exactly num2 times if num2 is positive, otherwise decrement num1
// by one exactly num2 times
func addByLoop(num1 int, num2 int) int {
	// if num2 < 0, it is negative
	if num2 < 0 {
		// loop until num2 == 0
		for num2 != 0 {
			// decrement num1 by one
			num1--
			// increment num2 by one until it equals zero
			num2++
		}
		return num1
	}
	// otherwise num2 is positive
	for num2 != 0 {
		// increment num1 by one
		num1++
		// decrement num2 until it equals zero
		num2--
	}
	// returns num1 if incremented by the above for loop or unchanged if num2 is already zero
	return num1
}

// addByHalfAdder() adds using half-adder logic, a "combination circuit" which adds two integers a bit at a time
func addByHalfAdder(num1 int, num2 int) int {
	// in the event num2 is zero, we can simply return num1 (num1 + 0 = num1); however, using recursion, the carry value
	// will be passed in as num2 and, when zero, the addition operation is done signifying the end of the recursive calls
	if num2 == 0 {
		return num1
	}
	// the intermediate sum is equal to the bitwise X-OR of num1 and num2
	sum := num1 ^ num2
	// the carry value is equal to the bitwise AND of num1 and num2
	// the result is bitwise-shifted to the left one position
	carry := num1 & num2 << 1
	return addByHalfAdder(sum, carry)
}

// addByLogExp() adds using a property of exponentiation and its opposite opperation logarithms
func addByLogExp(num1 int, num2 int) int {
	// when a base b to the nth power is multiplied by base b to the mth power, the result is equal to the base b to the
	// n + m power (i.e. b^n * b^m = b^(n+m))
	// as logarithms are the opposite operation of exponentiation, taking the logarithm with base b of b^(n+m) yeilds n+m
	// both math.Log and math.Exp use a base of e (Euler's Number)
	return int(math.Log(math.Exp(float64(num1)) * math.Exp(float64(num2))))
}

// addByPrintf() takes advantage of format specifiers and fmt.Printf() returning the number of bytes written to print a one
// byte character num1 + num2 times - this only works if both numbers are positive
func addByPrintf(num1 int, num2 int) (int, error) {
	// the * in %*s prompts fmt.Printf() to look for a width value among the arguments passed into it
	// the width, if positive, pads the variable with spaces on the left ("right justified") and, if negative, pads the
	// variable with spaces on the right ("left justified")
	// using the empty string as the variable to pad, the only bytes written are the spaces used in the padding
	// here we use the empty string twice, once padded with num1 spaces and once with num2 spaces (num1 + num2 padding spaces)
	return fmt.Printf("%*s%*s", num1, "", num2, "")
}
