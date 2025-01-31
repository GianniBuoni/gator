package main

import (
	"log"
	"os"

	"github.com/GianniBuoni/gator/internal/commands"
	"github.com/GianniBuoni/gator/internal/config"
)

func main() {
	state := &commands.State{}
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	state.Config = cfg

	commandsList := commands.Commands{
		Registry: map[string]func(*commands.State, commands.Command) error{},
	}

	commandsList.Register("login", commands.HandlerLogin)

	input := os.Args
	if len(input) < 2 {
		log.Fatalln("error: expecting command name and command argument.")
	}

	command := commands.Command{
		Name: input[1],
		Args: input[2:],
	}

	if err := commandsList.Run(state, command); err != nil {
		log.Fatalf("issue running command: %v", err)
	}
}
