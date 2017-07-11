package main

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

func (f *MetaFile) saveFile() {

	if f.checkUniq() {
		createFolder()
		log.Println(f.bildurl(), f.buildthumb(), f.Hash, "downloaded!")
		return
	}

	// 	out, err := os.Create("output.txt")
	// defer out.Close()
	// ...
	// resp, err := http.Get("http://example.com/")
	// defer resp.Body.Close()
	// ...
	// n, err := io.Copy(out, resp.Body)

}

func createFolder() string {
	t := time.Now().Format("2006-01-02")
	path := "files/" + t + "/"
	path, err := filepath.Abs(path)
	// if e
	if _, err = os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0755)
		log.Println(path, "created!")
	}
	return path
}

func downloadMeta(post Posts) error {
	return nil
}
