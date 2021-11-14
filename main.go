package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// get
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
	getHelp := getCmd.Bool("h", false, "help")
	getVer := getCmd.Bool("v", false, "version")
	getAll := getCmd.Bool("l", false, "list all items ( un-completed )")
	getComp := getCmd.Bool("c", true, "list completed items")
	getDellID := getCmd.String("d", "", "delete item")
	getID := getCmd.String("id", "", "Todos ID")

	// add
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addID := addCmd.String("id", "", "Todos ID")
	addItem := addCmd.String("item", "", "Todos Item")
	addDate := addCmd.String("date", "", "Todos Date")
	addStatus := addCmd.String("status", "", "Todos Status")
	addMark := addCmd.String("m", "", "isComplete")

	if len(os.Args) < 2 {
		fmt.Println("expected 'get' , 'add' or 'del' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "get":
		HandleGet(getCmd, getAll, getID, getComp, getDellID, getHelp, getVer)
	case "add":
		HandleAdd(addCmd, addID, addItem, addDate, addStatus, addMark)
	default:
	}

}

func HandleGet(getCmd *flag.FlagSet, l *bool, id *string, c *bool, d *string, h *bool, v *bool) {

	getCmd.Parse(os.Args[2:])

	// if *l == false && *id == "" {
	// 	fmt.Print("id is required or specify -l for all todos")
	// 	getCmd.PrintDefaults()
	// 	os.Exit(1)
	// }

	if *l {
		//return all videos
		todos := getTodos()

		fmt.Printf("ID \t Item \t Date \t Status \n")
		for _, todo := range todos {
			if todo.Status == "false" {
				fmt.Printf("%v \t %v \t %v \t %v \n", todo.Id, todo.Item, todo.Date, todo.Status)
			}
		}

		return
	}

	if *c {

		todos := getTodos()
		for _, todo := range todos {
			if todo.Status == "true" {
				fmt.Printf("ID \t Item \t Date \t Status \n")
				fmt.Printf("%v \t %v \t %v  \t %v  \n", todo.Id, todo.Item, todo.Date, todo.Status)
			}
		}
	}

	if *d != "" {
		todos := getTodos()
		id := *d
		for index, todo := range todos {
			if id == todo.Id {
				todos = append(todos[0:index], todos[index+1:]...)
			}
		}
	}

	if *id != "" {
		todos := getTodos()
		id := *id
		for _, todo := range todos {
			if id == todo.Id {
				fmt.Printf("ID \t Item \t Date \t Status \n")
				fmt.Printf("%v \t %v \t %v  \t %v  \n", todo.Id, todo.Item, todo.Date, todo.Status)

			}
		}
	}

	if *h {
		fmt.Println(" get -h : help ")
		fmt.Println(" get -v : version ")
		fmt.Println(" get -l : list all items (un-completed) ")
		fmt.Println(" get -c : list completed items ")
		fmt.Println(" add -a : add new item : { -id <id_number> -item <item_name> -date <date_name> -status <status_name> }")
		fmt.Println(" add -m : mark as complete ")
		fmt.Println(" get -d : delete item ")
	}

	if *v {
		fmt.Println(" module todo / go 1.17 ")
	}

}

func ValidateTodo(addCmd *flag.FlagSet, id *string, item *string, date *string, status *string) {
	if *id == "" || *item == "" || *date == "" || *status == "" {
		fmt.Print("all fields are required for adding a todo")
		fmt.Println()
		addCmd.PrintDefaults()
		os.Exit(1)
	}

}

func HandleAdd(addCmd *flag.FlagSet, id *string, item *string, date *string, status *string, m *string) {

	addCmd.Parse(os.Args[2:])

	// everything add control
	ValidateTodo(addCmd, id, item, date, status)

	todo := todo{
		Id:     *id,
		Item:   *item,
		Date:   *date,
		Status: *status,
	}
	// read data on .json file
	todos := getTodos()
	// save old and new data together
	todos = append(todos, todo)
	// dave data on .json file
	saveTodos(todos)

	if *m != "" {
		todos := getTodos()
		id := *m
		for _, todo := range todos {
			if id == todo.Id {
				todo.Status = "true"
				todos = append(todos, todo)
				saveTodos(todos)
			}
		}
	}

}
