//+build AWS

package main

import "testing"

func TestCreateDeleteRequestResp(t *testing.T) {

	s := new(slackResp)
	fl := new(fileList)
	fl.Files = make([]fileInfo, 3)
	for i := range fl.Files {
		switch i {
		case 0:
			fl.Files[i].Title = "James and the Giant Peach"
		case 1:
			fl.Files[i].Title = "The Hunt for Red October"
		case 2:
			fl.Files[i].Title = "Enders Game"
		default:
			fl.Files[i].Title = "Unknown"
		}
	}
	createDeleteRequestResp(s, fl)
}
