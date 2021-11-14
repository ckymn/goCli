package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	//'videos get' subcommand
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)

	// inputs for `videos get` command
	getAll := getCmd.Bool("all", false, "Get all todos")
	getID := getCmd.String("id", "", "Todos ID")

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)

	addID := addCmd.String("id", "", "Todos ID")
	addItem := addCmd.String("item", "", "Todos Item")
	addDate := addCmd.String("date", "", "Todos Date")

	if len(os.Args) < 2 {
		fmt.Println("expected 'get' or 'add' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "get":
		HandleGet(getCmd, getAll, getID)
	case "add":
		HandleAdd(addCmd, addID, addItem, addDate)
	default:
	}

}

func HandleGet(getCmd *flag.FlagSet, all *bool, id *string) {

	getCmd.Parse(os.Args[2:])

	if *all == false && *id == "" {
		fmt.Print("id is required or specify --all for all todos")
		getCmd.PrintDefaults()
		os.Exit(1)
	}

	if *all {
		//return all videos
		todos := getTodos()

		fmt.Printf("ID \t Item \t Date \n")
		for _, todo := range todos {
			fmt.Printf("%v \t %v \t %v \n", todo.Id, todo.Item, todo.Date)
		}

		return
	}

	if *id != "" {
		todos := getTodos()
		id := *id
		for _, todo := range todos {
			if id == todo.Id {
				fmt.Printf("ID \t Item \t Date \n")
				fmt.Printf("%v \t %v \t %v \n", todo.Id, todo.Item, todo.Date)

			}
		}
	}

}

func ValidateTodo(addCmd *flag.FlagSet, id *string, item *string, date *string) {
	if *id == "" || *item == "" || *date == "" {
		fmt.Print("all fields are required for adding a todo")
		addCmd.PrintDefaults()
		os.Exit(1)
	}

}

func HandleAdd(addCmd *flag.FlagSet, id *string, item *string, date *string) {
	ValidateTodo(addCmd, id, item, date)

	todo := todo{
		Id:   *id,
		Item: *item,
		Date: *date,
	}

	todos := getTodos()
	todos = append(todos, todo)

	saveTodos(todos)
}
