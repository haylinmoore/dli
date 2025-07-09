package rr_types

import (
	"fmt"
	"strconv"
	"time"

	"github.com/libdns/libdns"
)

// MXRecordParser handles MX record parsing
type MXRecordParser struct{}

func (p MXRecordParser) Parse(args []string) (libdns.RR, error) {
	if len(args) < 1 || len(args) > 3 {
		return libdns.RR{}, fmt.Errorf("MX record requires 1-3 arguments: <name> [preference] [target] or <name> [data]")
	}

	name := args[0]
	var data string

	if len(args) == 1 {
		// Just name - for deletion
		data = ""
	} else if len(args) == 2 {
		// Name and data - for specific deletion
		data = args[1]
	} else if len(args) == 3 {
		// Full record: name, preference, target
		preferenceStr := args[1]
		target := args[2]

		// Parse preference
		preference, err := strconv.ParseUint(preferenceStr, 10, 16)
		if err != nil {
			return libdns.RR{}, fmt.Errorf("invalid preference value '%s': %v", preferenceStr, err)
		}

		// Format MX data
		data = fmt.Sprintf("%d %s", preference, target)
	}

	return libdns.RR{
		Type: "MX",
		Name: name,
		Data: data,
		TTL:  time.Duration(300) * time.Second,
	}, nil
}

func (p MXRecordParser) GetUsage() string {
	return "mx <name> <preference> <target>"
}

func (p MXRecordParser) GetShortDescription() string {
	return "an MX record"
}

func (p MXRecordParser) GetLongDescription() string {
	return "an MX record to specify mail servers for a domain"
}

// Register the MX record parser
func init() {
	RegisterRecordParser("mx", MXRecordParser{})
}
