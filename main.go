package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const toke = "ADD TOKEN"
const slackfilesListURL = "https://slack.com/api/files.list"
const slackfilesDelURL = "https://slack.com/api/files.delete"

var age = flag.Int("age", 30, "Age of flags")
var count = flag.Int("c", 100, "Number of files to return")

type fileList struct {
	Files []fileInfo `json:"files"`
}

type fileInfo struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Size    uint32 `json:"size"`
	Created int64  `json:"created"`
}

func init() {
	flag.Parse()
}

func main() {

	data, err := getFiles(*age, *count)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, f := range data.Files {
		fmt.Printf(" %v \t%v KiB Created: %v\n", f.Title, f.Size/1024, time.Unix(f.Created, 0))
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Confirm Delete: ")
	var delete bool
	for scanner.Scan() {
		text := scanner.Text()
		if strings.EqualFold(text, "ok") || strings.EqualFold(text, "y") {
			delete = true
		}
		break
	}

	if delete {
		if ok := delFiles(data); ok != nil {
			fmt.Println(ok)
		} else {
			fmt.Printf("Successfully deleted %v files\n", len(data.Files))
		}
	} else {
		fmt.Println("Aborting")
	}
}

func getFiles(age, count int) (*fileList, error) {

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

func delFiles(f *fileList) error {

	// Delete all files
	for _, a := range f.Files {
		ok := delFile(a.ID)
		if ok != nil {
			return fmt.Errorf("Failed to delete file %v", a.Title)
		}
	}

	return nil
}

type okStruct struct {
	OK bool `json:"ok"`
}

func delFile(id string) error {

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
