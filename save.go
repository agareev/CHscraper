package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func (f *MetaFile) saveFile() {

	if f.checkUniq() {

		response, e := http.Get(f.bildurl())
		if e != nil {
			log.Fatal(e)
		}

		defer response.Body.Close()
		file, err := os.Create(f.buildpath())
		if err != nil {
			log.Fatal(err)
		}

		_, err = io.Copy(file, response.Body)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()

		// f.saveMeta()
		log.Println(f.bildurl(), f.buildthumb(), " downloaded!")
		return
	}

}

func (f *MetaFile) buildpath() string {
	t := time.Now().Format("2006-01-02")
	createFolder()
	path, err := filepath.Abs("files/" + t + "/" + strconv.Itoa(f.Name) + ".webm")
	if err != nil {
		log.Fatalln(err)
	}
	return path
}

func createFolder() string {
	t := time.Now().Format("2006-01-02")
	path, err := filepath.Abs("files/" + t + "/")
	if err != nil {
		log.Fatal(err)
	}
	if _, err = os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0755)
		log.Println(path, "created!")
	}
	return path
}
