package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const maxTries = 3
const delay = 2 * time.Second

func main() {

	showIntroduction()

	for {

		showMenu()

		var cmd int = readCommand()

		switch cmd {
		case 1:
			startUrlMonitor()
		case 2:
		case 0:
			os.Exit(0)
		default:
			fmt.Println("Command not found")
			os.Exit(-1)
		}
	}
}

func showIntroduction() {

	var version float32 = 1.1

	fmt.Println("Website status monitor")
	fmt.Println("Version:", version)
}

func showMenu() {
	fmt.Println("\nMenu\n1 - Start monitoring\n2 - Show logs\n0 - Exit program\n")
}

func readCommand() int {

	var cmd int
	_, err := fmt.Scan(&cmd)

	if err != nil {
		fmt.Println("Error on reading command:", err.Error())
		return 0
	}

	return cmd
}

func startUrlMonitor() {

	fmt.Println("Monitoring urls...\n")

	urls := []string{
		"https://random-status-code.herokuapp.com/",
		"https://random-status-code.herokuapp.com/",
		"https://random-status-code.herokuapp.com/",
		"https://random-status-code.herokuapp.com/",
	}

	for i := 0; i < maxTries; i++ {
		for _, url := range urls {
			testUrl(url)
		}

		fmt.Println("\nWaiting for next tests...\n")
		time.Sleep(delay)
	}
}

func testUrl(url string) {

	var res, _ = http.Get(url)

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		fmt.Println("Url:", url, "Success:", res.StatusCode)
	} else {
		fmt.Println("Url:", url, "Error: ", res.StatusCode)
	}
}
