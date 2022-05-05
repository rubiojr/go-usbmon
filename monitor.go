package usbmon

import (
	"context"

	"github.com/jochenvg/go-udev"
)

type Monitor interface {
	DeviceChan(context.Context) (<-chan *udev.Device, error)
}

type UdevMonitor struct {
	monitor *udev.Monitor
}

func (m *UdevMonitor) DeviceChan(ctx context.Context) (<-chan *udev.Device, error) {
	return m.monitor.DeviceChan(ctx)
}

func NewUdevMonitor() *UdevMonitor {
	u := udev.Udev{}
	m := u.NewMonitorFromNetlink("udev")
	_ = m.FilterAddMatchTag("seat")
	_ = m.FilterAddMatchSubsystem("usb")

	return &UdevMonitor{monitor: m}
}
