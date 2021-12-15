package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args[1:]) != 1 {
		fmt.Println("Expected 1 argument")
		os.Exit(1)
	}
	f, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Printf("%.2f feets = %.2f meters\n", f, feetToMeters(f))
}

const meterScale = 0.3048

func feetToMeters(f float64) float64 {
	return f * meterScale
}
