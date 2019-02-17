package main

import (
	"math"
)

// Addition function takes in a variable number of inputs and adds
// them using bit manipulation.
func Addition(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum = BitwiseAdd(sum, number)
	}
	return sum
}

// BitwiseAdd is the subroutine used in Addition. It adds two ints using bit
// manipulation.
func BitwiseAdd(x, y int) int {
	for y != 0 {
		carry := x & y
		x = x ^ y
		y = carry << 1
	}
	return x
}

// Subtraction runs a series of BitwiseSubtractFast on a variable number of
// ints provided as input.
func Subtraction(numbers ...int) int {
	minus := numbers[0]
	for _, number := range numbers[1:] {
		minus = BitwiseSubtract(minus, number)
	}
	return minus
}

// BitwiseSubtract is slower
func BitwiseSubtract(x, y int) int {
	for y != 0 {
		carry := y &^ x
		x = x ^ y
		y = carry << 1
	}
	return x
}

// BitwiseSubtractFast is faster
func BitwiseSubtractFast(x, y int) int {
	y = y * -1
	for y != 0 {
		carry := x & y
		x = x ^ y
		y = carry << 1
	}
	return x
}

// SlowMultiplication is a naive implementation of multiplication through
// addition.
func SlowMultiplication(x, y int) int {
	result := 0
	for i := 0; i < y; i++ {
		result += x
	}
	return result
}

// LongMultiplication multiplies 2 ints using long multiplication technique
// that is typically used by people to multiply 2 numbers by hand.
func LongMultiplication(x, y int) int {
	return LongMultiplicationHelper(MakeIntArrayFromInt(x), MakeIntArrayFromInt(y))
}

// LongMultiplicationHelper is a helper function for LongMultiplication that
// accepts numbers in the form of int arrays. e.g. 24 * 35 becomes [4,2] * [5,3]
// IntArray[0] = as dummy value, IntArray[1] = unit position (10^1), IntArray[2] = tenth position (10^2)
// IntArray[3] = hundredth position (10^3) ...
func LongMultiplicationHelper(x, y []int) int {
	numOfDigitsInX := len(x) - 1
	numOfDigitsInY := len(y) - 1
	product := make([]int, numOfDigitsInX+numOfDigitsInY+1)

	for yi := 1; yi < len(y); yi++ {
		carry := 0
		for xi := 1; xi < len(x); xi++ {
			product[xi+yi-1] += carry + x[xi]*y[yi]
			carry = product[xi+yi-1] / 10
			product[xi+yi-1] = product[xi+yi-1] % 10
		}
		product[yi+numOfDigitsInX] += carry
	}
	return MakeIntFromIntArray(product)
}

// KaratsubaMultiplication is an implementation of multiplication for large
// numbers. More on the algorithm can be found here:
// https://en.wikipedia.org/wiki/Karatsuba_algorithm
func KaratsubaMultiplication(x, y int) int {
	//karatsuba
	if x < 10 || y < 10 {
		return x * y
	}
	maxNumberOfDigits := MaxBetween(CountNumDigits(x), CountNumDigits(y))
	m := maxNumberOfDigits / 2

	tenToThePowerM := ToThePowerInt(10, m)

	high1 := x / tenToThePowerM
	low1 := x % tenToThePowerM
	high2 := y / tenToThePowerM
	low2 := y % tenToThePowerM

	a := KaratsubaMultiplication(high1, high2)
	c := KaratsubaMultiplication(low1, low2)
	d := KaratsubaMultiplication((low1 + high1), (low2 + high2))

	return (a * ToThePowerInt(10, 2*m)) + ((d - a - c) * ToThePowerInt(10, m)) + c
}

// KaratsubaMultiplicationFast uses CountNumDigitsFast instead of
// CountNumDigits. This allows the algorithm to perform better.
func KaratsubaMultiplicationFast(x, y int) int {
	if x < 10 || y < 10 {
		return x * y
	}

	m := MaxBetween(CountNumDigitsFast(x), CountNumDigitsFast(y)) / 2

	tenToThePowerM := ToThePowerInt(10, m)

	a := x / tenToThePowerM
	b := x % tenToThePowerM
	c := y / tenToThePowerM
	d := y % tenToThePowerM

	ac := KaratsubaMultiplication(a, c)
	bd := KaratsubaMultiplication(b, d)
	adplusbc := KaratsubaMultiplication((a+b), (c+d)) - ac - bd

	return (ac * ToThePowerInt(10, 2*m)) + (adplusbc * ToThePowerInt(10, m)) + bd
}

// SlowDivision is a naive implementation of division through subtraction.
func SlowDivision(dividend, divisor int) int {
	quotient := 0
	for dividend >= divisor {
		dividend -= divisor
		quotient++
	}
	return quotient
}

// LongDivision implements the common long division technique of dividing 2
// numbers that is typically done by hand.
func LongDivision(dividend, divisor int) int {
	sign := -1
	if (dividend > 0 && divisor > 0) || (dividend < 0 && divisor < 0) {
		sign = 1
	}

	return sign * LongDivisionHelper(uint(Abs(dividend)), uint(Abs(divisor)))
}

// LongDivisionHelper is a helper function for LongDivision
func LongDivisionHelper(dividend, divisor uint) int {

	denominator := divisor
	current := 1
	answer := 0

	if denominator > dividend {
		return 0
	}

	if denominator == dividend {
		return 1
	}

	for denominator <= dividend {
		denominator <<= 1
		current <<= 1
	}

	denominator >>= 1
	current >>= 1

	for current != 0 {
		if dividend >= denominator {
			dividend -= denominator
			answer |= current
		}
		current >>= 1
		denominator >>= 1
	}
	return answer
}

// Permutation computes and returns nPk.
func Permutation(n, k int) int {
	if k > n {
		panic("k must be less than n")
	}
	product := 1
	for i := n; i > n-k; i-- {
		product = product * i
	}
	return product
}

// Combination computes and returns nCk.
func Combination(n, k int) int {
	if k > n {
		panic("k must be less than n")
	}

	var limit int
	var factorial int

	if k > n-k {
		limit = k
		factorial = n - k
	} else {
		limit = n - k
		factorial = k
	}

	// This computation is done in float to prevent overflow of integers
	product := float64(1) / Factorial(float64(factorial))

	for i := n; i > limit; i-- {
		product = product * float64(i)
	}
	return int(math.Ceil(product))
}

// Factorial computes and returns n!
func Factorial(n float64) float64 {
	var absFactorial float64
	if n == 1 || n == 0 {
		return 1
	} else if n > 1 {
		return n * Factorial(n-1)
	} else { // n < 1
		n = n * -1.0
		absFactorial = n * Factorial(n-1)
		return -1.0 * absFactorial
	}
}

// Pi approximates the value of the constant Pi through Gregory Leibniz series
// up to the nth term.
func Pi(n int) float64 {
	piValue := 1.0
	nextGregoryLeibnizSeriesValue := -3.0
	for i := 0; i < n; i++ {
		piValue += 1 / nextGregoryLeibnizSeriesValue
		if nextGregoryLeibnizSeriesValue < 0 {
			nextGregoryLeibnizSeriesValue = (nextGregoryLeibnizSeriesValue * -1) + 2
		} else {
			nextGregoryLeibnizSeriesValue = (nextGregoryLeibnizSeriesValue * -1) - 2
		}
	}
	return piValue * 4
}

// Abs returns the absolute version of an int, x.
func Abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

// AbsFloat returns the absolute version of a float, x.
func AbsFloat(x float64) float64 {
	if x < 0 {
		return x * -1
	}
	return x
}

// ToThePowerInt computes and returns base^power.
func ToThePowerInt(base, power int) int {
	if power == 0 {
		return 1
	}

	temp := ToThePowerInt(base, int(power/2))

	if power%2 == 0 {
		return temp * temp
	} else if power > 0 {
		return base * temp * temp
	} else {
		return (temp * temp) / base
	}
}

// ToThePowerFloat is a naive implementation of raising power through
// multiplication to compute and return base^power.
func ToThePowerFloat(base, power float64) float64 {
	product := 1.0
	for i := 0.0; i < power; i++ {
		product = product * base
	}
	return product
}

// NewtonianSquareRoot computes the square root of a number, x through binary
// search within the margin of error n.
func NewtonianSquareRoot(x float64, n float64) float64 {
	m := x / 2
	for AbsFloat(m*m-x) > n {
		if m*m > x {
			m /= 2
		} else {
			m += m / 2
		}
	}
	return m
}

// HeronsSquareRoot computes the square root of a number through convergence
// using the Heron or Babsylonian method. It computes square root of x within the
// margin of error, n.
func HeronsSquareRoot(x float64, n float64) float64 {
	guess := x / 2.0
	for AbsFloat(guess*guess-x) > n {
		guess = 0.5 * (guess + x/guess)
	}
	return guess
}

// Exponent computes e^x using Taylor Series expanded to the n-th term.
func Exponent(x float64, n int) float64 {
	exponentValue := 1.0 + x
	for i := 0; i < n; i++ {
		exponentValue += ToThePowerFloat(x, float64(i+2)) / Factorial(float64(i+2))
	}
	return exponentValue
}

// NaturalLog computes ln(x) using a Taylor Series expanded to the a-th term.
func NaturalLog(x float64, a int) float64 {
	halfOfNaturalLogValue := (x - 1) / (x + 1)
	n := 3.0
	for i := 0; i < a; i++ {
		halfOfNaturalLogValue += (1 / n) * ToThePowerFloat(halfOfNaturalLogValue, n)
		n += 2
	}
	return 2 * halfOfNaturalLogValue
}

// LogBaseTen computes the value of log(x) by conversion of ln(x). More on this
// method can be found here:
// http://mathonweb.com/help_ebook/html/algorithms.htm#log
func LogBaseTen(x float64, a int) float64 {
	return 0.43429448 * NaturalLog(x, 100)
}
