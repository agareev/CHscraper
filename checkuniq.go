package main

import "log"

func checkUniq(md5 string) bool {
	rows, err := db.Query("select md5 from meta_webms count 1 where md5=?", md5)
	if err != nil {
		log.Println(err, "checked!")
		return true
	}
	log.Println(rows, "wqr")
	return false
}
