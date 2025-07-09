package rr_types

import (
	"fmt"
	"net/netip"
	"time"

	"github.com/libdns/libdns"
)

// ARecordParser handles A record parsing
type ARecordParser struct{}

func (p ARecordParser) ParseForOperation(operation string, args []string) (libdns.RR, error) {
	if operation == "delete" {
		// Delete: <name> [ipv4]
		if len(args) < 1 || len(args) > 2 {
			return libdns.RR{}, fmt.Errorf("delete A requires 1-2 arguments: <name> [ipv4]")
		}

		name := args[0]
		var data string

		if len(args) == 2 {
			// Validate IPv4 if provided
			ip, err := netip.ParseAddr(args[1])
			if err != nil {
				return libdns.RR{}, fmt.Errorf("invalid IP address '%s': %v", args[1], err)
			}
			if !ip.Is4() {
				return libdns.RR{}, fmt.Errorf("'%s' is not an IPv4 address (use 'delete aaaa' for IPv6)", args[1])
			}
			data = ip.String()
		}

		return libdns.RR{
			Type: "A",
			Name: name,
			Data: data,
		}, nil
	} else {
		// Set/Append: <name> <ipv4>
		if len(args) != 2 {
			return libdns.RR{}, fmt.Errorf("%s A requires 2 arguments: <name> <ipv4>", operation)
		}

		name := args[0]
		ipStr := args[1]

		// Parse IPv4 address
		ip, err := netip.ParseAddr(ipStr)
		if err != nil {
			return libdns.RR{}, fmt.Errorf("invalid IP address '%s': %v", ipStr, err)
		}

		// Ensure it's IPv4
		if !ip.Is4() {
			return libdns.RR{}, fmt.Errorf("'%s' is not an IPv4 address (use '%s aaaa' for IPv6)", ipStr, operation)
		}

		return libdns.RR{
			Type: "A",
			Name: name,
			Data: ip.String(),
			TTL:  time.Duration(300) * time.Second, // Default TTL, will be overridden by command flags
		}, nil
	}
}

func (p ARecordParser) GetUsage() string {
	return "a <name> <ipv4>"
}

func (p ARecordParser) GetDeleteUsage() string {
	return "a <name> [ipv4]"
}

func (p ARecordParser) GetShortDescription() string {
	return "an A record"
}

func (p ARecordParser) GetLongDescription() string {
	return "an A record (IPv4) for a domain name"
}

// Register the A record parser
func init() {
	RegisterRecordParser("a", ARecordParser{})
}
