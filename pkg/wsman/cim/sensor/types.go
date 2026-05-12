/*********************************************************************
 * Copyright (c) Intel Corporation 2026
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package sensor

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
)

type (
	// Response wraps a single CIM_Sensor SOAP envelope. The embedded
	// *client.Message exposes the raw XMLInput / XMLOutput byte streams for
	// callers that need to log or replay the wire format; the Body field
	// carries the parsed payload (one of PackageResponse, PullResponse, or
	// EnumerateResponse depending on which method was invoked).
	Response struct {
		*client.Message
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}

	// Body is a discriminated union of the three response shapes that the
	// CIM_Sensor resource can return: a single PackageResponse for Get, an
	// EnumerateResponse carrying the EnumerationContext, or a PullResponse
	// carrying zero-or-more PackageResponse items.
	Body struct {
		XMLName           xml.Name `xml:"Body"`
		PullResponse      PullResponse
		EnumerateResponse common.EnumerateResponse
		PackageResponse   PackageResponse
	}

	// PullResponse is the payload of a wsman Pull. AMT firmware returns one
	// CIM_Sensor per Pull call and increments the enumeration context until
	// the sequence is exhausted (signalled by <g:EndOfSequence/>), so callers
	// that want every sensor must loop until SensorItems is empty.
	PullResponse struct {
		XMLName     xml.Name          `xml:"PullResponse"`
		SensorItems []PackageResponse `xml:"Items>CIM_Sensor"`
	}

	// PackageResponse is a single CIM_Sensor instance. The five Key fields
	// (DeviceID, CreationClassName, SystemName, SystemCreationClassName) plus
	// the resource's resource URI form the WSMAN identity used for selector-
	// based addressing.
	PackageResponse struct {
		XMLName xml.Name `xml:"CIM_Sensor"`
		// DeviceID is a Key. On AMT it is formatted as "Sensor N", where N is
		// the 1-based index of the ASF_ALERTDATA structure in the ASF_ALRT
		// ACPI table. MaxLen 64.
		DeviceID string `xml:"DeviceID,omitempty"`
		// CreationClassName is a Key. Always "CIM_Sensor" on instances
		// returned by AMT (the firmware does not subclass). MaxLen 11.
		CreationClassName string `xml:"CreationClassName,omitempty"`
		// SystemName is a Key propagated from CIM_System.Name — the scoping
		// computer system, e.g. "ManagedSystem". MaxLen 256.
		SystemName string `xml:"SystemName,omitempty"`
		// SystemCreationClassName is a Key propagated from
		// CIM_System.CreationClassName — typically "CIM_ComputerSystem".
		// MaxLen 19.
		SystemCreationClassName string `xml:"SystemCreationClassName,omitempty"`
		// ElementName is a user-friendly name for the sensor. On AMT this is
		// often the generic literal "Sensor"; do not rely on it for
		// identification (use DeviceID instead). MaxLen 256.
		ElementName string `xml:"ElementName,omitempty"`
		// OperationalStatus is an indexed array of the current statuses of
		// the sensor. The first element conveys the primary status (see
		// OperationalStatus constants). AMT typically reports a single
		// element here; values beyond index 0 are reserved.
		OperationalStatus []OperationalStatus `xml:"OperationalStatus,omitempty"`
		// HealthState is a coarse 0–30 health indicator independent of
		// OperationalStatus: 0=Unknown, 5=OK, 10=Degraded/Warning,
		// 15=Minor failure, 20=Major failure, 25=Critical failure,
		// 30=Non-recoverable error.
		HealthState HealthState `xml:"HealthState,omitempty"`
		// EnabledState reflects whether the sensor is executing commands or
		// processing requests. AMT-side sensors are typically "Not
		// Applicable" (5) because the host cannot enable or disable them.
		EnabledState EnabledState `xml:"EnabledState,omitempty"`
		// RequestedState is the last state transition requested through any
		// mechanism. AMT-side sensors report "Not Applicable" (12) when
		// transitions are not supported.
		RequestedState RequestedState `xml:"RequestedState,omitempty"`
		// SensorType identifies what physical property the sensor measures.
		// On AMT this is mapped from the ASF_ALERTDATA.EventSensorType byte:
		// 01h -> 02h (Temperature), 02h -> 03h (Voltage), 03h -> 04h
		// (Current). Other ASF event types are reported as 0 (Unknown).
		SensorType SensorType `xml:"SensorType,omitempty"`
		// PossibleStates enumerates the string outputs the sensor can
		// report. For temperature, voltage, and current sensors AMT
		// publishes exactly {"Bad", "Good", "Unknown"} (MaxLen 128 each).
		PossibleStates []string `xml:"PossibleStates,omitempty"`
		// CurrentState is the sensor's current reading and is always one
		// member of PossibleStates. MaxLen 128.
		CurrentState string `xml:"CurrentState,omitempty"`
	}
)

// OperationalStatus enumerates the values defined for
// CIM_ManagedSystemElement.OperationalStatus. The full DMTF value map is
// captured in the constants and operationalStatusMap in decoder.go.
type OperationalStatus int

// HealthState is the 0–30 health continuum defined by
// CIM_ManagedSystemElement.HealthState (see decoder.go for the discrete
// values AMT reports).
type HealthState int

// EnabledState enumerates the values defined for
// CIM_EnabledLogicalElement.EnabledState. AMT sensors typically report
// "Not Applicable" (5).
type EnabledState int

// RequestedState enumerates the values defined for
// CIM_EnabledLogicalElement.RequestedState. AMT sensors typically report
// "Not Applicable" (12).
type RequestedState int

// SensorType enumerates the physical property a sensor measures (see
// decoder.go for the full DMTF value map and notes on the ASF_ALERTDATA
// mapping AMT applies).
type SensorType int
