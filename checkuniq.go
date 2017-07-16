package main

import (
	"database/sql"
	"log"
)

func (f *MetaFile) checkUniq() bool {
	var id int64
	err := DB.QueryRow("select id from meta_webms where md5=?", f.Hash).Scan(&id)
	switch {
	case err == sql.ErrNoRows:
		// log.Println(f.Hash, " is uniq")
		return true
	case err != nil:
		log.Fatal(err)
	default:
		log.Println(id, " is not iniq")
		return false
	}

	// update row popularity
	return false
}
