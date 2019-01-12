package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	userInfo := make(map[string]string)
	userInfo["name"] = readString("Enter your name: ")
	userInfo["address"] = readString("Enter your address: ")
	barr, err := json.Marshal(userInfo)
	if err != nil {
		panic(errors.New("error to marshal specified map"))
	}
	fmt.Println("JSON object from the map:\n" + string(barr))
}

func readString(prompt string) string {
	fmt.Printf(prompt)
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		panic(errors.New("failed to read from stdin"))
	}
	return strings.TrimSuffix(str, "\n")
}
