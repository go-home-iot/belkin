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
	require.Equal(t, val, 8)

	err = dev.TurnOff()
	require.Nil(t, err)

	val, err = dev.FetchBinaryState()
	require.Nil(t, err)
	require.Equal(t, val, 0)

	_, err = dev.FetchAttributes()
	require.Equal(t, err, belkin.ErrUnsupportedAction)
}

func TestParseAttributeList(t *testing.T) {
	s := "<attributeList> " +
		"<attribute> " +
		"<name>Switch</name> " +
		"<value>10</value> " +
		"</attribute> " +
		"<attribute> " +
		"<name>Sensor</name> " +
		"<value>20</value> " +
		"</attribute> " +
		"<attribute> " +
		"<name>SwitchMode</name> " +
		"<value>30</value> " +
		"</attribute> " +
		"<attribute> " +
		"<name>SensorPresent</name> " +
		"<value>40</value> " +
		"</attribute> " +
		"</attributeList> "

	attrs := belkin.ParseAttributeList(s)

	require.NotNil(t, attrs)
	require.Equal(t, 10, *attrs.Switch)
	require.Equal(t, 20, *attrs.Sensor)
	require.Equal(t, 30, *attrs.SwitchMode)
	require.Equal(t, 40, *attrs.SensorPresent)
}

func TestParseAttributeListBadInput(t *testing.T) {
	s := "<attributeList>THIS IS NOT VALID</attributeList>"

	attrs := belkin.ParseAttributeList(s)
	require.Nil(t, attrs)
}
