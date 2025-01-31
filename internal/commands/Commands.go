package commands

import (
	"fmt"

	"github.com/GianniBuoni/gator/internal/config"
	"github.com/GianniBuoni/gator/internal/database"
)

type State struct {
	Config   *config.Config
	Database *database.Queries
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
