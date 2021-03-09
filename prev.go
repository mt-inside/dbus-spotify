package main

type prevCmd struct {
}

var prevOpts prevCmd

func init() {
	if _, err := flagParser.AddCommand(
		"prev",
		"Plays the previous track",
		"Plays the previous track",
		&prevOpts,
	); err != nil {
		panic(err)
	}
}

func (c *prevCmd) Execute(args []string) error {
	conn := getDBusSessionConnection()
	obj := getDBusObj(conn)
	call := obj.Call(
		"org.mpris.MediaPlayer2.Player.Previous", // Method
		0,                                        // Flags
		// Args (to ...Previous)...
	)
	if call.Err != nil {
		panic(call.Err)
	}

	return nil
}
