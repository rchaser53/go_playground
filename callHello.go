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

	resp, err := http.Get("https://google.com/")

	if err != nil {
		log.Fatal(err)
	}

	byteArray, _ := ioutil.ReadAll(resp.Body)
	f.Println(string(byteArray))
	defer resp.Body.Close()
}
