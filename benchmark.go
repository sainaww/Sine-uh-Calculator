package main

import (
	"fmt"
	"math"
	"time"
)

// RunBenchmark runs benchmark code for arithmetic and trigonometry functions.
func RunBenchmark() {
	BenchmarkArithmeticFunctions()
	BenchMarkTrigonometryFunctions()
}

// BenchmarkArithmeticFunctions benchmarks arithmetic functions in calculator.
func BenchmarkArithmeticFunctions() {
	fmt.Println("===============================================================")
	fmt.Println("|   Math Benchmark                                            |")
	fmt.Println("===============================================================")

	x := 12
	y := 24

	// Compare add functions
	start := time.Now()
	Add(x, y)
	fmt.Printf("%d + %d                              took %s\n", x, y, time.Now().Sub(start))
	start = time.Now()
	BitwiseAdd(x, y)
	fmt.Printf("BitwiseAdd(%d, %d)                   took %s\n\n", x, y, time.Now().Sub(start))

	// Compare subtract functions
	start = time.Now()
	Subtract(x, y)
	fmt.Printf("%d - %d                              took %s\n", x, y, time.Now().Sub(start))
	start = time.Now()
	BitwiseSubtractFast(x, y)
	fmt.Printf("BitwiseSubtractFast(%d, %d)          took %s\n\n", x, y, time.Now().Sub(start))

	// Compare multiply functions
	start = time.Now()
	Multiply(x, y)
	fmt.Printf("%d * %d                              took %s\n", x, y, time.Now().Sub(start))
	start = time.Now()
	LongMultiplication(x, y)
	fmt.Printf("LongMultiplication(%d, %d)           took %s\n", x, y, time.Now().Sub(start))
	start = time.Now()
	KaratsubaMultiplicationFast(x, y)
	fmt.Printf("KaratsubaMultiplicationFast(%d, %d)  took %s\n\n", x, y, time.Now().Sub(start))

	// Compare division functions
	start = time.Now()
	Divide(y, x)
	fmt.Printf("%d / %d                              took %s\n", y, x, time.Now().Sub(start))
	start = time.Now()
	LongDivision(y, x)
	fmt.Printf("LongDivision(%d, %d)                 took %s\n\n", y, x, time.Now().Sub(start))
}

// BenchMarkTrigonometryFunctions computes and prints time taken for every
// trigonometry function implemented in the calculator.
func BenchMarkTrigonometryFunctions() {
	fmt.Println("===============================================================")
	fmt.Println("|   Trigonometry Benchmark                                    |")
	fmt.Println("===============================================================")

	x := 0.25
	accuracy := 9

	// Compare sin(x) functions
	start := time.Now()
	math.Sin(x)
	fmt.Printf("math.Sin(%.2f)      took %s\n", x, time.Now().Sub(start))
	start = time.Now()
	Sine(x, accuracy)
	fmt.Printf("Sine(%.2f)          took %s\n\n", x, time.Now().Sub(start))

	// Compare arcsin(x) functions
	start = time.Now()
	math.Asin(x)
	fmt.Printf("math.Asin(%.2f)     took %s\n", x, time.Now().Sub(start))
	start = time.Now()
	InverseSine(x, accuracy)
	fmt.Printf("InverseSine(%.2f)   took %s\n\n", x, time.Now().Sub(start))

	// Compare cos(x) functions
	start = time.Now()
	math.Cos(x)
	fmt.Printf("math.Cos(%.2f)      took %s\n", x, time.Now().Sub(start))
	start = time.Now()
	Cosine(x, accuracy)
	fmt.Printf("Cosine(%.2f)        took %s\n\n", x, time.Now().Sub(start))

	// Compare arccos(x) functions
	start = time.Now()
	math.Acos(x)
	fmt.Printf("math.arccos(%.2f)   took %s\n", x, time.Now().Sub(start))
	start = time.Now()
	InverseCosine(x, accuracy)
	fmt.Printf("InverseCosine(%.2f) took %s\n\n", x, time.Now().Sub(start))

	// Compare tan(x) functions
	start = time.Now()
	math.Tan(x)
	fmt.Printf("math.tan(%.2f)      took %s\n", x, time.Now().Sub(start))
	start = time.Now()
	Tangent(x, accuracy)
	fmt.Printf("Tangent(%.2f)       took %s\n\n", x, time.Now().Sub(start))

	// Compare arctan(x) functions
	start = time.Now()
	math.Atan(x)
	fmt.Printf("math.arctan(%.2f)   took %s\n", x, time.Now().Sub(start))
	start = time.Now()
	InverseTangent(x, accuracy)
	fmt.Printf("InverseTanget(%.2f) took %s\n\n", x, time.Now().Sub(start))
}
