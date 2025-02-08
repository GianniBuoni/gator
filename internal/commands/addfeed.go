package commands

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/GianniBuoni/gator/internal/database"
	"github.com/google/uuid"
)

var addfeed CommandData = CommandData{
	name:    "addfeed",
	handler: HandlerAddFeed,
}

func HandlerAddFeed(s *State, cmd Command) error {
	if len(cmd.Args) != 2 {
		return errors.New("addfeed takes two arguments. $1 name of the feed, $2 url of the feed.")
	}
	ctx := context.Background()
	user, err := s.Database.GetUser(ctx, s.Config.CurrentUserName)
	if err != nil {
		return err
	}
	params := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
		UserID:    user.ID,
	}
	feed, err := s.Database.CreateFeed(ctx, params)
	fmt.Printf("⭐ Feed for %s added! %s : %s", user.Name, feed.Name, feed.Url)
	return nil
}
