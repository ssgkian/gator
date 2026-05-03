package main

import (
	"context"
	"fmt"
)

func scrapeFeeds(s *state) {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	err = s.db.MarkFeedFetched(context.Background(), nextFeed.ID)
	if err != nil {
		fmt.Println(err)
	}
	rssFeed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		fmt.Println(err)
	}
	for _, item := range rssFeed.Channel.Item {
		fmt.Println(item.Title)
	}

}
