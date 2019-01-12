package main

import (
	"fmt"
)

func main() {
	var floatNum float32
	fmt.Printf("Enter floating point number: ")
	fmt.Scan(&floatNum)
	fmt.Printf("Floating point number: %f\n", floatNum)
	result := int(floatNum)
	fmt.Printf("Integer number: %d\n", result)
}
