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
	ctx, cancel := context.WithCancel(context.Background())
	// Print device properties when plugged in
	devs, err := usbmon.ListenFiltered(ctx, usbmon.Add) // use usbmon.Listen to monitor both addition/removals
	if err != nil {
		panic(err)
	}

	for dev := range devs {
		fmt.Println(dev.Serial())
		fmt.Println(dev.Path())
		fmt.Println(dev.Vendor())
	}
	cancel()
}
```

## Building

### Requirements

* libudev

Ubuntu: `apt install libudev-dev`

Fedora: `dnf install systemd-devel`
