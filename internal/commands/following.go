package commands

import (
	"context"
	"fmt"

	"github.com/GianniBuoni/gator/internal/database"
	"github.com/GianniBuoni/gator/internal/lib"
)

var following CommandData = CommandData{
	name:    "following",
	handler: middlewareLoggedIn(HandlerFollowing),
}

func HandlerFollowing(s *State, cmd Command, user database.User) error {
	ctx := context.Background()
	feeds, err := s.Database.GetFeedFollowsForUser(ctx, user.Name)
	if err != nil {
		return fmt.Errorf("issue getting user feeds: %w", err)
	}
	fmt.Println()
	w := lib.NewFeedTable()
	for _, feed := range feeds {
		row := fmt.Sprintf("%s\t%s\t%s\t", feed.FeedName, feed.FeedUrl, feed.UserName)
		fmt.Fprintln(w, row)
	}
	w.Flush()
	return nil
}
