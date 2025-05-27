/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package card

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
)

type Package struct {
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
		PullResponse      PullResponse
		EnumerateResponse common.EnumerateResponse
		PackageResponse   PackageResponse
	}

	PullResponse struct {
		XMLName   xml.Name          `xml:"PullResponse"`
		CardItems []PackageResponse `xml:"Items>CIM_Card"`
	}

	PackageResponse struct {
		XMLName           xml.Name            `xml:"CIM_Card"`
		CanBeFRUed        bool                `xml:"CanBeFRUed"`        // Boolean that indicates whether this PhysicalElement can be FRUed (TRUE) or not (FALSE).
		CreationClassName string              `xml:"CreationClassName"` // CreationClassName indicates the name of the class or the subclass used in the creation of an instance.
		ElementName       string              `xml:"ElementName"`
		Manufacturer      string              `xml:"Manufacturer"`      // The name of the organization responsible for producing the PhysicalElement.
		Model             string              `xml:"Model"`             // The name by which the PhysicalElement is generally known.
		OperationalStatus []OperationalStatus `xml:"OperationalStatus"` // Indicates the current statuses of the element
		PackageType       PackageType         `xml:"PackageType"`       // Enumeration defining the type of the PhysicalPackage. Note that this enumeration expands on the list in the Entity MIB (the attribute, entPhysicalClass). The numeric values are consistent with CIM's enum numbering guidelines, but are slightly different than the MIB's values.
		SerialNumber      string              `xml:"SerialNumber"`      // A manufacturer-allocated number used to identify the Physical Element.
		Tag               string              `xml:"Tag"`               // An arbitrary string that uniquely identifies the Physical Element and serves as the key of the Element.
		Version           string              `xml:"Version"`           // A string that indicates the version of the PhysicalElement.
	}

	// OperationalStatus is the current statuses of the element.
	OperationalStatus int
	// PackageType is the type of the PhysicalPackage.
	PackageType int
)
