package lib

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/GianniBuoni/gator/internal/database"
)

func ScrapeFeeds(store *database.Queries) error {
	ctx := context.Background()
	// get next feed
	feed, err := store.NextFeedToFetch(ctx)
	if err != nil {
		return fmt.Errorf("issue getting next feed: %w", err)
	}
	// mark feed as fetched
	markedParams := database.MarkFeedFetchedParams{
		UpdatedAt:     time.Now(),
		LastFetchedAt: sql.NullTime{Time: time.Now()},
		ID:            feed.ID,
	}
	err = store.MarkFeedFetched(ctx, markedParams)
	if err != nil {
		return fmt.Errorf("issue updating feed fetch time: %w", err)
	}
	// fetch feed data
	fmt.Println()
	fmt.Println(feed.Name)
	fmt.Println("------")
	feedData, err := FetchFeed(ctx, feed.Url)
	if err != nil {
		return fmt.Errorf("issue getting feed data: %w", err)
	}
	for _, rssItem := range feedData.Channel.Item {
		fmt.Println(rssItem.Title)
	}
	return nil
}
