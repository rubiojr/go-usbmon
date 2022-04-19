package main

import (
	"context"
	"fmt"

	"github.com/rubiojr/go-usbmon"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	// Print device properties when plugged in
	devs, err := usbmon.ListenFiltered(ctx, usbmon.Add)
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
