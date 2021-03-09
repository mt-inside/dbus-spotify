package main

import (
	"fmt"

	"github.com/godbus/dbus/v5"
	"github.com/mt-inside/go-usvc"
)

type summaryCmd struct {
}

var summaryOpts summaryCmd

type watchCmd struct {
}

var watchOpts watchCmd

func init() {
	if _, err := flagParser.AddCommand(
		"summary",
		"Prints summary output",
		"Prints summary output suitable for eg a polybar custom/script module",
		&summaryOpts,
	); err != nil {
		panic(err)
	}

	if _, err := flagParser.AddCommand(
		"watch",
		"Blocks, and prints updates summaries as they change",
		"Prints summary outputs suitable for eg a polybar custom/script module",
		&watchOpts,
	); err != nil {
		panic(err)
	}
}

func (c *summaryCmd) Execute(args []string) error {
	conn := getDBusSessionConnection()
	obj := getDBusObj(conn)

	printSummary(getPlaybackStatus(obj), getMetadata(obj))

	return nil
}

func (c *watchCmd) Execute(args []string) error {
	conn := getDBusSessionConnection()
	obj := getDBusObj(conn)

	// Print initial values
	printSummary(getPlaybackStatus(obj), getMetadata(obj))

	// Watch
	sigs := watchPropertiesChanged(conn, obj)
	var iface string
	var props map[string]interface{}
	var unknown []string

	for s := range sigs {
		dbus.Store(s.Body, &iface, &props, &unknown) // Have to use Store(), because it's Variants all the way down (props, props["Metadata"]), and Store() unpacks those so we don't have to keep going through .Value()
		status := props["PlaybackStatus"].(string)
		metadata := props["Metadata"].(map[string]interface{})

		printSummary(status, metadata)
	}

	return nil
}

func printSummary(status string, metadata map[string]interface{}) {
	if status != "Playing" {
		fmt.Printf("%s: ", status)
	}
	summary := fmt.Sprintf("%s %s", metadata["xesam:artist"], metadata["xesam:title"]) // Hack: array singleton printing looks neat

	usvc.PrintUpdateLn(summary)
}
