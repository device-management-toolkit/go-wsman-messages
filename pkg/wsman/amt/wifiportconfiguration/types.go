/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package wifiportconfiguration

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/models"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/wifi"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
)

type Service struct {
	base message.Base
}

// OUTPUT
// Response Types.
type (
	Response struct {
		*client.Message
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body
	}

	Body struct {
		XMLName                      xml.Name `xml:"Body"`
		WiFiPortConfigurationService WiFiPortConfigurationServiceResponse
		PullResponse                 PullResponse
		EnumerateResponse            common.EnumerateResponse
		AddWiFiSettingsOutput        AddWiFiSettings_OUTPUT
	}
	WiFiPortConfigurationServiceResponse struct {
		XMLName                            xml.Name                           `xml:"AMT_WiFiPortConfigurationService"`
		RequestedState                     RequestedState                     `xml:"RequestedState"`                            // RequestedState is an integer enumeration that indicates the last requested or desired state for the element, irrespective of the mechanism through which it was requested.
		EnabledState                       EnabledState                       `xml:"EnabledState"`                              // EnabledState is an integer enumeration that indicates the enabled and disabled states of an element.
		HealthState                        HealthState                        `xml:"HealthState"`                               // Indicates the current health of the element.
		ElementName                        string                             `xml:"ElementName,omitempty"`                     // A user-friendly name for the object. This property allows each instance to define a user-friendly name in addition to its key properties, identity data, and description information. Note that the Name property of ManagedSystemElement is also defined as a user-friendly name. But, it is often subclassed to be a Key. It is not reasonable that the same property can convey both identity and a user-friendly name, without inconsistencies. Where Name exists and is not a Key (such as for instances of LogicalDevice), the same information can be present in both the Name and ElementName properties. Note that if there is an associated instance of CIM_EnabledLogicalElementCapabilities, restrictions on this properties may exist as defined in ElementNameMask and MaxElementNameLen properties defined in that class.
		SystemCreationClassName            string                             `xml:"SystemCreationClassName,omitempty"`         // The CreationClassName of the scoping System.
		SystemName                         string                             `xml:"SystemName,omitempty"`                      // The Name of the scoping System.
		CreationClassName                  string                             `xml:"CreationClassName,omitempty"`               // CreationClassName indicates the name of the class or the subclass that is used in the creation of an instance. When used with the other key properties of this class, this property allows all instances of this class and its subclasses to be uniquely identified.
		Name                               string                             `xml:"Name,omitempty"`                            // The Name property uniquely identifies the Service and provides an indication of the functionality that is managed. This functionality is described in more detail in the Description property of the object.
		LocalProfileSynchronizationEnabled LocalProfileSynchronizationEnabled `xml:"localProfileSynchronizationEnabled"`        // Administrator's policy regarding enablement of local profile synchronization.Remote profile synchronization is always enabled.
		LastConnectedSsidUnderMeControl    string                             `xml:"LastConnectedSsidUnderMeControl,omitempty"` // The SSID of the Wireless network that was last connected in ME Control state
		NoHostCsmeSoftwarePolicy           NoHostCsmeSoftwarePolicy           `xml:"NoHostCsmeSoftwarePolicy"`                  // Setting Policy regarding no HOST CSME software.
		UEFIWiFiProfileShareEnabled        bool                               `xml:"UEFIWiFiProfileShareEnabled"`               // Enables or disables UEFI/CSME Wi-Fi Profile Sharing. The feature is available from Intel® CSME 16.0. The feature can be disabled even if the value of AMT_BootCapabilities.UEFIWiFiCoExistenceAndProfileShare is False.
	}
	PullResponse struct {
		XMLName                    xml.Name                               `xml:"PullResponse"`
		WiFiPortConfigurationItems []WiFiPortConfigurationServiceResponse `xml:"Items>AMT_WiFiPortConfigurationService"`
	}

	// ValueMap={0, 1, 2, 3, 4, .., 32768..65535}
	//
	// Values={Completed with No Error, Not Supported, Failed, Invalid Parameter, Invalid Reference, Method Reserved, Vendor Specific}.
	AddWiFiSettings_OUTPUT struct {
		XMLName xml.Name `xml:"AddWiFiSettings_OUTPUT"`
		// not concerned with these entries on OUTPUT

		// ClientCredential     *ClientCredential         `xml:"g:ClientCredential,omitempty"`
		// CACredential         *CACredential             `xml:"g:CACredential,omitempty"`
		ReturnValue ReturnValue `xml:"ReturnValue"`
	}
)

type (
	// Administrator's policy regarding enablement of local profile synchronization.Remote profile synchronization is always enabled.
	//
	// 1) 'localProfileSynchronizationEnabled' is only supported in Intel AMT Release 6.0 and later releases.
	//
	// ValueMap={0, 1, 2, 3, 4..}
	//
	// Values={Local synchronization disabled, Local user profile synchronization enabled, Vendor Reserved, Unrestricted synchronization, Reserved}.
	LocalProfileSynchronizationEnabled int
	// Setting Policy regarding no HOST CSME software.
	//
	// ValueMap={0, 1, 2}
	//
	// Values={NoHostCsmeSoftwareRelaxedPolicy, NoHostCsmeSoftwareAggressivePolicy, Reserved}.
	NoHostCsmeSoftwarePolicy int
	// RequestedState is an integer enumeration that indicates the last requested or desired state for the element, irrespective of the mechanism through which it was requested. The actual state of the element is represented by EnabledState. This property is provided to compare the last requested and current enabled or disabled states. Note that when EnabledState is set to 5 ("Not Applicable"), then this property has no meaning. Refer to the EnabledState property description for explanations of the values in the RequestedState enumeration.
	//
	// "Unknown" (0) indicates the last requested state for the element is unknown.
	//
	// Note that the value "No Change" (5) has been deprecated in lieu of indicating the last requested state is "Unknown" (0). If the last requested or desired state is unknown, RequestedState should have the value "Unknown" (0), but may have the value "No Change" (5).Offline (6) indicates that the element has been requested to transition to the Enabled but Offline EnabledState.
	//
	// It should be noted that there are two new values in RequestedState that build on the statuses of EnabledState. These are "Reboot" (10) and "Reset" (11). Reboot refers to doing a "Shut Down" and then moving to an "Enabled" state. Reset indicates that the element is first "Disabled" and then "Enabled". The distinction between requesting "Shut Down" and "Disabled" should also be noted. Shut Down requests an orderly transition to the Disabled state, and might involve removing power, to completely erase any existing state. The Disabled state requests an immediate disabling of the element, such that it will not execute or accept any commands or processing requests.
	//
	// This property is set as the result of a method invocation (such as Start or StopService on CIM_Service), or can be overridden and defined as WRITEable in a subclass. The method approach is considered superior to a WRITEable property, because it allows an explicit invocation of the operation and the return of a result code.
	//
	// If knowledge of the last RequestedState is not supported for the EnabledLogicalElement, the property shall be NULL or have the value 12 "Not Applicable".
	//
	// ValueMap={0, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, .., 32768..65535}
	//
	// Values={Unknown, Enabled, Disabled, Shut Down, No Change, Offline, Test, Deferred, Quiesce, Reboot, Reset, Not Applicable, DMTF Reserved, Vendor Reserved}.
	RequestedState int
	// EnabledState is an integer enumeration that indicates the enabled and disabled states of an element. It can also indicate the transitions between these requested states. For example, shutting down (value=4) and starting (value=10) are transient states between enabled and disabled. The following text briefly summarizes the various enabled and disabled states:
	//
	// Enabled (2) indicates that the element is or could be executing commands, will process any queued commands, and queues new requests.
	//
	// Disabled (3) indicates that the element will not execute commands and will drop any new requests.
	//
	// Shutting Down (4) indicates that the element is in the process of going to a Disabled state.
	//
	// Not Applicable (5) indicates the element does not support being enabled or disabled.
	//
	// Enabled but Offline (6) indicates that the element might be completing commands, and will drop any new requests.
	//
	// Test (7) indicates that the element is in a test state.
	//
	// Deferred (8) indicates that the element might be completing commands, but will queue any new requests.
	//
	// Quiesce (9) indicates that the element is enabled but in a restricted mode.
	//
	// Starting (10) indicates that the element is in the process of going to an Enabled state. New requests are queued.
	//
	// ValueMap={0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11..32767, 32768..65535}
	//
	// Values={Unknown, Other, Enabled, Disabled, Shutting Down, Not Applicable, Enabled but Offline, In Test, Deferred, Quiesce, Starting, DMTF Reserved, Vendor Reserved}.
	EnabledState int
	// Indicates the current health of the element. This attribute expresses the health of this element but not necessarily that of its subcomponents. The possible values are 0 to 30, where 5 means the element is entirely healthy and 30 means the element is completely non-functional. The following continuum is defined:
	//
	// "Non-recoverable Error" (30) - The element has completely failed, and recovery is not possible. All functionality provided by this element has been lost.
	//
	// "Critical Failure" (25) - The element is non-functional and recovery might not be possible.
	//
	// "Major Failure" (20) - The element is failing. It is possible that some or all of the functionality of this component is degraded or not working.
	//
	// "Minor Failure" (15) - All functionality is available but some might be degraded.
	//
	// "Degraded/Warning" (10) - The element is in working order and all functionality is provided. However, the element is not working to the best of its abilities. For example, the element might not be operating at optimal performance or it might be reporting recoverable errors.
	//
	// "OK" (5) - The element is fully functional and is operating within normal operational parameters and without error.
	//
	// "Unknown" (0) - The implementation cannot report on HealthState at this time.
	//
	// DMTF has reserved the unused portion of the continuum for additional HealthStates in the future.
	//
	// ValueMap={0, 5, 10, 15, 20, 25, 30, .., 32768..65535}
	//
	// Values={Unknown, OK, Degraded/Warning, Minor failure, Major failure, Critical failure, Non-recoverable error, DMTF Reserved, Vendor Specific}.
	HealthState int
	// ReturnValue is an integer enumeration that indicates the completion status of the method. This value shall be 0 if the method was completed successfully. A non-zero value indicates an error condition.
	ReturnValue int
)

// INPUT
// Request Types.
type (
	AddWiFiSettings_INPUT struct {
		XMLName              xml.Name `xml:"h:AddWiFiSettings_INPUT"`
		H                    string   `xml:"xmlns:h,attr"`
		WifiEndpoint         WiFiEndpoint
		WiFiEndpointSettings wifi.WiFiEndpointSettingsRequest
		IEEE8021xSettings    *models.IEEE8021xSettings `xml:"h:IEEE8021xSettingsInput,omitempty"`
		ClientCredential     *ClientCredentialRequest  `xml:"h:ClientCredential,omitempty"`
		CACredential         *CACredentialRequest      `xml:"h:CACredential,omitempty"`
	}
	WiFiPortConfigurationServiceRequest struct {
		XMLName                            xml.Name                           `xml:"h:AMT_WiFiPortConfigurationService"`
		H                                  string                             `xml:"xmlns:h,attr"`
		RequestedState                     RequestedState                     `xml:"h:RequestedState,omitempty"`                  // RequestedState is an integer enumeration that indicates the last requested or desired state for the element, irrespective of the mechanism through which it was requested.
		EnabledState                       EnabledState                       `xml:"h:EnabledState,omitempty"`                    // EnabledState is an integer enumeration that indicates the enabled and disabled states of an element.
		HealthState                        HealthState                        `xml:"h:HealthState,omitempty"`                     // Indicates the current health of the element.
		ElementName                        string                             `xml:"h:ElementName,omitempty"`                     // A user-friendly name for the object. This property allows each instance to define a user-friendly name in addition to its key properties, identity data, and description information. Note that the Name property of ManagedSystemElement is also defined as a user-friendly name. But, it is often subclassed to be a Key. It is not reasonable that the same property can convey both identity and a user-friendly name, without inconsistencies. Where Name exists and is not a Key (such as for instances of LogicalDevice), the same information can be present in both the Name and ElementName properties. Note that if there is an associated instance of CIM_EnabledLogicalElementCapabilities, restrictions on this properties may exist as defined in ElementNameMask and MaxElementNameLen properties defined in that class.
		SystemCreationClassName            string                             `xml:"h:SystemCreationClassName,omitempty"`         // The CreationClassName of the scoping System.
		SystemName                         string                             `xml:"h:SystemName,omitempty"`                      // The Name of the scoping System.
		CreationClassName                  string                             `xml:"h:CreationClassName,omitempty"`               // CreationClassName indicates the name of the class or the subclass that is used in the creation of an instance. When used with the other key properties of this class, this property allows all instances of this class and its subclasses to be uniquely identified.
		Name                               string                             `xml:"h:Name,omitempty"`                            // The Name property uniquely identifies the Service and provides an indication of the functionality that is managed. This functionality is described in more detail in the Description property of the object.
		LocalProfileSynchronizationEnabled LocalProfileSynchronizationEnabled `xml:"h:localProfileSynchronizationEnabled"`        // Administrator's policy regarding enablement of local profile synchronization.Remote profile synchronization is always enabled.
		LastConnectedSsidUnderMeControl    string                             `xml:"h:LastConnectedSsidUnderMeControl,omitempty"` // The SSID of the Wireless network that was last connected in ME Control state
		NoHostCsmeSoftwarePolicy           NoHostCsmeSoftwarePolicy           `xml:"h:NoHostCsmeSoftwarePolicy,omitempty"`        // Setting Policy regarding no HOST CSME software.
		UEFIWiFiProfileShareEnabled        bool                               `xml:"h:UEFIWiFiProfileShareEnabled,omitempty"`     // Enables or disables UEFI/CSME Wi-Fi Profile Sharing. The feature is available from Intel® CSME 16.0. The feature can be disabled even if the value of AMT_BootCapabilities.UEFIWiFiCoExistenceAndProfileShare is False.
	}

	// a Reference to an AMT_PublicKeyCertificate, which represents the CA certificate.
	CACredentialRequest struct {
		XMLName             xml.Name            `xml:"h:CACredential,omitempty"`
		H                   string              `xml:"xmlns:a,attr"`
		Address             string              `xml:"a:Address,omitempty"`
		ReferenceParameters ReferenceParameters `xml:"a:ReferenceParameters,omitempty"`
	}

	// a Reference to an AMT_PublicKeyCertificate, which represents the client certificate.
	ClientCredentialRequest struct {
		XMLName             xml.Name            `xml:"h:ClientCredential,omitempty"`
		H                   string              `xml:"xmlns:a,attr"`
		Address             string              `xml:"a:Address,omitempty"`
		ReferenceParameters ReferenceParameters `xml:"a:ReferenceParameters,omitempty"`
	}

	ReferenceParameters struct {
		XMLName     xml.Name    `xml:"a:ReferenceParameters"`
		H           string      `xml:"xmlns:c,attr"`
		ResourceURI string      `xml:"c:ResourceURI,omitempty"`
		SelectorSet SelectorSet `xml:"c:SelectorSet,omitempty"`
	}

	SelectorSet struct {
		H        string   `xml:"xmlns:c,attr"`
		XMLName  xml.Name `xml:"c:SelectorSet,omitempty"`
		Selector []Selector
	}

	Selector struct {
		H       string   `xml:"xmlns:c,attr"`
		XMLName xml.Name `xml:"c:Selector,omitempty"`
		Name    string   `xml:"Name,attr"`
		Value   string   `xml:",chardata"`
	}
)

// The endpoint to associate the new settings with.
type WiFiEndpoint struct {
	XMLName             xml.Name            `xml:"h:WiFiEndpoint,omitempty"`
	Address             string              `xml:"a:Address,omitempty"`
	ReferenceParameters ReferenceParameters `xml:"a:ReferenceParameters,omitempty"`
}

// A reference to a new CIM_WiFiEndpointSettings instance that shall be created by the method using the property values in WiFiEndpointSettingsInput.
type AddWiFiSettingsResponse struct {
	XMLName                xml.Name `xml:"Body"`
	AddWiFiSettings_OUTPUT AddWiFiSettings_OUTPUT
}
