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
	"strconv"
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
	records = make([]RawEventData, len(eventlogdata))

	for idx, eventRecord := range eventlogdata {
		decodedEventRecord, err := base64.StdEncoding.DecodeString(eventRecord)
		if err != nil {
			return records, err
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

		records[idx] = eventData
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
	case 5:
		if eventOffset == 0 {
			return "Case intrusion"
		}
	case 6:
		value := int(eventDataField[1]) + (int(eventDataField[2]) << 8)

		return fmt.Sprintf("Authentication failed %d times. The system may be under attack.", value)
	case 15:
		{
			if eventDataField[0] == 235 {
				return InvalidData
			}

			if eventOffset == 0 {
				return SystemFirmwareError[int(eventDataField[1])]
			}

			if eventOffset == 3 {
				if eventDataField[0] == 170 && eventDataField[1] == 48 {
					return fmt.Sprintf("One Click Recovery: %s", OCRErrorEvents[int(eventDataField[2])])
				} else if eventDataField[0] == 170 && eventDataField[1] == 64 {
					return PlatformEraseErrorEvents[int(eventDataField[2])]
				}

				return OEMSpecificFirmwareErrorEvent
			}

			if eventOffset == 5 {
				if eventDataField[0] == 170 && eventDataField[1] == 48 {
					if eventDataField[2] == 1 {
						return fmt.Sprintf("One Click Recovery: CSME Boot Option %d:%s added successfully", eventDataField[3], OCRSource[int(eventDataField[3])])
					}
					if eventDataField[2] < 7 {
						return fmt.Sprintf("One Click Recovery: %s", OCRProgressEvents[int(eventDataField[2])])
					}

					return fmt.Sprintf("One Click Recovery: Unknown progress event %d", eventDataField[2])
				}
				if eventDataField[0] == 170 && eventDataField[1] == 64 {
					if eventDataField[2] == 1 {
						if eventDataField[3] == 2 {
							return "Started erasing Device SSD"
						}
						if eventDataField[3] == 3 {
							return "Started erasing Device TPM"
						}
						if eventDataField[3] == 5 {
							return "Started erasing Device BIOS Reload of Golden Config"
						}
					}
					if eventDataField[2] == 2 {
						if eventDataField[3] == 2 {
							return "Erasing Device SSD ended successfully"
						}
						if eventDataField[3] == 3 {
							return "Erasing Device TPM ended successfully"
						}
						if eventDataField[3] == 5 {
							return "Erasing Device BIOS Reload of Golden Config ended successfully"
						}
					}
					if eventDataField[2] == 3 {
						return "Beginning Platform Erase"
					}
					if eventDataField[2] == 4 {
						return "Clear Reserved Parameters"
					}
					if eventDataField[2] == 5 {
						return "All setting decremented"
					}
				}

				return OEMSpecificFirmwareErrorEvent
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
		if eventDataField[0] == 64 {
			return "BIOS POST (Power On Self-Test) Watchdog Timeout." // 64,2,252,84,89,0,0,0
		}

		return "System boot failure"
	case 36:
		var handle uint32
		handle = (uint32(eventDataField[1]) << 24) + (uint32(eventDataField[2]) << 16) + (uint32(eventDataField[3]) << 8) + uint32(eventDataField[4])

		var nic = "#" + strconv.Itoa(int(eventDataField[0]))

		if eventDataField[0] == 0xAA {
			nic = "wired"
		}
		// TODO: Add wireless *****
		//if (eventDataField[0] == 0xAA) nic = "wireless";

		if handle == 4294967293 {
			return fmt.Sprintf("All received packet filter was matched on %s interface.", nic)
		}

		if handle == 4294967292 {
			return fmt.Sprintf("All outbound packet filter was matched on %s interface.", nic)
		}

		if handle == 4294967290 {
			return fmt.Sprintf("Spoofed packet filter was matched on %s interface.", nic)
		}

		return fmt.Sprintf("Filter %d was matched on %s interface.", handle, nic)
	case 37:
		return "System firmware started (at least one CPU is properly executing)."
	case 192:
		if eventOffset == 0 && eventDataField[0] == 170 && eventDataField[1] == 48 {
			return SOLIDERStatus[int(eventDataField[2])]
		}
		if eventDataField[2] == 0 || eventDataField[2] == 2 {
			return SecurityPolicyEvent[int(eventDataField[2])]
		} else {
			return "Security policy invoked."
		}
	case 193:
		if (eventDataField[0] == 0xAA) && (eventDataField[1] == 0x30) && (eventDataField[2] == 0x00) && (eventDataField[3] == 0x00) {
			return "User request for remote connection."
		}
		if (eventDataField[0] == 0xAA) && (eventDataField[1] == 0x20) && (eventDataField[2] == 0x03) && (eventDataField[3] == 0x01) {
			return "EAC error: attempt to get posture while NAC in Intel® AMT is disabled."
		} // eventDataField = 0xAA20030100000000
		if (eventDataField[0] == 0xAA) && (eventDataField[1] == 0x20) && (eventDataField[2] == 0x04) && (eventDataField[3] == 0x00) {
			return "HWA Error: general error"
		} // Used to be "Certificate revoked." but don"t know the source of this.
	}
	return fmt.Sprintf("Unknown Sensor Type #%d", eventSensorType)
}

const InvalidData = "Invalid Data"
const OEMSpecificFirmwareErrorEvent = "OEM Specific Firmware Error event"

var SecurityPolicyEvent = map[int]string{
	0: "Security policy invoked. Some or all network traffic (TX) was stopped.",
	2: "Security policy invoked. Some or all network traffic (RX) was stopped.",
}

var SOLIDERStatus = map[int]string{
	0: "A remote Serial Over LAN session was established.",
	1: "Remote Serial Over LAN session finished.  User control was restored.",
	2: "A remote IDE-Redirection session was established.",
	3: "Remote IDE-Redirection session finished. User control was restored.",
}

var PlatformEraseErrorEvents = map[int]string{
	1: "Got an error erasing Device SSD",
	2: "Erasing Device TPM is not supported",
	3: "Reached Max Counter",
}

var OCRProgressEvents = map[int]string{
	0: "Boot parameters received from CSME",
	1: "CSME Boot Option % added successfully",
	2: "HTTPS URI name resolved",
	3: "HTTPS connected successfully",
	4: "HTTPSBoot download is completed",
	5: "Attempt to boot",
	6: "Exit boot services",
}

var OCRSource = map[int]string{
	1: "",
	2: "HTTPS",
	4: "Local PBA",
	8: "WinRE",
}

var OCRErrorEvents = map[int]string{
	0: "",
	1: "No network connection available",
	2: "Name resolution of URI failed",
	3: "Connect to URI failed",
	4: "OEM app not found at local URI",
	5: "HTTPS TLS Auth failed",
	6: "HTTPS Digest Auth failed",
	7: "Verified boot failed (bad image)",
	8: "HTTPS Boot File not found",
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
