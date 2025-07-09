package rr_types

import (
	"fmt"
	"time"

	"github.com/libdns/libdns"
)

// TXTRecordParser handles TXT record parsing
type TXTRecordParser struct{}

func (p TXTRecordParser) Parse(args []string) (libdns.RR, error) {
	if len(args) < 1 || len(args) > 2 {
		return libdns.RR{}, fmt.Errorf("TXT record requires 1-2 arguments: <name> [text]")
	}

	name := args[0]
	var data string

	if len(args) == 2 {
		data = args[1]
	}

	return libdns.RR{
		Type: "TXT",
		Name: name,
		Data: data,
		TTL:  time.Duration(300) * time.Second, // Default TTL, will be overridden by command flags
	}, nil
}

func (p TXTRecordParser) GetUsage() string {
	return "txt <name> <text>"
}

func (p TXTRecordParser) GetShortDescription() string {
	return "a TXT record"
}

func (p TXTRecordParser) GetLongDescription() string {
	return "a TXT record to add arbitrary text data to a name in a DNS zone"
}

// Register the TXT record parser
func init() {
	RegisterRecordParser("txt", TXTRecordParser{})
}
