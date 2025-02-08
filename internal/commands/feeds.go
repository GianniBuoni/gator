package commands

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"
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
	padding := 3
	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', tabwriter.Debug)
	fmt.Fprintln(w, "NAME\tURL\tUSER\t")
	for _, feed := range feedList {
		row := fmt.Sprintf("%s\t%s\t%s\t", feed.Name, feed.Url, feed.User)
		fmt.Fprintln(w, row)
	}
	w.Flush()

	return nil
}
