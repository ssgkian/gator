package main

import "errors"

type command struct {
	name string
	args []string
}

type commands struct {
	handlers map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.handlers[cmd.name]
	if ok {
		return handler(s, cmd)
	}
	return errors.New("function does not exist")
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.handlers[name] = f

}
