package usbmon

import (
	"context"
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

func (d *Device) Action() string {
	return d.properties["ACTION"]
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
	return ListenFiltered(ctx)
}

// ListenFiltered returns the usb storage devices that match all the filters passed
// as arguments.
//
// Filters are additive, meaning every device needs to match all the filter arguments.
//
// Example:
func ListenFiltered(ctx context.Context, filters ...Filter) (chan *Device, error) {
	m := NewUdevMonitor()
	devchan := make(chan *Device)
	ch, ech, err := m.DeviceChan(ctx)
	if err != nil {
		return nil, err
	}

	var lerr error
	go func() {
	Loop:
		for {
			select {
			case <-ctx.Done():
				close(devchan)
				return
			case lerr = <-ech:
				break Loop
			case d := <-ch:
				dev := &Device{
					properties: d.Properties(),
				}

				if filters == nil {
					devchan <- dev
					continue
				}

				match := true
				for _, f := range filters {
					if !f.Matches(dev) {
						match = false
						break
					}
				}

				if match {
					devchan <- dev
				}
			}
		}
	}()

	return devchan, lerr
}
