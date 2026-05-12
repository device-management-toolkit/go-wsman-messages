/*********************************************************************
 * Copyright (c) Intel Corporation 2026
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package fan

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
)

// Response Types.
type (
	// Response is the decoded SOAP envelope returned for any CIM_Fan operation.
	// The embedded *client.Message exposes the raw XML in/out for callers that
	// need to inspect or replay the wire payload.
	Response struct {
		*client.Message
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}

	// Body holds the operation-specific payload. Exactly one of the contained
	// fields is populated per response, depending on which method was invoked
	// (Get → FanResponse, Enumerate → EnumerateResponse, Pull → PullResponse).
	Body struct {
		XMLName           xml.Name `xml:"Body"`
		PullResponse      PullResponse
		EnumerateResponse common.EnumerateResponse
		FanResponse       FanResponse
	}

	// PullResponse carries the items returned by a Pull on an enumeration
	// previously opened with Enumerate. AMT firmware emits one <CIM_Fan>
	// element per fan in the system.
	PullResponse struct {
		XMLName  xml.Name      `xml:"PullResponse"`
		FanItems []FanResponse `xml:"Items>CIM_Fan"`
	}

	// FanResponse is the AMT representation of a single CIM_Fan instance.
	//
	// Capabilities and management of a Fan CoolingDevice. The four key
	// properties (SystemCreationClassName, SystemName, CreationClassName,
	// DeviceID) together uniquely identify the instance among all CIM
	// instances.
	FanResponse struct {
		XMLName xml.Name `xml:"CIM_Fan"`

		// ActiveCooling indicates that the Cooling Device provides active (as
		// opposed to passive) cooling.
		ActiveCooling bool `xml:"ActiveCooling"`

		// CreationClassName indicates the name of the class or subclass used in
		// the creation of an instance. Combined with the other key properties,
		// it allows all instances of this class and its subclasses to be
		// uniquely identified.
		//
		// Key. MaxLen=10.
		CreationClassName string `xml:"CreationClassName"`

		// DesiredSpeed is the currently requested fan speed, in Revolutions per
		// Minute, when a variable speed fan is supported (VariableSpeed=true).
		// The current speed is reported by a CIM_Tachometer sensor associated
		// with the fan through CIM_AssociatedSensor.
		//
		// Units=Revolutions per Minute.
		DesiredSpeed uint64 `xml:"DesiredSpeed"`

		// DeviceID is an address or other identifying information that uniquely
		// names the LogicalDevice. AMT emits this as "Fan N", where N is the
		// zero-based index of the fan among all fans in the system.
		//
		// Key. MaxLen=64.
		DeviceID string `xml:"DeviceID"`

		// ElementName is a user-friendly name for the object, in addition to
		// its key properties.
		//
		// MaxLen=8.
		ElementName string `xml:"ElementName"`

		// EnabledState indicates the enabled and disabled states of the
		// element, including transient states between them.
		//
		// ValueMap={0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		//
		// Values={Unknown, Other, Enabled, Disabled, Shutting Down, Not
		// Applicable, Enabled but Offline, In Test, Deferred, Quiesce, Starting}
		EnabledState EnabledState `xml:"EnabledState"`

		// HealthState indicates the current health of the element. Values are
		// 0..30 on a continuum from "entirely healthy" (5) to "completely
		// non-functional" (30); 0 means "unknown".
		//
		// AMT firmware derives this from SMBIOS type 27 "Device Type and
		// Status"[7:5]:
		//   001b (Other)           → 0  (Unknown)
		//   010b (Unknown)         → 0  (Unknown)
		//   011b (OK)              → 5  (OK)
		//   100b (Non-critical)    → 15 (Minor failure)
		//   101b (Critical)        → 25 (Critical failure)
		//   110b (Non-recoverable) → 30 (Non-recoverable error)
		//
		// ValueMap={0, 5, 10, 15, 20, 25, 30}
		//
		// Values={Unknown, OK, Degraded/Warning, Minor failure, Major failure,
		// Critical failure, Non-recoverable error}
		HealthState HealthState `xml:"HealthState"`

		// OperationalStatus indicates the current statuses of the element. The
		// first entry is the primary status.
		//
		// AMT firmware derives this from SMBIOS type 27 "Device Type and
		// Status"[7:5]:
		//   001b (Other)           → 0 (Unknown)
		//   010b (Unknown)         → 0 (Unknown)
		//   011b (OK)              → 2 (OK)
		//   100b (Non-critical)    → 6 (Error)
		//   101b (Critical)        → 6 (Error)
		//   110b (Non-recoverable) → 7 (Non-recoverable error)
		//
		// ValueMap={0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16,
		// 17, 18, 19}
		//
		// Values={Unknown, Other, OK, Degraded, Stressed, Predictive Failure,
		// Error, Non-Recoverable Error, Starting, Stopping, Stopped, In
		// Service, No Contact, Lost Communication, Aborted, Dormant, Supporting
		// Entity in Error, Completed, Power Mode, Relocating}
		OperationalStatus []OperationalStatus `xml:"OperationalStatus"`

		// RequestedState indicates the last requested or desired state for the
		// element, irrespective of the mechanism through which it was
		// requested. The actual state is reflected by EnabledState.
		//
		// If the last requested state is unknown, this property is 0
		// ("Unknown"). If the EnabledLogicalElement does not track requested
		// states, the property is 12 ("Not Applicable") — which is what AMT
		// reports for CIM_Fan instances.
		//
		// ValueMap={0, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
		//
		// Values={Unknown, Enabled, Disabled, Shut Down, No Change, Offline,
		// Test, Deferred, Quiesce, Reboot, Reset, Not Applicable}
		RequestedState RequestedState `xml:"RequestedState"`

		// SystemCreationClassName is the scoping System's CreationClassName
		// (propagated from CIM_System.CreationClassName).
		//
		// Key. MaxLen=20.
		SystemCreationClassName string `xml:"SystemCreationClassName"`

		// SystemName is the scoping System's Name (propagated from
		// CIM_System.Name).
		//
		// Key. MaxLen=256.
		SystemName string `xml:"SystemName"`

		// VariableSpeed indicates whether the fan supports variable speeds.
		// When false, DesiredSpeed has no meaning.
		VariableSpeed bool `xml:"VariableSpeed"`
	}

	// OperationalStatus is an integer enumeration of the current statuses of
	// the element. See FanResponse.OperationalStatus for the value mapping.
	OperationalStatus int

	// EnabledState is an integer enumeration of the enabled/disabled state of
	// the element. See FanResponse.EnabledState for the value mapping.
	EnabledState int

	// RequestedState is an integer enumeration of the last requested or desired
	// state for the element. See FanResponse.RequestedState for the value
	// mapping.
	RequestedState int

	// HealthState is an integer enumeration of the current health of the
	// element on a 0..30 continuum. See FanResponse.HealthState for the value
	// mapping, including the SMBIOS type 27 translation AMT applies.
	HealthState int
)
