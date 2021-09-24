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
)

const numMonitoring = 5
const delay = 2

func main() {
	for {
		showMenu()
		input := readMenuInput()

		switch input {
		case 1:
			startMonitor()
		case 2:
			showLogs()
		case 0:
			quit()
		default:
			fmt.Println("Invalid input!")
		}
	}
}

func showMenu() {
	fmt.Println("1- Start monitor")
	fmt.Println("2- Show Logs")
	fmt.Println("0- Exit")
}

func readMenuInput() int {
	var read int
	fmt.Scan(&read)
	return read
}

func startMonitor() 
{
	fmt.Println("Monitoring......")

	sites := getSites()

	for i := 0; i < numMonitoring; i++ {
		fmt.Println("TEST", i+1, "/", numMonitoring)
		for _, site := range sites {
			resp, err := http.Get(site)

			if err != nil {
				showError(&err)
			}
			//fmt.Println("Index:", i, "->", site, "->", resp.StatusCode)
			fmt.Println(resp.StatusCode, "->", site)
			generateLog(site, resp.StatusCode)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
	fmt.Println("FINISHED!")
	fmt.Println("")
}

func showLogs() {
	fmt.Println("Showing logs......")
	file, err := ioutil.ReadFile("logs.txt")

	if err != nil {
		showError(&err)
	}

	fmt.Println(string(file))
}

func quit() {
	fmt.Println("Exiting...")
	os.Exit(2)
}

func getSites() []string {

	var sites []string

	fileName := "sites.txt"

	file, err := os.Open(fileName)
	//file, err := ioutil.ReadFile(fileName)

	if err != nil {
		showError(&err)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		sites = append(sites, line)

		if err == io.EOF {
			break
		}
	}

	file.Close()

	return sites
}

func showError(err *error) {
	fmt.Println("Error:", *err)
}

func generateLog(site string, status int) {
	file, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		showError(&err)
	}
	date := time.Now().Format("[02/01/2006 15:04:05]: ")
	file.WriteString(date + strconv.Itoa(status) + " -> " + site + "\n")
	file.Close()
}