package commands

import (
	"context"
	"fmt"
)

var reset CommandData = CommandData{
	name:    "reset",
	handler: HandlerReset,
}

func HandlerReset(s *State, c Command) error {
	err := s.Database.Reset(context.Background())
	if err != nil {
		return fmt.Errorf("issue resetting db: %w", err)
	}

	fmt.Println("💀 Database reset")

	return nil
}
