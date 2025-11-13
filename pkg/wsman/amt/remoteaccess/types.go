/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package remoteaccess

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
)

type PolicyAppliesToMPS struct {
	base message.Base
}

// OUTPUTS
// Response Types.
type (
	Response struct {
		*client.Message
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}
	Body struct {
		XMLName                                   xml.Name `xml:"Body"`
		RemoteAccessServiceGetResponse            RemoteAccessServiceResponse
		RemoteAccessPolicyRuleGetResponse         RemoteAccessPolicyRuleResponse
		RemoteAccessPolicyAppliesToMPSGetResponse RemoteAccessPolicyAppliesToMPSResponse
		EnumerateResponse                         common.EnumerateResponse
		PullResponse                              PullResponse
		AddMpServerResponse                       AddMpServerResponse
		AddRemotePolicyRuleResponse               AddRemoteAccessPolicyRuleResponse
	}
	RemoteAccessServiceResponse struct {
		XMLName                      xml.Name `xml:"AMT_RemoteAccessService"`
		CreationClassName            string   `xml:"CreationClassName,omitempty"`            // CreationClassName indicates the name of the class or the subclass that is used in the creation of an instance. When used with the other key properties of this class, this property allows all instances of this class and its subclasses to be uniquely identified.
		ElementName                  string   `xml:"ElementName,omitempty"`                  // A user-friendly name for the object. This property allows each instance to define a user-friendly name in addition to its key properties, identity data, and description information. Note that the Name property of ManagedSystemElement is also defined as a user-friendly name. But, it is often subclassed to be a Key. It is not reasonable that the same property can convey both identity and a user-friendly name, without inconsistencies. Where Name exists and is not a Key (such as for instances of LogicalDevice), the same information can be present in both the Name and ElementName properties. Note that if there is an associated instance of CIM_EnabledLogicalElementCapabilities, restrictions on this properties may exist as defined in ElementNameMask and MaxElementNameLen properties defined in that class.
		Name                         string   `xml:"Name,omitempty"`                         // The Name property uniquely identifies the Service and provides an indication of the functionality that is managed. This functionality is described in more detail in the Description property of the object.
		SystemCreationClassName      string   `xml:"SystemCreationClassName,omitempty"`      // The CreationClassName of the scoping System.
		SystemName                   string   `xml:"SystemName,omitempty"`                   // The Name of the scoping System.
		IsRemoteTunnelConnected      bool     `xml:"IsRemoteTunnelConnected,omitempty"`      // Reflects the connection status of the remote tunnel. Supported starting from Intel CSME 17.
		RemoteTunnelKeepAliveTimeout int      `xml:"RemoteTunnelKeepAliveTimeout,omitempty"` // Reflects the keep-alive timeout value of the remote tunnel (in seconds). Supported starting from Intel CSME 17.
	}
	RemoteAccessPolicyRuleResponse struct {
		XMLName                 xml.Name `xml:"AMT_RemoteAccessPolicyRule"`
		CreationClassName       string   `xml:"CreationClassName,omitempty"`       // CreationClassName indicates the name of the class or the subclass used in the creation of an instance. When used with the other key properties of this class, this property allows all instances of this class and its subclasses to be uniquely identified.
		ElementName             string   `xml:"ElementName,omitempty"`             // A user-friendly name for the object. This property allows each instance to define a user-friendly name in addition to its key properties, identity data, and description information. Note that the Name property of ManagedSystemElement is also defined as a user-friendly name. But, it is often subclassed to be a Key. It is not reasonable that the same property can convey both identity and a user-friendly name, without inconsistencies. Where Name exists and is not a Key (such as for instances of LogicalDevice), the same information can be present in both the Name and ElementName properties. Note that if there is an associated instance of CIM_EnabledLogicalElementCapabilities, restrictions on this properties may exist as defined in ElementNameMask and MaxElementNameLen properties defined in that class.
		ExtendedData            string   `xml:"ExtendedData,omitempty"`            // Data associated with the policy, up to 32 bytes. The data should be in a network order. The extended data for a policy with a periodic trigger contains first a periodic type and after that the data for that type. For periodic type 0 [Interval - The CIRA connection will be established every fixed number of seconds] - the data should include a uint32 value that indicates the time period in seconds between tunnel establishments. For periodic type 1 [Daily - The CIRA connection will be established every day in a specific pre-defined time (hour and minutes)] - the data should include two uint32 values which define the wanted hour of the day and minutes of that hour. For the other triggers extended data is not defined and not needed. The length and data should be zero.
		PolicyRuleName          string   `xml:"PolicyRuleName,omitempty"`          // A user-friendly name of this PolicyRule. In Intel AMT Release 6.0 and later releases value is "%s %d" where %s is the policy type ("User Initiated" / "Alert" / "Periodic") and %d is the policy handle
		SystemCreationClassName string   `xml:"SystemCreationClassName,omitempty"` // The scoping System's CreationClassName.
		SystemName              string   `xml:"SystemName,omitempty"`              // The scoping System's Name.
		Trigger                 Trigger  `xml:"Trigger"`                           // The event that will trigger the establishment of the remote connection to the MpServer.
		TunnelLifeTime          int      `xml:"TunnelLifeTime"`                    // Defines the tunnel lifetime in seconds, 0 means no lifetime- the tunnel should stay open until it is closed by CloseRemoteAccessConnection or when a different policy with higher priority needs to be processed.
	}
	RemoteAccessPolicyAppliesToMPSResponse struct {
		XMLName        xml.Name               `xml:"AMT_RemoteAccessPolicyAppliesToMPS"`
		ManagedElement ManagedElementResponse `xml:"ManagedElement"` // The MpServer to which the policy applies.
		MpsType        MPSType                `xml:"MpsType"`        // This field indicates if the MpServer is to be used inside or outside of the organization, or both. Default is outside (0).
		OrderOfAccess  int                    `xml:"OrderOfAccess"`  // This field indicates in what order will the Intel® AMT subsystem attempt to connect to the referenced MpServer when the referenced Remote Access policy is triggered.
		PolicySet      PolicySetResponse      `xml:"PolicySet"`      // The Remote Access policy rule that is currently applied to the MpServer.
	}
	PullResponse struct {
		XMLName               xml.Name                                 `xml:"PullResponse"`
		RemoteAccessItems     []RemoteAccessServiceResponse            `xml:"Items>AMT_RemoteAccessService"`
		RemotePolicyRuleItems []RemoteAccessPolicyRuleResponse         `xml:"Items>AMT_RemoteAccessPolicyRule"`
		PolicyAppliesItems    []RemoteAccessPolicyAppliesToMPSResponse `xml:"Items>AMT_RemoteAccessPolicyAppliesToMPS"`
	}
	AddMpServerResponse struct {
		XMLName     xml.Name    `xml:"AddMpServer_OUTPUT"`
		MpServer    MpServer    `xml:"MpServer"` // A reference to the created MPS if the operation succeeded.
		ReturnValue ReturnValue `xml:"ReturnValue,omitempty"`
	}
	AddRemoteAccessPolicyRuleResponse struct {
		XMLName            xml.Name           `xml:"AddRemoteAccessPolicyRule_OUTPUT"`
		PolicyRuleResponse PolicyRuleResponse `xml:"PolicyRule"`
		ReturnValue        ReturnValue        `xml:"ReturnValue"` // ValueMap={0, 1, 36, 38, 2058} Values={PT_STATUS_SUCCESS, PT_STATUS_INTERNAL_ERROR, PT_STATUS_INVALID_PARAMETER, PT_STATUS_FLASH_WRITE_LIMIT_EXCEEDED, PT_STATUS_DUPLICATE}
	}
	MpServer struct {
		XMLName             xml.Name                    `xml:"MpServer"`
		Address             string                      `xml:"Address"`
		ReferenceParameters ReferenceParametersResponse `xml:"ReferenceParameters"`
	}
	ManagedElementResponse struct {
		XMLName             xml.Name                    `xml:"ManagedElement"`
		Address             string                      `xml:"Address"`
		ReferenceParameters ReferenceParametersResponse `xml:"ReferenceParameters"`
	}
	PolicySetResponse struct {
		XMLName             xml.Name                    `xml:"PolicySet"`
		Address             string                      `xml:"Address"`
		ReferenceParameters ReferenceParametersResponse `xml:"ReferenceParameters"`
	}
	PolicyRuleResponse struct {
		XMLName             xml.Name                    `xml:"PolicyRule"`
		Address             string                      `xml:"Address"`
		ReferenceParameters ReferenceParametersResponse `xml:"ReferenceParameters"`
	}
	ReferenceParametersResponse struct {
		XMLName     xml.Name            `xml:"ReferenceParameters"`
		ResourceURI string              `xml:"ResourceURI"`
		SelectorSet SelectorSetResponse `xml:"SelectorSet"`
	}
	SelectorSetResponse struct {
		XMLName   xml.Name           `xml:"SelectorSet"`
		Selectors []SelectorResponse `xml:"Selector"`
	}
	SelectorResponse struct {
		XMLName xml.Name `xml:"Selector"`
		Name    string   `xml:"Name,attr"`
		Text    string   `xml:",chardata"`
	}
)

// INPUTS
// Request Types.
type (
	AddMpServerRequest struct {
		XMLName     xml.Name           `xml:"h:AddMpServer_INPUT"`
		H           string             `xml:"xmlns:h,attr"`
		AccessInfo  string             `xml:"h:AccessInfo"`            // A string holding the IP address or FQDN of the server
		InfoFormat  MPServerInfoFormat `xml:"h:InfoFormat"`            // An enumerated integer describing the format and interpretation of the AccessInfo property.
		Port        int                `xml:"h:Port"`                  // The port to be used to establish a tunnel with the MPS.
		AuthMethod  MPServerAuthMethod `xml:"h:AuthMethod"`            // Authentication method to be used when the Intel® AMT subsystem opens a tunnel to the MpServer
		Username    string             `xml:"h:Username,omitempty"`    // A Username to be used for the connection with the MPS if Username-Pwd authentication is used. Limited to 16 alphanumeric characters
		Password    string             `xml:"h:Password,omitempty"`    // A Password to be used for the connection with the MPS if Username-Pwd authentication is used. Limited to 16 characters
		CommonName  string             `xml:"h:CN"`                    // A common name used when AccessInfo is an IP address.
		Certificate string             `xml:"h:Certificate,omitempty"` // A reference to a certificate. Required if AuthMethod is set to mutual authentication
	}
	RemoteAccessPolicyRuleRequest struct {
		XMLName        xml.Name `xml:"h:AddRemoteAccessPolicyRule_INPUT"`
		H              string   `xml:"xmlns:h,attr"`
		Trigger        Trigger  `xml:"h:Trigger"`        // The event that will trigger the establishment of the remote connection to the MpServer.
		TunnelLifeTime int      `xml:"h:TunnelLifeTime"` // Defines the tunnel lifetime in seconds, 0 means no lifetime- the tunnel should stay open until it is closed by CloseRemoteAccessConnection or when a different policy with higher priority needs to be processed.
		ExtendedData   string   `xml:"h:ExtendedData"`   // Data associated with the policy, up to 32 bytes. The data should be in a network order. The extended data for a policy with a periodic trigger contains first a periodic type and after that the data for that type. For periodic type 0 [Interval - The CIRA connection will be established every fixed number of seconds] - the data should include a uint32 value that indicates the time period in seconds between tunnel establishments. For periodic type 1 [Daily - The CIRA connection will be established every day in a specific pre-defined time (hour and minutes)] - the data should include two uint32 values which define the wanted hour of the day and minutes of that hour. For the other triggers extended data is not defined and not needed. The length and data should be zero.
	}
	RemoteAccessPolicyAppliesToMPSRequest struct {
		XMLName        xml.Name       `xml:"h:AMT_RemoteAccessPolicyAppliesToMPS"`
		H              string         `xml:"xmlns:h,attr"`
		ManagedElement ManagedElement `xml:"h:ManagedElement"` // The MpServer to which the policy applies.
		OrderOfAccess  int            `xml:"h:OrderOfAccess"`  // This field indicates in what order will the Intel® AMT subsystem attempt to connect to the referenced MpServer when the referenced Remote Access policy is triggered.
		MPSType        MPSType        `xml:"h:MpsType"`        // This field indicates if the MpServer is to be used inside or outside of the organization, or both. Default is outside (0).
		PolicySet      PolicySet      `xml:"h:PolicySet"`      // The Remote Access policy rule that is currently applied to the MpServer.
	}
	ManagedElement struct {
		Address             string              `xml:"b:Address"`
		B                   string              `xml:"xmlns:b,attr"`
		ReferenceParameters ReferenceParameters `xml:"b:ReferenceParameters"`
	}
	ReferenceParameters struct {
		ResourceURI string      `xml:"c:ResourceURI"`
		C           string      `xml:"xmlns:c,attr"`
		SelectorSet SelectorSet `xml:"c:SelectorSet"`
	}
	SelectorSet struct {
		Selectors []Selector `xml:"c:Selector"`
	}
	Selector struct {
		Name string `xml:"Name,attr"`
		Text string `xml:",chardata"`
	}
	PolicySet struct {
		Address             string              `xml:"b:Address"`
		B                   string              `xml:"xmlns:b,attr"`
		ReferenceParameters ReferenceParameters `xml:"b:ReferenceParameters"`
	}
)

// Property Types.
type (
	RemoteAccessPolicyRuleSelector message.Selector
	PolicyDecisionStrategy         int // First Matching:1 | All:2
	// An enumerated integer describing the format and interpretation of the AccessInfo property.
	//
	// ValueMap={3, 4, 201}
	//
	// Values={IPv4 Address, IPv6 Address, FQDN}.
	MPServerInfoFormat int
	// Authentication method to be used when the Intel® AMT subsystem opens a tunnel to the MpServer
	//
	// ValueMap={1, 2}
	//
	// Values={Mutual Authentication, Username Password Authentication}.
	MPServerAuthMethod int
	// The event that will trigger the establishment of the remote connection to the MpServer.
	//
	// ValueMap={0, 1, 2, 3}
	//
	// Values={User Initiated, Alert, Periodic, Home Provisioning}.
	Trigger int
	// This field indicates if the MpServer is to be used inside or outside of the organization, or both. Default is outside (0).
	//
	// ValueMap={0, 1, 2}
	//
	// Values={External MPS, Internal MPS, Both}.
	MPSType int

	// ReturnValue is an integer enumeration that indicates the completion status of the method. A value of 0 indicates success. A non-zero value indicates an error.
	ReturnValue int
)
