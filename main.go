package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/ssgkian/gator/internal/config"
	"github.com/ssgkian/gator/internal/database"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	newState := state{
		db:  dbQueries,
		cfg: &cfg,
	}

	cmds := commands{
		handlers: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerGetUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", handlerAddFeed)
	cmds.register("feeds", handlerListFeeds)
	cmds.register("follow", handlerFollows)
	cmds.register("following", handlerFollowing)

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
