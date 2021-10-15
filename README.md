# golangtutorial

# Load string to JSON
```go

// load JSON package
import (
"encoding/json"
)

// string you want to convert to JSON object
var data string = `{"Name": "John", "Age": 12}`

// define structure of the JSON (struct in GoLang)
type MyStruct struct {
Name string `json:"Name"`
Age int `json:"Age"`
}

// load this data in struct
var MyJsonData Mystruct
err := json.Unmarshall([]byte(data), &MyJsonData) // convert string data to []bytes
if err != nil {
// handle error here
}


```