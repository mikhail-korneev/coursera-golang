package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	fmt.Println("Enter values for acceleration (a), initial velocity (v0) and displacement (s0)")
	var a, v0, s0 float64
	if _, err := fmt.Scan(&a, &v0, &s0); err != nil {
		log.Fatal(err)
	}
	fn := GenDisplaceFn(a, v0, s0)
	fmt.Println("Value for time (t): ")
	var t float64
	if _, err := fmt.Scan(&t); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Displacement: %f", fn(t))
}

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return (a*math.Pow(t, 2))/2 + v0*t + s0
	}
}
