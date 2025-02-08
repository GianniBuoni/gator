package commands

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/GianniBuoni/gator/internal/database"
	"github.com/google/uuid"
)

var register CommandData = CommandData{
	name:    "register",
	handler: HandlerRegister,
}

func HandlerRegister(s *State, c Command) error {
	if len(c.Args) < 0 {
		return errors.New("register expects an argument: name")
	}

	name := c.Args[0]
	ctx := context.Background()

	if _, err := s.Database.GetUser(
		ctx, name,
	); err == nil {
		return fmt.Errorf("Name: %s, is already registered", name)
	}

	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      c.Args[0],
	}

	newUser, err := s.Database.CreateUser(
		ctx,
		params,
	)
	if err != nil {
		return fmt.Errorf("Issue creating new user: %w", err)
	}

	s.Config.SetUser(newUser.Name)

	fmt.Printf("New user: %s, %d", newUser.Name, newUser.ID)
	return nil
}
