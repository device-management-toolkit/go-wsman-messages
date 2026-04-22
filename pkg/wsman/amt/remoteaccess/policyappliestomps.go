/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package remoteaccess facilitiates communication with Intel® AMT devices to access and configure Remote Access Policy Applies to MPS, Remote Access Policy Rules, and Remote Access Service.
//
// Remote Access Policy Applies To MPS:
// This class associates a Management Presence Server with a Remote Access Policy rule.
// When a Policy Rule is triggered, the Intel® AMT subsystem will attempt to connect to the MpServers associated with the triggered policy in the order by which the associations were created.
// This order is indicated in the OrderOfAccess field where lower numbers indicate a higher priority.
//
// Remote Access Policy Rule:
// Represents a Remote Access policy.
// The policy defines a condition that will trigger the establishment of a tunnel between the Intel® AMT subsystem and a remote MpServer.
// The policy also defines parameters for the connection such as TunnelLifeTime in seconds.
//
// Remote Access Service:
// Represents the Remote Access Service in the Intel® AMT subsystem.
package remoteaccess

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

// NewRemoteAccessPolicyAppliesToMPSWithClient instantiates a new PolicyAppliesToMPS.
func NewRemoteAccessPolicyAppliesToMPSWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) PolicyAppliesToMPS {
	return PolicyAppliesToMPS{
		base.NewService[Response](wsmanMessageCreator, AMTRemoteAccessPolicyAppliesToMPS, client),
	}
}

// Put overrides the generic Put because the target instance must be addressed
// via two EPR selectors (ManagedElement and PolicySet), which the generic Put
// does not provide.
func (policyAppliesToMPS PolicyAppliesToMPS) Put(remoteAccessPolicyAppliesToMPS *RemoteAccessPolicyAppliesToMPSRequest) (response Response, err error) {
	selectors := []message.Selector{
		{
			Name:  "ManagedElement",
			Value: `<EndpointReference xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing"><Address xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing">http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</Address><ReferenceParameters xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing"><ResourceURI xmlns="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd">http://intel.com/wbem/wscim/1/amt-schema/1/AMT_ManagementPresenceRemoteSAP</ResourceURI><SelectorSet xmlns="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd"><Selector Name="CreationClassName">AMT_ManagementPresenceRemoteSAP</Selector><Selector Name="Name">Intel(r) AMT:Management Presence Server 0</Selector><Selector Name="SystemCreationClassName">CIM_ComputerSystem</Selector><Selector Name="SystemName">Intel(r) AMT</Selector></SelectorSet></ReferenceParameters></EndpointReference>`,
		},
		{
			Name:  "PolicySet",
			Value: `<EndpointReference xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing"><Address xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing">http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</Address><ReferenceParameters xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing"><ResourceURI xmlns="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd">http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyRule</ResourceURI><SelectorSet xmlns="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd"><Selector Name="CreationClassName">AMT_RemoteAccessPolicyRule</Selector><Selector Name="PolicyRuleName">Periodic</Selector><Selector Name="SystemCreationClassName">CIM_ComputerSystem</Selector><Selector Name="SystemName">Intel(r) AMT</Selector></SelectorSet></ReferenceParameters></EndpointReference>`,
		},
	}

	response = Response{
		Message: &client.Message{
			XMLInput: policyAppliesToMPS.Base.Put(remoteAccessPolicyAppliesToMPS, true, selectors),
		},
	}
	// send the message to AMT
	err = policyAppliesToMPS.Base.Execute(response.Message)
	if err != nil {
		return response, err
	}

	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	return response, err
}

// Delete removes the specified instance.
func (policyAppliesToMPS PolicyAppliesToMPS) Delete(handle string) (response Response, err error) {
	selector := message.Selector{Name: "Name", Value: handle}
	response = Response{
		Message: &client.Message{
			XMLInput: policyAppliesToMPS.Base.Delete(selector),
		},
	}
	// send the message to AMT
	err = policyAppliesToMPS.Base.Execute(response.Message)
	if err != nil {
		return response, err
	}

	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	return response, err
}
