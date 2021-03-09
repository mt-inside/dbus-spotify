package main

import (
	"fmt"
)

type statusCmd struct {
}

var statusOpts statusCmd

func init() {
	if _, err := flagParser.AddCommand(
		"status",
		"Prints current playback status",
		"Prints current playback status",
		&statusOpts,
	); err != nil {
		panic(err)
	}
}

func (c *statusCmd) Execute(args []string) error {
	conn := getDBusSessionConnection()
	obj := getDBusObj(conn)
	fmt.Println(getPlaybackStatus(obj))

	return nil
}
