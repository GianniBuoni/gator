package rss

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	// create request
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("issue creating client request: %w", err)
	}
	req.Header.Set("User-Agent", "gator")

	// make request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("issue making http request: %w", err)
	}
	defer res.Body.Close()

	// process response
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("status code %s", res.Status)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("issue reading response body: %w", err)
	}
	feed := &RSSFeed{}
	err = xml.Unmarshal(body, feed)
	if err != nil {
		return nil, fmt.Errorf("issue unmarshaling response body %w", err)
	}

	// process html escape sequences
	html.UnescapeString(feed.Channel.Title)
	html.UnescapeString(feed.Channel.Description)
	for _, item := range feed.Channel.Item {
		html.UnescapeString(item.Title)
		html.UnescapeString(item.Description)
	}
	return feed, nil
}
