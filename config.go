package main

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

// Config contains the app's configuration
type Config struct {
	Port             int    `required:"false" split_words:"true" default:"8000"`
	DBProjectID      string `required:"true" split_words:"true"`
	DBCollectionName string `required:"true" split_words:"true"`
}

var conf Config

func getConfig() (*Config, error) {
	err := envconfig.Process("rebred", &conf)
	if err != nil {
		fmt.Println("Error reading config:", err)
	}
	return &conf, nil
}

// Back-End:
// TODO: Fix bug where multiple routes are allowed to be registered but the server blows up trying to register them
// TODO: Use a better logging/error handling library like Zap
// Front-End:
// TODO: Clean up the Javascript
// TODO: Submit button should be grayed out and do nothing until the URL bar is populated
// TODO: If URL is invalid, display an error on screen
// TODO: If shortcut is already taken, display an error on screen
