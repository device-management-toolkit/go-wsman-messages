/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package messagelog

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCapabilities_String(t *testing.T) {
	tests := []struct {
		state    Capabilities
		expected string
	}{
		{CapabilitiesUnknown, "Unknown"},
		{CapabilitiesOther, "Other"},
		{CapabilitiesWriteRecordSupported, "WriteRecordSupported"},
		{CapabilitiesDeleteRecordSupported, "DeleteRecordSupported"},
		{CapabilitiesCanMoveBackwardInLog, "CanMoveBackwardInLog"},
		{CapabilitiesFreezeLogSupported, "FreezeLogSupported"},
		{CapabilitiesClearLogSupported, "ClearLogSupported"},
		{CapabilitiesSupportsAddressingByOrdinalRecordNumber, "SupportsAddressingByOrdinalRecordNumber"},
		{CapabilitiesVariableLengthRecordsSupported, "VariableLengthRecordsSupported"},
		{CapabilitiesVariableFormatsForRecords, "VariableFormatsForRecords"},
		{CapabilitiesCanFlagRecordsForOverwrite, "CanFlagRecordsForOverwrite"},
		{Capabilities(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestCharacterSet_String(t *testing.T) {
	tests := []struct {
		state    CharacterSet
		expected string
	}{
		{CharacterSetUnknown, "Unknown"},
		{CharacterSetOther, "Other"},
		{CharacterSetASCII, "ASCII"},
		{CharacterSetUnicode, "Unicode"},
		{CharacterSetISO2022, "ISO2022"},
		{CharacterSetISO8859, "ISO8859"},
		{CharacterSetExtendedUNIXCode, "ExtendedUNIXCode"},
		{CharacterSetUTF8, "UTF8"},
		{CharacterSetUCS2, "UCS2"},
		{CharacterSetBitmappedData, "BitmappedData"},
		{CharacterSetOctetString, "OctetString"},
		{CharacterSetDefinedByIndividualRecords, "DefinedByIndividualRecords"},
		{CharacterSet(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestEnabledDefault_String(t *testing.T) {
	tests := []struct {
		state    EnabledDefault
		expected string
	}{
		{EnabledDefaultEnabled, "Enabled"},
		{EnabledDefaultDisabled, "Disabled"},
		{EnabledDefaultNotApplicable, "NotApplicable"},
		{EnabledDefaultEnabledButOffline, "EnabledButOffline"},
		{EnabledDefaultNoDefault, "NoDefault"},
		{EnabledDefaultQuiesce, "Quiesce"},
		{EnabledDefault(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestEnabledState_String(t *testing.T) {
	tests := []struct {
		state    EnabledState
		expected string
	}{
		{EnabledStateUnknown, "Unknown"},
		{EnabledStateOther, "Other"},
		{EnabledStateEnabled, "Enabled"},
		{EnabledStateDisabled, "Disabled"},
		{EnabledStateShuttingDown, "ShuttingDown"},
		{EnabledStateNotApplicable, "NotApplicable"},
		{EnabledStateEnabledButOffline, "EnabledButOffline"},
		{EnabledStateInTest, "InTest"},
		{EnabledStateDeferred, "Deferred"},
		{EnabledStateQuiesce, "Quiesce"},
		{EnabledStateStarting, "Starting"},
		{EnabledState(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestHealthState_String(t *testing.T) {
	tests := []struct {
		state    HealthState
		expected string
	}{
		{HealthStateUnknown, "Unknown"},
		{HealthStateOK, "OK"},
		{HealthStateDegraded, "Degraded"},
		{HealthStateMinorFailure, "MinorFailure"},
		{HealthStateMajorFailure, "MajorFailure"},
		{HealthStateCriticalFailure, "CriticalFailure"},
		{HealthStateNonRecoverableError, "NonRecoverableError"},
		{HealthState(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestLastChange_String(t *testing.T) {
	tests := []struct {
		state    LastChange
		expected string
	}{
		{LastChangeUnknown, "Unknown"},
		{LastChangeAdd, "Add"},
		{LastChangeDelete, "Delete"},
		{LastChangeModify, "Modify"},
		{LastChangeLogCleared, "LogCleared"},
		{LastChange(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestLogState_String(t *testing.T) {
	tests := []struct {
		state    LogState
		expected string
	}{
		{LogStateUnknown, "Unknown"},
		{LogStateNormal, "Normal"},
		{LogStateErasing, "Erasing"},
		{LogStateNotApplicable, "NotApplicable"},
		{LogState(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestOperationalStatus_String(t *testing.T) {
	tests := []struct {
		state    OperationalStatus
		expected string
	}{
		{OperationalStatusUnknown, "Unknown"},
		{OperationalStatusOther, "Other"},
		{OperationalStatusOK, "OK"},
		{OperationalStatusDegraded, "Degraded"},
		{OperationalStatusStressed, "Stressed"},
		{OperationalStatusPredictiveFailure, "PredictiveFailure"},
		{OperationalStatusError, "Error"},
		{OperationalStatusNonRecoverableError, "NonRecoverableError"},
		{OperationalStatusStarting, "Starting"},
		{OperationalStatusStopping, "Stopping"},
		{OperationalStatusStopped, "Stopped"},
		{OperationalStatusInService, "InService"},
		{OperationalStatusNoContact, "NoContact"},
		{OperationalStatusLostCommunication, "LostCommunication"},
		{OperationalStatusAborted, "Aborted"},
		{OperationalStatusDormant, "Dormant"},
		{OperationalStatusSupportingEntityInError, "SupportingEntityInError"},
		{OperationalStatusCompleted, "Completed"},
		{OperationalStatusPowerMode, "PowerMode"},
		{OperationalStatusRelocating, "Relocating"},
		{OperationalStatus(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestOverwritePolicy_String(t *testing.T) {
	tests := []struct {
		state    OverwritePolicy
		expected string
	}{
		{OverwritePolicyUnknown, "Unknown"},
		{OverwritePolicyOther, "Other"},
		{OverwritePolicyWrapWhenFull, "WrapWhenFull"},
		{OverwritePolicyClearWhenFull, "ClearWhenFull"},
		{OverwritePolicyOverwriteOutdatedWhenNeeded, "OverwriteOutdatedWhenNeeded"},
		{OverwritePolicyRemoveOutdatedRecords, "RemoveOutdatedRecords"},
		{OverwritePolicyOverwriteSpecificRecords, "OverwriteSpecificRecords"},
		{OverwritePolicyNeverOverwrite, "NeverOverwrite"},
		{OverwritePolicy(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestRequestedState_String(t *testing.T) {
	tests := []struct {
		state    RequestedState
		expected string
	}{
		{RequestedStateUnknown, "Unknown"},
		{RequestedStateEnabled, "Enabled"},
		{RequestedStateDisabled, "Disabled"},
		{RequestedStateShutDown, "ShutDown"},
		{RequestedStateNoChange, "NoChange"},
		{RequestedStateOffline, "Offline"},
		{RequestedStateTest, "Test"},
		{RequestedStateDeferred, "Deferred"},
		{RequestedStateQuiesce, "Quiesce"},
		{RequestedStateReboot, "Reboot"},
		{RequestedStateReset, "Reset"},
		{RequestedStateNotApplicable, "NotApplicable"},
		{RequestedState(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestGetRecordsReturnValue_String(t *testing.T) {
	tests := []struct {
		state    GetRecordsReturnValue
		expected string
	}{
		{GetRecordsReturnValueCompletedWithNoError, "CompletedWithNoError"},
		{GetRecordsReturnValueNotSupported, "NotSupported"},
		{GetRecordsReturnValueInvalidRecordPointed, "InvalidRecordPointed"},
		{GetRecordsReturnValueNoRecordExistsInLog, "NoRecordExistsInLog"},
		{GetRecordsReturnValue(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestPositionToFirstRecordReturnValue_String(t *testing.T) {
	tests := []struct {
		state    PositionToFirstRecordReturnValue
		expected string
	}{
		{PositionToFirstRecordReturnValueCompletedWithNoError, "CompletedWithNoError"},
		{PositionToFirstRecordReturnValueNotSupported, "NotSupported"},
		{PositionToFirstRecordReturnValueNoRecordExists, "NoRecordExists"},
		{PositionToFirstRecordReturnValue(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestConvertToEventLogResult(t *testing.T) {
	records := []string{"Y8iYZf8GbwVoEP8mYaoKAAAAAAAA", "IgYBZf8PbwJoAf8iAEAHAAAAAAAA", "IgYBZf8PbwJoAf8iAEAHAAAAAAAA"}
	expectedResult := []RawEventData{{TimeStamp: 0x6598c863, DeviceAddress: 0xff, EventSensorType: 0x6, EventType: 0x6f, EventOffset: 0x5, EventSourceType: 0x68, EventSeverity: 0x10, SensorNumber: 0xff, Entity: 0x26, EntityInstance: 0x61, EventData: []uint8{0xaa, 0xa, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}}, {TimeStamp: 0x65010622, DeviceAddress: 0xff, EventSensorType: 0xf, EventType: 0x6f, EventOffset: 0x2, EventSourceType: 0x68, EventSeverity: 0x1, SensorNumber: 0xff, Entity: 0x22, EntityInstance: 0x0, EventData: []uint8{0x40, 0x7, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}}, {TimeStamp: 0x65010622, DeviceAddress: 0xff, EventSensorType: 0xf, EventType: 0x6f, EventOffset: 0x2, EventSourceType: 0x68, EventSeverity: 0x1, SensorNumber: 0xff, Entity: 0x22, EntityInstance: 0x0, EventData: []uint8{0x40, 0x7, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}}}
	expectedDecodedResult := []RefinedEventData{{TimeStamp: time.Unix(int64(0x6598c863), 0), DeviceAddress: 0xff, Description: "Authentication failed 10 times. The system may be under attack.", Entity: "Intel(r) ME", EntityInstance: 0x61, EventData: []uint8{0xaa, 0xa, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, EventSensorType: 0x6, EventType: 0x6f, EventOffset: 0x5, EventSourceType: 0x68, EventSeverity: "Critical condition", SensorNumber: 0xff}, {TimeStamp: time.Unix(int64(0x65010622), 0), DeviceAddress: 0xff, Description: "PCI resource configuration", Entity: "BIOS", EntityInstance: 0x0, EventData: []uint8{0x40, 0x7, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, EventSensorType: 0xf, EventType: 0x6f, EventOffset: 0x2, EventSourceType: 0x68, EventSeverity: "Monitor", SensorNumber: 0xff}, {TimeStamp: time.Unix(int64(0x65010622), 0), DeviceAddress: 0xff, Description: "PCI resource configuration", Entity: "BIOS", EntityInstance: 0x0, EventData: []uint8{0x40, 0x7, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, EventSensorType: 0xf, EventType: 0x6f, EventOffset: 0x2, EventSourceType: 0x68, EventSeverity: "Monitor", SensorNumber: 0xff}}

	result, err := parseEventLogResult(records)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	decodedResult := decodeEventRecord(result)

	assert.Equal(t, expectedResult, result)
	assert.Equal(t, expectedDecodedResult, decodedResult)
}

func TestDecodeEventDetailString(t *testing.T) {
	tests := []struct {
		eventSensorType uint8
		eventOffset     uint8
		eventDataField  []uint8
		expected        string
	}{
		{6, 0, []uint8{0, 5, 0}, "Authentication failed 5 times. The system may be under attack."},
		{6, 0, []uint8{0, 1, 1}, "Authentication failed 257 times. The system may be under attack."},
		{15, 0, []uint8{235}, "Invalid Data"},
		{15, 0, []uint8{0, 1}, "No system memory is physically installed in the system."},
		{15, 1, []uint8{0, 2}, "Starting hard-disk initialization and test"},
		{18, 0, []uint8{170, 1, 2, 3, 4, 5, 6, 1}, "Agent watchdog 4321-65-... changed to Not Started"},
		{30, 0, nil, "No bootable media"},
		{32, 0, nil, "Operating system lockup or power interrupt"},
		{35, 0, nil, "System boot failure"},
		{37, 0, nil, "System firmware started (at least one CPU is properly executing)."},
		{0, 0, nil, "Unknown Sensor Type #0"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("eventSensorType=%d/eventOffset=%d", test.eventSensorType, test.eventOffset), func(t *testing.T) {
			result := decodeEventDetailString(test.eventSensorType, test.eventOffset, test.eventDataField)
			if result != test.expected {
				t.Errorf("Expected %q but got %q", test.expected, result)
			}
		})
	}
}

func TestParseEventLogResultSkipsInvalidRecord(t *testing.T) {
	records := []string{"Y8iYZf8GbwVoEP8mYaoKAAAAAAAA", "not base64!", "IgYBZf8PbwJoAf8iAEAHAAAAAAAA"}

	result, err := parseEventLogResult(records)

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, uint32(0x6598c863), result[0].TimeStamp)
	assert.Equal(t, uint32(0x65010622), result[1].TimeStamp)
}
