package cmd

import (
	"time"

	"dli/cmd/rr_types"
	"github.com/libdns/libdns"
	"github.com/spf13/cobra"
)

// Set up the integration between rr_types and main cmd package
func init() {
	// Set up all record commands
	parentCmds := map[rr_types.RecordOperation]*cobra.Command{
		rr_types.OperationSet:    setCmd,
		rr_types.OperationAppend: appendCmd,
		rr_types.OperationDelete: deleteCmd,
	}

	rr_types.SetupAllRecordCommands(parentCmds, addRecordCommandToParent, executeRecordOperationFromParser)
}

// executeRecordOperationFromParser executes a record operation using a parser
func executeRecordOperationFromParser(operation rr_types.RecordOperation, parser rr_types.RecordParser, args []string) {
	// Parse the record using the parser
	record, err := parser.ParseForOperation(string(operation), args)
	if err != nil {
		OutputError("Failed to parse record", err)
		return
	}

	// Override TTL if set via flags
	if recordTTL > 0 {
		record.TTL = time.Duration(recordTTL) * time.Second
	}

	// Execute the appropriate operation
	switch operation {
	case rr_types.OperationSet:
		executeRecordSetWithRR(record)
	case rr_types.OperationAppend:
		executeRecordAppendWithRR(record)
	case rr_types.OperationDelete:
		executeRecordDeleteWithRR(record)
	}
}

// executeRecordSetWithRR executes a set operation with a pre-parsed RR
func executeRecordSetWithRR(record libdns.RR) {
	executeRecordSet(record.Type, record.Name, record.Data)
}

// executeRecordAppendWithRR executes an append operation with a pre-parsed RR
func executeRecordAppendWithRR(record libdns.RR) {
	executeRecordAppend(record.Type, record.Name, record.Data)
}

// executeRecordDeleteWithRR executes a delete operation with a pre-parsed RR
func executeRecordDeleteWithRR(record libdns.RR) {
	executeRecordDelete(record.Type, record.Name, record.Data)
}
