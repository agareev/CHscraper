package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var (
	datasource = "root:root@tcp(127.0.0.1:3306)/database"
	// DB is a database connector
	DB *sql.DB
)

// day_id = SELECT MAX(id) FROM days;
// SELECT name FROM files WHERE day_id=day_id;
// META_webms - пока не используется

func init() {
	log.Println("Start serving")
	flag.StringVar(&datasource, "c", datasource, "connect")
	flag.Parse()
	var err error
	DB, err = sql.Open("mysql", datasource)
	if err != nil {
		log.Fatal(err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}
}

func getDayNames() string {
	id := ""
	err := DB.QueryRow("SELECT MAX(id) FROM days;").Scan(&id)
	switch {
	case err == sql.ErrNoRows:
		// log.Println(f.Hash, " is uniq")
		return ""
	case err != nil:
		log.Fatal(err)
	default:
		log.Println(id, " is not iniq")
		return id
	}
	return id
}

func main() {
	http.HandleFunc("/", postHandler)
	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":3000", nil))
	defer DB.Close()
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello\n")
	fmt.Fprintf(w, getDayNames())
}
