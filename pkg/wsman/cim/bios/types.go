/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package bios

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
)

type Element struct {
	base message.Base
}

// Response Types.
type (
	Response struct {
		*client.Message
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}

	Body struct {
		XMLName           xml.Name `xml:"Body"`
		GetResponse       BiosElement
		EnumerateResponse common.EnumerateResponse
		PullResponse      PullResponse
	}

	BiosElement struct {
		XMLName               xml.Name              `xml:"CIM_BIOSElement"`
		TargetOperatingSystem TargetOperatingSystem `xml:"TargetOperatingSystem"` // The TargetOperatingSystem property specifies the element's operating system environment.
		SoftwareElementID     string                `xml:"SoftwareElementID"`     // This is an identifier for the SoftwareElement and is designed to be used in conjunction with other keys to create a unique representation of the element.
		SoftwareElementState  SoftwareElementState  `xml:"SoftwareElementState"`  // The SoftwareElementState is defined in this model to identify various states of a SoftwareElement's life cycle.
		Name                  string                `xml:"Name"`                  // The name used to identify this SoftwareElement.
		OperationalStatus     []OperationalStatus   `xml:"OperationalStatus"`     // Indicates the current statuses of the element.
		ElementName           string                `xml:"ElementName"`           // A user-friendly name for the object. This property allows each instance to define a user-friendly name in addition to its key properties, identity data, and description information. Note that the Name property of ManagedSystemElement is also defined as a user-friendly name. But, it is often subclassed to be a Key. It is not reasonable that the same property can convey both identity and a user-friendly name, without inconsistencies. Where Name exists and is not a Key (such as for instances of LogicalDevice), the same information can be present in both the Name and ElementName properties. Note that if there is an associated instance of CIM_EnabledLogicalElementCapabilities, restrictions on this properties may exist as defined in ElementNameMask and MaxElementNameLen properties defined in that class.
		Version               string                `xml:"Version"`               // The version of the BIOS software image.
		Manufacturer          string                `xml:"Manufacturer"`          // The manufacturer of the BIOS software image.
		PrimaryBIOS           bool                  `xml:"PrimaryBIOS"`           // If true, this is the primary BIOS of the ComputerSystem.
		ReleaseDate           Time                  `xml:"ReleaseDate"`           // Date that this BIOS was released.
	}

	Time struct {
		DateTime string `xml:"Datetime"`
	}

	PullResponse struct {
		XMLName          xml.Name      `xml:"PullResponse"`
		BiosElementItems []BiosElement `xml:"Items>CIM_BIOSElement"`
	}

	// TargetOperatingSystem is the element's operating system environment.
	TargetOperatingSystem int
	// SoftwareElementState is defined in this model to identify various states of a SoftwareElement's life cycle.
	SoftwareElementState int
	// OperationalStatus indicates the current statuses of the element.
	OperationalStatus int
)
