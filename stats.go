package main

import (
	"math"
)

// Min returns the smallest element in a slice
func Min(data []float64) float64 {
	minSoFar := math.Inf(1)
	for _, value := range data {
		if value < minSoFar {
			minSoFar = value
		}
	}
	return minSoFar
}

// Max returns the largest element in a slice.
func Max(data []float64) float64 {
	maxSoFar := math.Inf(-1)
	for _, value := range data {
		if value > maxSoFar {
			maxSoFar = value
		}
	}
	return maxSoFar
}

// Mean returns the average value of all elements in a slice.
func Mean(data []float64) float64 {
	numberOfDataPoints := float64(len(data))
	sum := Sum(data)
	return sum / numberOfDataPoints
}

// Variance returns the expectation of the squared deviation of a random
// variable from its mean given a set of data values.
func Variance(data []float64) float64 {
	mean := Mean(data)
	diffFromMeanSquaredSum := 0.0
	for _, value := range data {
		diffFromMeanSquaredSum += ToThePowerFloat(value-mean, 2)
	}
	return diffFromMeanSquaredSum / float64(len(data))
}

// StandardDeviation returns the amount of variation or dispersion of a set of
// data values.
func StandardDeviation(data []float64) float64 {
	variance := Variance(data)
	return HeronsSquareRoot(variance, 0.00001)
}

// QuickSort sorts a set of data values in place. Theory behind QuickSort can
// be found here: https://en.wikipedia.org/wiki/Quicksort
func QuickSort(data []float64) {
	QuickSortHelper(data, 0, len(data)-1)
}

// QuickSortHelper is a helper function for QuickSort.
func QuickSortHelper(data []float64, low, high int) {
	if low < high {
		partitionIndex := QuickSortPartition(data, low, high)
		QuickSortHelper(data, low, partitionIndex-1)
		QuickSortHelper(data, partitionIndex+1, high)
	}
}

// QuickSortPartition takes last elements as pivot and puts it in its correct
// position in a sorted array. Then, it proceeds to put all value less than or
// equal it to its left and all values greater than it on its right.
func QuickSortPartition(data []float64, low, high int) int {
	i := low - 1
	pivot := data[high]

	for j := low; j < high; j++ {
		if data[j] <= pivot {
			i++
			temp := data[i]
			data[i] = data[j]
			data[j] = temp
		}
	}

	temp := data[i+1]
	data[i+1] = data[high]
	data[high] = temp
	return i + 1
}

// Median returns the value separating the higher half from the lower half of a
// data sample.
func Median(data []float64) float64 {
	n := len(data)
	QuickSort(data)
	if n%2 == 0 {
		return (data[n/2] + data[(n/2)-1]) / 2.0
	}
	return data[n/2]
}

// Sum returns the sum of all elements in a set of data values.
func Sum(data []float64) float64 {
	sum := 0.0
	for _, value := range data {
		sum += value
	}
	return sum
}

// Mode returns the value that appears most often in a set of data values.
func Mode(data []float64) float64 {
	frequencyMap := make(map[float64]int)
	for _, dataPoint := range data {
		frequencyMap[dataPoint]++
	}

	mode := data[0]
	for dataPoint, frequency := range frequencyMap {
		if frequency > frequencyMap[mode] {
			mode = dataPoint
		}
	}
	return mode
}

// NormalDistributionPdf computes and returns the probability density function
// of a data set at x i.e. pdf(x) in a normal distribution.
// Y = { 1/[ σ * sqrt(2π) ] } * e^-((x - μ)^2)/(2*σ^2)
func NormalDistributionPdf(data []float64, x float64) float64 {
	return NormalDistributionPdfHelper(Mean(data), StandardDeviation(data), x)
}

// NormalDistributionPdfHelper is a helper function for NormalDistributionPdf
// which computes pdf(x) given mean, sd and x.
func NormalDistributionPdfHelper(mean, standardDeviation, x float64) float64 {
	n := 1.0 / HeronsSquareRoot(standardDeviation*standardDeviation*2*math.Pi, 0.00001)
	ex := -1.0 * (((x - mean) * (x - mean)) / (2 * standardDeviation * standardDeviation))
	y := math.Pow(math.E, ex)
	return n * y
}
