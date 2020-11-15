package main

import (
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
		return

	case "POST":
		// sanitize the URL
		body, _ := ioutil.ReadAll(r.Body)
		fmt.Println("raw body:", string(body))
		rawURL := strings.Split(string(body), "=")[1]
		fmt.Println("rawURL:", rawURL)
		var URL string
		if strings.Contains(rawURL, "http://") || strings.Contains(rawURL, "https://") {
			_, err := http.Get(rawURL)
			if err != nil {
				fmt.Println("URL not reachable:", err)
				return
			}
			fmt.Println("Final URL:", rawURL)
		} else {
			fmt.Println("Missing http/https. Adding it in...")
			URL = "https://" + rawURL
			_, err := http.Get(URL)
			if err != nil {
				fmt.Println("URL not reachable:", err)
				return
			}
			fmt.Println("Final URL:", URL)
		}
		fmt.Println("We good!")
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
