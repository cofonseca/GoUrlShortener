package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

// this should be an API

// User will paste in a long URL
// Give user option to create their own short URL sub-path
// If the sub-path is not available, pop an error on the page and let them know
// Optionally, just randomly generate a sub-path
// We need to store all of these sub-paths as well as the URLs that they should go to
// We need a router that can read these sub-paths from a db and  redirect to the real URL

type urlMap struct {
	FullURL  string
	Shortcut string
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
		return

	case "POST":
		// Unmarshal JSON
		body, _ := ioutil.ReadAll(r.Body)
		var urlMap urlMap
		json.Unmarshal(body, &urlMap)
		fmt.Println("rawURL:", urlMap.FullURL)

		// Validate URL from User
		var URL string
		if strings.Contains(urlMap.FullURL, "http://") || strings.Contains(urlMap.FullURL, "https://") {
			_, err := http.Get(urlMap.FullURL)
			if err != nil {
				fmt.Println("URL not reachable:", err)
				return
			}
			fmt.Println("Final URL:", urlMap.FullURL)
		} else {
			fmt.Println("Missing http/https. Adding it in...")
			URL = "https://" + urlMap.FullURL
			_, err := http.Get(URL)
			if err != nil {
				fmt.Println("URL not reachable:", err)
				return
			}
			fmt.Println("Final URL:", URL)
			// I learned how to split my project into multiple files! (see below)
			// printSomething("test")
		}
		if urlMap.Shortcut != "" {
			fmt.Println("User requested shortcut:", urlMap.Shortcut)
		} else {
			fmt.Println("User did not request a shortcut. Generating one...")
			shortcut := generateRandString()
			fmt.Println("Generated shortcut:", shortcut)
		}

		// check the short url that the user entered, and see if it's already in use.
		// if it is already in use
		// display an error on the screen,
		//and give the user the option to pick a new one or have the app generate a new one automatically
		// if not, create a new mapping between the short url and the long url
		// then, save that in the db and create a new handler

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}

func main() {
	rand.Seed(time.Now().Unix())
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8000", nil)
}
