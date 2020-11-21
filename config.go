package main

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type config struct {
	Port             int    `required:"false" split_words:"true" default:"8000"`
	DBProjectID      string `required:"true" split_words:"true"`
	DBCollectionName string `required:"true" split_words:"true"`
}

func getConfig() (*config, error) {
	var conf config
	err := envconfig.Process("rebred", &conf)
	if err != nil {
		fmt.Println("Error reading config:", err)
		return nil, err
	}
	return &conf, nil
}
