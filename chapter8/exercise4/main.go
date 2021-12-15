package main

import (
	"fmt"

	"github.com/santiago-rodrig/gointro/chapter8/exercise4/math"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := math.Average(xs)
	fmt.Println(avg)
}
