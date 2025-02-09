package commands

import (
	"context"
	"fmt"

	"github.com/GianniBuoni/gator/internal/lib"
)

var agg CommandData = CommandData{
	name:    "agg",
	handler: HandlerAgg,
}

func HandlerAgg(s *State, cmd Command) error {
	feed, err := lib.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("FetchFeed: %w", err)
	}
	fmt.Println(*feed)
	return nil
}
