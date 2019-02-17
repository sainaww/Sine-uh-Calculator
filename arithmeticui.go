package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

// PromptBasicArithmeticValuesAndCompute seeks input for trigonometry functions.
func PromptBasicArithmeticValuesAndCompute(function string, reader *bufio.Reader) {
	firstPromptString, secondPromptString := GetBasicArithmeticPromptString(function)
	fmt.Printf("%s = ", firstPromptString)

	// Seek first input
	xStr := "xStr"
	for {
		xStr, _ = reader.ReadString('\n')
		xStr = strings.TrimSpace(xStr)
		if IsInt(xStr) {
			break
		}
		PrintRetryPrompt(firstPromptString, "int")
	}

	// Seek second input
	yStr := "yStr"
	for {
		yStr, _ = reader.ReadString('\n')
		yStr = strings.TrimSpace(yStr)
		if IsInt(yStr) {
			break
		}
		PrintRetryPrompt(secondPromptString, "int")
	}

	x, _ := strconv.Atoi(xStr)
	y, _ := strconv.Atoi(yStr)
	v := DetermineBasicArithmeticResult(function, x, y)
	PrintBasicArithmeticResult(function, x, y, v)
}

func PrintBasicArithmeticResult(function string, x, y, v int) {
	fmt.Printf("%s(%d, %d) = %d\n", function, x, y, v)
	fmt.Println("===============================================================")
}

// DetermineBasicArithmeticResult finds the appropriate function handler for
// user prompt and computes the result accordingly.
func DetermineBasicArithmeticResult(function string, x, y int) int {
	switch function {
	case "add", "+":
		return BitwiseAdd(x, y)
	case "subtract", "-":
		return BitwiseSubtractFast(x, y)
	case "divide", "/":
		return LongDivision(x, y)
	case "multiply", "*":
		return KaratsubaMultiplicationFast(x, y)
	case "permutation", "p":
		return Permutation(x, y)
	case "combination", "cs":
		return Combination(x, y)
	default:
		return ToThePowerInt(x, y)
	}
}

// GetBasicArithmeticPromptString prints the appropriate prompt depending on the
// user input.
func GetBasicArithmeticPromptString(function string) (string, string) {
	switch function {
	case "permutation", "p", "combination", "c":
		return "n", "r"
	default:
		return "x", "y"
	}
}

// PromptBasicArithmeticForSingleInput seeks input for arithmetic functions that
// accept only one input.
func PromptBasicArithmeticForSingleInput(function string, reader *bufio.Reader) {
	// Seek input
	fmt.Print("x = ")
	xStr := "data"
	for {
		xStr, _ = reader.ReadString('\n')
		xStr = strings.TrimSpace(xStr)
		if IsInt(xStr) {
			break
		}
		PrintRetryPrompt("x", "int")
	}

	x, _ := strconv.Atoi(xStr)
	v := DetermineBasicArithmeticResultForSingleInput(function, x)
	PrintBasicArithmeticResultForSingleInput(function, x, v)
}

// PrintBasicArithmeticResultForSingleInput pretty prints result.
func PrintBasicArithmeticResultForSingleInput(function string, x, v int) {
	fmt.Printf("%s(%d) = %d\n", function, x, v)
	fmt.Println("===============================================================")
}

// DetermineBasicArithmeticResultForSingleInput finds the appropriate function
// in calculator to compute result.
func DetermineBasicArithmeticResultForSingleInput(function string, x int) int {
	switch function {
	case "factorial", "!":
		return int(Factorial(float64(x)))
	default:
		return Abs(x)
	}
}

// PromptComplexArithmetic seeks input for complex arithmetic functions,
// validates the input and then computes the result.
func PromptComplexArithmetic(function string, reader *bufio.Reader) {
	PrintComplexArithmeticPromptHeader()

	// Seek first input
	fmt.Print("x = ")
	xStr := "xStr"
	for {
		xStr, _ = reader.ReadString('\n')
		xStr = strings.TrimSpace(xStr)
		if IsFloat(xStr) {
			break
		}
		PrintRetryPrompt("x", "float")
	}

	// Seek second input
	fmt.Print("n = ")
	nStr := "nStr"
	for {
		nStr, _ = reader.ReadString('\n')
		nStr = strings.TrimSpace(nStr)
		if function == "sqrt" {
			if IsFloat(nStr) {
				break
			}
			PrintRetryPrompt("n", "float")
		} else {
			if IsInt(nStr) {
				break
			}
			PrintRetryPrompt("n", "int")
		}
	}

	x, _ := strconv.ParseFloat(xStr, 64)
	n, _ := strconv.Atoi(nStr)
	v := DetermineComplexArithmeticResult(function, nStr, x, n)
	PrintComplexArithmeticResult(function, xStr, nStr, v)
}

// PrintComplexArithmeticResult pretty prints complex arithmetic results.
func PrintComplexArithmeticResult(function, xStr, nStr string, v float64) {
	fmt.Printf("%s(%s, %s) = %.5f\n", function, xStr, nStr, v)
	fmt.Println("===============================================================")
}

// DetermineComplexArithmeticResult finds the appropriate complex arithmetic
// function corresponding to user input and computes the result.
func DetermineComplexArithmeticResult(function, nStr string, x float64, n int) float64 {
	switch function {
	case "ln":
		return NaturalLog(x, n)
	case "log":
		return LogBaseTen(x, n)
	case "exponent", "e":
		return Exponent(x, n)
	default:
		a, _ := strconv.ParseFloat(nStr, 64)
		return HeronsSquareRoot(x, a)
	}
}

// PrintComplexArithmeticPromptHeader prints a pretty prompt at the start of
// complex functions explaining what they require as input.
func PrintComplexArithmeticPromptHeader() {
	fmt.Println("===============================================================")
	fmt.Println("| This functions require 2 input values.                      |")
	fmt.Println("| x : the value to compute.                                   |")
	fmt.Println("| n : The number of terms to expand on Taylor Series to compu-|")
	fmt.Println("|     te the value. In case of sqrt, the n stands for margin  |")
	fmt.Println("|     of error in float within which to calculate the sqrt.   |")
	fmt.Println("===============================================================")
}
