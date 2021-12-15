package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args[1:]) != 2 {
		fmt.Println("Expected 2 arguments")
		os.Exit(1)
	}
	ips := []*int{new(int), new(int)}
	var err error
	for i, n := range os.Args[1:] {
		*ips[i], err = strconv.Atoi(n)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
	}
	a, b := ips[0], ips[1]
	fmt.Println("Before swapping.")
	fmt.Printf("a = %d; b = %d\n", *a, *b)
	*a, *b = *b, *a
	fmt.Println("After swapping")
	fmt.Printf("a = %d; b = %d\n", *a, *b)
}
