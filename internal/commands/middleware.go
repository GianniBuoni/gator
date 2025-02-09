package commands

import (
	"context"

	"github.com/GianniBuoni/gator/internal/database"
)

func middlewareLoggedIn(
	handler func(s *State, cmd Command, user database.User) error,
) func(*State, Command) error {
	return func(s *State, c Command) error {
		ctx := context.Background()
		user, err := s.Database.GetUser(ctx, s.Config.CurrentUserName)
		if err != nil {
			return err
		}
		err = handler(s, c, user)
		if err != nil {
			return err
		}
		return nil
	}
}
