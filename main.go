package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

//TODO: There's a bug that allows the same shortcut to be registerd twice. Handler doesn't like that.

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

var conf *config

func main() {
	conf, err := getConfig()
	if err != nil {
		fmt.Println("Error getting config.")
		return
	}
	//TODO: Get port from config
	rand.Seed(time.Now().Unix())
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", indexHandler)
	fmt.Println("Listening on port", conf.Port)
	http.ListenAndServe(fmt.Sprintf("%s:%d", "", conf.Port), nil)
}
