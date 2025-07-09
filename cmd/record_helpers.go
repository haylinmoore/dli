package cmd

import (
	"context"
	"fmt"
	"strings"
	"time"

	"dli/providers"
	"github.com/libdns/libdns"
	"github.com/spf13/cobra"
)

// RecordConfig holds common configuration for all record types
type RecordConfig struct {
	TTL      int
	Zone     string
	Provider string
}

// Common flags that all record commands will inherit
var recordTTL int

// setupRecordCommand configures common flags and validation for record commands
func setupRecordCommand(cmd *cobra.Command) {
	cmd.Flags().IntVar(&recordTTL, "ttl", 0, "TTL in seconds (default: 300)")
}

// addRecordCommand adds a command with both uppercase and lowercase versions to the set command
func addRecordCommand(cmd *cobra.Command) {
	addRecordCommandToParent(setCmd, cmd)
}

// addRecordCommandWithFlags adds a command with both cases and custom flag setup to the set command
func addRecordCommandWithFlags(cmd *cobra.Command, setupFlags func(*cobra.Command)) {
	addRecordCommandWithFlagsToParent(setCmd, cmd, setupFlags)
}

// addAppendRecordCommand adds a command with both uppercase and lowercase versions to the append command
func addAppendRecordCommand(cmd *cobra.Command) {
	addRecordCommandToParent(appendCmd, cmd)
}

// addAppendRecordCommandWithFlags adds a command with both cases and custom flag setup to the append command
func addAppendRecordCommandWithFlags(cmd *cobra.Command, setupFlags func(*cobra.Command)) {
	addRecordCommandWithFlagsToParent(appendCmd, cmd, setupFlags)
}

// addDeleteRecordCommand adds a command with both uppercase and lowercase versions to the delete command
func addDeleteRecordCommand(cmd *cobra.Command) {
	addRecordCommandToParent(deleteCmd, cmd)
}

// addDeleteRecordCommandWithFlags adds a command with both cases and custom flag setup to the delete command
func addDeleteRecordCommandWithFlags(cmd *cobra.Command, setupFlags func(*cobra.Command)) {
	addRecordCommandWithFlagsToParent(deleteCmd, cmd, setupFlags)
}

// validateRecordCommand performs common validation for record commands
func validateRecordCommand() error {
	if zone == "" {
		return fmt.Errorf("--zone flag is required")
	}
	return nil
}

// executeRecordOperation handles the common record operation logic
func executeRecordOperation(operation, recordType, name, data string) {
	if err := validateRecordCommand(); err != nil {
		OutputError("Validation failed", err)
		return
	}

	recordTTLDuration := 300
	if recordTTL > 0 {
		recordTTLDuration = recordTTL
	}

	record := libdns.RR{
		Type: recordType,
		Name: name,
		Data: data,
		TTL:  time.Duration(recordTTLDuration) * time.Second,
	}

	providerInstance, err := providers.GetProvider(provider)
	if err != nil {
		OutputError("Failed to get provider", err)
		return
	}

	// Execute the operation
	ctx := context.Background()
	var updatedRecords []libdns.Record

	switch strings.ToLower(operation) {
	case "setting":
		updatedRecords, err = providerInstance.SetRecords(ctx, zone, []libdns.Record{record})
	case "appending":
		updatedRecords, err = providerInstance.AppendRecords(ctx, zone, []libdns.Record{record})
	default:
		OutputError("Unknown operation", fmt.Errorf("operation %s not supported", operation))
		return
	}

	if err != nil {
		OutputError(fmt.Sprintf("Failed to %s record", strings.ToLower(operation)), err)
		return
	}

	var updatedRRs []map[string]interface{}
	for _, record := range updatedRecords {
		rr := record.RR()
		updatedRRs = append(updatedRRs, map[string]interface{}{
			"name": rr.Name,
			"type": rr.Type,
			"ttl":  int(rr.TTL.Seconds()),
			"data": rr.Data,
		})
	}

	result := map[string]interface{}{
		"operation":       strings.ToLower(operation),
		"record_type":     recordType,
		"name":            name,
		"data":            data,
		"zone":            zone,
		"provider":        provider,
		"updated_records": updatedRRs,
	}
	OutputSuccess(fmt.Sprintf("Successfully %s %s record", strings.ToLower(operation), recordType), result)
}

// executeRecordSet handles the common record setting logic
func executeRecordSet(recordType, name, data string) {
	executeRecordOperation("Setting", recordType, name, data)
}

// executeRecordAppend handles the common record appending logic
func executeRecordAppend(recordType, name, data string) {
	executeRecordOperation("Appending", recordType, name, data)
}

// executeRecordDelete handles the common record deletion logic
func executeRecordDelete(recordType, name, data string) {
	if err := validateRecordCommand(); err != nil {
		OutputError("Validation failed", err)
		return
	}

	providerInstance, err := providers.GetProvider(provider)
	if err != nil {
		OutputError("Failed to get provider", err)
		return
	}

	ctx := context.Background()

	if data != "" {

		record := libdns.RR{
			Type: recordType,
			Name: name,
			Data: data,
		}

		deletedRecords, err := providerInstance.DeleteRecords(ctx, zone, []libdns.Record{record})
		if err != nil {
			OutputError("Failed to delete record", err)
			return
		}

		var deletedRRs []map[string]interface{}
		for _, deletedRecord := range deletedRecords {
			rr := deletedRecord.RR()
			deletedRRs = append(deletedRRs, map[string]interface{}{
				"name": rr.Name,
				"type": rr.Type,
				"ttl":  int(rr.TTL.Seconds()),
				"data": rr.Data,
			})
		}

		result := map[string]interface{}{
			"operation":       "delete",
			"record_type":     recordType,
			"name":            name,
			"data":            data,
			"zone":            zone,
			"provider":        provider,
			"deleted_count":   len(deletedRecords),
			"deleted_records": deletedRRs,
		}
		OutputSuccess("Successfully deleted specific record", result)
	} else {
		allRecords, err := providerInstance.GetRecords(ctx, zone)
		if err != nil {
			OutputError("Failed to get records", err)
			return
		}

		// Filter records matching type and name
		var recordsToDelete []libdns.Record
		for _, record := range allRecords {
			rr := record.RR()
			if rr.Type == recordType && rr.Name == name {
				recordsToDelete = append(recordsToDelete, record)
			}
		}

		if len(recordsToDelete) == 0 {
			result := map[string]interface{}{
				"operation":       "delete",
				"record_type":     recordType,
				"name":            name,
				"zone":            zone,
				"provider":        provider,
				"deleted_count":   0,
				"deleted_records": []map[string]interface{}{},
			}
			OutputSuccess("No matching records found", result)
			return
		}

		deletedRecords, err := providerInstance.DeleteRecords(ctx, zone, recordsToDelete)
		if err != nil {
			OutputError("Failed to delete records", err)
			return
		}

		var deletedRRs []map[string]interface{}
		for _, deletedRecord := range deletedRecords {
			rr := deletedRecord.RR()
			deletedRRs = append(deletedRRs, map[string]interface{}{
				"name": rr.Name,
				"type": rr.Type,
				"ttl":  int(rr.TTL.Seconds()),
				"data": rr.Data,
			})
		}

		result := map[string]interface{}{
			"operation":       "delete",
			"record_type":     recordType,
			"name":            name,
			"zone":            zone,
			"provider":        provider,
			"deleted_count":   len(deletedRecords),
			"deleted_records": deletedRRs,
		}
		OutputSuccess("Successfully deleted all matching records", result)
	}
}

// createRecordCommand creates a record command for the specified operation
func createRecordCommand(operation, recordType, use, short, long string, args int, runFunc func(cmd *cobra.Command, args []string)) *cobra.Command {
	return &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,
		Args:  cobra.ExactArgs(args),
		Run:   runFunc,
	}
}

// addRecordCommandToParent adds a command with both cases to a parent command
func addRecordCommandToParent(parentCmd *cobra.Command, cmd *cobra.Command) {
	// Add lowercase version (hidden from help)
	cmd.Hidden = true
	parentCmd.AddCommand(cmd)
	setupRecordCommand(cmd)

	// Create uppercase version (visible in help)
	upperCmd := &cobra.Command{
		Use:    strings.ToUpper(cmd.Use),
		Short:  cmd.Short,
		Long:   cmd.Long,
		Args:   cmd.Args,
		Run:    cmd.Run,
		Hidden: false,
	}
	parentCmd.AddCommand(upperCmd)
	setupRecordCommand(upperCmd)
}

// addRecordCommandWithFlagsToParent adds a command with both cases and custom flags to a parent command
func addRecordCommandWithFlagsToParent(parentCmd *cobra.Command, cmd *cobra.Command, setupFlags func(*cobra.Command)) {
	// Set up flags for lowercase version
	setupFlags(cmd)

	// Add lowercase version (hidden from help)
	cmd.Hidden = true
	parentCmd.AddCommand(cmd)
	setupRecordCommand(cmd)

	// Create uppercase version (visible in help)
	upperCmd := &cobra.Command{
		Use:    strings.ToUpper(cmd.Use),
		Short:  cmd.Short,
		Long:   cmd.Long,
		Args:   cmd.Args,
		Run:    cmd.Run,
		Hidden: false,
	}
	parentCmd.AddCommand(upperCmd)
	setupRecordCommand(upperCmd)

	// Set up flags for uppercase version
	setupFlags(upperCmd)
}
