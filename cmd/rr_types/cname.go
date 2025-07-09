package rr_types

import (
	"fmt"
	"time"

	"github.com/libdns/libdns"
)

// CNAMERecordParser handles CNAME record parsing
type CNAMERecordParser struct{}

func (p CNAMERecordParser) Parse(args []string) (libdns.RR, error) {
	if len(args) < 1 || len(args) > 2 {
		return libdns.RR{}, fmt.Errorf("CNAME record requires 1-2 arguments: <name> [target]")
	}

	name := args[0]
	var data string

	if len(args) == 2 {
		data = args[1]
	}

	return libdns.RR{
		Type: "CNAME",
		Name: name,
		Data: data,
		TTL:  time.Duration(300) * time.Second,
	}, nil
}

func (p CNAMERecordParser) GetUsage() string {
	return "cname <name> <target>"
}

func (p CNAMERecordParser) GetShortDescription() string {
	return "a CNAME record"
}

func (p CNAMERecordParser) GetLongDescription() string {
	return "a CNAME record to delegate authority to another name"
}

// Register the CNAME record parser
func init() {
	RegisterRecordParser("cname", CNAMERecordParser{})
}
