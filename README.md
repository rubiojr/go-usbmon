[![Go Reference](https://pkg.go.dev/badge/github.com/rubiojr/go-usbmon.svg)](https://pkg.go.dev/github.com/rubiojr/go-usbmon)

# USBMon

A lightweight Go wrapper around libudev to simplify monitoring USB device add/remove events. This package abstracts away the complexities of the low-level udev API, providing a simple interface for detecting and responding to USB devices being connected or disconnected.

```Go
// monitor USB hotplug events
package main

import (
	"context"
	"fmt"

	"github.com/rubiojr/go-usbmon"
)

func main() {
	// Print device properties when plugged in or unplugged
	filter := &usbmon.ActionFilter{Action: usbmon.ActionAll}
	devs, err := usbmon.ListenFiltered(context.Background(), filter)
	if err != nil {
		panic(err)
	}

	for dev := range devs {
		fmt.Printf("-- Device %s\n", dev.Action())
		fmt.Println("Serial: " + dev.Serial())
		fmt.Println("Path: " + dev.Path())
		fmt.Println("Vendor: " + dev.Vendor())
	}
}
```

## Building

### Requirements

* libudev

Ubuntu/Debian: `apt install libudev-dev`

Fedora/RHEL: `dnf install systemd-devel`

Arch Linux: `pacman -S systemd-libs`

## Additional examples

See [examples](_examples) for more usage examples.

## License

MIT - See [LICENSE](LICENSE) for details.
