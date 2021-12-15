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
	fmt.Printf("%.2f°F = %.2f°C\n", f, FahrenheitToCelsius(f))
}

func FahrenheitToCelsius(f float64) (c float64) {
	c = (f - 32) * 5 / 9
	return
}
