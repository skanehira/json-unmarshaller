# json-unmarshaler
A wrapper that makes it easy to understand the location of json's parse error.

# Usage
```go
package main

import (
	"fmt"

	json "github.com/skanehira/json-unmarshaller"
)

func main() {
	var s struct {
		Name      string   `json:"name"`
		Age       int      `json:"age"`
		IsHuman   bool     `json:"is_human"`
		Languages []string `json:"languages"`
	}

	data := `
{
	"name": "gorilla,
	"age": 26,
	"is_human": false,
	"languages": [
		"Go",
		"PHP",
		"Java",
		"Java Script",
		"HTML"
	]
}`

	b := []byte(data)

	if err := json.Unmarshal(b, &s); err != nil {
		panic(err)
	}

	fmt.Printf("%#+v", s)
}
```
