package rr_types

import (
	"fmt"
	"time"

	"github.com/libdns/libdns"
)

// CNAMERecordParser handles CNAME record parsing
type CNAMERecordParser struct{}

func (p CNAMERecordParser) ParseForOperation(operation string, args []string) (libdns.RR, error) {
	if operation == "delete" {
		// Delete: <name> [target]
		if len(args) < 1 || len(args) > 2 {
			return libdns.RR{}, fmt.Errorf("delete CNAME requires 1-2 arguments: <name> [target]")
		}

		name := args[0]
		var data string

		if len(args) == 2 {
			data = args[1]
		}

		return libdns.RR{
			Type: "CNAME",
			Name: name,
			Data: data,
		}, nil
	} else {
		// Set/Append: <name> <target>
		if len(args) != 2 {
			return libdns.RR{}, fmt.Errorf("%s CNAME requires 2 arguments: <name> <target>", operation)
		}

		name := args[0]
		target := args[1]

		return libdns.RR{
			Type: "CNAME",
			Name: name,
			Data: target,
			TTL:  time.Duration(300) * time.Second,
		}, nil
	}
}

func (p CNAMERecordParser) GetUsage() string {
	return "cname <name> <target>"
}

func (p CNAMERecordParser) GetDeleteUsage() string {
	return "cname <name> [target]"
}

func (p CNAMERecordParser) GetShortDescription() string {
	return "a CNAME record"
}

func (p CNAMERecordParser) GetLongDescription() string {
	return "a CNAME record to delegate authority to another name"
}

// Register the CNAME record parser
func init() {
	RegisterRecordParser("cname", CNAMERecordParser{})
}
