//+build !AWS, !GCE

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"
)

var age = flag.Int("age", 30, "Age of flags")
var count = flag.Int("c", 100, "Number of files to return")

func init() {
	flag.Parse()
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
}

func main() {

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		return
	}

	token := viper.GetString("token")

	data, err := getFiles(*age, *count, token)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(data.Files) == 0 {
		fmt.Println("Nothing to delete")
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
		if ok := delFiles(data, token); ok != nil {
			fmt.Println(ok)
		} else {
			fmt.Printf("Successfully deleted %v files\n", len(data.Files))
		}
	} else {
		fmt.Println("Aborting")
	}
}
