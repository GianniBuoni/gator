package lib

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/GianniBuoni/gator/internal/database"
	"github.com/google/uuid"
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
	fmt.Printf("Fetched from '%s'\n", feed.Name)
	feedData, err := FetchFeed(ctx, feed.Url)
	if err != nil {
		return fmt.Errorf("issue getting feed data: %w", err)
	}
	for _, rssItem := range feedData.Channel.Item {
		publishedAt, err := time.Parse(time.RFC1123Z, rssItem.PubDate)
		if err != nil {
			return err
		}
		postParam := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       rssItem.Title,
			Description: rssItem.Description,
			Url:         rssItem.Link,
			PublishedAt: publishedAt,
			FeedID:      feed.ID,
		}
		_, err = store.CreatePost(ctx, postParam)
		if err != nil {
			return fmt.Errorf("issue creating post: %w", err)
		}
	}
	return nil
}
