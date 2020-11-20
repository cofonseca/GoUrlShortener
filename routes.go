package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var shortcut string

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
		if strings.Contains(urlMap.FullURL, "http://") || strings.Contains(urlMap.FullURL, "https://") {
			_, err := http.Get(urlMap.FullURL)
			if err != nil {
				fmt.Println("URL not reachable:", err)
				return
			}
			urlMap.FullURL = urlMap.FullURL
			fmt.Println("Final URL:", urlMap.FullURL)
		} else {
			fmt.Println("Missing http/https. Adding it in...")
			urlMap.FullURL = "https://" + urlMap.FullURL
			_, err := http.Get(urlMap.FullURL)
			if err != nil {
				fmt.Println("URL not reachable:", err)
				return
			}
			fmt.Println("Final URL:", urlMap.FullURL)
		}

		// Create Shortcut for URL
		if urlMap.Shortcut != "" {
			fmt.Println("User requested shortcut:", urlMap.Shortcut)
		} else {
			fmt.Println("User did not request a shortcut. Generating one...")
			urlMap.Shortcut = generateRandString()
			fmt.Println("Generated shortcut:", urlMap.Shortcut)
		}

		// Save the URL and Shortcut in a map to be read later
		// This will eventually be stored in a DB
		WriteURLMap(urlMap.FullURL, urlMap.Shortcut)
		http.HandleFunc(("/" + urlMap.Shortcut), redirectHandler(urlMap.Shortcut))
		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "text/html")
		//TODO: This should be a proper JSON response instead of just a string.
		json.NewEncoder(w).Encode(urlMap.Shortcut)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Println("User is making a request with a method that isn't allowed.")
	}

}
