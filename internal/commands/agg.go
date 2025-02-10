package commands

import (
	"fmt"
	"time"

	"github.com/GianniBuoni/gator/internal/lib"
)

var agg CommandData = CommandData{
	name:    "agg",
	handler: HandlerAgg,
}

func HandlerAgg(s *State, cmd Command) error {
	var duration string
	if len(cmd.Args) != 1 {
		duration = "30s"
	} else {
		duration = cmd.Args[0]
	}
	timeBetweenReq, err := time.ParseDuration(duration)
	if err != nil {
		return fmt.Errorf("issue parsing '%s': %w", cmd.Args[0], err)
	}
	ticker := time.NewTicker(timeBetweenReq)
	defer ticker.Stop()
	for ; ; <-ticker.C {
		err := lib.ScrapeFeeds(s.Database)
		if err != nil {
			return err
		}
	}
}
