package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonstring := `
		{
			"name": "ijtt",
			"age" : 30
		}
	`
	data := make(map[string]interface{})

	fmt.Println(data)
	err := json.Unmarshal([]byte(jsonstring), &data)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(data["name"])
}
