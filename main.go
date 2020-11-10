package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// this should be an API

// User will paste in a long URL
// Give user option to create their own short URL sub-path
// If the sub-path is not available, pop an error on the page and let them know
// Optionally, just randomly generate a sub-path
// We need to store all of these sub-paths as well as the URLs that they should go to
// We need a router that can read these sub-paths from a db and  redirect to the real URL

type redirect struct {
	FullURL string
	SubPath string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// serve the home page
		http.ServeFile(w, r, "index.html")
	case "POST":
		body, _ := ioutil.ReadAll(r.Body)
		fmt.Println(string(body))
	}
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8000", nil)
}
