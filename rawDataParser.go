package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"runtime"
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

	log.Info(runtime.Version())

	//////////////////////////////
	// Loading json from string //
	//////////////////////////////
	// string you want to convert to JSON object
	var stringData = `{"Name": "John", "Age": 12}`

	// define structure of the JSON (struct in GoLang)
	type MyStruct struct {
		Name string `json:"Name"`
		Age  int    `json:"Age"`
	}

	// load this stringData in struct
	var MyJsonData MyStruct
	err := json.Unmarshal([]byte(stringData), &MyJsonData) // convert string data to []bytes
	if err != nil {
		// handle error here
		println("Got an error while loading string into JSON")
	}

	///////////////////////////////////////////
	// Loading unstructured json from string //
	///////////////////////////////////////////
	var unstructuredData string = `
{
  "users": [
    {
      "name": "Elliot",
      "type": "Reader",
      "age": 23,
      "social": {
        "facebook": "https://facebook.com",
        "twitter": "https://twitter.com"
      }
    },
    {
      "name": "Fraser",
      "type": "Author",
      "age": 17,
      "social": {
        "facebook": "https://facebook.com",
        "twitter": "https://twitter.com"
      }
    }
  ]
}
`
	var unstructuredJSON map[string]interface{}
	err = json.Unmarshal([]byte(unstructuredData), &unstructuredJSON)
	if err != nil {
		return
	}

	////////////////////////////
	// reading file in golang //
	////////////////////////////
	//file, err := os.Open("./testFile1.txt")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//println("Successfully Opened file")
	//defer file.Close()
	var file, _ = ioutil.ReadFile("./testFile1.txt")
	var fileData string = string(file)
	fmt.Printf("There are : %d rows in the file\n", len(fileData))
	fmt.Printf("Data is :\n%s\n", fileData)
	var ans []string = strings.Split(fileData, "\n")
	fmt.Printf("There  are %d rows in file\n", len(ans))
	for _, val := range ans {
		fmt.Printf("\nSplitted Data is %s\n", val)
	}

	////////////////////////////
	// Reading file in golang //
	////////////////////////////
	var fileName string = "./testFile2.txt"
	log.Infof("Reading file %s", fileName)
	f, _ := ioutil.ReadFile(fileName)
	//var fData string = string(f)
	//log.Infof("File Contents are : %s", fData)
	log.Infof("File Contents 2nd are : %s", f)

	/////////////////////////////
	// Read each line in a file /
	/////////////////////////////
	fileName = "./testFile2.txt"

	//
	//
	//
	//
	//////////////////////////////////
	//// Parsing datetime in Golang //
	//////////////////////////////////
	//// use the dateparse library in golang - equivalent to ciso8601 library
	//ans, _ := dateparse.ParseAny("2021-08-12 12:34:54")
	//println(ans.String())

	ReadFile("./data.csv")

}
