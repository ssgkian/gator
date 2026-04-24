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
	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return err
	}
	fmt.Printf("username '%s' has logged in successfully\n", cmd.args[0])
	return nil

}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("username is required")
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

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return err
	}
	fmt.Printf("reset complete")
	return nil
}

func handlerGetUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}
	for _, user := range users {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}
	return nil
}
