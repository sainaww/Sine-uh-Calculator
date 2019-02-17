package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

/**
This file contains all the common functions shared across components in this
package.
*/

// MakeIntArrayFromInt takes an int and returns an array of integer represen-
// tation of the int where most significant bit is to the farthest right
// position. 0th index always contains 0 as a dummy value.
// i.e., 23 -> {0, 3, 2}
func MakeIntArrayFromInt(x int) []int {
	intArr := []int{0}
	for x != 0 {
		intArr = append(intArr, x%10)
		x /= 10
	}
	return intArr
}

// MakeIntFromIntArray takes an array of integers and returns an int. The array
// representation of the int has the most significant bit in its right most
// index position. 0th index always contains 0 as a dummy value.
// i.e., {0, 3, 2} -> 23
func MakeIntFromIntArray(x []int) int {
	number := 0
	for i := len(x) - 1; i > 0; i-- {
		number += x[i] * ToThePowerInt(10, i-1)
	}
	return number
}

// CountNumDigits is a relatively slow implementation to count the number of
// digits in an integer.
func CountNumDigits(x int) (count int) {
	if x == 0 {
		return 1
	}
	for x != 0 {
		x = x / 10
		count++
	}
	return count
}

// CountNumDigitsFast is a faster implementation to count the number of digits
// in an integer.
func CountNumDigitsFast(x int) int {
	stringOfX := strconv.Itoa(x)
	return len(stringOfX)
}

// MaxBetween returns takes in 2 integers as input and returns the larger value.
func MaxBetween(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// PrintRetryPrompt is used in UI code to prompt user to renter a value for a
// function input.
func PrintRetryPrompt(s string, t string) {
	fmt.Printf("Expected %s. Pleasy try again!\n", t)
	fmt.Printf("%s = ", s)
}

// IsFloat checks if the string, x can be represented as a float.
func IsFloat(x string) bool {
	_, err := strconv.ParseFloat(strings.TrimSpace(x), 64)
	return err == nil
}

// IsInt checks if the string, x can be represented as an int.
func IsInt(x string) bool {
	_, err := strconv.Atoi(strings.TrimSpace(x))
	return err == nil
}

// IsFloatArray checks if the string array can be represnted as an array of
// floats.
func IsFloatArray(data []string) bool {
	for _, v := range data {
		if !IsFloat(v) {
			return false
		}
	}
	return true
}

// ParseStringArrayToFloatArray converts an array of strings to an array of
// floats.
func ParseStringArrayToFloatArray(data []string) []float64 {
	floatArr := make([]float64, 0)
	for _, v := range data {
		x, _ := strconv.ParseFloat(strings.TrimSpace(v), 64)
		floatArr = append(floatArr, x)
	}
	return floatArr
}

// ParseInputToArray turns a string representation of an array into an array of
// strings. e.g. "1,2,3" -> {"1", "2", "3"}
func ParseInputToArray(input string) []string {
	return strings.Split(strings.TrimSpace(input), ",")
}

/**
The functions below were written to benchmark against my implementations.
*/

//GoPi is go's implementation of pi.
func GoPi(n int) float64 {
	ch := make(chan float64)
	for k := 0; k <= n; k++ {
		go Term(ch, float64(k))
	}
	f := 0.0
	for k := 0; k <= n; k++ {
		f += <-ch
	}
	return f
}

// Term is a helper function for GoPi
func Term(ch chan float64, k float64) {
	ch <- 4 * math.Pow(-1, k) / (2*k + 1)
}

// Add returns x + y
func Add(x, y int) int {
	return x + y
}

// Subtract returns x - y
func Subtract(x, y int) int {
	return x - y
}

// Multiply returns x * y
func Multiply(x, y int) int {
	return x * y
}

// Divide returns x / y
func Divide(x, y int) int {
	return x / y
}
