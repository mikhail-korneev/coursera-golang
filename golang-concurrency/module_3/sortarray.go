package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"sync"
)

const partCnt = 4

func main() {
	// Input initial slice.
	input := inputSlice()
	fmt.Printf("Initial slice: %v\n", input)

	parts := partition(input, partCnt)

	// Sorting parts of the array.
	var wg sync.WaitGroup
	wg.Add(partCnt)
	for _, part := range parts {
		go sortRoutine(part, &wg)
	}
	wg.Wait()

	// Merging sorted parts into one large sorted one.
	res := make([]int, 0, len(input))
	for _, part := range parts {
		res = append(res, part...)
	}
	sort.Ints(res)

	// Sorted result slice.
	fmt.Println(res)
}

func inputSlice() []int {
	slice := make([]int, 0)
	fmt.Println("Please type in a sequence of integers (X to break):")
	for {
		fmt.Printf("> ")
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
	}
	return slice
}

func readString() string {
	var input string
	fmt.Scan(&input)
	return input
}

func partition(array []int, partCnt int) [][]int {
	partSize := round(float64(len(array)) / float64(partCnt))
	var parts [][]int
	for startIdx := 0; startIdx < partSize*partCnt; startIdx += partSize {
		endIdx := startIdx + partSize

		// Size correction for the last part.
		if (endIdx == partSize*partCnt) && (endIdx != len(array)) {
			endIdx = len(array)
		}

		parts = append(parts, array[startIdx:endIdx])
	}

	return parts
}

// Rounds a float number to the closest int number.
// Golang doesn't provide round() function with the same behaviour at the moment.
func round(x float64) int {
	if x-math.Trunc(x) >= 0.5 {
		return int(math.Ceil(x))
	}
	return int(math.Floor(x))
}

func sortRoutine(slice []int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Sorting subslice: %v\n", slice)
	sort.Ints(slice)
}
