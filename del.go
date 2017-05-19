package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const slackfilesListURL = "https://slack.com/api/files.list"
const slackfilesDelURL = "https://slack.com/api/files.delete"

type fileList struct {
	Files []fileInfo `json:"files"`
}

type fileInfo struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Size    uint32 `json:"size"`
	Created int64  `json:"created"`
}

type okStruct struct {
	OK bool `json:"ok"`
}

func getFiles(age, count int, token string) (*fileList, error) {

	// Create time stamp
	stamp := time.Now().AddDate(0, 0, -1*age).Unix()

	vals := url.Values{}
	vals.Add("token", token)
	vals.Add("ts_to", fmt.Sprintf("%v", stamp))
	vals.Add("count", fmt.Sprintf("%v", count))
	queryURL := fmt.Sprintf("%v?%v", slackfilesListURL, vals.Encode())
	resp, err := http.Get(queryURL)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	js, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data := new(fileList)
	if err := json.Unmarshal(js, data); err != nil {
		return nil, err
	}

	return data, nil
}

func delFiles(f *fileList, token string) error {

	// Delete all files
	for _, a := range f.Files {
		ok := delFile(a.ID, token)
		if ok != nil {
			return fmt.Errorf("Failed to delete file %v", a.Title)
		}
	}

	return nil
}

func delFile(id string, token string) error {

	vals := url.Values{}
	vals.Add("token", token)
	vals.Add("file", id)
	queryURL := fmt.Sprintf("%v?%v", slackfilesDelURL, vals.Encode())
	resp, err := http.Get(queryURL)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	js, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	data := new(okStruct)
	if err := json.Unmarshal(js, data); err != nil {
		return err
	}

	if !data.OK {
		return fmt.Errorf("DELETE FAIL")
	}

	return nil
}
