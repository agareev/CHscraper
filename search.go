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

func (f *MetaFile) bildurl() string {
	return fileurl + strconv.Itoa(f.Name) + ".webm"
}

func (f *MetaFile) buildthumb() string {
	return thumburl + strconv.Itoa(f.Name) + "s.jpg"
}

func hasFile(number int) bool {
	if number == 0 {
		return false
	}
	return true
}

func getPosts(url string) []MetaFile {
	var processing Posts
	var output []MetaFile
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
			fileToPrepare := MetaFile{post.Tim, post.Md5, 0}
			output = append(output, fileToPrepare)
		}
	}
	return output
}
