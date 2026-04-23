package main

import (
	"github.com/ssgkian/gator/internal/config"
	"github.com/ssgkian/gator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}
