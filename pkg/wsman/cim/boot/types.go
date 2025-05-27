/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package boot

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
)

type ConfigSetting struct {
	base message.Base
}

type SourceSetting struct {
	base message.Base
}

type Service struct {
	base message.Base
}

type Source string

// Response Types.
type (
	Response struct {
		*client.Message
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}

	Body struct {
		XMLName                   xml.Name          `xml:"Body"`
		ConfigSettingGetResponse  BootConfigSetting `xml:"CIM_BootConfigSetting"`
		SourceSettingGetResponse  BootSourceSetting `xml:"CIM_BootSourceSetting"`
		ServiceGetResponse        BootService       `xml:"CIM_BootService"`
		EnumerateResponse         common.EnumerateResponse
		PullResponse              PullResponse             `xml:"PullResponse"`
		ChangeBootOrder_OUTPUT    ChangeBootOrder_OUTPUT   `xml:"ChangeBootOrder_OUTPUT"`
		SetBootConfigRole_OUTPUT  SetBootConfigRole_OUTPUT `xml:"SetBootConfigRole_OUTPUT"`
		RequestStateChange_OUTPUT common.ReturnValue
	}

	BootConfigSetting struct {
		XMLName     xml.Name `xml:"CIM_BootConfigSetting"`
		InstanceID  string   `xml:"InstanceID"`  // Within the scope of the instantiating Namespace, InstanceID opaquely and uniquely identifies an instance of this class. To ensure uniqueness within the NameSpace, the value of InstanceID should be constructed using the following "preferred" algorithm: <OrgID>:<LocalID> Where <OrgID> and <LocalID> are separated by a colon (:), and where <OrgID> must include a copyrighted, trademarked, or otherwise unique name that is owned by the business entity that is creating or defining the InstanceID or that is a registered ID assigned to the business entity by a recognized global authority. (This requirement is similar to the <Schema Name>_<Class Name> structure of Schema class names.) In addition, to ensure uniqueness, <OrgID> must not contain a colon (:). When using this algorithm, the first colon to appear in InstanceID must appear between <OrgID> and <LocalID>. <LocalID> is chosen by the business entity and should not be reused to identify different underlying (real-world) elements. If the above "preferred" algorithm is not used, the defining entity must assure that the resulting InstanceID is not reused across any InstanceIDs produced by this or other providers for the NameSpace of this instance.	For DMTF-defined instances, the "preferred" algorithm must be used with the <OrgID> set to CIM.
		ElementName string   `xml:"ElementName"` // The user-friendly name for this instance of SettingData. In addition, the user-friendly name can be used as an index property for a search or query. (Note: The name does not have to be unique within a namespace.)
	}

	BootSourceSetting struct {
		XMLName              xml.Name             `xml:"CIM_BootSourceSetting"`
		ElementName          string               `xml:"ElementName"`          // he user-friendly name for this instance of SettingData. In addition, the user-friendly name can be used as an index property for a search or query. (Note: The name does not have to be unique within a namespace.)
		InstanceID           string               `xml:"InstanceID"`           // Within the scope of the instantiating Namespace, InstanceID opaquely and uniquely identifies an instance of this class. To ensure uniqueness within the NameSpace, the value of InstanceID should be constructed using the following "preferred" algorithm: <OrgID>:<LocalID> Where <OrgID> and <LocalID> are separated by a colon (:), and where <OrgID> must include a copyrighted, trademarked, or otherwise unique name that is owned by the business entity that is creating or defining the InstanceID or that is a registered ID assigned to the business entity by a recognized global authority. (This requirement is similar to the <Schema Name>_<Class Name> structure of Schema class names.) In addition, to ensure uniqueness, <OrgID> must not contain a colon (:). When using this algorithm, the first colon to appear in InstanceID must appear between <OrgID> and <LocalID>. <LocalID> is chosen by the business entity and should not be reused to identify different underlying (real-world) elements. If the above "preferred" algorithm is not used, the defining entity must assure that the resulting InstanceID is not reused across any InstanceIDs produced by this or other providers for the NameSpace of this instance. For DMTF-defined instances, the "preferred" algorithm must be used with the <OrgID> set to CIM.
		StructuredBootString string               `xml:"StructuredBootString"` // A string identifying the boot source using the format "<OrgID>:<identifier>:<index>", in which neither <OrgID>, <identifier> or <index> contains a colon (":"). The value of <OrgID> is a copyrighted, trademarked or otherwise unique name that is owned by the entity defining the <identifier>, or is a registered ID that is assigned to the entity by a recognized global authority. For DMTF defined identifiers, the <OrgID> is set to 'CIM'. The <identifiers> are "Floppy", "Hard-Disk", "CD/DVD", "Network", "PCMCIA", "USB". The value of <index> shall be a non-zero integer.
		BIOSBootString       string               `xml:"BIOSBootString"`       // BIOS description of the boot option. UTF8 null-terminated string, relevant only to One-Click Recovery WinRE and PBA (pre-boot application) boot options.
		BootString           string               `xml:"BootString"`           // BIOS description of the EFI device path. UTF8 null-terminated string, relevant only to One-Click Recovery WinRE and PBA (pre-boot application) boot options.
		FailThroughSupported FailThroughSupported `xml:"FailThroughSupported"` // An enumeration indicating the behavior when the attempt to boot using the boot source fails (no media, timeout).
	}

	BootService struct {
		XMLName                 xml.Name            `xml:"CIM_BootService"`
		Name                    string              `xml:"Name"`                    // The Name property uniquely identifies the Service and provides an indication of the functionality that is managed. This functionality is described in more detail in the Description property of the object.
		CreationClassName       string              `xml:"CreationClassName"`       // CreationClassName indicates the name of the class or the subclass that is used in the creation of an instance. When used with the other key properties of this class, this property allows all instances of this class and its subclasses to be uniquely identified.
		SystemName              string              `xml:"SystemName"`              // The Name of the scoping System.
		SystemCreationClassName string              `xml:"SystemCreationClassName"` // The CreationClassName of the scoping System.
		ElementName             string              `xml:"ElementName"`             // A user-friendly name for the object. This property allows each instance to define a user-friendly name in addition to its key properties, identity data, and description information. Note that the Name property of ManagedSystemElement is also defined as a user-friendly name. But, it is often subclassed to be a Key. It is not reasonable that the same property can convey both identity and a user-friendly name, without inconsistencies. Where Name exists and is not a Key (such as for instances of LogicalDevice), the same information can be present in both the Name and ElementName properties. Note that if there is an associated instance of CIM_EnabledLogicalElementCapabilities, restrictions on this properties may exist as defined in ElementNameMask and MaxElementNameLen properties defined in that class.
		OperationalStatus       []OperationalStatus `xml:"OperationalStatus"`       // Indicates the current statuses of the element. Various operational statuses are defined.
		EnabledState            EnabledState        `xml:"EnabledState"`            // EnabledState is an integer enumeration that indicates the enabled and disabled states of an element.
		RequestedState          RequestedState      `xml:"RequestedState"`          // RequestedState is an integer enumeration that indicates the last requested or desired state for the element, irrespective of the mechanism through which it was requested.
	}

	PullResponse struct {
		BootSourceSettingItems []BootSourceSetting `xml:"Items>CIM_BootSourceSetting"`
		BootConfigSettingItems []BootConfigSetting `xml:"Items>CIM_BootConfigSetting"`
		BootServiceItems       []BootService       `xml:"Items>CIM_BootService"`
	}

	ChangeBootOrder_OUTPUT struct {
		ReturnValue ReturnValue `xml:"ReturnValue"`
	}

	SetBootConfigRole_OUTPUT struct {
		ReturnValue ReturnValue `xml:"ReturnValue"`
	}

	// FailThroughSupported is an enumeration indicating the behavior when the attempt to boot using the boot source fails (no media, timeout).
	FailThroughSupported int
	// ReturnValue is an enumeration indicating the return value of the operation.
	ReturnValue int
	// OperationalStatus is an enumeration indicating the current statuses of the element.
	OperationalStatus int
	// EnabledState is an integer enumeration that indicates the enabled and disabled states of an element.
	EnabledState int
	// RequestedState is an integer enumeration that indicates the last requested or desired state for the element, irrespective of the mechanism through which it was requested.
	RequestedState int
)
