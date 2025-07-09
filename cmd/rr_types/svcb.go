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

func (p SVCBRecordParser) ParseForOperation(operation string, args []string) (libdns.RR, error) {
	if operation == "delete" {
		// Delete: <name> [data]
		if len(args) < 1 || len(args) > 2 {
			return libdns.RR{}, fmt.Errorf("delete SVCB requires 1-2 arguments: <name> [data]")
		}

		name := args[0]
		var data string

		if len(args) == 2 {
			data = args[1]
		}

		return libdns.RR{
			Type: "SVCB", // Default to SVCB for delete
			Name: name,
			Data: data,
		}, nil
	} else {
		// Set/Append: <scheme> <name> <priority> <target>
		if len(args) != 4 {
			return libdns.RR{}, fmt.Errorf("%s SVCB requires 4 arguments: <scheme> <name> <priority> <target>", operation)
		}

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
	}
}

func (p SVCBRecordParser) GetUsage() string {
	return "svcb <scheme> <name> <priority> <target>"
}

func (p SVCBRecordParser) GetDeleteUsage() string {
	return "svcb <name> [data]"
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
