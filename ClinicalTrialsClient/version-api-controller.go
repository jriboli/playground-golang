package main

import "fmt"

func GetVersion() {
	url := baseUrl + "/version"

	var version Version

	err := GetJson(url, &version)
	if err != nil {
		fmt.Printf("Error -- %s\n", err.Error())
	} else {
		fmt.Printf("Response: %s", version)
	}
}