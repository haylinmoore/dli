package rr_types

import (
	"fmt"
	"net/netip"
	"time"

	"github.com/libdns/libdns"
)

// AAAARecordParser handles AAAA record parsing
type AAAARecordParser struct{}

func (p AAAARecordParser) ParseForOperation(operation string, args []string) (libdns.RR, error) {
	if operation == "delete" {
		// Delete: <name> [ipv6]
		if len(args) < 1 || len(args) > 2 {
			return libdns.RR{}, fmt.Errorf("delete AAAA requires 1-2 arguments: <name> [ipv6]")
		}

		name := args[0]
		var data string

		if len(args) == 2 {
			// Validate IPv6 if provided
			ip, err := netip.ParseAddr(args[1])
			if err != nil {
				return libdns.RR{}, fmt.Errorf("invalid IP address '%s': %v", args[1], err)
			}
			if !ip.Is6() {
				return libdns.RR{}, fmt.Errorf("'%s' is not an IPv6 address (use 'delete a' for IPv4)", args[1])
			}
			data = ip.String()
		}

		return libdns.RR{
			Type: "AAAA",
			Name: name,
			Data: data,
		}, nil
	} else {
		// Set/Append: <name> <ipv6>
		if len(args) != 2 {
			return libdns.RR{}, fmt.Errorf("%s AAAA requires 2 arguments: <name> <ipv6>", operation)
		}

		name := args[0]
		ipStr := args[1]

		// Parse IPv6 address
		ip, err := netip.ParseAddr(ipStr)
		if err != nil {
			return libdns.RR{}, fmt.Errorf("invalid IP address '%s': %v", ipStr, err)
		}

		// Ensure it's IPv6
		if !ip.Is6() {
			return libdns.RR{}, fmt.Errorf("'%s' is not an IPv6 address (use '%s a' for IPv4)", ipStr, operation)
		}

		return libdns.RR{
			Type: "AAAA",
			Name: name,
			Data: ip.String(),
			TTL:  time.Duration(300) * time.Second, // Default TTL, will be overridden by command flags
		}, nil
	}
}

func (p AAAARecordParser) GetUsage() string {
	return "aaaa <name> <ipv6>"
}

func (p AAAARecordParser) GetDeleteUsage() string {
	return "aaaa <name> [ipv6]"
}

func (p AAAARecordParser) GetShortDescription() string {
	return "an AAAA record"
}

func (p AAAARecordParser) GetLongDescription() string {
	return "an AAAA record (IPv6) for a domain name"
}

// Register the AAAA record parser
func init() {
	RegisterRecordParser("aaaa", AAAARecordParser{})
}
