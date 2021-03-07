package main

import (
	"errors"
	"os"

	"github.com/godbus/dbus/v5"
	"github.com/jessevdk/go-flags"
)

var opts struct {
}
var (
	flagParser = flags.NewParser(&opts, flags.Default)
)

func main() {
	if _, err := flagParser.Parse(); err != nil {
		var e *flags.Error
		if errors.As(err, &e) {
			/* If it's flags' error type, the error message has already been printed */

			if e.Type == flags.ErrHelp {
				os.Exit(0)
			}
			os.Exit(1)
		}
		panic(err)
	}
}

func getDBusObj() dbus.BusObject {
	conn, err := dbus.SessionBus()
	if err != nil {
		panic(err)
	}
	// defer conn.Close() - static lifetime

	obj := conn.Object(
		"org.mpris.MediaPlayer2.spotify", // bus address
		"/org/mpris/MediaPlayer2",        // object path
	)

	return obj
}

func getMetadata(obj dbus.BusObject) map[string]interface{} {
	var metadata map[string]interface{}

	err := obj.Call(
		"org.freedesktop.DBus.Properties.Get", // Method
		0,                                     // Flags
		// Args (to ...Get)...
		"org.mpris.MediaPlayer2.Player", // Interface
		"Metadata",                      // Property
	).Store(&metadata)
	if err != nil {
		panic(err)
	}

	return metadata
}

func getPlaybackStatus(obj dbus.BusObject) string {
	var status string

	err := obj.Call(
		"org.freedesktop.DBus.Properties.Get", // Method
		0,                                     // Flags
		// Args (to ...Get)...
		"org.mpris.MediaPlayer2.Player", // Interface
		"PlaybackStatus",                // Property
	).Store(&status)
	if err != nil {
		panic(err)
	}

	return status
}
