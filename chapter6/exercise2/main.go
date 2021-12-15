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
	half, even := halfInt(n)
	fmt.Printf("Half = %d; Even? %t\n", half, even)
}

func halfInt(n int) (int, bool) {
	return n / 2, n%2 == 0
}
