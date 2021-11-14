package main

import (
	"encoding/json"
	"io/ioutil"
)

type todo struct {
	Id     string
	Item   string
	Date   string
	Status string
}

func getTodos() (todos []todo) {

	fileBytes, err := ioutil.ReadFile("./todo.json")

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(fileBytes, &todos)

	if err != nil {
		panic(err)
	}

	return todos
}

func saveTodos(todos []todo) {

	todoBytes, err := json.Marshal(todos)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("./todo.json", todoBytes, 0644)
	if err != nil {
		panic(err)
	}

}
