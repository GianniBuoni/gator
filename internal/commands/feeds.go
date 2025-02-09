package commands

import (
	"context"
	"fmt"

	"github.com/GianniBuoni/gator/internal/lib"
)

var feeds CommandData = CommandData{
	name:    "feeds",
	handler: HandlerFeeds,
}

func HandlerFeeds(s *State, cmd Command) error {
	feedList, err := s.Database.GetFeeds(context.Background())
	if err != nil {
		return err
	}
	fmt.Println()
	w := lib.NewFeedTable()
	for _, feed := range feedList {
		row := fmt.Sprintf("%s\t%s\t%s\t", feed.Name, feed.Url, feed.User)
		fmt.Fprintln(w, row)
	}
	w.Flush()

	return nil
}
