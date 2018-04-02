package handlers

import "io/ioutil"

func GetData() ([]byte, error) {
	raw, err := ioutil.ReadFile("structures/examples/jira_issue_updated.json")
	if err != nil {
		panic(err)
	}
	return raw, err
}


