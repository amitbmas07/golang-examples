package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var strarray = []string{"lorem", "ipsum", "dolor", "sit", "amet"}

func main() {
	jsonData, err := json.Marshal(strarray)

	if err != nil {
		os.Exit(2)
	}

	fmt.Println(string(jsonData))
}