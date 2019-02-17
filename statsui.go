package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

// PromptDefaultStatValuesAndCompute seeks input for standard stats functions,
// validates them and the proceeds to compute results.
func PromptDefaultStatValuesAndCompute(function string, reader *bufio.Reader) {
	PrintStatDefaultPromptHeader()

	// seek input for data
	fmt.Print("data = ")
	dataStr := "data"
	for {
		dataStr, _ = reader.ReadString('\n')
		dataStr = strings.TrimSpace(dataStr)
		if IsFloatArray(ParseInputToArray(dataStr)) {
			break
		}
		PrintRetryPrompt("data", "array of floats")
	}

	data := ParseStringArrayToFloatArray(ParseInputToArray(dataStr))
	v := DetermineStatsResult(function, data)
	PrintStatsResult(function, dataStr, v)
}

// PrintStatsResult prints the result in a pretty manner.
func PrintStatsResult(function, dataStr string, v float64) {
	fmt.Printf("%s([%s]) = %.5f\n", function, dataStr, v)
	fmt.Println("===============================================================")
}

// DetermineStatsResult finds the appropriate function corresponding to the
// user request and computes the result.
func DetermineStatsResult(function string, data []float64) float64 {
	switch function {
	case "min":
		return Min(data)
	case "max":
		return Max(data)
	case "mean":
		return Mean(data)
	case "sd", "standard deviation":
		return StandardDeviation(data)
	case "median":
		return Median(data)
	case "sum":
		return Sum(data)
	default:
		return Mode(data)
	}
}

// PrintStatDefaultPromptHeader prints a pretty prompt to let users know what
// it expects as input.
func PrintStatDefaultPromptHeader() {
	fmt.Println("===============================================================")
	fmt.Println("| This statistic function requires only 1 input value.        |")
	fmt.Println("| data: the set of values to analyze. e.g. 1,2,3,4,5          |")
	fmt.Println("|       all values must be floats and comma separated.        |")
	fmt.Println("===============================================================")
}

// PromptPdfStatValuesAndCompute seeks input for pdf function, validates inputs
// and proceeds to compute and print the pdf.
func PromptPdfStatValuesAndCompute(function string, reader *bufio.Reader) {
	PrintStatPromptForPdf()

	// Seek input for data
	fmt.Print("data = ")
	dataStr := "data"
	for {
		dataStr, _ = reader.ReadString('\n')
		dataStr = strings.TrimSpace(dataStr)
		if IsFloatArray(ParseInputToArray(dataStr)) {
			break
		}
		PrintRetryPrompt("data", "array of floats")
	}

	// Seek input for x
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

	data := ParseStringArrayToFloatArray(ParseInputToArray(dataStr))
	x, _ := strconv.ParseFloat(xStr, 64)
	v := NormalDistributionPdf(data, x)
	PrintPdfResult(function, dataStr, xStr, v)
}

// PrintPdfResult prints the result of pdf function in a pretty manner.
func PrintPdfResult(function, dataStr, xStr string, v float64) {
	fmt.Printf("%s([%s], %s) = %.5f\n", function, dataStr, xStr, v)
	fmt.Println("===============================================================")
}

// PrintStatPromptForPdf prints a nice prompt to explain what inputs it expects.
func PrintStatPromptForPdf() {
	fmt.Println("===============================================================")
	fmt.Println("| The probability distribution function (pdf) requires 2      |")
	fmt.Println("| input values.                                               |")
	fmt.Println("| data: the set of values to analyze. e.g. 1,2,3,4,5          |")
	fmt.Println("|       all values must be floats and comma separated.        |")
	fmt.Println("| x   : x is the point  at which the pdf is evaluated.        |")
	fmt.Println("===============================================================")
}
