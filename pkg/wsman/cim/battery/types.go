/*********************************************************************
 * Copyright (c) Intel Corporation 2026
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package battery

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
)

// Response Types.
type (
	// Response is the decoded SOAP envelope returned for any CIM_Battery
	// operation. The embedded *client.Message exposes the raw XMLInput
	// and XMLOutput captured by the transport so callers can audit the
	// exact bytes on the wire.
	Response struct {
		*client.Message
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}

	// Body holds the union of possible response bodies. For a given call
	// only one of PackageResponse (Get), EnumerateResponse (Enumerate),
	// or PullResponse (Pull) will be populated.
	Body struct {
		XMLName           xml.Name `xml:"Body"`
		PullResponse      PullResponse
		EnumerateResponse common.EnumerateResponse
		PackageResponse   PackageResponse
	}

	// PullResponse carries the batch of CIM_Battery instances returned
	// by a Pull. EndOfSequence is populated on the final page of an
	// enumeration; if it is empty the caller should issue another Pull
	// to drain remaining instances.
	PullResponse struct {
		XMLName       xml.Name          `xml:"PullResponse"`
		BatteryItems  []PackageResponse `xml:"Items>CIM_Battery"`
		EndOfSequence xml.Name          `xml:"EndOfSequence"`
	}

	// PackageResponse is a single CIM_Battery instance. The four key
	// properties (DeviceID, CreationClassName, SystemName,
	// SystemCreationClassName) together uniquely identify a battery in
	// the managed system.
	PackageResponse struct {
		XMLName xml.Name `xml:"CIM_Battery"`
		// DeviceID is the key that uniquely names the LogicalDevice (Key,
		// MaxLen=14).
		DeviceID string `xml:"DeviceID,omitempty"`
		// CreationClassName indicates the class or subclass used to create
		// the instance. For this class the value is "CIM_Battery" (Key,
		// MaxLen=12).
		CreationClassName string `xml:"CreationClassName,omitempty"`
		// SystemName is the scoping System's Name, propagated from
		// CIM_System.Name (Key, MaxLen=15).
		SystemName string `xml:"SystemName,omitempty"`
		// SystemCreationClassName is the scoping System's
		// CreationClassName, propagated from CIM_System.CreationClassName
		// (Key, MaxLen=20).
		SystemCreationClassName string `xml:"SystemCreationClassName,omitempty"`
		// ElementName is an optional user-friendly name for the instance
		// (MaxLen=60). May be empty.
		ElementName string `xml:"ElementName,omitempty"`
		// OperationalStatus reports the current statuses of the element.
		// Indexed array — the first entry is the primary status.
		OperationalStatus []OperationalStatus `xml:"OperationalStatus,omitempty"`
		// HealthState reports the current health on a 0–30 continuum
		// (5 = OK, 30 = non-recoverable).
		HealthState HealthState `xml:"HealthState,omitempty"`
		// EnabledState reports the actual enabled/disabled state. Compare
		// with RequestedState to detect transitions.
		EnabledState EnabledState `xml:"EnabledState,omitempty"`
		// RequestedState reports the last requested or desired state for
		// the element. See EnabledState for the actual state.
		RequestedState RequestedState `xml:"RequestedState,omitempty"`
		// BatteryStatus describes the battery charge status. Values 6–9
		// (Charging variants) are deprecated in favour of ChargingStatus;
		// value 10 (Undefined) is deprecated in favour of 2 (Unknown) and
		// indicates "no battery installed" in DMI — in that case the
		// instance should not exist at all.
		BatteryStatus BatteryStatus `xml:"BatteryStatus,omitempty"`
		// EstimatedChargeRemaining is the percentage of full charge
		// remaining (0–100).
		EstimatedChargeRemaining uint16 `xml:"EstimatedChargeRemaining,omitempty"`
		// Chemistry describes the battery chemistry (e.g. Lithium-ion).
		Chemistry Chemistry `xml:"Chemistry,omitempty"`
		// DesignCapacity is the battery's design capacity in mWatt-hours.
		// 0 means the property is not supported.
		DesignCapacity uint32 `xml:"DesignCapacity,omitempty"`
		// FullChargeCapacity is the current full-charge capacity in
		// mWatt-hours. End-of-life is typically signalled when this drops
		// below 80% of DesignCapacity. 0 means the property is not
		// supported.
		FullChargeCapacity uint32 `xml:"FullChargeCapacity,omitempty"`
		// DesignVoltage is the design voltage in mVolts. 0 means the
		// property is not supported.
		DesignVoltage uint64 `xml:"DesignVoltage,omitempty"`
		// ChargingStatus reports whether the battery is currently
		// charging, discharging, or idle. Supersedes the deprecated
		// "Charging*" variants of BatteryStatus.
		ChargingStatus ChargingStatus `xml:"ChargingStatus,omitempty"`
	}
)

// OperationalStatus is the CIM operational-status enumeration. The first
// entry of CIM_Battery.OperationalStatus is the primary status; trailing
// entries qualify it (e.g. "Completed" combined with "OK").
type OperationalStatus int

// HealthState is the CIM health-state enumeration on a 0–30 continuum,
// where 0 = Unknown, 5 = OK, and 30 = Non-recoverable.
type HealthState int

// EnabledState is the CIM enabled-state enumeration describing whether
// the element is currently enabled, disabled, or in a transitional
// state. Compare with RequestedState to spot in-flight transitions.
type EnabledState int

// RequestedState is the CIM requested-state enumeration describing the
// last requested or desired state. When EnabledState is "Not Applicable"
// this value has no meaning.
type RequestedState int

// BatteryStatus describes the charge status of the battery. Several
// values (6–9, 10) are deprecated by the CIM schema — see the
// per-constant documentation in decoder.go.
type BatteryStatus int

// Chemistry enumerates the battery's chemistry (Lead Acid, Lithium-ion,
// etc).
type Chemistry int

// ChargingStatus reports whether the battery is currently charging,
// discharging, or idle. Supersedes the deprecated "Charging*" variants
// of BatteryStatus.
type ChargingStatus int
