package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	slice := make([]int, 3)
	fmt.Printf("Initial slice: %v\n", slice)
	for {
		fmt.Printf("Enter an integer (X to exit): ")
		inputString := readString()
		if inputString == "X" {
			break
		}
		i, e := strconv.Atoi(inputString)
		if e != nil {
			fmt.Println("Invalid value, please try again")
			continue
		}
		slice = append(slice, i)
		sort.Ints(slice)
		fmt.Printf("Sorted slice: %v\n", slice)
	}
}

func readString() string {
	var input string
	fmt.Scan(&input)
	return input
}
