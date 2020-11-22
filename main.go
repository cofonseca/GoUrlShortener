package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

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

func main() {
	// Get Config
	conf, err := getConfig()
	if err != nil {
		fmt.Println("Error getting config.")
		return
	}

	// Start Server
	rand.Seed(time.Now().Unix())
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", indexHandler)
	fmt.Println("Listening on port", conf.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), nil)
}
