package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	catalogurl = ""
	threadurl  = ""
)

// Thread is a container for posts
type Thread struct {
	Number int `json:"no"`
}

// Page is a container for threads
type Page struct {
	Page    int      `json:"page"`
	Threads []Thread `json:"threads"`
}

func init() {
	fmt.Println("asd")
}

func getThreadNumbers() []int {
	var processing []Page
	var output []int
	client := &http.Client{}
	req, _ := http.NewRequest("GET", catalogurl, nil)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("couldn't connect ", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("couldn't read ", err)
	}
	json.Unmarshal(body, &processing)
	for _, page := range processing {
		for _, thread := range page.Threads {
			output = append(output, thread.Number)
		}
	}
	return output
}

func buildThreadURL(number int) string {
	return threadurl + string(number)
}

func getWebmPosts(url string) []string {

}

func main() {
	for _, threads := range getThreadNumbers() {
		fmt.Println(buildThreadURL(threads))
	}
}
