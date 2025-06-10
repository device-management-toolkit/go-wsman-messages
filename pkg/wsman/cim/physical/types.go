/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package physical

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/card"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/chassis"
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
		MemoryResponse    PhysicalMemory
	}

	PullResponse struct {
		XMLName xml.Name `xml:"PullResponse"`
		common.EnumerateResponse
		MemoryItems     []PhysicalMemory          `xml:"Items>CIM_PhysicalMemory"`
		Card            []card.PackageResponse    `xml:"Items>CIM_Card"`
		PhysicalPackage []PhysicalPackage         `xml:"Items>CIM_PhysicalPackage"`
		Chassis         []chassis.PackageResponse `xml:"Items>CIM_Chassis"`
		EndOfSequence   xml.Name                  `xml:"EndOfSequence"`
	}

	PhysicalMemory struct {
		XMLName                    xml.Name            `xml:"CIM_PhysicalMemory"`
		PartNumber                 string              `xml:"PartNumber"`        // The part number assigned by the organization that is responsible for producing or manufacturing the PhysicalElement.
		SerialNumber               string              `xml:"SerialNumber"`      // A manufacturer-allocated number used to identify the Physical Element.
		Manufacturer               string              `xml:"Manufacturer"`      // The name of the organization responsible for producing the PhysicalElement. This organization might be the entity from whom the Element is purchased, but this is not necessarily true. The latter information is contained in the Vendor property of CIM_Product.
		ElementName                string              `xml:"ElementName"`       // 'ElementName' is constant. In CIM_Chip instances its value is set to 'Managed System Memory Chip'.
		CreationClassName          string              `xml:"CreationClassName"` // CreationClassName indicates the name of the class or the subclass used in the creation of an instance. When used with the other key properties of this class, this property allows all instances of this class and its subclasses to be uniquely identified.
		Tag                        string              `xml:"Tag"`               // An arbitrary string that uniquely identifies the Physical Element and serves as the key of the Element. The Tag property can contain information such as asset tag or serial number data. The key for PhysicalElement is placed very high in the object hierarchy in order to independently identify the hardware or entity, regardless of physical placement in or on Cabinets, Adapters, and so on. For example, a hotswappable or removable component can be taken from its containing (scoping) Package and be temporarily unused. The object still continues to exist and can even be inserted into a different scoping container. Therefore, the key for Physical Element is an arbitrary string and is defined independently of any placement or location-oriented hierarchy.
		OperationalStatus          []OperationalStatus `xml:"OperationalStatus"` // Indicates the current statuses of the element. Various operational statuses are defined.
		FormFactor                 int                 `xml:"FormFactor,omitempty"`
		MemoryType                 MemoryType          `xml:"MemoryType,omitempty"`                 // The type of PhysicalMemory. Synchronous DRAM is also known as SDRAM Cache DRAM is also known as CDRAM CDRAM is also known as Cache DRAM SDRAM is also known as Synchronous DRAM BRAM is also known as Block RAM
		Speed                      int                 `xml:"Speed,omitempty"`                      // The speed of the PhysicalMemory, in nanoseconds.
		Capacity                   int                 `xml:"Capacity,omitempty"`                   // The total capacity of this PhysicalMemory, in bytes.
		BankLabel                  string              `xml:"BankLabel,omitempty"`                  // A string identifying the physically labeled bank where the Memory is located - for example, 'Bank 0' or 'Bank A'.
		ConfiguredMemoryClockSpeed int                 `xml:"ConfiguredMemoryClockSpeed,omitempty"` // The configured clock speed (in MHz) of PhysicalMemory.
		IsSpeedInMhz               bool                `xml:"IsSpeedInMhz,omitempty"`               // The IsSpeedInMHz property is used to indicate if the Speed property or the MaxMemorySpeed contains the value of the memory speed. A value of TRUE shall indicate that the speed is represented by the MaxMemorySpeed property. A value of FALSE shall indicate that the speed is represented by the Speed property.
		MaxMemorySpeed             int                 `xml:"MaxMemorySpeed,omitempty"`             // The maximum speed (in MHz) of PhysicalMemory.
	}

	PhysicalPackage struct {
		XMLName              xml.Name `xml:"CIM_PhysicalPackage"`
		CanBeFRUed           bool     `xml:"CanBeFRUed"`
		VendorEquipmentType  string   `xml:"VendorEquipmentType"`
		ManufactureDate      string   `xml:"ManufactureDate"`
		OtherIdentifyingInfo string   `xml:"OtherIdentifyingInfo"`
		SerialNumber         string   `xml:"SerialNumber"`
		SKU                  string   `xml:"SKU"`
		Model                string   `xml:"Model"`
		Manufacturer         string   `xml:"Manufacturer"`
		ElementName          string   `xml:"ElementName"`
		CreationClassName    string   `xml:"CreationClassName"`
		Tag                  string   `xml:"Tag"`
		OperationalStatus    []int    `xml:"OperationalStatus"`
		PackageType          int      `xml:"PackageType"`
	}

	// MemoryType is an enumeration that describes the type of memory.
	MemoryType int
	// OperationalStatus indicates the current statuses of the element. Various operational statuses are defined.
	OperationalStatus int
)
