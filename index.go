package main

import (
	f "fmt"
	"goPlayground/nyanHello"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

func main() {
	// nyanHello.Nyan()
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/hoge", serveTemplate)
	http.HandleFunc("/git", serveGitRequest)

	r := chi.NewRouter()
	r.Route("/nest/", func(r chi.Router) {
		r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("welcome"))
		})
		r.Get("/gyan", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("ma"))
		})
	})

	http.ListenAndServe(":3000", r)
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
