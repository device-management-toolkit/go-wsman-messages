/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package processor

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
)

type (
	Response struct {
		*client.Message
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}

	Body struct {
		XMLName           xml.Name `xml:"Body"`
		PullResponse      PullResponse
		EnumerateResponse common.EnumerateResponse
		PackageResponse   PackageResponse
	}

	PullResponse struct {
		XMLName      xml.Name          `xml:"PullResponse"`
		PackageItems []PackageResponse `xml:"Items>CIM_Processor"`
	}

	PackageResponse struct {
		XMLName                 xml.Name            `xml:"CIM_Processor"`
		DeviceID                string              `xml:"DeviceID,omitempty"`                // An address or other identifying information to uniquely name the LogicalDevice.
		CreationClassName       string              `xml:"CreationClassName,omitempty"`       // CreationClassName indicates the name of the class or the subclass used in the creation of an instance. When used with the other key properties of this class, this property allows all instances of this class and its subclasses to be uniquely identified.
		SystemName              string              `xml:"SystemName,omitempty"`              // The scoping System's Name.
		SystemCreationClassName string              `xml:"SystemCreationClassName,omitempty"` // The scoping System's CreationClassName.
		ElementName             string              `xml:"ElementName,omitempty"`             // A user-friendly name for the object. This property allows each instance to define a user-friendly name in addition to its key properties, identity data, and description information. Note that the Name property of ManagedSystemElement is also defined as a user-friendly name. But, it is often subclassed to be a Key. It is not reasonable that the same property can convey both identity and a user-friendly name, without inconsistencies. Where Name exists and is not a Key (such as for instances of LogicalDevice), the same information can be present in both the Name and ElementName properties. Note that if there is an associated instance of CIM_EnabledLogicalElementCapabilities, restrictions on this properties may exist as defined in ElementNameMask and MaxElementNameLen properties defined in that class.
		OperationalStatus       []OperationalStatus `xml:"OperationalStatus,omitempty"`       // Indicates the current statuses of the element.
		HealthState             HealthState         `xml:"HealthState,omitempty"`             // Indicates the current health of the element.
		EnabledState            EnabledState        `xml:"EnabledState,omitempty"`            // EnabledState is an integer enumeration that indicates the enabled and disabled states of an element.
		RequestedState          RequestedState      `xml:"RequestedState,omitempty"`          // RequestedState is an integer enumeration that indicates the last requested or desired state for the element, irrespective of the mechanism through which it was requested.
		Role                    string              `xml:"Role,omitempty"`                    // A free-form string that describes the role of the Processor, for example, "Central Processor" or "Math Processor".
		Family                  int                 `xml:"Family,omitempty"`                  // The Processor family type. For example, values include "Pentium(R) processor with MMX(TM) technology" (value=14) and "68040" (value=96).
		OtherFamilyDescription  string              `xml:"OtherFamilyDescription,omitempty"`  // A string that describes the Processor Family type. It is used when the Family property is set to 1 ("Other"). This string should be set to NULL when the Family property is any value other than 1.
		UpgradeMethod           UpgradeMethod       `xml:"UpgradeMethod,omitempty"`           // CPU socket information that includes data on how this Processor can be upgraded (if upgrades are supported). This property is an integer enumeration.
		MaxClockSpeed           int                 `xml:"MaxClockSpeed,omitempty"`           // The maximum speed (in MHz) of this Processor.
		CurrentClockSpeed       int                 `xml:"CurrentClockSpeed,omitempty"`       // The current speed (in MHz) of this Processor.
		Stepping                string              `xml:"Stepping,omitempty"`                // Stepping is a free-form string that indicates the revision level of the Processor within the Processor.Family.
		CPUStatus               CPUStatus           `xml:"CPUStatus,omitempty"`               // The CPUStatus property that indicates the current status of the Processor.
		ExternalBusClockSpeed   int                 `xml:"ExternalBusClockSpeed,omitempty"`   // The speed (in MHz) of the external bus interface (also known as the front side bus).
	}
)

// OperationalStatus is an enumeration of the possible values for the OperationalStatus property.
type OperationalStatus int

// HealthState is an enumeration of the possible values for the HealthState property.
type HealthState int

// EnabledState is an enumeration of the possible values for the EnabledState property.
type EnabledState int

// RequestedState is an enumeration of the possible values for the RequestedState property.
type RequestedState int

// UpgradeMethod is an enumeration of the possible values for the UpgradeMethod property.
type UpgradeMethod int

// CPUStatus is an enumeration of the possible values for the CPUStatus property.
type CPUStatus int
