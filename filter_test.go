package usbmon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestActionFilter(t *testing.T) {
	d := Device{
		properties: map[string]string{"ACTION": "add"},
	}

	f := ActionFilter{Action: ActionAdd}
	assert.True(t, f.Matches(&d))
	d.properties["ACTION"] = "all"
	assert.False(t, f.Matches(&d))
	d.properties["ACTION"] = "remove"
	assert.False(t, f.Matches(&d))

	f = ActionFilter{Action: ActionAll}
	d.properties["ACTION"] = "all"
	assert.True(t, f.Matches(&d))
	d.properties["ACTION"] = "add"
	assert.True(t, f.Matches(&d))
	d.properties["ACTION"] = "remove"
	assert.True(t, f.Matches(&d))
	d.properties["ACTION"] = "unbind"
	assert.False(t, f.Matches(&d))
	d.properties["ACTION"] = "bind"
	assert.False(t, f.Matches(&d))

	f = ActionFilter{Action: ActionRemove}
	d.properties["ACTION"] = "all"
	assert.False(t, f.Matches(&d))
	d.properties["ACTION"] = "add"
	assert.False(t, f.Matches(&d))
	d.properties["ACTION"] = "remove"
	assert.True(t, f.Matches(&d))
}

func TestSerialFilter(t *testing.T) {
	d := Device{
		properties: map[string]string{"ACTION": "add", "ID_SERIAL_SHORT": "1234"},
	}

	f := SerialFilter{Serial: "1234"}
	assert.True(t, f.Matches(&d))
	d.properties["ACTION"] = "bind"
	assert.False(t, f.Matches(&d))
	d.properties["ACTION"] = "remove"
	assert.True(t, f.Matches(&d))
	d.properties["ACTION"] = "unbind"
	assert.False(t, f.Matches(&d))
	f.Serial = ""
	assert.False(t, f.Matches(&d))
	f.Serial = "123"
	assert.False(t, f.Matches(&d))
}
