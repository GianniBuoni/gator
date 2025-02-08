package commands

import (
	"context"
	"fmt"

	"github.com/GianniBuoni/gator/internal/rss"
)

func HandlerAgg(s *State, cmd Command) error {
	feed, err := rss.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("FetchFeed: %w", err)
	}
	fmt.Println(*feed)
	return nil
}
