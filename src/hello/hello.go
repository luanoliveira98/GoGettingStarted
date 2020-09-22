package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	showIntroduction()
	for {
		showMenu()

		command := getCommand()
		switch command {
		case 1:
			startMoniroting()
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

func startMoniroting() {

	fmt.Println("Monitoring...")
	var sites [4]string
	sites[0] = "https://random-status-code.herokuapp.com/"
	sites[1] = "https://www.alura.com.br"
	sites[2] = "https://www.caelum.com.br"

	fmt.Println(sites)

	site := "https://random-status-code.herokuapp.com/"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "was successfully loaded!")
	} else {
		fmt.Println("Site:", site, "is in trouble. Status Code:", resp.StatusCode)
	}

}
