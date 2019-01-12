package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("Enter string: ")
	inputStr := readString()
	if len(inputStr) >= 3 && isMatchingString(inputStr) {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}

func readString() string {
	inputReader := bufio.NewReader(os.Stdin)
	inputStr, _ := inputReader.ReadString('\n')
	return strings.TrimRight(inputStr, "\n")
}

func isMatchingString(str string) bool {
	str = strings.ToLower(str)
	return strings.HasPrefix(str, "i") &&
		strings.HasSuffix(str, "n") &&
		strings.ContainsAny(str, "a")
}
