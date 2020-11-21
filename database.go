package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
)

// WriteURLMap saves a mapping of Shortcut to URL
func WriteURLMap(URL string, Shortcut string) {
	ctx := context.Background()

	client, err := firestore.NewClient(ctx, conf.DBProjectID)
	if err != nil {
		//TODO: Kill the app if we can't connect to DB.
		fmt.Println("Can't create Firestore client:", err)
	}

	defer client.Close()

	fmt.Println("WRITING TO DB")
	_, err = client.Collection(conf.DBCollectionName).Doc(Shortcut).Create(ctx, map[string]interface{}{
		"FullURL":  URL,
		"Shortcut": Shortcut,
	})
	if err != nil {
		fmt.Println("Error writing to DB:", err)
	}
}

// ReadURLMap returns the full URL for a given shortcut
func ReadURLMap(Shortcut string) string {
	ctx := context.Background()

	client, err := firestore.NewClient(ctx, conf.DBProjectID)
	if err != nil {
		//TODO: Kill the app if we can't connect to DB.
		fmt.Println("Can't create Firestore client:", err)
	}

	defer client.Close()

	fmt.Println("READING FROM DB")
	// if this returns an error, then the shortcut doesn't exist, and we should handle that appropriately.
	doc, err := client.Collection(conf.DBCollectionName).Doc(Shortcut).Get(ctx)
	if err != nil {
		fmt.Println("Error getting the document:", err)
	}
	// only run this block if the block above was successful
	result, err := doc.DataAt("FullURL")
	if err != nil {
		fmt.Println("Error reading the document:", err)
	}
	fmt.Println(result)
	return fmt.Sprintf("%s", result)
}
