package main

import "fmt"

func main() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	smallest := x[0]
	for _, n := range x[1:] {
		if n < smallest {
			smallest = n
		}
	}
	fmt.Printf("The smalles number in the list is %d\n", smallest)
}
