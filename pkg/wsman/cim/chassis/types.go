/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package chassis

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
)

type Package struct {
	base message.Base
}

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
		PackageItems []PackageResponse `xml:"Items>CIM_Chassis"`
	}

	PackageResponse struct {
		XMLName            xml.Name            `xml:"CIM_Chassis"`
		Version            string              `xml:"Version"`      // A string that indicates the version of the PhysicalElement.
		SerialNumber       string              `xml:"SerialNumber"` // A manufacturer-allocated number used to identify the Physical Element.
		Model              string              `xml:"Model"`        // The name by which the PhysicalElement is generally known.
		Manufacturer       string              `xml:"Manufacturer"` // The name of the organization responsible for producing the PhysicalElement. This organization might be the entity from whom the Element is purchased, but this is not necessarily true. The latter information is contained in the Vendor property of CIM_Product.
		ElementName        string              `xml:"ElementName"`
		CreationClassName  string              `xml:"CreationClassName"`  // CreationClassName indicates the name of the class or the subclass used in the creation of an instance. When used with the other key properties of this class, this property allows all instances of this class and its subclasses to be uniquely identified.
		Tag                string              `xml:"Tag"`                // An arbitrary string that uniquely identifies the Physical Element and serves as the key of the Element. The Tag property can contain information such as asset tag or serial number data. The key for PhysicalElement is placed very high in the object hierarchy in order to independently identify the hardware or entity, regardless of physical placement in or on Cabinets, Adapters, and so on. For example, a hotswappable or removable component can be taken from its containing (scoping) Package and be temporarily unused. The object still continues to exist and can even be inserted into a different scoping container. Therefore, the key for Physical Element is an arbitrary string and is defined independently of any placement or location-oriented hierarchy.
		OperationalStatus  []OperationalStatus `xml:"OperationalStatus"`  // Indicates the current statuses of the element.
		PackageType        PackageType         `xml:"PackageType"`        // Enumeration defining the type of the PhysicalPackage. Note that this enumeration expands on the list in the Entity MIB (the attribute, entPhysicalClass). The numeric values are consistent with CIM's enum numbering guidelines, but are slightly different than the MIB's values.
		ChassisPackageType ChassisPackageType  `xml:"ChassisPackageType"` // ChassisPackageType indicates the physical form factor for the type of Chassis.
	}

	// ChassisPackageType is an enumeration defining the type of the PhysicalPackage.
	ChassisPackageType int
	// OperationalStatus is the current statuses of the element.
	OperationalStatus int
	// PackageType is the type of the PhysicalPackage.
	PackageType int
)
