package commands

import (
	"context"
	"fmt"
	"regexp"
	"strconv"

	"github.com/GianniBuoni/gator/internal/database"
)

var browse CommandData = CommandData{
	name:    "browse",
	handler: middlewareLoggedIn(handlerBrowse),
}

func handlerBrowse(s *State, cmd Command, user database.User) error {
	var limit int64
	limit = 2
	if len(cmd.Args) == 1 {
		var err error
		limit, err = strconv.ParseInt(cmd.Args[0], 10, 64)
		if err != nil {
			return err
		}
	}

	ctx := context.Background()
	ags := database.GetPostsForUserParams{
		ID:    user.ID,
		Limit: int32(limit),
	}
	posts, err := s.Database.GetPostsForUser(ctx, ags)
	if err != nil {
		return err
	}

	for _, post := range posts {
		fmt.Println()
		fmt.Println(post.Title)
		fmt.Println(post.PublishedAt)
		match := regexp.MustCompile(`\<.+\>`)
		desc := match.ReplaceAllString(post.Description[:80], "") + "..."
		fmt.Println(desc)
	}

	return nil
}
