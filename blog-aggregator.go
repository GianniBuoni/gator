package main

import (
	"fmt"
	"log"

	"github.com/GianniBuoni/blog-aggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	cfg.SetUser("Jon")

	newCfg, err := config.Read()
	fmt.Print(*newCfg)
}
