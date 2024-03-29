.PHONY: run
run:
	go run ./... dump-metadata

.PHONY: check
check:
	build/check-go

.PHONY: gen
gen:
	go generate ./...

# This isn't a "static" build, but that's ok.
# The go runtime gets statically linked anyway, so this only depends on libc, libsensors, etc
# ie it has reasonable deps, but doesn't need go to be installed on the user's system
# Static linking when using cgo is really hard, not least you need STATIC versions of the c libraries (like libsensors), which themselves have libc baked in
dbus-spotify:
	go build -o $@ -a -tags netgo -ldflags "-w" ./...
