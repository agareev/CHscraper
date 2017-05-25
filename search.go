package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

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
	return threadurl + strconv.Itoa(number) + ".json"
}

func buldFileURL(number int) string {
	return fileurl + strconv.Itoa(number) + ".webm"
}

func hasFile(number int) bool {
	if number == 0 {
		return false
	}
	return true
}

func getPosts(url string) []int {
	var processing Posts
	var output []int
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
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
	for _, post := range processing.Post {
		if hasFile(post.Tim) {
			output = append(output, post.Tim)
		}
	}
	return output
}
