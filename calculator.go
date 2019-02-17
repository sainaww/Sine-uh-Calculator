package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	PrintWelcomeHeader()
	PrintHelp()
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Please enter the name of function you would like to use.")
		fmt.Print(">")
		input, _ := reader.ReadString('\n')
		fmt.Println(input)
		ParseAndExecute(input, reader)
	}
}

// ParseAndExecute routes user input to appropriate function handlers.
func ParseAndExecute(input string, reader *bufio.Reader) {
	input = strings.TrimSpace(strings.ToLower(input))
	switch input {
	case "tests", "test", "t":
		RunTests()
	case "bm", "benchmark":
		RunBenchmark()
	case "h", "help":
		PrintHelp()
	case "add", "+", "subtract", "-", "divide", "/", "multiply", "*", "permutation", "p", "combination", "c", "pow":
		PromptBasicArithmeticValuesAndCompute(input, reader)
	case "factorial", "!", "abs":
		PromptBasicArithmeticForSingleInput(input, reader)
	case "ln", "log", "exponent", "e", "sqrt":
		PromptComplexArithmetic(input, reader)
	case "sin", "arcsin", "cos", "arccos", "tan", "arctan":
		PromptTrigValuesAndCompute(input, reader)
	case "min", "max", "mean", "sd", "standard deviation", "mode", "median", "sum":
		PromptDefaultStatValuesAndCompute(input, reader)
	case "probability density function", "pdf":
		PromptPdfStatValuesAndCompute(input, reader)
	case "exit":
		os.Exit(3)
	}
}

// PrintWelcomeHeader prints a pretty welcome header.
func PrintWelcomeHeader() {
	fmt.Println(" __________    |")
	fmt.Println("| ________ |   |")
	fmt.Println("||12345678||   | WELCOME TO")
	fmt.Println("|\"\"\"\"\"\"\"\"\"\"|   | Sine-uh")
	fmt.Println("|[M|#|C][-]|   | CALCULATOR")
	fmt.Println("|[7|8|9][+]|   |")
	fmt.Println("|[1|2|3][%]|   |")
	fmt.Println("|[.|O|:][=]|   |")
	fmt.Println(" ----------    |")
}

// PrintHelp prints a pretty help message for the user.
func PrintHelp() {
	fmt.Println("===============================================================")
	fmt.Println("|                            Help                             |")
	fmt.Println("===============================================================")
	fmt.Println("| 1. Arithmetic Functions:                                    |")
	fmt.Println("|    * add (+)       * sqrt        * factorial (!)            |")
	fmt.Println("|    * subtract (-)  * abs         * permutation (P)          |")
	fmt.Println("|    * divide (/)    * pow         * combination (C)          |")
	fmt.Println("|    * multiply (*)  * ln          * exponent (e)             |")
	fmt.Println("|                    * log                                    |")
	fmt.Println("===============================================================")
	fmt.Println("| 2. Trigonometry Functions:                                  |")
	fmt.Println("|    * sin           * cos         * tan                      |")
	fmt.Println("|    * arcsin        * arccos      * arctan                   |")
	fmt.Println("===============================================================")
	fmt.Println("| 3. Statistical Functions:                                   |")
	fmt.Println("|    * min           * mode        * standard deviation (sd)  |")
	fmt.Println("|    * max           * median      * probability density      |")
	fmt.Println("|    * mean          * sum           function (pdf)           |")
	fmt.Println("===============================================================")
	fmt.Println("|    [help/h]    [tests/t]    [benchmark/bm]    [exit]        |")
	fmt.Println("===============================================================")
}
