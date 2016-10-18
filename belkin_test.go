package belkin_test

import (
	"fmt"
	"testing"

	"github.com/go-home-iot/belkin"
	"github.com/stretchr/testify/require"
)

func TestMakerScan(t *testing.T) {
	devices, err := belkin.Scan(belkin.DTMaker, 5)
	require.Nil(t, err)
	require.NotEqual(t, 0, len(devices), "no maker device found after scanning")

	dev := devices[0]
	err = dev.Load()
	require.Nil(t, err)

	if testing.Verbose() {
		fmt.Printf("%#v\n\n", devices)
		fmt.Printf("%#v\n", dev)
	}

	testDevice(t, dev)
}

func TestInsightScan(t *testing.T) {
	devices, err := belkin.Scan(belkin.DTInsight, 5)
	require.Nil(t, err)
	require.NotEqual(t, 0, len(devices), "no insight device found after scanning")

	dev := devices[0]
	err = dev.Load()
	require.Nil(t, err)

	if testing.Verbose() {
		fmt.Printf("%#v\n\n", devices)
		fmt.Printf("%#v\n", dev)
	}

	testDevice(t, dev)
}

func testDevice(t *testing.T, d *belkin.Device) {
	err := d.TurnOn()
	require.Nil(t, err)

	err = d.TurnOff()
	require.Nil(t, err)
}
