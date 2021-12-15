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
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	odds := make([]int, n)
	generator := makeOddGenerator()
	for i := range odds {
		odds[i] = generator()
	}
	fmt.Printf("The following odd numbers were generated:\n|\n|---> %v\n", odds)
}

func makeOddGenerator() func() int {
	n := -1
	return func() int {
		n += 2
		return n
	}
}
