package main

import "github.com/davecgh/go-spew/spew"

type metadataCmd struct {
}

var metadataOpts metadataCmd

func init() {
	if _, err := flagParser.AddCommand(
		"dump-metadata",
		"Prints playback metadata",
		"Prints playback metadata",
		&metadataOpts,
	); err != nil {
		panic(err)
	}
}

func (c *metadataCmd) Execute(args []string) error {
	obj := getDBusObj()
	spew.Dump(getMetadata(obj))

	return nil
}
