package main

import (
	"fmt"
)

// This map is temporary until I wire up my DB.
var db = make(map[string]string)

// WriteURLMap saves a mapping of Shortcut to URL
func WriteURLMap(URL string, Shortcut string) {
	db[Shortcut] = URL
	fmt.Println("Added element to map.")
}

// ReadURLMap returns the full URL for a given shortcut
func ReadURLMap(Shortcut string) string {
	URL, found := db[Shortcut]
	if found == false {
		fmt.Println("Error finding element in map:", found)
	}
	return URL
}
