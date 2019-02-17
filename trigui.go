package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

// PromptTrigValuesAndCompute seeks input for trigonometry functions, validates
// these inputs and eventually, computes the result and prints it.
func PromptTrigValuesAndCompute(function string, reader *bufio.Reader) {
	PrintTrigPromptHeader()

	// Seek valid input for x
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

	// Seek valid input for n
	fmt.Print("n = ")
	nStr := "nStr"
	for {
		nStr, _ = reader.ReadString('\n')
		nStr = strings.TrimSpace(nStr)
		if IsInt(nStr) {
			break
		}
		PrintRetryPrompt("n", "int")
	}

	x, _ := strconv.ParseFloat(xStr, 64)
	n, _ := strconv.Atoi(nStr)

	if !IsTrigInputValid(function, x) {
		return
	}
	v := DetermineTrigResult(function, x, n)
	PrintTrigResult(function, xStr, nStr, v)
}

// IsTrigInputValid verifies the value of x is -1 < x < 1 for sin and arcsin.
func IsTrigInputValid(function string, x float64) bool {
	if function == "arccos" || function == "arcsin" {
		if x > 1 || x < -1 {
			fmt.Printf("ERROR: Domain of %s is between -1 < x < 1 inclusive\n", function)
			return false
		}
	}
	return true
}

// DetermineTrigResult calls the appropriate function that maps to user request.
func DetermineTrigResult(function string, x float64, n int) float64 {
	switch function {
	case "sin":
		return Sine(x, n)
	case "arcsin":
		return InverseSine(x, n)
	case "cos":
		return Cosine(x, n)
	case "arccos":
		return InverseCosine(x, n)
	case "tan":
		return Tangent(x, n)
	default:
		return InverseTangent(x, n)
	}
}

// PrintTrigResult pretty prints the result.
func PrintTrigResult(function, xStr, nStr string, v float64) {
	fmt.Printf("%s(%s, %s) = %.5f\n", function, xStr, nStr, v)
	fmt.Println("===============================================================")
}

// PrintTrigPromptHeader prints a pretty prompt before requesting user input.
func PrintTrigPromptHeader() {
	fmt.Println("===============================================================")
	fmt.Println("| All trigonometry functions require 2 input values.          |")
	fmt.Println("| x: the value to compute in radians.                         |")
	fmt.Println("| n: the number of terms you would like to expand in the      |")
	fmt.Println("|    Taylor Series. A lower value for n yields a better       |")
	fmt.Println("|    performance, and vice-versa.                             |")
	fmt.Println("===============================================================")
}
