package main

import (
	"fmt"
	f "fmt"
	"goPlayground/nyanHello"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
)

func main() {
	str := `<a href=""></a>
	<a href=""></a>
	<a href=""></a>
	`
	rep := regexp.MustCompile(`(?m)href=""`)
	str = rep.ReplaceAllString(str, "nya-n")

	println(str) //

}

func hoge() {
	fmt.Println(123)
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

func serveGitRequest(w http.ResponseWriter, r *http.Request) {
	jsonStr := tryGet("https://api.github.com/repos/rchaser53/vue-table-playground/commits")
	w.Header().Set("Content-Type", "application/json")
	f.Fprint(w, jsonStr)
}

func tryGet(url string) string {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	byteArray, _ := ioutil.ReadAll(resp.Body)
	bodyStr := string(byteArray)
	// f.Println(bodyStr)
	defer resp.Body.Close()

	return bodyStr
}
