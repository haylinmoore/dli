package rr_types

import (
	"github.com/libdns/libdns"
	"github.com/spf13/cobra"
)

// Global variable for SVCB params
var svcbParams []string

// RecordParser defines how to parse arguments into a libdns.RR
type RecordParser interface {
	// Parse parses args into a libdns.RR
	// Should handle both complete records (name + data) and partial records (name only)
	Parse(args []string) (libdns.RR, error)

	// GetUsage returns the usage string
	GetUsage() string

	// GetShortDescription returns a short description
	GetShortDescription() string

	// GetLongDescription returns a long description
	GetLongDescription() string
}

// RecordOperation represents the type of operation being performed
type RecordOperation string

const (
	OperationSet    RecordOperation = "set"
	OperationAppend RecordOperation = "append"
	OperationDelete RecordOperation = "delete"
)

// SetupRecordCommands creates set/append/delete commands for a record type
func SetupRecordCommands(parser RecordParser, parentCmds map[RecordOperation]*cobra.Command,
	addToParent func(*cobra.Command, *cobra.Command),
	executeOp func(RecordOperation, RecordParser, []string)) {

	recordType := parser.GetShortDescription()

	// Create set command
	setCmd := &cobra.Command{
		Use:   parser.GetUsage(),
		Short: "Set " + recordType,
		Long:  "Set " + parser.GetLongDescription(),
		Args:  cobra.MinimumNArgs(1), // Will be validated in parser
		Run: func(cmd *cobra.Command, args []string) {
			executeOp(OperationSet, parser, args)
		},
	}

	// Create append command
	appendCmd := &cobra.Command{
		Use:   parser.GetUsage(),
		Short: "Append " + recordType,
		Long:  "Append " + parser.GetLongDescription(),
		Args:  cobra.MinimumNArgs(1), // Will be validated in parser
		Run: func(cmd *cobra.Command, args []string) {
			executeOp(OperationAppend, parser, args)
		},
	}

	// Create delete command
	deleteCmd := &cobra.Command{
		Use:   parser.GetUsage() + " (data optional for delete)",
		Short: "Delete " + recordType,
		Long:  "Delete " + parser.GetLongDescription() + ". If specific data is omitted, all records of this type and name will be deleted.",
		Args:  cobra.MinimumNArgs(1), // Will be validated in parser
		Run: func(cmd *cobra.Command, args []string) {
			executeOp(OperationDelete, parser, args)
		},
	}

	// Add special flags for SVCB record type
	if parser.GetShortDescription() == "an SVCB or HTTPS record" {
		// Import the svcbParams from the svcb package
		setCmd.Flags().StringSliceVar(&svcbParams, "params", []string{}, "Service parameters (e.g., alpn=h2,h3 port=443)")
		appendCmd.Flags().StringSliceVar(&svcbParams, "params", []string{}, "Service parameters (e.g., alpn=h2,h3 port=443)")
		deleteCmd.Flags().StringSliceVar(&svcbParams, "params", []string{}, "Service parameters (e.g., alpn=h2,h3 port=443)")
	}

	// Add to parent commands
	addToParent(parentCmds[OperationSet], setCmd)
	addToParent(parentCmds[OperationAppend], appendCmd)
	addToParent(parentCmds[OperationDelete], deleteCmd)
}

// Registry for record parsers
var recordParsers = make(map[string]RecordParser)

// RegisterRecordParser registers a record parser for a record type
func RegisterRecordParser(recordType string, parser RecordParser) {
	recordParsers[recordType] = parser
}

// GetRecordParser returns the parser for a record type
func GetRecordParser(recordType string) (RecordParser, bool) {
	parser, exists := recordParsers[recordType]
	return parser, exists
}

// GetAllRecordTypes returns all registered record types
func GetAllRecordTypes() []string {
	types := make([]string, 0, len(recordParsers))
	for recordType := range recordParsers {
		types = append(types, recordType)
	}
	return types
}

// SetupAllRecordCommands sets up commands for all registered record types
func SetupAllRecordCommands(parentCmds map[RecordOperation]*cobra.Command,
	addToParent func(*cobra.Command, *cobra.Command),
	executeOp func(RecordOperation, RecordParser, []string)) {

	for _, parser := range recordParsers {
		SetupRecordCommands(parser, parentCmds, addToParent, executeOp)
	}
}
