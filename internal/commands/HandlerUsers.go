package commands

import (
	"context"
	"fmt"
)

func HandlerUsers(s *State, c Command) error {
	users, err := s.Database.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("issue getting users: %w", err)
	}

	for _, user := range users {
		name := user
		if user == s.Config.CurrentUserName {
			name += " (current)"
		}
		fmt.Println(name)
	}

	return nil
}
