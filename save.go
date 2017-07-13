package main

import "log"

func saveFile(name int, hash string) {
	log.Println(buldFileURL(name), hash, " downloaded!")
	// 	out, err := os.Create("output.txt")
	// defer out.Close()
	// ...
	// resp, err := http.Get("http://example.com/")
	// defer resp.Body.Close()
	// ...
	// n, err := io.Copy(out, resp.Body)
}

func createFolder() {
	log.Println("test")
}
