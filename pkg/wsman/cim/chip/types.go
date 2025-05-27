/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package chip

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
		ChipItems []PackageResponse `xml:"Items>CIM_Chip"`
	}

	PackageResponse struct {
		XMLName           xml.Name            `xml:"CIM_Chip"`
		CanBeFRUed        bool                `xml:"CanBeFRUed"`        // Boolean that indicates whether this PhysicalElement can be FRUed (TRUE) or not (FALSE).
		CreationClassName string              `xml:"CreationClassName"` // CreationClassName indicates the name of the class or the subclass used in the creation of an instance. When used with the other key properties of this class, this property allows all instances of this class and its subclasses to be uniquely identified.
		ElementName       string              `xml:"ElementName"`
		Manufacturer      string              `xml:"Manufacturer"`      // The name of the organization responsible for producing the PhysicalElement. This organization might be the entity from whom the Element is purchased, but this is not necessarily true. The latter information is contained in the Vendor property of CIM_Product.
		OperationalStatus []OperationalStatus `xml:"OperationalStatus"` // Indicates the current statuses of the element.
		Tag               string              `xml:"Tag"`               // An arbitrary string that uniquely identifies the Physical Element and serves as the key of the Element. The Tag property can contain information such as asset tag or serial number data. The key for PhysicalElement is placed very high in the object hierarchy in order to independently identify the hardware or entity, regardless of physical placement in or on Cabinets, Adapters, and so on. For example, a hotswappable or removable component can be taken from its containing (scoping) Package and be temporarily unused. The object still continues to exist and can even be inserted into a different scoping container. Therefore, the key for Physical Element is an arbitrary string and is defined independently of any placement or location-oriented hierarchy.
		Version           string              `xml:"Version"`           // A string that indicates the version of the PhysicalElement.
	}

	// OperationalStatus is the current statuses of the element.
	OperationalStatus int
)
