package main

import (
	"math"
)

// Sine returns the value of sin(x) where x is in radians. x can be negative or
// positive. sin(x) is calculated by expressing the sin value in a Taylor Series
// for all real numbers x (where x is the angle in radians). The series will be
// expanded to n terms to compute the value of sin(x). More on the series can
// found here: https://en.wikipedia.org/wiki/Sine#Series_definition
func Sine(x float64, n int) float64 {
	sineValue := x
	power := 3.0
	denominator := -3.0

	for i := 0; i < n; i++ {
		sineValue += ToThePowerFloat(x, power) / Factorial(denominator)
		if denominator < 0 {
			denominator = (denominator * -1) + 2
		} else {
			denominator = (denominator * -1) - 2
		}
		power += 2
	}
	return sineValue
}

// InverseSine returns the value of arcsin(x) where x is in radians. x can be
// negative or positive. arcsin(x) is calculated by expressing the arcsin value
// in a Taylor Series for all real numbers x (where x is the angle in radians).
// The series will be expanded to n terms to compute the value of arcsin(x).
// More on the series can found here:
// https://www.mathportal.org/formulas/pdf/taylor-series-formulas.pdf
func InverseSine(x float64, n int) float64 {
	if x > 1 || x < -1 {
		panic("Domain of arcsin is between -1 < x < 1 inclusive")
	}

	inverseSineValue := x
	numeratorBase := 1.0
	numeratorProduct := 1.0
	denominatorBase := 2.0
	denominatorProduct := 2.0
	power := 3.0

	for i := 1; i <= n; i++ {
		a := numeratorProduct / denominatorProduct
		b := ToThePowerFloat(x, power) / power
		inverseSineValue += a * b

		power += 2.0
		numeratorBase += 2.0
		numeratorProduct *= numeratorBase
		denominatorBase += 2.0
		denominatorProduct *= denominatorBase
	}
	return inverseSineValue
}

// Tangent returns the value of tan(x) where x is in radians.
func Tangent(x float64, n int) float64 {
	return Sine(x, n) / Cosine(x, n)
}

// InverseTangent returns the value of arctan(x) where x is in radians. x can be
// negative or positive. arctan(x) is calculated by expressing the arctan value
// in a Taylor Series for all real numbers x (where x is the angle in radians).
// The series will be expanded to n terms to compute the value of arctan(x).
// * http://mathworld.wolfram.com/MaclaurinSeries.html
// * https://www.mathportal.org/formulas/pdf/taylor-series-formulas.pdf
func InverseTangent(x float64, n int) float64 {
	v := 0.0
	power := 3.0
	denominator := 3.0

	for i := 0; i < n; i++ {
		v += ToThePowerFloat(x, power) / denominator
		if denominator < 0 {
			denominator = (denominator * -1) + 2
		} else {
			denominator = (denominator * -1) - 2
		}
		power += 2
	}

	if x > -1 && x < 1 {
		return x - v
	} else if x >= 1 {
		return v + (math.Pi / 2) - (1 / x)
	} else {
		return v - (math.Pi / 2) - (1 / x)
	}
}

// Cosine returns the value of cos(x) where x is in radians. x can be negative
// or positive. cos(x) is calculated by expressing the cos value in a Taylor
// Series for all real numbers x (where x is the angle in radians). The series
// will be expanded to n terms to compute the value of cos(x). More on the
// series can found here:
// http://people.math.sc.edu/girardi/m142/handouts/10sTaylorPolySeries.pdf
func Cosine(x float64, n int) float64 {

	cosineValue := 1.0
	power := 2.0
	denominator := -2.0

	for i := 0; i < n; i++ {
		cosineValue += ToThePowerFloat(x, power) / Factorial(denominator)
		if denominator < 0 {
			denominator = (denominator * -1) + 2
		} else {
			denominator = (denominator * -1) - 2
		}
		power += 2.0
	}
	return cosineValue
}

// InverseCosine computes and returns arccos(x) where -1 <= x <= 1.
func InverseCosine(x float64, n int) float64 {
	return math.Pi/2 - InverseSine(x, n)
}

// ConvertToRadian converts degree to radian.
func ConvertToRadian(x float64) float64 {
	return (x / 180.0) * math.Pi
}
