package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	name string
}

func (c Cow) Eat()   { fmt.Println("grass") }
func (c Cow) Move()  { fmt.Println("walk") }
func (c Cow) Speak() { fmt.Println("moo") }

type Bird struct {
	name string
}

func (c Bird) Eat()   { fmt.Println("worms") }
func (c Bird) Move()  { fmt.Println("fly") }
func (c Bird) Speak() { fmt.Println("peep") }

type Snake struct {
	name string
}

func (c Snake) Eat()   { fmt.Println("mice") }
func (c Snake) Move()  { fmt.Println("slither") }
func (c Snake) Speak() { fmt.Println("hss") }

var animalTypeMap = createMap("cow", "bird", "snake")
var infoTypeMap = createMap("eat", "move", "speak")

var animalMap = make(map[string]Animal)

func main() {
	for {
		fmt.Print("> ")
		request := readString()
		if request == "exit" {
			break
		}
		performAction(request)
	}
}

func readString() string {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		panic(errors.New("failed to read from stdin"))
	}
	return strings.TrimRight(str, "\n")
}

func performAction(request string) {
	params := strings.Split(request, " ")
	validateSize(params)
	switch params[0] {
	case "newanimal":
		createAnimal(params[1], params[2])
	case "query":
		getAnimalInfo(params[1], params[2])
	default:
		fmt.Println("Unknown command, please try again.")
	}
}

func validateSize(params []string) {
	if len(params) != 3 {
		panic(errors.New("number of request parameters is not valid"))
	}
}

func createAnimal(name string, animalType string) {
	validateAnimalType(animalType)
	if _, ok := animalMap[name]; !ok {
		switch animalType {
		case "cow":
			animalMap[name] = Cow{name}
		case "bird":
			animalMap[name] = Bird{name}
		case "snake":
			animalMap[name] = Snake{name}
		}
		fmt.Println("Created it!")
	} else {
		fmt.Printf("Name '%s' is already used, please try another name\n", name)
	}
}

func validateAnimalType(animalType string) {
	if _, valid := animalTypeMap[animalType]; !valid {
		panic(errors.New("animal type request parameter is not valid"))
	}
}

func getAnimalInfo(name string, infoType string) {
	validateInfoType(infoType)
	animal := animalMap[name]
	switch infoType {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	}
}

func validateInfoType(infoType string) {
	if _, valid := infoTypeMap[infoType]; !valid {
		panic(errors.New("animal info type request parameter is not valid"))
	}
}

func createMap(param ...string) map[string]string {
	m := make(map[string]string)
	for _, p := range param {
		m[p] = ""
	}
	return m
}
