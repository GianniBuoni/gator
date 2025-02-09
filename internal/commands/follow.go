package commands

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/GianniBuoni/gator/internal/database"
	"github.com/GianniBuoni/gator/internal/lib"
	"github.com/google/uuid"
)

var follow CommandData = CommandData{
	name:    "follow",
	handler: HandlerFollow,
}

func HandlerFollow(s *State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return errors.New("follow takes one argument. $1, url.")
	}
	ctx := context.Background()
	feed, err := s.Database.GetFeed(ctx, cmd.Args[0])
	if err != nil {
		return fmt.Errorf("'%s' issue getting feed: %w\n", cmd.Args[0], err)
	}
	user, err := s.Database.GetUser(ctx, s.Config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("issue getting current user: %w\n", err)
	}
	params := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedID:    feed.ID,
		UserID:    user.ID,
	}
	feedFollow, err := s.Database.CreateFeedFollow(ctx, params)
	fmt.Println()
	w := lib.NewFeedTable()
	row := fmt.Sprintf("%s\t%s\t%s\t", feedFollow.FeedName, feed.Url, feedFollow.UserName)
	fmt.Fprintln(w, row)
	w.Flush()

	return nil
}
