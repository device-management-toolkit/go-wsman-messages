/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package managementpresence

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
)

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
		XMLName           xml.Name `xml:"Body"`
		GetResponse       ManagementRemoteResponse
		EnumerateResponse common.EnumerateResponse
		PullResponse      PullResponse
	}

	PullResponse struct {
		XMLName               xml.Name                   `xml:"PullResponse"`
		ManagementRemoteItems []ManagementRemoteResponse `xml:"Items>AMT_ManagementPresenceRemoteSAP"`
	}

	ManagementRemoteResponse struct {
		XMLName                 xml.Name   `xml:"AMT_ManagementPresenceRemoteSAP"`
		AccessInfo              string     `xml:"AccessInfo,omitempty"`              // Access or addressing information or a combination of this information for a remote connection. This information can be a host name, network address, or similar information.
		CN                      string     `xml:"CN,omitempty"`                      // A common name used when AccessInfo is an IP address.
		CreationClassName       string     `xml:"CreationClassName,omitempty"`       // CreationClassName indicates the name of the class or the subclass used in the creation of an instance. When used with the other key properties of this class, this property allows all instances of this class and its subclasses to be uniquely identified.
		ElementName             string     `xml:"ElementName,omitempty"`             // A user-friendly name for the object. This property allows each instance to define a user-friendly name in addition to its key properties, identity data, and description information. Note that the Name property of ManagedSystemElement is also defined as a user-friendly name. But, it is often subclassed to be a Key. It is not reasonable that the same property can convey both identity and a user-friendly name, without inconsistencies. Where Name exists and is not a Key (such as for instances of LogicalDevice), the same information can be present in both the Name and ElementName properties. Note that if there is an associated instance of CIM_EnabledLogicalElementCapabilities, restrictions on this properties may exist as defined in ElementNameMask and MaxElementNameLen properties defined in that class.
		InfoFormat              InfoFormat `xml:"InfoFormat,omitempty"`              // An enumerated integer that describes the format and interpretation of the AccessInfo property.
		Name                    string     `xml:"Name,omitempty"`                    // The Name property uniquely identifies the ServiceAccessPoint and provides an indication of the functionality that is managed. This functionality is described in more detail in the Description property of the object.
		Port                    int        `xml:"Port,omitempty"`                    // The port to be used to establish a tunnel with the MPS.
		SystemCreationClassName string     `xml:"SystemCreationClassName,omitempty"` // The CreationClassName of the scoping System.
		SystemName              string     `xml:"SystemName,omitempty"`              // The Name of the scoping System.
	}

	InfoFormat int
)
