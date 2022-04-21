# USBMon

Thin udev wrapper to simplify usb device add/remove monitoring.

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

Ubuntu: `apt install libudev-dev`

Fedora: `dnf install systemd-devel`
