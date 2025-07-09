package rr_types

import (
	"fmt"
	"strconv"
	"time"

	"github.com/libdns/libdns"
)

// CAARecordParser handles CAA record parsing
type CAARecordParser struct{}

func (p CAARecordParser) Parse(args []string) (libdns.RR, error) {
	if len(args) < 1 || len(args) > 4 {
		return libdns.RR{}, fmt.Errorf("CAA record requires 1-4 arguments: <name> [flags] [tag] [value] or <name> [data]")
	}

	name := args[0]
	var data string

	if len(args) == 1 {
		// Just name - for deletion
		data = ""
	} else if len(args) == 2 {
		// Name and data - for specific deletion
		data = args[1]
	} else if len(args) == 4 {
		// Full record: name, flags, tag, value
		flagsStr := args[1]
		tag := args[2]
		value := args[3]

		// Parse flags
		flags, err := strconv.ParseUint(flagsStr, 10, 8)
		if err != nil {
			return libdns.RR{}, fmt.Errorf("invalid flags value '%s': %v", flagsStr, err)
		}

		// Validate flags (only 0 and 128 are valid)
		if flags != 0 && flags != 128 {
			return libdns.RR{}, fmt.Errorf("CAA flags must be 0 or 128, got %d", flags)
		}

		// Format CAA data
		data = fmt.Sprintf(`%d %s %q`, flags, tag, value)
	} else {
		return libdns.RR{}, fmt.Errorf("CAA record requires either 1-2 arguments (for deletion) or 4 arguments (for creation): <name> <flags> <tag> <value>")
	}

	return libdns.RR{
		Type: "CAA",
		Name: name,
		Data: data,
		TTL:  time.Duration(300) * time.Second,
	}, nil
}

func (p CAARecordParser) GetUsage() string {
	return "caa <name> <flags> <tag> <value>"
}

func (p CAARecordParser) GetShortDescription() string {
	return "a CAA record"
}

func (p CAARecordParser) GetLongDescription() string {
	return "a CAA record to specify which certificate authorities are allowed to issue certificates for a domain"
}

// Register the CAA record parser
func init() {
	RegisterRecordParser("caa", CAARecordParser{})
}
