package main

import (
	"log"
	"os"

	"github.com/ssgkian/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	newState := state{
		cfg: &cfg,
	}

	cmds := commands{
		handlers: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		log.Fatal("not enough arguments provided")
	}
	cmd := command{
		name: os.Args[1],
		args: os.Args[2:],
	}
	err = cmds.run(&newState, cmd)
	if err != nil {
		log.Fatal(err)
	}

}
