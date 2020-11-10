package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
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

func generateRandString() string {
	const alpha = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := make([]byte, 6)
	for i := range bytes {
		bytes[i] = alpha[rand.Intn(len(alpha))]
	}
	return string(bytes)
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
	rand.Seed(time.Now().Unix())
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8000", nil)
}
