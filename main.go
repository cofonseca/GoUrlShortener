package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

//TODO: There's a bug that allows the same shortcut to be registerd twice. Handler doesn't like that.

type urlMap struct {
	FullURL  string
	Shortcut string
}

var shortcut string

func generateRandString() string {
	const alpha = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := make([]byte, 6)
	for i := range bytes {
		bytes[i] = alpha[rand.Intn(len(alpha))]
	}
	return string(bytes)
}

func redirectHandler(shortcut string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		redirectURL := ReadURLMap(shortcut)
		http.Redirect(w, r, redirectURL, 302)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// Serve the homepage
		http.ServeFile(w, r, "./static/index.html")
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
			URL = urlMap.FullURL
		} else {
			fmt.Println("Missing http/https. Adding it in...")
			URL = "https://" + urlMap.FullURL
			_, err := http.Get(URL)
			if err != nil {
				fmt.Println("URL not reachable:", err)
				return
			}
			fmt.Println("Final URL:", URL)
		}

		// Create Shortcut for URL
		if urlMap.Shortcut != "" {
			shortcut = urlMap.Shortcut
			fmt.Println("User requested shortcut:", urlMap.Shortcut)
		} else {
			fmt.Println("User did not request a shortcut. Generating one...")
			shortcut = generateRandString()
			fmt.Println("Generated shortcut:", shortcut)
		}

		// Save the URL and Shortcut in a map to be read later
		// This will eventually be stored in a DB
		WriteURLMap(URL, shortcut)
		http.HandleFunc(("/" + shortcut), redirectHandler(shortcut))
		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "text/html")
		//TODO: This should be a proper JSON response instead of just a string.
		json.NewEncoder(w).Encode(shortcut)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Println("User is making a request with a method that isn't allowed.")
	}

}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	rand.Seed(time.Now().Unix())
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", indexHandler)
	fmt.Println("Listening on port", port)
	http.ListenAndServe((":" + port), nil)
}
