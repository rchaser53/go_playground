package main

import (
	f "fmt"
	nyanHello "goPlayground/nyanHello"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	nyanHello.Nyan()

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})
	http.ListenAndServe(":3000", nil)
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
