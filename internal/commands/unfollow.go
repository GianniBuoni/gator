package commands

import (
	"context"
	"errors"
	"fmt"

	"github.com/GianniBuoni/gator/internal/database"
	"github.com/GianniBuoni/gator/internal/lib"
)

var unfollow CommandData = CommandData{
	name:    "unfollow",
	handler: middlewareLoggedIn(HandlerUnfollow),
}

func HandlerUnfollow(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) != 1 {
		return errors.New("expects one argument: $1, feed url")
	}
	// check feed exists
	ctx := context.Background()
	getFfParams := database.GetFeedFollowParams{
		Name: user.Name,
		Url:  cmd.Args[0],
	}
	feedFollow, err := s.Database.GetFeedFollow(ctx, getFfParams)
	if err != nil {
		return fmt.Errorf("issue finding feed for user: %w", err)
	}
	fmt.Println()
	fmt.Println("Deleting...")
	w := lib.NewFeedTable()
	row := fmt.Sprintf("%s\t%s\t%s\t", feedFollow.FeedName, feedFollow.FeedUrl, feedFollow.UserName)
	fmt.Fprintln(w, row)
  w.Flush()
	err = s.Database.DeleteFeedFollow(ctx, feedFollow.ID)
	if err != nil {
		return fmt.Errorf("issue deleting feed: %w", err)
	}
	return nil
}
