package main

import (
	"database/sql"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func (f *MetaFile) saveFile() {

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

	log.Println(f.bildurl(), f.buildthumb(), " downloaded!")
	return

}

func createDay(day string) int64 {
	// return 11
	stmt, err := DB.Prepare("INSERT INTO days(day) VALUES(?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(day)
	if err != nil {
		log.Fatal(err)
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	return lastID
}

func newMeta(md5 string, name int, day_id int64) {
	// SAVE MD5
	stmt, err := DB.Prepare("INSERT INTO meta_webms(md5) VALUES(?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(md5)
	if err != nil {
		log.Fatal(err)
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(lastID)
	// SAVE FILE META
	fstmt, err := DB.Prepare("INSERT INTO files(name,day_id,meta_id) VALUES(?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = fstmt.Exec(name, day_id, lastID)
	if err != nil {
		log.Fatal(err)
	}
}

func (f *MetaFile) saveMeta() {
	var dayID int64
	t := time.Now().Format("2006-01-02")

	err := DB.QueryRow("select id from days where day=?", t).Scan(&dayID)
	switch {
	case err == sql.ErrNoRows:
		dayID = createDay(t)
		log.Println(t)
		newMeta(f.Hash, f.Name, dayID)
	case err != nil:
		log.Fatal(err)
	default:
		newMeta(f.Hash, f.Name, dayID)
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

func createFolder() {
	t := time.Now().Format("2006-01-02")
	path, err := filepath.Abs("files/" + t + "/")
	if err != nil {
		log.Fatal(err)
	}
	if _, err = os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0755)
		log.Println(path, "created!")
	}
}
