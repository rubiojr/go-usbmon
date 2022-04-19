package usbmon

import (
	"context"

	"github.com/jochenvg/go-udev"
)

type ActionFilter string

const (
	Add    ActionFilter = "add"
	Remove ActionFilter = "remove"
	All    ActionFilter = "all"
)

type Device struct {
	properties map[string]string
}

func (d *Device) Serial() string {
	return d.properties["ID_SERIAL_SHORT"]
}

func (d *Device) Properties() map[string]string {
	return d.properties
}

func (d *Device) Vendor() string {
	return d.properties["ID_VENDOR"]
}

func (d *Device) Major() string {
	return d.properties["MAJOR"]
}

func (d *Device) Minor() string {
	return d.properties["MINOR"]
}

func (d *Device) Path() string {
	return d.properties["DEVPATH"]
}

func Listen(ctx context.Context) (chan *Device, error) {
	return ListenFiltered(ctx, All)
}

func ListenFiltered(ctx context.Context, filter ActionFilter) (chan *Device, error) {
	u := udev.Udev{}
	m := u.NewMonitorFromNetlink("udev")
	m.FilterAddMatchTag("seat")
	m.FilterAddMatchSubsystem("usb")

	devchan := make(chan *Device)
	ch, err := m.DeviceChan(ctx)
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				close(devchan)
				return
			case d := <-ch:
				dev := &Device{
					properties: map[string]string{},
				}
				dev.properties = d.Properties()

				action := ActionFilter(d.Properties()["ACTION"])
				if filter == All && (action == Add || action == Remove) {
					devchan <- dev
					continue
				}

				if action == filter {
					devchan <- dev
				}
			}
		}
	}()

	return devchan, nil
}
