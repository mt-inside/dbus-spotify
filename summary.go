package main

import (
	"fmt"
)

type summaryCmd struct {
}

var summaryOpts summaryCmd

func init() {
	if _, err := flagParser.AddCommand(
		"summary",
		"Prints summary output",
		"Prints summary output suitable for eg a polybar custom/script module",
		&summaryOpts,
	); err != nil {
		panic(err)
	}
}

func (c *summaryCmd) Execute(args []string) error {
	obj := getDBusObj()
	metadata := getMetadata(obj)
	status := getPlaybackStatus(obj)

	if status != "Playing" {
		fmt.Printf("%s: ", status)
	}
	fmt.Printf("%s %s\n", metadata["xesam:artist"], metadata["xesam:title"]) // Hack: array singleton printing looks neat
	return nil
}
