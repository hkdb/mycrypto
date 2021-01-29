package main

import (
	"fmt"
	"log"
)

func confirm() bool {
	var response string
	_, err := fmt.Scanln(&response)
	if err != nil {
		if err.Error() == "unexpected newline" {
			return true
		} else {
			log.Fatal(err)
		}
	}

	okayResponses := []string{"y", "Y", "yes", "Yes", "YES", ""}
	nokayResponses := []string{"n", "N", "no", "No", "NO"}
	if containsString(okayResponses, response) {
		return true
	} else if containsString(nokayResponses, response) {
		return false
	} else {
		return true
	}
}

// If slice does not contain element, returns -1.
func posString(slice []string, element string) int {
	for index, elem := range slice {
		if elem == element {
			return index
		}
	}
	return -1
}

// containsString returns true if slice contains element
func containsString(slice []string, element string) bool {
	return !(posString(slice, element) == -1) // Read CSV
}
