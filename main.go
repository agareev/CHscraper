package main

import (
	"database/sql"
	"flag"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	catalogurl = "https://a.4cdn.org/gif/catalog.json"
	threadurl  = "https://a.4cdn.org/gif/thread/"
	fileurl    = "https://i.4cdn.org/gif/"
	thumburl   = "https://t.4cdn.org/gif/"
	datasource = "root:root@tcp(127.0.0.1:3306)/database"
	// DB is a database connector
	DB *sql.DB
)

func init() {
	log.Println("Start downloading")
	flag.StringVar(&datasource, "c", datasource, "connect")
	flag.Parse()
}

func main() {
	log.Println(datasource)
	DB, err := sql.Open("mysql", datasource)
	for _, threadNUM := range getThreadNumbers() {
		Dthreadurl := buildThreadURL(threadNUM)
		for _, i := range getPosts(Dthreadurl) {
			i.saveFile()
		}
	}
	log.Println(err)
	defer DB.Close()
}
