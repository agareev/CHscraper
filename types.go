package main

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

// MetaFile is ready for download file
type MetaFile struct {
	Name int
	Hash string
}
