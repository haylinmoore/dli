package rr_types

import (
	"fmt"
	"net/netip"
	"time"

	"github.com/libdns/libdns"
)

// AAAARecordParser handles AAAA record parsing
type AAAARecordParser struct{}

func (p AAAARecordParser) Parse(args []string) (libdns.RR, error) {
	if len(args) < 1 || len(args) > 2 {
		return libdns.RR{}, fmt.Errorf("AAAA record requires 1-2 arguments: <name> [ipv6]")
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
			return libdns.RR{}, fmt.Errorf("'%s' is not an IPv6 address (use 'a' for IPv4)", args[1])
		}
		data = ip.String()
	}

	return libdns.RR{
		Type: "AAAA",
		Name: name,
		Data: data,
		TTL:  time.Duration(300) * time.Second, // Default TTL, will be overridden by command flags
	}, nil
}

func (p AAAARecordParser) GetUsage() string {
	return "aaaa <name> <ipv6>"
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
