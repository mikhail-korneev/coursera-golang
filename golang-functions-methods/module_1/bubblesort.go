package main

import (
	"fmt"
	"strconv"
)

func main() {
	slice := make([]int, 0)
	fmt.Println("Please type in a sequence of up to 10 integers")
	for {
		fmt.Printf("Enter an integer (X to break): ")
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
		if len(slice) == 10 {
			break
		}
	}
	fmt.Printf("\nInitial slice: %v\n", slice)
	BubbleSort(slice)
	fmt.Printf("Sorted slice: %v\n", slice)
}

func readString() string {
	var input string
	fmt.Scan(&input)
	return input
}

func BubbleSort(slice []int) {
	for j := len(slice) - 1; j > 0; j-- {
		for i := 0; i < j; i++ {
			if slice[i] > slice[i+1] {
				Swap(slice, i)
			}
		}
	}
}

func Swap(slice []int, i int) {
	tmp := slice[i]
	slice[i] = slice[i+1]
	slice[i+1] = tmp
}
