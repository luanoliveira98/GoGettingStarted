package main

import "fmt"

func main() {
	name := "Luan"
	version := 2.0

	fmt.Println("Hello, mr.", name)
	fmt.Println("-----------------------")

	fmt.Println("1 - Start Monitoring")
	fmt.Println("2 - Show Logs")
	fmt.Println("0 - Exit Program")

	fmt.Println("-----------------------")
	fmt.Println("Version:", version)

	var command int
	fmt.Scan(&command)

	if command == 1 {
		fmt.Println("Monitoring...")
	} else if command == 2 {
		fmt.Println("Showing Logs...")
	} else if command == 0 {
		fmt.Println("Exiting program!")
	} else {
		fmt.Println("I don't know this command!")
	}
}
