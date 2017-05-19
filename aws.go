//+build AWS

package main

import (
	"fmt"
	"net/url"
	"os"
)

func main() {

	if len(os.Args) != 2 { // <executable name> <url-encoded string>
		fmt.Fprintln(os.Stderr, "Error: Invalid number of arguments")
		return
	}

	decodedMap, err := url.ParseQuery(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: url.ParseQuery")
		return
	}

	token := decodedMap["token"]
	user := decodedMap["user"]

	switch text[0] {
	case "DELETE": // User selected to delete the files previously sent
		delFiles(text[1]) // Second argument is expected to be the file list
	default:
		getFiles() // Get all the files and return a confirmation message to the user
	}
}
