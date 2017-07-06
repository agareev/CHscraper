package main

var (
	catalogurl = "https://a.4cdn.org/gif/catalog.json"
	threadurl  = "https://a.4cdn.org/gif/thread/"
	fileurl    = "https://i.4cdn.org/gif/"
	thumburl   = "https://t.4cdn.org/gif/"
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

// Post is a container for messages or/and files
type Post struct {
	Number   int    `json:"no"`
	Filename string `json:"filename"`
	Ext      string `json:"ext"`
	Tim      int    `json:"tim"`
	Md5      string `json:"md5"`
}

// Posts is a container for post
type Posts struct {
	Post []Post `json:"posts"`
}

func init() {
	// log.Println("Start downloading")
}

func main() {
	for _, threadNUM := range getThreadNumbers() {
		Dthreadurl := buildThreadURL(threadNUM)
		for downloadURL, hash := range getPosts(Dthreadurl) {
			saveFile(buldFileURL(downloadURL), hash)
		}
	}
}
