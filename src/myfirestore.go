package main

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type Post struct {
	PostID      string `json:"postID"`
	Title       string `json:"title"`
	Src         string `json:"src"`
	MyType      string `json:"myType"`
	Description string `json:"description"`
	Link        string `json:"link"`
}

func createClient(ctx context.Context) *firestore.Client {
	// Sets your Google Cloud Platform project ID.
	projectID := "alvardevlp07"
	databaseID := "site"

	client, err := firestore.NewClientWithDatabase(ctx, projectID, databaseID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	// Close client when done with
	// defer client.Close()
	return client
}

func getFirestorePosts(ctx context.Context) []Post {
	client := createClient(ctx)
	defer client.Close()

	posts := []Post{}
	iter := client.Collection("posts").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}

		var p Post
		doc.DataTo(&p)
		p.PostID = doc.Ref.ID
		posts = append(posts, p)
	}

	return posts
}
