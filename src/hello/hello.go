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

	switch command {
	case 1:
		fmt.Println("Monitoring...")
	case 2:
		fmt.Println("Showing Logs...")
	case 0:
		fmt.Println("Exiting program!")
	default:
		fmt.Println("I don't know this command!")
	}
}
