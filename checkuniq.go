package main

import "log"

func (f *MetaFile) checkUniq() bool {
	rows, err := DB.Query("select * from meta_webms where md5=? LIMIT 1;", f.Hash)
	if err != nil {
		log.Println(err, "checked!")
		return true
	}
	defer rows.Close()
	log.Println("wqr")
	return false
}
