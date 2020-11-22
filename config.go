package main

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

// Config contains the app's configuration
type config struct {
	Port             int    `required:"false" split_words:"true" default:"8000"`
	DBProjectID      string `required:"true" split_words:"true"`
	DBCollectionName string `required:"true" split_words:"true"`
}

var conf config

func getConfig() (*config, error) {
	err := envconfig.Process("rebred", &conf)
	if err != nil {
		fmt.Println("Error reading config:", err)
	}
	return &conf, nil
}

// Back-End:
// TODO: Fix bug where multiple routes are allowed to be registered but the server blows up trying to register them
// TODO: Use a better logging/error handling library like Zap
// TODO: If you're on reb.red/<random characters> but youre looking at the homepage, and you try to register a shortcut, it appends the shortcut to whatever sub-URL you're already on.
// TODO: If someone tries to go to reb.red/something that doesn't exist, it needs to redirect back to /
// TODO: Shortcut should only be lowercase and uppercase
// Front-End:
// TODO: Clean up the Javascript
// TODO: Submit button should be grayed out and do nothing until the URL bar is populated
// TODO: If URL is invalid, display an error on screen
// TODO: If shortcut is already taken, display an error on screen
// TODO: Shortcut should only be lowercase and uppercase
