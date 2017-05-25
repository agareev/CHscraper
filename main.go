package main

import "log"

var (
	catalogurl = ""
	threadurl  = ""
	fileurl    = ""
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

// Posts is a container for messages or/and files
type Posts struct {
	Post []struct {
		Number   int    `json:"no"`
		Filename string `json:"filename"`
		Ext      string `json:"ext"`
		Tim      int    `json:"tim"`
	} `json:"posts"`
}

func init() {
	log.Println("Start downloading")
}

func main() {
	for _, threadNUM := range getThreadNumbers() {
		Dthreadurl := buildThreadURL(threadNUM)
		for _, downloadURL := range getPosts(Dthreadurl) {
			log.Println(buldFileURL(downloadURL))
		}
	}

}
