package main

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

func saveFile(name int, hash string) {
	// 	out, err := os.Create("output.txt")
	// defer out.Close()
	// ...
	// resp, err := http.Get("http://example.com/")
	// defer resp.Body.Close()
	// ...
	// n, err := io.Copy(out, resp.Body)

	if checkUniq(hash) {
		createFolder()
		log.Println(buldFileURL(name), buildThumbURL(name), hash, "downloaded!")
		return
	}

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
