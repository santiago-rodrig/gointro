package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args[1:]) < 1 {
		fmt.Println("Expected at least one argument")
		os.Exit(1)
	}
	ns := make([]float64, len(os.Args[1:]))
	for i, arg := range os.Args[1:] {
		n, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		ns[i] = n
	}
	fmt.Printf("The greatest number of %v is %.2f\n", ns, findGreatest(ns...))
}

func findGreatest(ns ...float64) float64 {
	greatest := ns[0]
	for _, n := range ns[1:] {
		if n > greatest {
			greatest = n
		}
	}
	return greatest
}
