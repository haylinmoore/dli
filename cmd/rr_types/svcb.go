package rr_types

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/libdns/libdns"
)

// SVCBRecordParser handles SVCB/HTTPS record parsing
type SVCBRecordParser struct{}

func (p SVCBRecordParser) Parse(args []string) (libdns.RR, error) {
	if len(args) < 1 || len(args) > 4 {
		return libdns.RR{}, fmt.Errorf("SVCB record requires 1-4 arguments: <scheme> <name> <priority> <target> or <name> [data]")
	}

	if len(args) == 1 {
		// Just name - for deletion
		return libdns.RR{
			Type: "SVCB", // Default to SVCB for delete
			Name: args[0],
			Data: "",
		}, nil
	} else if len(args) == 2 {
		// Name and data - for specific deletion
		return libdns.RR{
			Type: "SVCB", // Default to SVCB for delete
			Name: args[0],
			Data: args[1],
		}, nil
	} else if len(args) == 4 {
		// Full record: scheme, name, priority, target
		scheme := args[0]
		name := args[1]
		priorityStr := args[2]
		target := args[3]

		// Parse priority
		priority, err := strconv.ParseUint(priorityStr, 10, 16)
		if err != nil {
			return libdns.RR{}, fmt.Errorf("invalid priority value '%s': %v", priorityStr, err)
		}

		// Determine record type and name based on scheme
		var recordType string
		var recordName string

		if scheme == "https" || scheme == "http" || scheme == "wss" || scheme == "ws" {
			recordType = "HTTPS"
			recordName = name
		} else {
			recordType = "SVCB"
			recordName = fmt.Sprintf("_%s.%s", scheme, name)
		}

		recordName = strings.TrimSuffix(recordName, ".@")

		// Format params
		var params string
		if priority == 0 && len(svcbParams) != 0 {
			// Alias mode - params should be empty
			params = ""
		} else {
			params = strings.Join(svcbParams, " ")
		}

		// Format SVCB/HTTPS data
		data := fmt.Sprintf("%d %s %s", priority, target, params)

		return libdns.RR{
			Type: recordType,
			Name: recordName,
			Data: data,
			TTL:  time.Duration(300) * time.Second,
		}, nil
	} else {
		return libdns.RR{}, fmt.Errorf("SVCB record requires either 1-2 arguments (for deletion) or 4 arguments (for creation): <scheme> <name> <priority> <target>")
	}
}

func (p SVCBRecordParser) GetUsage() string {
	return "svcb <scheme> <name> <priority> <target>"
}

func (p SVCBRecordParser) GetShortDescription() string {
	return "an SVCB or HTTPS record"
}

func (p SVCBRecordParser) GetLongDescription() string {
	return "an SVCB record (or HTTPS record for https/http schemes) to provide service binding information"
}

// Register the SVCB record parser
func init() {
	RegisterRecordParser("svcb", SVCBRecordParser{})
}
