/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package messagelog

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"time"
)

const (
	AMTMessageLog         string = "AMT_MessageLog"
	GetRecords            string = "GetRecords"
	PositionToFirstRecord string = "PositionToFirstRecord"
	ValueNotFound         string = "Value not found in map"
)

const (
	CapabilitiesUnknown Capabilities = iota
	CapabilitiesOther
	CapabilitiesWriteRecordSupported
	CapabilitiesDeleteRecordSupported
	CapabilitiesCanMoveBackwardInLog
	CapabilitiesFreezeLogSupported
	CapabilitiesClearLogSupported
	CapabilitiesSupportsAddressingByOrdinalRecordNumber
	CapabilitiesVariableLengthRecordsSupported
	CapabilitiesVariableFormatsForRecords
	CapabilitiesCanFlagRecordsForOverwrite
)

// capabilitiesString is a map of the capabilities to their string representation.
var capabilitiesString = map[Capabilities]string{
	CapabilitiesUnknown:                                 "Unknown",
	CapabilitiesOther:                                   "Other",
	CapabilitiesWriteRecordSupported:                    "WriteRecordSupported",
	CapabilitiesDeleteRecordSupported:                   "DeleteRecordSupported",
	CapabilitiesCanMoveBackwardInLog:                    "CanMoveBackwardInLog",
	CapabilitiesFreezeLogSupported:                      "FreezeLogSupported",
	CapabilitiesClearLogSupported:                       "ClearLogSupported",
	CapabilitiesSupportsAddressingByOrdinalRecordNumber: "SupportsAddressingByOrdinalRecordNumber",
	CapabilitiesVariableLengthRecordsSupported:          "VariableLengthRecordsSupported",
	CapabilitiesVariableFormatsForRecords:               "VariableFormatsForRecords",
	CapabilitiesCanFlagRecordsForOverwrite:              "CanFlagRecordsForOverwrite",
}

// String returns the string representation of the capabilities.
func (c Capabilities) String() string {
	if value, exists := capabilitiesString[c]; exists {
		return value
	}

	return ValueNotFound
}

const (
	CharacterSetUnknown CharacterSet = iota
	CharacterSetOther
	CharacterSetASCII
	CharacterSetUnicode
	CharacterSetISO2022
	CharacterSetISO8859
	CharacterSetExtendedUNIXCode
	CharacterSetUTF8
	CharacterSetUCS2
	CharacterSetBitmappedData
	CharacterSetOctetString
	CharacterSetDefinedByIndividualRecords
)

// characterSetString is a map of the character set to their string representation.
var characterSetString = map[CharacterSet]string{
	CharacterSetUnknown:                    "Unknown",
	CharacterSetOther:                      "Other",
	CharacterSetASCII:                      "ASCII",
	CharacterSetUnicode:                    "Unicode",
	CharacterSetISO2022:                    "ISO2022",
	CharacterSetISO8859:                    "ISO8859",
	CharacterSetExtendedUNIXCode:           "ExtendedUNIXCode",
	CharacterSetUTF8:                       "UTF8",
	CharacterSetUCS2:                       "UCS2",
	CharacterSetBitmappedData:              "BitmappedData",
	CharacterSetOctetString:                "OctetString",
	CharacterSetDefinedByIndividualRecords: "DefinedByIndividualRecords",
}

// ConvertCharacterSetToString returns the string representation of the character set.
func (c CharacterSet) String() string {
	if value, exists := characterSetString[c]; exists {
		return value
	}

	return ValueNotFound
}

const (
	EnabledDefaultEnabled           EnabledDefault = 2
	EnabledDefaultDisabled          EnabledDefault = 3
	EnabledDefaultNotApplicable     EnabledDefault = 5
	EnabledDefaultEnabledButOffline EnabledDefault = 6
	EnabledDefaultNoDefault         EnabledDefault = 7
	EnabledDefaultQuiesce           EnabledDefault = 9
)

// enabledDefaultString is a map of the enabled default to their string representation.
var enabledDefaultString = map[EnabledDefault]string{
	EnabledDefaultEnabled:           "Enabled",
	EnabledDefaultDisabled:          "Disabled",
	EnabledDefaultNotApplicable:     "NotApplicable",
	EnabledDefaultEnabledButOffline: "EnabledButOffline",
	EnabledDefaultNoDefault:         "NoDefault",
	EnabledDefaultQuiesce:           "Quiesce",
}

// String returns the string representation of the enabled default.
func (e EnabledDefault) String() string {
	if value, exists := enabledDefaultString[e]; exists {
		return value
	}

	return ValueNotFound
}

const (
	EnabledStateUnknown EnabledState = iota
	EnabledStateOther
	EnabledStateEnabled
	EnabledStateDisabled
	EnabledStateShuttingDown
	EnabledStateNotApplicable
	EnabledStateEnabledButOffline
	EnabledStateInTest
	EnabledStateDeferred
	EnabledStateQuiesce
	EnabledStateStarting
)

// enabledStateString is a map of the enabled state to their string representation.
var enabledStateString = map[EnabledState]string{
	EnabledStateUnknown:           "Unknown",
	EnabledStateOther:             "Other",
	EnabledStateEnabled:           "Enabled",
	EnabledStateDisabled:          "Disabled",
	EnabledStateShuttingDown:      "ShuttingDown",
	EnabledStateNotApplicable:     "NotApplicable",
	EnabledStateEnabledButOffline: "EnabledButOffline",
	EnabledStateInTest:            "InTest",
	EnabledStateDeferred:          "Deferred",
	EnabledStateQuiesce:           "Quiesce",
	EnabledStateStarting:          "Starting",
}

// String returns the string representation of the enabled state.
func (e EnabledState) String() string {
	if value, exists := enabledStateString[e]; exists {
		return value
	}

	return ValueNotFound
}

const (
	HealthStateUnknown             HealthState = 0
	HealthStateOK                  HealthState = 5
	HealthStateDegraded            HealthState = 10
	HealthStateMinorFailure        HealthState = 20
	HealthStateMajorFailure        HealthState = 25
	HealthStateCriticalFailure     HealthState = 30
	HealthStateNonRecoverableError HealthState = 35
)

// healthStateString is a map of the health state to their string representation.
var healthStateString = map[HealthState]string{
	HealthStateUnknown:             "Unknown",
	HealthStateOK:                  "OK",
	HealthStateDegraded:            "Degraded",
	HealthStateMinorFailure:        "MinorFailure",
	HealthStateMajorFailure:        "MajorFailure",
	HealthStateCriticalFailure:     "CriticalFailure",
	HealthStateNonRecoverableError: "NonRecoverableError",
}

// String returns the string representation of the health state.
func (h HealthState) String() string {
	if value, exists := healthStateString[h]; exists {
		return value
	}

	return ValueNotFound
}

const (
	LastChangeUnknown LastChange = iota
	LastChangeAdd
	LastChangeDelete
	LastChangeModify
	LastChangeLogCleared
)

// lastChangeString is a map of the last change to their string representation.
var lastChangeString = map[LastChange]string{
	LastChangeUnknown:    "Unknown",
	LastChangeAdd:        "Add",
	LastChangeDelete:     "Delete",
	LastChangeModify:     "Modify",
	LastChangeLogCleared: "LogCleared",
}

// String returns the string representation of the last change.
func (l LastChange) String() string {
	if value, exists := lastChangeString[l]; exists {
		return value
	}

	return ValueNotFound
}

const (
	LogStateUnknown       LogState = 0
	LogStateNormal        LogState = 2
	LogStateErasing       LogState = 3
	LogStateNotApplicable LogState = 5
)

// logStateString is a map of the log state to their string representation.
var logStateString = map[LogState]string{
	LogStateUnknown:       "Unknown",
	LogStateNormal:        "Normal",
	LogStateErasing:       "Erasing",
	LogStateNotApplicable: "NotApplicable",
}

// String returns the string representation of the log state.
func (l LogState) String() string {
	if value, exists := logStateString[l]; exists {
		return value
	}

	return ValueNotFound
}

const (
	OperationalStatusUnknown OperationalStatus = iota
	OperationalStatusOther
	OperationalStatusOK
	OperationalStatusDegraded
	OperationalStatusStressed
	OperationalStatusPredictiveFailure
	OperationalStatusError
	OperationalStatusNonRecoverableError
	OperationalStatusStarting
	OperationalStatusStopping
	OperationalStatusStopped
	OperationalStatusInService
	OperationalStatusNoContact
	OperationalStatusLostCommunication
	OperationalStatusAborted
	OperationalStatusDormant
	OperationalStatusSupportingEntityInError
	OperationalStatusCompleted
	OperationalStatusPowerMode
	OperationalStatusRelocating
)

// operationalStatusString is a map of the operational status to their string representation.
var operationalStatusString = map[OperationalStatus]string{
	OperationalStatusUnknown:                 "Unknown",
	OperationalStatusOther:                   "Other",
	OperationalStatusOK:                      "OK",
	OperationalStatusDegraded:                "Degraded",
	OperationalStatusStressed:                "Stressed",
	OperationalStatusPredictiveFailure:       "PredictiveFailure",
	OperationalStatusError:                   "Error",
	OperationalStatusNonRecoverableError:     "NonRecoverableError",
	OperationalStatusStarting:                "Starting",
	OperationalStatusStopping:                "Stopping",
	OperationalStatusStopped:                 "Stopped",
	OperationalStatusInService:               "InService",
	OperationalStatusNoContact:               "NoContact",
	OperationalStatusLostCommunication:       "LostCommunication",
	OperationalStatusAborted:                 "Aborted",
	OperationalStatusDormant:                 "Dormant",
	OperationalStatusSupportingEntityInError: "SupportingEntityInError",
	OperationalStatusCompleted:               "Completed",
	OperationalStatusPowerMode:               "PowerMode",
	OperationalStatusRelocating:              "Relocating",
}

// String returns the string representation of the operational status.
func (o OperationalStatus) String() string {
	if value, exists := operationalStatusString[o]; exists {
		return value
	}

	return ValueNotFound
}

const (
	OverwritePolicyUnknown OverwritePolicy = iota
	OverwritePolicyOther
	OverwritePolicyWrapWhenFull
	OverwritePolicyClearWhenFull
	OverwritePolicyOverwriteOutdatedWhenNeeded
	OverwritePolicyRemoveOutdatedRecords
	OverwritePolicyOverwriteSpecificRecords
	OverwritePolicyNeverOverwrite
)

// overwritePolicyString is a map of the overwrite policy to their string representation.
var overwritePolicyString = map[OverwritePolicy]string{
	OverwritePolicyUnknown:                     "Unknown",
	OverwritePolicyOther:                       "Other",
	OverwritePolicyWrapWhenFull:                "WrapWhenFull",
	OverwritePolicyClearWhenFull:               "ClearWhenFull",
	OverwritePolicyOverwriteOutdatedWhenNeeded: "OverwriteOutdatedWhenNeeded",
	OverwritePolicyRemoveOutdatedRecords:       "RemoveOutdatedRecords",
	OverwritePolicyOverwriteSpecificRecords:    "OverwriteSpecificRecords",
	OverwritePolicyNeverOverwrite:              "NeverOverwrite",
}

// String returns the string representation of the overwrite policy.
func (o OverwritePolicy) String() string {
	if value, exists := overwritePolicyString[o]; exists {
		return value
	}

	return ValueNotFound
}

const (
	RequestedStateUnknown       RequestedState = 0
	RequestedStateEnabled       RequestedState = 2
	RequestedStateDisabled      RequestedState = 3
	RequestedStateShutDown      RequestedState = 4
	RequestedStateNoChange      RequestedState = 5
	RequestedStateOffline       RequestedState = 6
	RequestedStateTest          RequestedState = 7
	RequestedStateDeferred      RequestedState = 8
	RequestedStateQuiesce       RequestedState = 9
	RequestedStateReboot        RequestedState = 10
	RequestedStateReset         RequestedState = 11
	RequestedStateNotApplicable RequestedState = 12
)

// requestedStateString is a map of the requested state to their string representation.
var requestedStateString = map[RequestedState]string{
	RequestedStateUnknown:       "Unknown",
	RequestedStateEnabled:       "Enabled",
	RequestedStateDisabled:      "Disabled",
	RequestedStateShutDown:      "ShutDown",
	RequestedStateNoChange:      "NoChange",
	RequestedStateOffline:       "Offline",
	RequestedStateTest:          "Test",
	RequestedStateDeferred:      "Deferred",
	RequestedStateQuiesce:       "Quiesce",
	RequestedStateReboot:        "Reboot",
	RequestedStateReset:         "Reset",
	RequestedStateNotApplicable: "NotApplicable",
}

// String returns the string representation of the requested state.
func (r RequestedState) String() string {
	if value, exists := requestedStateString[r]; exists {
		return value
	}

	return ValueNotFound
}

const (
	GetRecordsReturnValueCompletedWithNoError GetRecordsReturnValue = 0
	GetRecordsReturnValueNotSupported         GetRecordsReturnValue = 1
	GetRecordsReturnValueInvalidRecordPointed GetRecordsReturnValue = 2
	GetRecordsReturnValueNoRecordExistsInLog  GetRecordsReturnValue = 3
)

// getRecordsReturnValueString is a map of the GetRecordsReturnValue to their string representation.
var getRecordsReturnValueString = map[GetRecordsReturnValue]string{
	GetRecordsReturnValueCompletedWithNoError: "CompletedWithNoError",
	GetRecordsReturnValueNotSupported:         "NotSupported",
	GetRecordsReturnValueInvalidRecordPointed: "InvalidRecordPointed",
	GetRecordsReturnValueNoRecordExistsInLog:  "NoRecordExistsInLog",
}

// String returns the string representation of the GetRecordsReturnValue value.
func (g GetRecordsReturnValue) String() string {
	if value, exists := getRecordsReturnValueString[g]; exists {
		return value
	}

	return ValueNotFound
}

const (
	PositionToFirstRecordReturnValueCompletedWithNoError PositionToFirstRecordReturnValue = 0
	PositionToFirstRecordReturnValueNotSupported         PositionToFirstRecordReturnValue = 1
	PositionToFirstRecordReturnValueNoRecordExists       PositionToFirstRecordReturnValue = 2
)

// positionToFirstRecordReturnValueString is a map of the PositionToFirstRecordReturnValue to their string representation.
var positionToFirstRecordReturnValueString = map[PositionToFirstRecordReturnValue]string{
	PositionToFirstRecordReturnValueCompletedWithNoError: "CompletedWithNoError",
	PositionToFirstRecordReturnValueNotSupported:         "NotSupported",
	PositionToFirstRecordReturnValueNoRecordExists:       "NoRecordExists",
}

// String returns the string representation of the PositionToFirstRecordReturnValue value.
func (p PositionToFirstRecordReturnValue) String() string {
	if value, exists := positionToFirstRecordReturnValueString[p]; exists {
		return value
	}

	return ValueNotFound
}

func parseEventLogResult(eventlogdata []string) (records []RawEventData, err error) {
	records = make([]RawEventData, 0, len(eventlogdata))

	for _, eventRecord := range eventlogdata {
		// skip a malformed record instead of failing the whole page
		decodedEventRecord, decodeErr := base64.StdEncoding.DecodeString(eventRecord)
		if decodeErr != nil {
			continue
		}

		eventData := RawEventData{}

		buf := bytes.NewReader(decodedEventRecord)

		_ = binary.Read(buf, binary.LittleEndian, &eventData.TimeStamp)
		_ = binary.Read(buf, binary.LittleEndian, &eventData.DeviceAddress)
		_ = binary.Read(buf, binary.LittleEndian, &eventData.EventSensorType)
		_ = binary.Read(buf, binary.LittleEndian, &eventData.EventType)
		_ = binary.Read(buf, binary.LittleEndian, &eventData.EventOffset)
		_ = binary.Read(buf, binary.LittleEndian, &eventData.EventSourceType)
		_ = binary.Read(buf, binary.LittleEndian, &eventData.EventSeverity)
		_ = binary.Read(buf, binary.LittleEndian, &eventData.SensorNumber)
		_ = binary.Read(buf, binary.LittleEndian, &eventData.Entity)
		_ = binary.Read(buf, binary.LittleEndian, &eventData.EntityInstance)

		for i := 13; i < 21; i++ {
			var b uint8

			_ = binary.Read(buf, binary.LittleEndian, &b)

			eventData.EventData = append(eventData.EventData, b)
		}

		records = append(records, eventData)
	}

	return records, err
}

func decodeEventRecord(eventLog []RawEventData) []RefinedEventData {
	refinedEventData := make([]RefinedEventData, len(eventLog))

	for idx, event := range eventLog {
		decodedEvent := RefinedEventData{
			TimeStamp:       time.Unix(int64(event.TimeStamp), 0),
			DeviceAddress:   event.DeviceAddress,
			Description:     decodeEventDetailString(event.EventSensorType, event.EventOffset, event.EventData),
			Entity:          SystemEntityTypes[int(event.Entity)],
			EntityInstance:  event.EntityInstance,
			EventData:       event.EventData,
			EventSensorType: event.EventSensorType,
			EventType:       event.EventType,
			EventOffset:     event.EventOffset,
			EventSourceType: event.EventSourceType,
			EventSeverity:   EventSeverity[int(event.EventSeverity)],
			SensorNumber:    event.SensorNumber,
		}
		refinedEventData[idx] = decodedEvent
	}

	return refinedEventData
}

func decodeEventDetailString(eventSensorType, eventOffset uint8, eventDataField []uint8) string {
	switch eventSensorType {
	case 6:
		value := int(eventDataField[1]) + (int(eventDataField[2]) << 8)

		return fmt.Sprintf("Authentication failed %d times. The system may be under attack.", value)
	case 15:
		{
			if eventDataField[0] == 235 {
				return "Invalid Data"
			}

			if eventOffset == 0 {
				return SystemFirmwareError[int(eventDataField[1])]
			}

			return SystemFirmwareProgress[int(eventDataField[1])]
		}
	case 18:
		// System watchdog event
		if eventDataField[0] == 170 {
			watchdog := fmt.Sprintf("%x%x%x%x-%x%x", eventDataField[4], eventDataField[3], eventDataField[2], eventDataField[1], eventDataField[6], eventDataField[5])
			watchdogCurrentState := WatchdogCurrentStates[int(eventDataField[7])]

			return fmt.Sprintf("Agent watchdog %s-... changed to %s", watchdog, watchdogCurrentState)
		}

		return "Unknown event data field"
	case 30:
		return "No bootable media"
	case 32:
		return "Operating system lockup or power interrupt"
	case 35:
		return "System boot failure"
	case 37:
		return "System firmware started (at least one CPU is properly executing)."
	default:
		return fmt.Sprintf("Unknown Sensor Type #%d", eventSensorType)
	}
}

var EventSeverity = map[int]string{
	0:  "Unspecified",
	1:  "Monitor",
	2:  "Information",
	4:  "OK",
	8:  "Non-critical condition",
	16: "Critical condition",
	32: "Non-recoverable condition",
}

var SystemEntityTypes = map[int]string{
	0:  "Unspecified",
	1:  "Other",
	2:  "Unknown",
	3:  "Processor",
	4:  "Disk",
	5:  "Peripheral",
	6:  "System management module",
	7:  "System board",
	8:  "Memory module",
	9:  "Processor module",
	10: "Power supply",
	11: "Add in card",
	12: "Front panel board",
	13: "Back panel board",
	14: "Power system board",
	15: "Drive backplane",
	16: "System internal expansion board",
	17: "Other system board",
	18: "Processor board",
	19: "Power unit",
	20: "Power module",
	21: "Power management board",
	22: "Chassis back panel board",
	23: "System chassis",
	24: "Sub chassis",
	25: "Other chassis board",
	26: "Disk drive bay",
	27: "Peripheral bay",
	28: "Device bay",
	29: "Fan cooling",
	30: "Cooling unit",
	31: "Cable interconnect",
	32: "Memory device",
	33: "System management software",
	34: "BIOS",
	35: "Intel(r) ME",
	36: "System bus",
	37: "Group",
	38: "Intel(r) ME",
	39: "External environment",
	40: "Battery",
	41: "Processing blade",
	42: "Connectivity switch",
	43: "Processor/memory module",
	44: "I/O module",
	45: "Processor I/O module",
	46: "Management controller firmware",
	47: "IPMI channel",
	48: "PCI bus",
	49: "PCI express bus",
	50: "SCSI bus",
	51: "SATA/SAS bus",
	52: "Processor front side bus",
}

var SystemFirmwareError = map[int]string{
	0:  "Unspecified.",
	1:  "No system memory is physically installed in the system.",
	2:  "No usable system memory, all installed memory has experienced an unrecoverable failure.",
	3:  "Unrecoverable hard-disk/ATAPI/IDE device failure.",
	4:  "Unrecoverable system-board failure.",
	5:  "Unrecoverable diskette subsystem failure.",
	6:  "Unrecoverable hard-disk controller failure.",
	7:  "Unrecoverable PS/2 or USB keyboard failure.",
	8:  "Removable boot media not found.",
	9:  "Unrecoverable video controller failure.",
	10: "No video device detected.",
	11: "Firmware (BIOS) ROM corruption detected.",
	12: "CPU voltage mismatch (processors that share same supply have mismatched voltage requirements)",
	13: "CPU speed matching failure",
}

var SystemFirmwareProgress = map[int]string{
	0:  "Unspecified.",
	1:  "Memory initialization.",
	2:  "Starting hard-disk initialization and test",
	3:  "Secondary processor(s) initialization",
	4:  "User authentication",
	5:  "User-initiated system setup",
	6:  "USB resource configuration",
	7:  "PCI resource configuration",
	8:  "Option ROM initialization",
	9:  "Video initialization",
	10: "Cache initialization",
	11: "SM Bus initialization",
	12: "Keyboard controller initialization",
	13: "Embedded controller/management controller initialization",
	14: "Docking station attachment",
	15: "Enabling docking station",
	16: "Docking station ejection",
	17: "Disabling docking station",
	18: "Calling operating system wake-up vector",
	19: "Starting operating system boot process",
	20: "Baseboard or motherboard initialization",
	21: "reserved",
	22: "Floppy initialization",
	23: "Keyboard test",
	24: "Pointing device test",
	25: "Primary processor initialization",
}

var WatchdogCurrentStates = map[int]string{
	1:  "Not Started",
	2:  "Stopped",
	4:  "Running",
	8:  "Expired",
	16: "Suspended",
}
