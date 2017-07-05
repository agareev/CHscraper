package main

import (
	"log"
	"os"
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
		log.Println(buldFileURL(name), hash, " downloaded!")
		return
	}

}

func createFolder() string {
	t := time.Now()
	path := "files/" + t.Format("2006-01-02") + "/"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0755)
		t = t.Add(-24 * time.Hour)
	}
	return path
}
