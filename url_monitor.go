package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const maxTries = 3
const delay = 2 * time.Second

const urlFilepath = "urls.txt"
const logFilepath = "log.txt"

func main() {

	showIntroduction()

	for {

		showMenu()

		var cmd int = readCommand()

		switch cmd {
		case 1:
			startUrlMonitor()
		case 2:
			showLogs()
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
		fmt.Println("Error on reading command:", err)
		return 0
	}

	return cmd
}

func startUrlMonitor() {

	fmt.Println("Monitoring urls...\n")

	urls := readUrlsFile()

	for i := 0; i < maxTries; i++ {
		for _, url := range urls {
			testUrl(url)
		}

		fmt.Println("\nWaiting for next tests...\n")
		time.Sleep(delay)
	}
}

func testUrl(url string) {

	var res, err = http.Get(url)

	if err != nil {
		fmt.Println("Error on requesting url", url)
		fmt.Println(err)
	} else {

		writeLog(url, res.StatusCode)

		if res.StatusCode >= 200 && res.StatusCode < 300 {
			fmt.Println("Url:", url, "Success:", res.StatusCode)
		} else if res.StatusCode > 200 && res.StatusCode < 400 {
			fmt.Println("Url:", url, "Redirect:", res.StatusCode)
		} else if res.StatusCode >= 400 && res.StatusCode < 500 {
			fmt.Println("Url:", url, "Client error:", res.StatusCode)
		} else {
			fmt.Println("Url:", url, "Server error: ", res.StatusCode)
		}
	}
}

func readUrlsFile() []string {

	fmt.Println("Reading file", urlFilepath, "\n")

	var urls []string

	file, err := os.Open(urlFilepath)

	if err != nil {
		fmt.Println("Error on opening file")
		fmt.Println(err)
		os.Exit(-1)
	}

	reader := bufio.NewReader(file)

	for {

		row, err := reader.ReadString('\n')

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("Error on reading line")
			fmt.Println(err)
		}

		urls = append(urls, strings.TrimSpace(row))
	}

	err = file.Close()

	if err != nil {
		fmt.Println("Error on closing file", file.Name())
		fmt.Println(err)
	}

	return urls
}

func writeLog(url string, statusCode int) {

	file := openFile(logFilepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)

	writeStringToFile(file, time.Now().Format(time.RFC3339)+" Url: "+url+" Status Code: "+strconv.Itoa(statusCode)+"\n")

	closeFile(file)
}

func showLogs() {
	fmt.Println("\nShowing logs...\n")

	bytes := readFile(logFilepath)

	fmt.Println(string(bytes))
}

func openFile(filepath string, flag int, perm os.FileMode) *os.File {
	file, err := os.OpenFile(filepath, flag, perm)
	checkError(err)

	return file
}

func readFile(filepath string) []byte {
	bytes, err := ioutil.ReadFile(filepath)
	checkError(err)

	return bytes
}

func closeFile(file *os.File) {
	err := file.Close()
	checkError(err)
}

func writeStringToFile(file *os.File, str string) {
	_, err := file.WriteString(str)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
