package main

import (
	"fmt"
	"my/jira-go-hooks/handlers"
	"my/jira-go-hooks/structures"
	"encoding/json"
)

func main() {
	contents, _ := handlers.GetData()
	var i  = structures.IssueFields{}

	err := json.Unmarshal([]byte(contents), &i)
	if err != nil {
		panic(err)
	}
	fmt.Println(err, i)
}
