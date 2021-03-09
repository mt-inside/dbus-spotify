package main

type nextCmd struct {
}

var nextOpts nextCmd

func init() {
	if _, err := flagParser.AddCommand(
		"next",
		"Play next track",
		"Play next track",
		&nextOpts,
	); err != nil {
		panic(err)
	}
}

func (c *nextCmd) Execute(args []string) error {
	conn := getDBusSessionConnection()
	obj := getDBusObj(conn)
	call := obj.Call(
		"org.mpris.MediaPlayer2.Player.Next", // Method
		0,                                    // Flags
		// Args (to ...Next)...
	)
	if call.Err != nil {
		panic(call.Err)
	}

	return nil
}
