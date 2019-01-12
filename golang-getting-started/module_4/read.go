package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	path := readFilePath("Enter source file path: ")
	file, err := os.Open(getAbsolutePath(path))
	checkError(err)
	defer file.Close()

	slice := make([]Name, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		trimSlice(s)
		slice = append(slice, Name{fname: s[0], lname: s[1]})
	}
	checkError(scanner.Err())

	printSlice(slice)
}

func readFilePath(prompt string) string {
	fmt.Print(prompt)
	var input string
	fmt.Scan(&input)
	return input
}

func getAbsolutePath(path string) string {
	absPath, err := filepath.Abs(path)
	checkError(err)
	return absPath
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func trimSlice(strings []string) {
	for idx, str := range strings {
		if len(str) > 20 {
			strings[idx] = str[:20]
		}
	}
}

func printSlice(slice []Name) {
	for _, name := range slice {
		fmt.Println(name.fname + " " + name.lname)
	}
}

type Name struct {
	fname string
	lname string
}
