package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/ssgkian/gator/internal/database"
)

func scrapeFeeds(s *state) {
	fmt.Println("Getting next feed...")
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	err = s.db.MarkFeedFetched(context.Background(), nextFeed.ID)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Fetching feed:", nextFeed.Url)
	rssFeed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Feed fetched successfully")
	for _, item := range rssFeed.Channel.Item {
		publishedAt := sql.NullTime{}
		t, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err == nil {
			publishedAt = sql.NullTime{
				Time:  t,
				Valid: true,
			}
		}
		_, err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title:     item.Title,
			Url:       item.Link,
			Description: sql.NullString{
				String: item.Description,
				Valid:  true,
			},
			PublishedAt: publishedAt,
			FeedID:      nextFeed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			} else {
				log.Printf("couldn't create post: %v", err)
			}
		}
	}

}
