//+build AWS

package main

import (
	"fmt"
	"net/url"
	"os"
	"time"
)

const (
	ephemeral = "ephemeral"
	inchannel = "in_channel"
)

func main() {

	resp := new(slackResp)

	if len(os.Args) != 2 { // <executable name> <url-encoded string>
		errorResp(resp, "Invalid number of arguments")
		return
	}

	decodedMap, err := url.ParseQuery(os.Args[1])
	if err != nil {
		errorResp(resp, "url.ParseQuery")
		return
	}

	token := decodedMap["token"]
	//user := decodedMap["user"] // FIXME user currently unused

	switch { // TODO add code to actually delete
	default:
		fl, err := getFiles(30, 20, token) // Get all the files and return a confirmation message to the user
		if err != nil {
			errorResp(resp, "No files match the criteria")
			return
		}
		createDeleteRequestResp(resp, fl)
	}
}

func errorResp(s *slackResp, e string) {
	s.RespType = ephemeral
	s.Text = fmt.Sprintf("Error: %v", e)
	fmt.Print(s)
}

func createDeleteRequestResp(s *slackResp, fl *fileList) {
	if len(fl.Files) == 0 {
		errorResp(s, "No files match the criteria")
		return
	}

	var fileList string
	var fileID string
	for _, f := range fl.Files {
		fileList = fmt.Sprintf("%v%v\t%v KiB\t%v\n", fileList, f.Title, f.Size/1024, time.Unix(f.Created, 0))
		fileID = fmt.Sprintf("%v %v", fileID, f.ID)
	}

	s.RespType = inchannel
	s.Text = "Would you like to delete these files?"
	s.Attachments = make([]attach, 1)
	for i := range s.Attachments {
		s.Attachments[i].Text = fileList
		s.Attachments[i].Actions = make([]action, 2) // Yes and No
		s.Attachments[i].Actions[0].Name = "yes"
		s.Attachments[i].Actions[0].Value = fileID
		s.Attachments[i].Actions[0].Type = "button"
		s.Attachments[i].Actions[0].Text = "Yes"
		s.Attachments[i].Actions[0].Confirm.Ok_text = "Deleted"
		s.Attachments[i].Actions[0].Confirm.Text = "Are you sure?"
		s.Attachments[i].Actions[1].Name = "no"
		s.Attachments[i].Actions[1].Value = "no"
		s.Attachments[i].Actions[1].Type = "button"
		s.Attachments[i].Actions[1].Text = "No"
	}

	// Dump the response
	fmt.Print(s)
}
