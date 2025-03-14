package main

import (
	"demoapp/demodb"
	"demoapp/employee"
	"demoapp/webapi"
	"encoding/json"
	"strings"
	"time"

	"fmt"
	"log"
)

func main() {
	var commandTimeout int = 60
	chCommand := make(chan string, 5)
	chPost := make(chan string, 5)

	webapi.SetCommandChannel(chCommand)
	webapi.SetPostChannel(chPost)

	isDbConnected, err := demodb.Connect()

	if err != nil {
		log.Fatal("Could not connect to DB")
	}

	if isDbConnected {
		defer demodb.Disconnect()

		go webapi.Start(":8080")

		if err != nil {
			log.Fatal(err)
		}

		exitMain := false
		cmd := ""
		var postBody string

		for !exitMain {
			select {
			case cmd = <-chCommand:
				if strings.ToLower(cmd) == "quit" {
					fmt.Println("cmd rx", cmd)
					exitMain = true
				} else {
					handleCommand(cmd)
				}

			case postBody = <-chPost:
				handlePost(postBody)

			case <-time.After(time.Duration(commandTimeout) * time.Second):
				fmt.Println("Cmd timeout exceeded")
				exitMain = true
			}
		}
	}

	fmt.Println("Main exiting - hit any key")
	fmt.Scanln()
}

func handleCommand(commandName string) (bool, error) {

	var resultBool bool = false
	var resultError error = nil

	commandName = strings.ToLower(commandName)

	fmt.Println("command: ", commandName)

	if commandName == "drop-employee-table" {
		demodb.DropEmployeeTable()
	} else if commandName == "create-employee-table" {
		demodb.CreateEmployeeTable()
	} else {
		fmt.Println("Unknown Command:", strings.ToLower(commandName))
	}

	return resultBool, resultError
}

func handlePost(postBody string) (bool, error) {
	var resultBool bool = false
	var resultError error = nil

	fmt.Println("POST IN", postBody)

	var employee employee.Employee

	err := json.Unmarshal([]byte(postBody), &employee)

	if err == nil {
		fmt.Println("POST Employee received")
		employee.Print()
		employee.Save()

	}

	return resultBool, resultError
}
