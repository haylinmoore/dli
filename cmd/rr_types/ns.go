package rr_types

import (
	"fmt"
	"time"

	"github.com/libdns/libdns"
)

// NSRecordParser handles NS record parsing
type NSRecordParser struct{}

func (p NSRecordParser) ParseForOperation(operation string, args []string) (libdns.RR, error) {
	if operation == "delete" {
		// Delete: <name> [target]
		if len(args) < 1 || len(args) > 2 {
			return libdns.RR{}, fmt.Errorf("delete NS requires 1-2 arguments: <name> [target]")
		}

		name := args[0]
		var data string

		if len(args) == 2 {
			data = args[1]
		}

		return libdns.RR{
			Type: "NS",
			Name: name,
			Data: data,
		}, nil
	} else {
		// Set/Append: <name> <target>
		if len(args) != 2 {
			return libdns.RR{}, fmt.Errorf("%s NS requires 2 arguments: <name> <target>", operation)
		}

		name := args[0]
		target := args[1]

		return libdns.RR{
			Type: "NS",
			Name: name,
			Data: target,
			TTL:  time.Duration(300) * time.Second,
		}, nil
	}
}

func (p NSRecordParser) GetUsage() string {
	return "ns <name> <target>"
}

func (p NSRecordParser) GetDeleteUsage() string {
	return "ns <name> [target]"
}

func (p NSRecordParser) GetShortDescription() string {
	return "an NS record"
}

func (p NSRecordParser) GetLongDescription() string {
	return "an NS record to specify authoritative nameservers for a zone"
}

// Register the NS record parser
func init() {
	RegisterRecordParser("ns", NSRecordParser{})
}
