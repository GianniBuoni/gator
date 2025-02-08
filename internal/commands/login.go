package commands

import (
	"context"
	"errors"
	"fmt"
)

var login CommandData = CommandData{
	name:    "login",
	handler: HandlerLogin,
}

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return errors.New("login extects a single argument: username")
	}

	name := cmd.Args[0]

	_, err := s.Database.GetUser(context.Background(), name)
	if err != nil {
		return fmt.Errorf("%s not found. Please register name.", name)
	}

	if err := s.Config.SetUser(name); err != nil {
		return err
	}

	fmt.Printf("%s logged in!", name)
	return nil
}
