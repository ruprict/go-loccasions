package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
)

func main() {

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", serveHtml)
	fmt.Println("**** Starting client *** ")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func serveHtml(w http.ResponseWriter, req *http.Request) {
	fp := path.Join("views", req.URL.Path)

	fmt.Println(fp)
	// Return a 404 if the template doesn't exist
	_, err := os.Stat(fp)
	if err != nil {
		if os.IsNotExist(err) {
			http.NotFound(w, req)
			return
		}
	}

	http.ServeFile(w, req, fp)
}
