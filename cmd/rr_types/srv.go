package rr_types

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/libdns/libdns"
)

// SRVRecordParser handles SRV record parsing
type SRVRecordParser struct{}

func (p SRVRecordParser) ParseForOperation(operation string, args []string) (libdns.RR, error) {
	if operation == "delete" {
		// Delete: <name> [data]
		if len(args) < 1 || len(args) > 2 {
			return libdns.RR{}, fmt.Errorf("delete SRV requires 1-2 arguments: <name> [data]")
		}

		name := args[0]
		var data string

		if len(args) == 2 {
			data = args[1]
		}

		return libdns.RR{
			Type: "SRV",
			Name: name,
			Data: data,
		}, nil
	} else {
		// Set/Append: <service> <protocol> <name> <priority> <weight> <port> <target>
		if len(args) != 7 {
			return libdns.RR{}, fmt.Errorf("%s SRV requires 7 arguments: <service> <protocol> <name> <priority> <weight> <port> <target>", operation)
		}

		service := args[0]
		protocol := args[1]
		name := args[2]
		priorityStr := args[3]
		weightStr := args[4]
		portStr := args[5]
		target := args[6]

		// Parse priority
		priority, err := strconv.ParseUint(priorityStr, 10, 16)
		if err != nil {
			return libdns.RR{}, fmt.Errorf("invalid priority value '%s': %v", priorityStr, err)
		}

		// Parse weight
		weight, err := strconv.ParseUint(weightStr, 10, 16)
		if err != nil {
			return libdns.RR{}, fmt.Errorf("invalid weight value '%s': %v", weightStr, err)
		}

		// Parse port
		port, err := strconv.ParseUint(portStr, 10, 16)
		if err != nil {
			return libdns.RR{}, fmt.Errorf("invalid port value '%s': %v", portStr, err)
		}

		// Build SRV record name
		var recordName string
		if service == "" && protocol == "" {
			recordName = name
		} else {
			recordName = fmt.Sprintf("_%s._%s.%s", service, protocol, name)
		}
		recordName = strings.TrimSuffix(recordName, ".@")

		// Format SRV data
		data := fmt.Sprintf("%d %d %d %s", priority, weight, port, target)

		return libdns.RR{
			Type: "SRV",
			Name: recordName,
			Data: data,
			TTL:  time.Duration(300) * time.Second,
		}, nil
	}
}

func (p SRVRecordParser) GetUsage() string {
	return "srv <service> <protocol> <name> <priority> <weight> <port> <target>"
}

func (p SRVRecordParser) GetDeleteUsage() string {
	return "srv <name> [data]"
}

func (p SRVRecordParser) GetShortDescription() string {
	return "an SRV record"
}

func (p SRVRecordParser) GetLongDescription() string {
	return "an SRV record to specify services available on a network"
}

// Register the SRV record parser
func init() {
	RegisterRecordParser("srv", SRVRecordParser{})
}
