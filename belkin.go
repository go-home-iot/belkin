// Package belkin provides support for Belkin devices, such as the WeMo Switch
package belkin

import (
	"fmt"
	"time"

	"github.com/fromkeith/gossdp"
)

// CREDIT: All the knowledge of how to control this product came from:
// https://github.com/timonreinhard/wemo-client

// Scan detects Belkin devices on the network
func Scan(dt DeviceType, waitTimeSeconds int) ([]*Device, error) {
	var responses []ScanResponse
	l := belkinListener{
		URN:       string(dt),
		Responses: &responses,
	}

	c, err := gossdp.NewSsdpClientWithLogger(l, l)
	if err != nil {
		return nil, fmt.Errorf("failed to start ssdp discovery client: %s", err)
	}

	defer c.Stop()
	go c.Start()
	err = c.ListenFor(string(dt))
	if err != nil {
		return nil, fmt.Errorf("discovery failed: %s", err)
	}

	time.Sleep(time.Duration(waitTimeSeconds) * time.Second)

	devices := make([]*Device, len(responses))
	for i, response := range responses {
		devices[i] = &Device{Scan: response}
	}
	return devices, nil
}
