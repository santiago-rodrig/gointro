// Package math provides functions for working with slices of numbers.
package math

// Gets the average of a set of numbers.
func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

// Gets the minimum in a set of numbers.
func Min(xs []float64) float64 {
	switch {
	case len(xs) == 0:
		return 0
	case len(xs) == 1:
		return xs[0]
	default:
		smallest := xs[0]
		for _, n := range xs {
			if n < smallest {
				smallest = n
			}
		}
		return smallest
	}
}

// Gets the maximum in a set of numbers.
func Max(xs []float64) float64 {
	switch {
	case len(xs) == 0:
		return 0
	case len(xs) == 1:
		return xs[0]
	default:
		greatest := xs[0]
		for _, n := range xs {
			if n > greatest {
				greatest = n
			}
		}
		return greatest
	}
}
