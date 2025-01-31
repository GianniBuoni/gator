package commands

import (
	"errors"
	"fmt"

	"github.com/GianniBuoni/gator/internal/config"
)

type State struct {
	Config *config.Config
}

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	Registry map[string]func(*State, Command) error
}

func (c *Commands) Register(name string, f func(*State, Command) error) {
	c.Registry[name] = f
}

func (c *Commands) Run(s *State, cmd Command) error {
	if _, ok := c.Registry[cmd.Name]; !ok {
		return fmt.Errorf("command not found: %s", cmd.Name)
	}
	if err := c.Registry[cmd.Name](s, cmd); err != nil {
		return err
	}
	return nil
}

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return errors.New("login extects a single argument: username")
	}

	if err := s.Config.SetUser(cmd.Args[0]); err != nil {
		return err
	}

	fmt.Printf("%s logged in!", cmd.Args[0])
	return nil
}
