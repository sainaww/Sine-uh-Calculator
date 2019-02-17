package main

import (
	"fmt"
	"math"
)

// RunTests run tests for all three key components of calculator.
func RunTests() {
	TestArithmeticFunctions()
	TestTrigonometryFunctions()
	TestStatsFunctions()
}

// TestArithmeticFunctions runs tests on all Arithmetic function
func TestArithmeticFunctions() {
	fmt.Println("===============================================================")
	fmt.Println("| Running Arithmetic Tests ...                                |")

	AssertOrPanicInt(Addition(24, 89, 34), 147)
	AssertOrPanicInt(BitwiseAdd(24, 89), 113)
	AssertOrPanicInt(Subtraction(89, 34, 21), 34)
	AssertOrPanicInt(BitwiseSubtract(89, 24), 65)
	AssertOrPanicInt(BitwiseSubtractFast(89, 24), 65)
	AssertOrPanicInt(LongMultiplication(24, 89), 2136)
	AssertOrPanicInt(KaratsubaMultiplicationFast(24, 89), 2136)
	AssertOrPanicInt(LongDivision(89, 24), 3)
	AssertOrPanicInt(Permutation(8, 4), 1680)
	AssertOrPanicInt(Combination(8, 4), 70)
	AssertOrPanic(Factorial(9), 362880)
	AssertOrPanic(Pi(500000), math.Pi)
	AssertOrPanicInt(Abs(-1), 1)
	AssertOrPanic(AbsFloat(-1.0), 1.0)
	AssertOrPanicInt(ToThePowerInt(2, 5), 32)
	AssertOrPanic(ToThePowerFloat(2, 5), 32)
	AssertOrPanic(NewtonianSquareRoot(9, 0.00001), 3)
	AssertOrPanic(HeronsSquareRoot(9, 0.00001), 3)
	AssertOrPanic(Exponent(1, 15), math.E)
	AssertLogIsClose(NaturalLog(2.5, 15), 0.91629073187)
	AssertLogIsClose(LogBaseTen(2.5, 15), 0.39794000867)

	PrintAllTestsOk()
}

// TestTrigonometryFunctions compares and ensures all output of Trigonometry
// functions implemented in the package are within reasonable margin of error to
// the math package implementation.
func TestTrigonometryFunctions() {
	fmt.Println("===============================================================")
	fmt.Println("| Running Trigonometry Tests ...                              |")

	x := 0.25
	accuracy := 9

	AssertOrPanic(math.Sin(x), Sine(x, accuracy))
	AssertOrPanic(math.Asin(x), InverseSine(x, accuracy))
	AssertOrPanic(math.Tan(x), Tangent(x, accuracy))
	AssertOrPanic(math.Atan(x), InverseTangent(x, accuracy))
	AssertOrPanic(math.Cos(x), Cosine(x, accuracy))
	AssertOrPanic(math.Acos(x), InverseCosine(x, accuracy))

	PrintAllTestsOk()
}

// TestStatsFunctions tests all statistical functions implemented in calculator.
func TestStatsFunctions() {
	fmt.Println("===============================================================")
	fmt.Println("| Running Stat Tests ...                                      |")

	data := []float64{1, 4, 3, 5, 2, 6, 4}

	AssertOrPanic(Min(data), 1)
	AssertOrPanic(Max(data), 6)
	AssertOrPanic(Mean(data), 3.571428571)
	AssertOrPanic(StandardDeviation(data), 1.5907898179514)
	AssertOrPanic(Median(data), 4)
	AssertOrPanic(Sum(data), 25)
	AssertOrPanic(Mode(data), 4)
	AssertOrPanic(NormalDistributionPdf(data, 2.5), 0.19989228)

	// Shuffling the data before testing sort. This is because sort has already
	// been called previously by the Median function.
	data = []float64{1, 4, 3, 5, 2, 6, 4}
	QuickSort(data)
	AssertSortWorked(data)

	PrintAllTestsOk()
}

// AssertLogIsClose ensures that the values for ln() and log() are within a
// reasonable margin of error.
func AssertLogIsClose(x, y float64) {
	if math.Abs(x-y) > 0.01 {
		panic("Function did not match expected output.")
	}
}

// AssertOrPanic ensures two floats are within a reasonable margin compared to
// each other.
func AssertOrPanic(x, y float64) {
	if math.Abs(x-y) > 0.00001 {
		panic("Function did not match expected output.")
	}
}

// AssertOrPanicInt ensures two ints are within a reasonable margin compared
// to each other.
func AssertOrPanicInt(x, y int) {
	if x != y {
		panic("Function did not match expected output.")
	}
}

// AssertSortWorked verifies that the data provided is sorted in ascending order
func AssertSortWorked(data []float64) {
	for i := 1; i < len(data); i++ {
		if data[i] < data[i-1] {
			panic("Data is not sorted!")
		}
	}
}

// PrintAllTestsOk prints that all tests are ok at the end.
func PrintAllTestsOk() {
	fmt.Println("| All tests Ok.                                               |")
	fmt.Println("===============================================================")
}
