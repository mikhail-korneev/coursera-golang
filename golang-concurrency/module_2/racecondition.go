package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var wg sync.WaitGroup
var sharedInt int

func producer() {
	defer wg.Done()
	for i := 1; i < 100; i++ {
		sharedInt = rand.Intn(100)
		fmt.Printf("Producer: %d\n", sharedInt)
	}
}

func consumer() {
	defer wg.Done()
	for i := 1; i < 100; i++ {
		fmt.Printf("Consumer: %d\n", sharedInt)
	}
}

func main() {
	wg.Add(2)
	go producer()
	go consumer()
	wg.Wait()
	fmt.Println("Done")
}
