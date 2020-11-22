package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
)

// WriteURLMap saves a mapping of Shortcut to URL
func WriteURLMap(URL string, Shortcut string) bool {
	ctx := context.Background()

	client, err := firestore.NewClient(ctx, conf.DBProjectID)
	if err != nil {
		//TODO: Kill the app if we can't connect to DB.
		fmt.Println("Can't create Firestore client:", err)
	}

	_, err = client.Collection(conf.DBCollectionName).Doc(Shortcut).Create(ctx, map[string]interface{}{
		"FullURL":  URL,
		"Shortcut": Shortcut,
	})
	if err != nil {
		fmt.Println("Error writing to DB:", err)
		return false
	}

	defer client.Close()

	return true
}

// ReadURLMap returns the full URL for a given shortcut
func ReadURLMap(Shortcut string) string {
	ctx := context.Background()

	client, err := firestore.NewClient(ctx, conf.DBProjectID)
	if err != nil {
		//TODO: Kill the app if we can't connect to DB.
		fmt.Println("Can't create Firestore client:", err)
	}

	var result interface{}
	doc, err := client.Collection(conf.DBCollectionName).Doc(Shortcut).Get(ctx)
	fmt.Println(doc)
	if err != nil {
		fmt.Println("Error getting the document:", err)
		return ""
	} else {
		result, err = doc.DataAt("FullURL")
		fmt.Println(result)
		if err != nil {
			fmt.Println("Error reading the document:", err)
			return ""
		}
	}

	defer client.Close()
	fmt.Println(result)
	return fmt.Sprintf("%s", result)

}
