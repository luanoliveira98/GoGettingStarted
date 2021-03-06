package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitorings = 3
const delay = 5

func main() {

	showIntroduction()
	for {
		showMenu()

		command := getCommand()
		switch command {
		case 1:
			startMoniroting()
		case 2:
			getLogs()
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
	version := 5.0

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

	fmt.Println()
	fmt.Println("Monitoring...")
	sites := getSitesArchive()

	for i := 0; i < monitorings; i++ {
		for _, site := range sites {
			testSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println()
	}

	fmt.Println()

}

func getSitesArchive() []string {
	var sites []string

	archive, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("An error has occurred", err)
	}

	reader := bufio.NewReader(archive)
	for {
		row, err := reader.ReadString('\n')
		row = strings.TrimSpace(row)

		sites = append(sites, row)

		if err == io.EOF {
			break
		}
	}

	archive.Close()

	return sites
}

func testSite(site string) {

	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("An error has occurred", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "was successfully loaded!")
		setLog(site, true)
	} else {
		fmt.Println("Site:", site, "is in trouble. Status Code:", resp.StatusCode)
		setLog(site, false)
	}

}

func setLog(site string, status bool) {

	archive, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("An error has occurred", err)
	}

	archive.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	archive.Close()
}

func getLogs() {

	archive, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println("An error has occurred", err)
	}

	fmt.Println(string(archive))
}
