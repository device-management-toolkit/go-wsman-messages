/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package setupandconfiguration

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/common"
)

type Service struct {
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
		XMLName                   xml.Name `xml:"Body"`
		GetResponse               SetupAndConfigurationServiceResponse
		EnumerateResponse         common.EnumerateResponse
		PullResponse              PullResponse
		GetUuid_OUTPUT            GetUuid_OUTPUT            `xml:"GetUuid_OUTPUT"`
		Unprovision_OUTPUT        Unprovision_OUTPUT        `xml:"Unprovision_OUTPUT"`
		PartialUnprovision_OUTPUT PartialUnprovision_OUTPUT `xml:"PartialUnprovision_OUTPUT"`
		CommitChanges_OUTPUT      CommitChanges_OUTPUT      `xml:"CommitChanges_OUTPUT"`
		SetMEBxPassword_OUTPUT    SetMEBxPassword_OUTPUT    `xml:"SetMEBxPassword_OUTPUT"`
	}

	SetupAndConfigurationServiceResponse struct {
		XMLName                       xml.Name               `xml:"AMT_SetupAndConfigurationService"`
		RequestedState                RequestedState         `xml:"RequestedState,omitempty"`                // RequestedState is an integer enumeration that indicates the last requested or desired state for the element, irrespective of the mechanism through which it was requested.
		EnabledState                  EnabledState           `xml:"EnabledState,omitempty"`                  // EnabledState is an integer enumeration that indicates the enabled and disabled states of an element.
		ElementName                   string                 `xml:"ElementName,omitempty"`                   // A user-friendly name for the object. This property allows each instance to define a user-friendly name in addition to its key properties, identity data, and description information. Note that the Name property of ManagedSystemElement is also defined as a user-friendly name. But, it is often subclassed to be a Key. It is not reasonable that the same property can convey both identity and a user-friendly name, without inconsistencies. Where Name exists and is not a Key (such as for instances of LogicalDevice), the same information can be present in both the Name and ElementName properties. Note that if there is an associated instance of CIM_EnabledLogicalElementCapabilities, restrictions on this properties may exist as defined in ElementNameMask and MaxElementNameLen properties defined in that class.
		SystemCreationClassName       string                 `xml:"SystemCreationClassName,omitempty"`       // The CreationClassName of the scoping System.
		SystemName                    string                 `xml:"SystemName,omitempty"`                    // The Name of the scoping System.
		CreationClassName             string                 `xml:"CreationClassName,omitempty"`             // CreationClassName indicates the name of the class or the subclass that is used in the creation of an instance. When used with the other key properties of this class, this property allows all instances of this class and its subclasses to be uniquely identified.
		Name                          string                 `xml:"Name,omitempty"`                          // The Name property uniquely identifies the Service and provides an indication of the functionality that is managed. This functionality is described in more detail in the Description property of the object.
		ProvisioningMode              ProvisioningModeValue  `xml:"ProvisioningMode,omitempty"`              // A Read-Only enumeration value that determines the behavior of Intel® AMT when it is deployed.
		ProvisioningState             ProvisioningStateValue `xml:"ProvisioningState,omitempty"`             // An enumeration value that indicates the state of the Intel® AMT subsystem in the provisioning process"Pre" - the setup operation has not started."In" - the setup operation is in progress."Post" - Intel® AMT is configured.
		ZeroTouchConfigurationEnabled bool                   `xml:"ZeroTouchConfigurationEnabled,omitempty"` // Indicates if Zero Touch Configuration (Remote Configuration) is enabled or disabled. This property affects only enterprise mode. It can be modified while in SMB mode
		ProvisioningServerOTP         string                 `xml:"ProvisioningServerOTP,omitempty"`         // A optional binary data value containing 8-32 characters,that represents a one-time password (OTP), used to authenticate the Intel® AMT to the configuration server. This property can be retrieved only in IN Provisioning state, nevertheless, it is settable also in POST provisioning state.
		ConfigurationServerFQDN       string                 `xml:"ConfigurationServerFQDN,omitempty"`       // The FQDN of the configuration server.
		PasswordModel                 PasswordModelValue     `xml:"PasswordModel,omitempty"`                 // An enumeration value that determines the password model of Intel® AMT.
		DhcpDNSSuffix                 string                 `xml:"DhcpDNSSuffix,omitempty"`                 // Domain name received from DHCP
		TrustedDNSSuffix              string                 `xml:"TrustedDNSSuffix,omitempty"`              // Trusted domain name configured in MEBX
	}
	PullResponse struct {
		XMLName                           xml.Name                               `xml:"PullResponse"`
		SetupAndConfigurationServiceItems []SetupAndConfigurationServiceResponse `xml:"Items>AMT_SetupAndConfigurationService"`
	}

	// UUID of the system. If the value is all FFh, the ID is not currently present in the system, but is settable. If the value is all 00h, the ID is not present in the system. Corresponds to the UUID field of the SMBIOS Type 1 structure.
	GetUuid_OUTPUT struct {
		XMLName xml.Name `xml:"GetUuid_OUTPUT"`
		UUID    string   `xml:"UUID"`
	}

	// ValueMap={0, 1, 16, 36, 2076}
	//
	// Values={PT_STATUS_SUCCESS, PT_STATUS_INTERNAL_ERROR, PT_STATUS_NOT_PERMITTED, PT_STATUS_INVALID_PARAMETER, PT_STATUS_BLOCKING_COMPONENT}.
	Unprovision_OUTPUT struct {
		XMLName     xml.Name `xml:"Unprovision_OUTPUT"`
		ReturnValue ReturnValue
	}
	// Values={PT_STATUS_SUCCESS, PT_STATUS_INTERNAL_ERROR, PT_STATUS_NOT_PERMITTED, PT_STATUS_BLOCKING_COMPONENT}.
	PartialUnprovision_OUTPUT struct {
		XMLName     xml.Name `xml:"PartialUnprovision_OUTPUT"`
		ReturnValue ReturnValue
	}
	// ValueMap={0, 1, 38, 2057}
	//
	// Values={PT_STATUS_SUCCESS, PT_STATUS_INTERNAL_ERROR, PT_STATUS_FLASH_WRITE_LIMIT_EXCEEDED, PT_STATUS_DATA_MISSING}.
	CommitChanges_OUTPUT struct {
		XMLName     xml.Name `xml:"CommitChanges_OUTPUT"`
		ReturnValue ReturnValue
	}

	// ValueMap={0, 1, 16, 2054}
	//
	// Values={PT_STATUS_SUCCESS, PT_STATUS_INTERNAL_ERROR, PT_STATUS_NOT_PERMITTED, PT_STATUS_INVALID_PASSWORD}.
	SetMEBxPassword_OUTPUT struct {
		XMLName     xml.Name `xml:"SetMEBxPassword_OUTPUT"`
		ReturnValue ReturnValue
	}
)

// Request Types.
type (
	SetupAndConfigurationServiceRequest struct {
		XMLName                       xml.Name               `xml:"h:AMT_SetupAndConfigurationService"`
		H                             string                 `xml:"xmlns:h,attr"`
		RequestedState                RequestedState         `xml:"h:RequestedState,omitempty"`                // RequestedState is an integer enumeration that indicates the last requested or desired state for the element, irrespective of the mechanism through which it was requested.
		EnabledState                  EnabledState           `xml:"h:EnabledState,omitempty"`                  // EnabledState is an integer enumeration that indicates the enabled and disabled states of an element.
		ElementName                   string                 `xml:"h:ElementName,omitempty"`                   // A user-friendly name for the object. This property allows each instance to define a user-friendly name in addition to its key properties, identity data, and description information. Note that the Name property of ManagedSystemElement is also defined as a user-friendly name. But, it is often subclassed to be a Key. It is not reasonable that the same property can convey both identity and a user-friendly name, without inconsistencies. Where Name exists and is not a Key (such as for instances of LogicalDevice), the same information can be present in both the Name and ElementName properties. Note that if there is an associated instance of CIM_EnabledLogicalElementCapabilities, restrictions on this properties may exist as defined in ElementNameMask and MaxElementNameLen properties defined in that class.
		SystemCreationClassName       string                 `xml:"h:SystemCreationClassName,omitempty"`       // The CreationClassName of the scoping System.
		SystemName                    string                 `xml:"h:SystemName,omitempty"`                    // The Name of the scoping System.
		CreationClassName             string                 `xml:"h:CreationClassName,omitempty"`             // CreationClassName indicates the name of the class or the subclass that is used in the creation of an instance. When used with the other key properties of this class, this property allows all instances of this class and its subclasses to be uniquely identified.
		Name                          string                 `xml:"h:Name,omitempty"`                          // The Name property uniquely identifies the Service and provides an indication of the functionality that is managed. This functionality is described in more detail in the Description property of the object.
		ProvisioningMode              ProvisioningModeValue  `xml:"h:ProvisioningMode,omitempty"`              // A Read-Only enumeration value that determines the behavior of Intel® AMT when it is deployed.
		ProvisioningState             ProvisioningStateValue `xml:"h:ProvisioningState,omitempty"`             // An enumeration value that indicates the state of the Intel® AMT subsystem in the provisioning process"Pre" - the setup operation has not started."In" - the setup operation is in progress."Post" - Intel® AMT is configured.
		ZeroTouchConfigurationEnabled bool                   `xml:"h:ZeroTouchConfigurationEnabled,omitempty"` // Indicates if Zero Touch Configuration (Remote Configuration) is enabled or disabled. This property affects only enterprise mode. It can be modified while in SMB mode
		ProvisioningServerOTP         string                 `xml:"h:ProvisioningServerOTP,omitempty"`         // A optional binary data value containing 8-32 characters,that represents a one-time password (OTP), used to authenticate the Intel® AMT to the configuration server. This property can be retrieved only in IN Provisioning state, nevertheless, it is settable also in POST provisioning state.
		ConfigurationServerFQDN       string                 `xml:"h:ConfigurationServerFQDN,omitempty"`       // The FQDN of the configuration server.
		PasswordModel                 PasswordModelValue     `xml:"h:PasswordModel,omitempty"`                 // An enumeration value that determines the password model of Intel® AMT.
		DhcpDNSSuffix                 string                 `xml:"h:DhcpDNSSuffix,omitempty"`                 // Domain name received from DHCP
		TrustedDNSSuffix              string                 `xml:"h:TrustedDNSSuffix,omitempty"`              // Trusted domain name configured in MEBX
	}

	// Password needs to be strong: Contain at least one of: upper-case, lower-case, digit and special character
	//
	// MinLen=8, MaxLen=32.
	MEBXPassword struct {
		XMLName  xml.Name `xml:"h:SetMEBxPassword_INPUT"`
		H        string   `xml:"xmlns:h,attr"`
		Password string   `xml:"h:Password,omitempty"`
	}

	// Indicates the provisioning mode (Enterprise , Small Business or Remote Connectivity) the device will enter following successful completion of the command. Starting from Release 6.0 only effective value is ProvisioningModeEnterprise
	//
	// ValueMap={0, 1, 2, 3}
	//
	// Values={ProvisioningModeCurrent, ProvisioningModeEnterprise, ProvisioningModeSmallBusiness, ProvisioningRemoteConnectivity}.
	ProvisioningMode struct {
		XMLName          xml.Name              `xml:"h:Unprovision_INPUT"`
		H                string                `xml:"xmlns:h,attr"`
		ProvisioningMode ProvisioningModeValue `xml:"h:ProvisioningMode,omitempty"`
	}

	// EnabledState is an integer enumeration that indicates the enabled and disabled states of an element. It can also indicate the transitions between these requested states. For example, shutting down (value=4) and starting (value=10) are transient states between enabled and disabled.
	EnabledState int

	// RequestedState is an integer enumeration that indicates the last requested or desired state for the element, irrespective of the mechanism through which it was requested. The actual state of the element is represented by EnabledState. This property is provided to compare the last requested and current enabled or disabled states. Note that when EnabledState is set to 5 ("Not Applicable"), then this property has no meaning. Refer to the EnabledState property description for explanations of the values in the RequestedState enumeration.
	RequestedState int

	// ProvisioningModeValue is a Read-Only enumeration value that determines the behavior of Intel® AMT when it is deployed. Starting from Release 7.0, this enumeration indicates whether AMT is deployed in "Admin control mode" or "Client control mode". In "Admin" mode, AMT functionality is on the same level of previous releases. In "Client" mode functionality is limited or requires user consent.
	ProvisioningModeValue int

	// ProvisioningStateValue is an enumeration value that indicates the state of the Intel® AMT subsystem in the provisioning process"Pre" - the setup operation has not started."In" - the setup operation is in progress."Post" - Intel® AMT is configured.
	ProvisioningStateValue int

	// PasswordModelValue is an enumeration value that determines the password model of Intel® AMT.
	PasswordModelValue int

	// ReturnValue is a 32-bit unsigned integer. It is used to indicate the success or failure of a method. A ReturnValue of zero indicates success. A non-zero value indicates failure. The specific values for ReturnValue are defined in the method description.
	ReturnValue int
)
