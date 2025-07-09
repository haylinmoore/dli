package rr_types

import (
	"fmt"
	"net/netip"
	"time"

	"github.com/libdns/libdns"
)

// ARecordParser handles A record parsing
type ARecordParser struct{}

func (p ARecordParser) Parse(args []string) (libdns.RR, error) {
	if len(args) < 1 || len(args) > 2 {
		return libdns.RR{}, fmt.Errorf("A record requires 1-2 arguments: <name> [ipv4]")
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
			return libdns.RR{}, fmt.Errorf("'%s' is not an IPv4 address (use 'aaaa' for IPv6)", args[1])
		}
		data = ip.String()
	}

	return libdns.RR{
		Type: "A",
		Name: name,
		Data: data,
		TTL:  time.Duration(300) * time.Second, // Default TTL, will be overridden by command flags
	}, nil
}

func (p ARecordParser) GetUsage() string {
	return "a <name> <ipv4>"
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
