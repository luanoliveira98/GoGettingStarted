package main

import (
	"fmt"
	"os"
)

func main() {

	showIntroduction()
	showMenu()
	command := getCommand()

	switch command {
	case 1:
		fmt.Println("Monitoring...")
	case 2:
		fmt.Println("Showing Logs...")
	case 0:
		fmt.Println("Exiting program!")
		os.Exit(0)
	default:
		fmt.Println("I don't know this command!")
		os.Exit(-1)
	}
}

func showIntroduction() {

	name := "Luan"
	version := 2.3

	fmt.Println("Hello, mr.", name)
	fmt.Println("Version:", version)
}

func showMenu() {
	fmt.Println("1 - Start Monitoring")
	fmt.Println("2 - Show Logs")
	fmt.Println("0 - Exit Program")
}

func getCommand() int {
	var command int
	fmt.Scan(&command)

	return command
}
