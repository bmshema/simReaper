# SimReaper

A Go application for interacting with GSM/3G/4G modems to retrieve SIM card information.

## Features

- Retrieve SIM identifiers (IMSI, ICCID, MSISDN)
- View stored contacts on the SIM card
- List SMS messages stored on the SIM
- Display APN settings

## Usage

```
go run cmd/simreaper/main.go [device_path]
```

If no device path is provided, it defaults to `/dev/ttyUSB2`.

## Project Structure

The project follows a standard Go project layout:

```
simReaper/
├── cmd/
│   └── simreaper/
│       └── main.go       # Application entry point
├── pkg/
│   ├── modem/
│   │   ├── client.go     # Modem communication client
│   │   └── commands.go   # AT command implementations
│   └── display/
│       ├── formatter.go  # Response formatting
│       └── ui.go         # Terminal UI functions
├── banner.txt           # ASCII art banner
├── go.mod              
└── go.sum
```

## Dependencies

- [github.com/tarm/serial](https://github.com/tarm/serial) - Serial port communication
