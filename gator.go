package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/GianniBuoni/gator/internal/commands"
	"github.com/GianniBuoni/gator/internal/config"
	"github.com/GianniBuoni/gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	state := &commands.State{}
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	state.Config = cfg

	conn, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		log.Fatal(err)
	}

	dbQueries := database.New(conn)

	state.Database = dbQueries

	commandsList := commands.Commands{
		Registry: map[string]func(*commands.State, commands.Command) error{},
	}

	commandsList.Register("login", commands.HandlerLogin)
	commandsList.Register("register", commands.HandlerRegister)
	commandsList.Register("reset", commands.HandlerReset)
	commandsList.Register("users", commands.HandlerUsers)
	commandsList.Register("agg", commands.HandlerAgg)

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
