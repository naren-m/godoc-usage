/*
Package calculator is tutorial that shows how to use godoc for creating documentation to you go code
*/
package calculator

import "math"

// Math constant pi
// Online Encyclopedia for Integer Sequences reference to Pi
const (
	Pi = 3.1415 // http://oeis.org/A000796
)

// Add calculates and returns sum of two integers passed as input parameters
//
func Add(a int, b int) int {
	return a + b
}

/*
Min returns the smaller of x or y.

Special cases are:

		Min(x, -Inf) = Min(-Inf, x) = -Inf
		Min(x, NaN) = Min(NaN, x) = NaN
		Min(-0, ±0) = Min(±0, -0) = -0
*/
func Min(x, y float64) float64 {
	// Used the code from golang math package to demonstrate
	// how indentation looks like using Special case scenario
	switch {
	case math.IsInf(x, -1) || math.IsInf(y, -1):
		return math.Inf(-1)
	case math.IsNaN(x) || math.IsNaN(y):
		return math.NaN()
	case x == 0 && x == y:
		if math.Signbit(x) {
			return x
		}
		return y
	}
	if x < y {
		return x
	}
	return y
}
