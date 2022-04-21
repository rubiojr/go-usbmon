package usbmon

type Filter interface {
	Matches(*Device) bool
}

type ActionEvent string

const (
	ActionAdd    ActionEvent = "add"
	ActionRemove ActionEvent = "remove"
	ActionAll    ActionEvent = "all"
)

type ActionFilter struct {
	Action ActionEvent
}

func (f *ActionFilter) Matches(dev *Device) bool {
	action := ActionEvent(dev.Properties()["ACTION"])

	if f.Action == ActionAll && (action == ActionAdd || action == ActionRemove) {
		return true
	}

	return action == f.Action
}

type SerialFilter struct {
	Serial string
}

func (f *SerialFilter) Matches(dev *Device) bool {
	action := ActionEvent(dev.Action())

	return f.Serial == dev.Serial() && (action == ActionAdd || action == ActionRemove)
}
