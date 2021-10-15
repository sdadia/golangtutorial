package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/araddon/dateparse"
	"os"
	"strings"
)

func ReadFile(fileName string) []string {

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened file")
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var newLine = scanner.Text()
		ans := strings.Split(newLine, "Search :")
		println(ans)
		os.Exit(-1)
	}

	return lines
}

func main() {

	//////////////////////////////
	// Loading json from string //
	//////////////////////////////
	// string you want to convert to JSON object
	var data = `{"Name": "John", "Age": 12}`

	// define structure of the JSON (struct in GoLang)
	type MyStruct struct {
		Name string `json:"Name"`
		Age  int    `json:"Age"`
	}

	// load this data in struct
	var MyJsonData MyStruct
	err := json.Unmarshal([]byte(data), &MyJsonData) // convert string data to []bytes
	if err != nil {
		// handle error here
		println("Got an error while loading string into JSON")
	}

	////////////////////////////////
	// Parsing datetime in Golang //
	////////////////////////////////
	// use the dateparse library in golang - equivalent to ciso8601 library
	ans, _ := dateparse.ParseAny("2021-08-12 12:34:54")
	println(ans.String())

	ReadFile("./data.csv")

}
