package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

var animals = map[string]Animal{
	"cow":   {food: "grass", locomotion: "walk", noise: "moo"},
	"bird":  {food: "worms", locomotion: "fly", noise: "peep"},
	"snake": {food: "mice", locomotion: "slither", noise: "hss"},
}

var animalMap = createMap("cow", "bird", "snake")
var actionMap = createMap("eat", "move", "speak")

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func (a Animal) PerformAction(action string) {
	switch action {
	case "eat":
		a.Eat()
	case "move":
		a.Move()
	case "speak":
		a.Speak()
	}
}

func main() {
	for {
		fmt.Print("> ")
		request := readString()
		if request == "exit" {
			break
		}
		reqAnimal, reqAction := getRequestParams(request)
		animal := animals[reqAnimal]
		animal.PerformAction(reqAction)
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

func getRequestParams(request string) (animal string, action string) {
	params := strings.Split(request, " ")
	validate(params)
	return params[0], params[1]
}

func validate(params []string) {
	validateSize(params)
	validateAnimal(params[0])
	validateAction(params[1])
}

func validateSize(params []string) {
	if len(params) != 2 {
		panic(errors.New("number of request parameters is not valid"))
	}
}

func validateAnimal(animal string) {
	if _, valid := animalMap[animal]; !valid {
		panic(errors.New("animal request parameter is not valid"))
	}
}

func validateAction(action string) {
	if _, valid := actionMap[action]; !valid {
		panic(errors.New("action request parameter is not valid"))
	}
}

func createMap(param ...string) map[string]string {
	m := make(map[string]string)
	for _, p := range param {
		m[p] = ""
	}
	return m
}
