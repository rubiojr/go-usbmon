// Example that filters USB devices by serial number
package main

import (
	"context"
	"fmt"

	"github.com/rubiojr/go-usbmon"
)

func main() {
	filter := &usbmon.SerialFilter{Serial: "S6TWNS0T214043V"}
	devs, err := usbmon.ListenFiltered(context.Background(), filter)
	if err != nil {
		panic(err)
	}

	for dev := range devs {
		fmt.Println("Device detected!")
		fmt.Println("Action: " + dev.Action())
		fmt.Println("Serial: " + dev.Serial())
		fmt.Println("Path: " + dev.Path())
		fmt.Println("Vendor: " + dev.Vendor())
	}
}
