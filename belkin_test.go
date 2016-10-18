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

	err = dev.TurnOn()
	require.Nil(t, err)

	_, err = dev.FetchBinaryState()
	require.Equal(t, err, belkin.ErrUnsupportedAction)

	_, err = dev.FetchAttributes()
	require.Nil(t, err)
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

	err = dev.TurnOn()
	require.Nil(t, err)

	val, err := dev.FetchBinaryState()
	require.Nil(t, err)
	require.Equal(t, val, belkin.BinaryState(belkin.BSOn))

	err = dev.TurnOff()
	require.Nil(t, err)

	val, err = dev.FetchBinaryState()
	require.Nil(t, err)
	require.Equal(t, val, belkin.BinaryState(belkin.BSOff))

	_, err = dev.FetchAttributes()
	require.Equal(t, err, belkin.ErrUnsupportedAction)
}
