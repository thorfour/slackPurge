//+build AWS GCE

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
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

	var a interface{}
	if err := json.Unmarshal([]byte(os.Args[1]), &a); err != nil {
		errorResp(resp, fmt.Sprintf("%v :: %v", err.Error(), os.Args[1]))
		return
	}

	m, ok := a.(map[string]interface{})
	if !ok {
		errorResp(resp, "map parse error")
		return
	}

	user, ok := m["user_id"].(string)
	if !ok { // This means it's a response from a button

		p, ok := m["payload"].(string)
		if !ok {
			errorResp(resp, "Unable to get Payload from map")
			return
		}

		pay := new(slackPayload)
		if err := json.Unmarshal([]byte(p), pay); err != nil {
			errorResp(resp, fmt.Sprintf("Unable to Unmarshal payload: %v", err))
			return
		}

		if pay.Actions[0].Name == "yes" {

			user = pay.User.ID
			files := strings.Split(strings.TrimPrefix(pay.Actions[0].Value, " "), " ")

			fl := new(fileList)
			fl.Files = make([]fileInfo, len(files))
			for i := range fl.Files { // Copy all fileID into file list
				fl.Files[i].ID = files[i]
			}

			if err := delFiles(fl, token); err != nil {
				errorResp(resp, fmt.Sprintf("DelFiles Err: %v", err.Error()))
				return
			}

			createDeletedResp(resp)
			return

		} else {
			createAbortDeleteResp(resp)
			return
		}
	}

	// Normal delete file list lookup
	fl, err := getFiles(30, 10, token, user) // Get all the files and return a confirmation message to the user
	if err != nil {
		errorResp(resp, "No files match the criteria")
		return
	}
	createDeleteRequestResp(resp, fl)
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
		s.Attachments[i].CallbackID = "del_req"
		s.Attachments[i].Actions = make([]action, 2) // Yes and No
		s.Attachments[i].Actions[0].Name = "yes"
		s.Attachments[i].Actions[0].Value = fileID
		s.Attachments[i].Actions[0].Type = "button"
		s.Attachments[i].Actions[0].Text = "Yes"
		s.Attachments[i].Actions[0].Confirm.Ok_text = "Delete"
		s.Attachments[i].Actions[0].Confirm.Text = "Are you sure?"
		s.Attachments[i].Actions[1].Name = "no"
		s.Attachments[i].Actions[1].Value = "no"
		s.Attachments[i].Actions[1].Type = "button"
		s.Attachments[i].Actions[1].Text = "No"
	}

	// Dump the response
	fmt.Print(s)
}

func createDeletedResp(s *slackResp) {
	s.RespType = inchannel
	s.Text = "Deleted"

	// Dump the response
	fmt.Print(s)
}

func createAbortDeleteResp(s *slackResp) {
	s.RespType = inchannel
	s.Text = "Aborted"

	// Dump the response
	fmt.Print(s)
}
