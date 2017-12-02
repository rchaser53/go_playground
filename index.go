package main

import (
	f "fmt"
	"goPlayground/nyanHello"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	// nyanHello.Nyan()
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/hoge", serveTemplate)

	http.ListenAndServe(":3000", nil)
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	indexFile := "static/index.html"

	info, err := os.Stat(indexFile)
	if err != nil {
		if os.IsNotExist(err) {
			nyanHello.Nyan()
			http.NotFound(w, r)
			return
		}
	}

	if info.IsDir() {
		http.NotFound(w, r)
		return
	}

	http.ServeFile(w, r, indexFile)
}

func tryGet() string {
	resp, err := http.Get("https://google.com/")

	if err != nil {
		log.Fatal(err)
	}

	byteArray, _ := ioutil.ReadAll(resp.Body)
	bodyStr := string(byteArray)
	f.Println(bodyStr)
	defer resp.Body.Close()

	return bodyStr
}
