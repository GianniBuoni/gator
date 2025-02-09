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

var addfeed CommandData = CommandData{
	name:    "addfeed",
	handler: middlewareLoggedIn(HandlerAddFeed),
}

func HandlerAddFeed(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) != 2 {
		return errors.New("addfeed takes two arguments. $1 name of the feed, $2 url of the feed.")
	}
	ctx := context.Background()
	params := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
		UserID:    user.ID,
	}
	feed, err := s.Database.CreateFeed(ctx, params)
	if err != nil {
		return err
	}
	ffParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}
	feedFollow, err := s.Database.CreateFeedFollow(ctx, ffParams)
	if err != nil {
		return err
	}
	fmt.Println()
	fmt.Println("⭐ New feed added.")
	w := lib.NewFeedTable()
	row := fmt.Sprintf("%s\t%s\t%s\t", feedFollow.FeedName, feed.Url, feedFollow.UserName)
	fmt.Fprintln(w, row)
	w.Flush()

	return nil
}
