package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args[1:]) != 1 {
		fmt.Println("Expected one argument")
		os.Exit(1)
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	if n < 1 {
		fmt.Println("Expected an integer greater than one")
		os.Exit(3)
	}
	fmt.Printf("The number at position %d of the fibonacci sequence is %d\n", n, fibonacci(n-1))
}

func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
