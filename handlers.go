package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/ssgkian/gator/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("username is required")
	}
	user, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}
	err = s.cfg.SetUser(user)
	if err != nil {
		return err
	}
	fmt.Printf("username '%s' has been set successfully\n", cmd.args[0])
	return nil

}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("name is required")
	}
	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	})
	if err != nil {
		return err
	}
	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return err
	}
	fmt.Printf("%s was created sucessfully", user.Name)
	fmt.Println(user)
	return nil
}
