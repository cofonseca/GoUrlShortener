package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
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
		// If the user requests a route other than /,
		// check for that route in the DB and redirect to the full URL if it exists,
		// else, redirect the user back to the homepage.
		if r.URL.Path != "/" && r.URL.Path != "/favicon.ico" {
			validPath := regexp.MustCompile(`^[a-zA-Z]+$`).MatchString
			path := strings.Replace(r.URL.Path, "/", "", 1)
			if validPath(path) {
				redirectURL := ReadURLMap(path)
				if redirectURL != "" {
					http.Redirect(w, r, redirectURL, 302)
				} else {
					fmt.Println(path, "is valid, but doesn't exist in the DB. Redirecting to /")
					http.Redirect(w, r, "/", 302)
				}
			} else {
				fmt.Println(path, "includes non-letter characters and is invalid. Redirecting to /")
				http.Redirect(w, r, "/", 302)
			}
		} else {
			http.ServeFile(w, r, "./static/index.html")
		}
		return

	case "POST":
		// Unmarshal JSON
		body, _ := ioutil.ReadAll(r.Body)
		var urlMap urlMap
		json.Unmarshal(body, &urlMap)

		// Validate URL from User
		if urlMap.FullURL == "" {
			http.Error(w, "URL cannot be empty", 400)
		} else {
			if strings.Contains(urlMap.FullURL, "http://") || strings.Contains(urlMap.FullURL, "https://") {
				_, err := http.Get(urlMap.FullURL)
				if err != nil {
					fmt.Println("URL not reachable:", err)
					http.Error(w, "URL is not reachable", 400)
					return
				}
				fmt.Println("User requested URL:", urlMap.FullURL)
			} else {
				fmt.Println("Missing http/https. Adding it in...")
				urlMap.FullURL = "https://" + urlMap.FullURL
				_, err := http.Get(urlMap.FullURL)
				if err != nil {
					fmt.Println("URL not reachable:", err)
					http.Error(w, "URL is not reachable", 400)
					return
				}
				fmt.Println("User requested URL:", urlMap.FullURL)
			}

			// Create Shortcut for URL
			if urlMap.Shortcut != "" {
				fmt.Println("User requested shortcut:", urlMap.Shortcut)
			} else {
				fmt.Println("User did not request a shortcut. Generating one...")
				urlMap.Shortcut = generateRandString()
				fmt.Println("Generated shortcut:", urlMap.Shortcut)
			}
		}

		// Save the URL and Shortcut in a map to be read later
		// This will eventually be stored in a DB
		success := WriteURLMap(urlMap.FullURL, urlMap.Shortcut)
		if !success {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]interface{}{
				"Success":  success,
				"Shortcut": "",
			})
		} else {
			http.HandleFunc(("/" + urlMap.Shortcut), redirectHandler(urlMap.Shortcut))
			w.WriteHeader(http.StatusOK)
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]interface{}{
				"Success":  success,
				"Shortcut": urlMap.Shortcut,
			})
		}

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Println("User is making a request with a method that isn't allowed.")
	}

}
