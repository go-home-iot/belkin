# belkin
A golang library to scan and control Belkin devices, such as the WeMo Maker, WeMo Insight

##Documentation

##Installation
```bash
go get github.com/go-home-iot/belkin
```

##Package
```go
import "github.com/go-home-iot/belkin"
```

##Testing
Run the unit tests to talk to actual devices.  The tests assume that you have real devices connected to the local network that can be used during testing

```bash
go test
```

or for more detailed responses from the devices
```bash
go test -v
```

##Version History
###0.1.0
Initial release, support for scanning for belkin devices and TurnOn/TurnOff

