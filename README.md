# dbus-spotify
Simple programme that prints info from Spotify over DBus.

[![Checks](https://github.com/mt-inside/dbus-spotify/actions/workflows/checks.yaml/badge.svg)](https://github.com/mt-inside/dbus-spotify/actions/workflows/checks.yaml)
[![GitHub Issues](https://img.shields.io/github/issues-raw/mt-inside/dbus-spotify)](https://github.com/mt-inside/dbus-spotify/issues)

[![Go Reference](https://pkg.go.dev/badge/github.com/mt-inside/dbus-spotify.svg)](https://pkg.go.dev/github.com/mt-inside/dbus-spotify)

## Setup
* Download [`dbus-spotity`](https://github.com/mt-inside/dbus-spotify/releases/latest/download/dbus-spotify_linux_amd64), `chmod 0755`, and enjoy!
  * Or, with the _go_ toolchain installed, build it manually with `go get github.com/mt-inside/dbus-sportify`

## Example

### Manual Execution
```
$ ./dbus-spotify summary
[Lorna Shore] Void
```

### Polybar Config
```
[module/spotify]
type = custom/script

exec = dbus-spotify summary 2> .local/share/polybar/script-dbus-spotify.log
interval = 2

click-left = dbus-spotify prev
click-middle = dbus-spotify playpause
click-right = dbus-spotify next
```
