package main

type playpauseCmd struct {
}

var playpauseOpts playpauseCmd

func init() {
	if _, err := flagParser.AddCommand(
		"playpause",
		"Toggle play/pause",
		"Toggle play/pause",
		&playpauseOpts,
	); err != nil {
		panic(err)
	}
}

func (c *playpauseCmd) Execute(args []string) error {
	conn := getDBusSessionConnection()
	obj := getDBusObj(conn)
	call := obj.Call(
		"org.mpris.MediaPlayer2.Player.PlayPause", // Method
		0, // Flags
		// Args (to ...PlayPause)...
	)
	if call.Err != nil {
		panic(call.Err)
	}

	return nil
}
