/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package managementpresence facilitiates communication with Intel® AMT devices to configure Management Presence Remote Service Access Points (or an MPS) to be accessed by the Intel® AMT subsystem from remote.
package managementpresence

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

// NewManagementPresenceRemoteSAPWithClient instantiates a new RemoteSAP.
func NewManagementPresenceRemoteSAPWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) RemoteSAP {
	return RemoteSAP{
		base: message.NewBaseWithClient(wsmanMessageCreator, AMTManagementPresenceRemoteSAP, client),
	}
}

// Get retrieves the representation of the instance.
func (remoteSAP RemoteSAP) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: remoteSAP.base.Get(nil),
		},
	}
	// send the message to AMT
	err = remoteSAP.base.Execute(response.Message)
	if err != nil {
		return
	}
	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}

// Enumerate returns an enumeration context which is used in a subsequent Pull call.
func (remoteSAP RemoteSAP) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: remoteSAP.base.Enumerate(),
		},
	}
	// send the message to AMT
	err = remoteSAP.base.Execute(response.Message)
	if err != nil {
		return
	}
	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}

// Pull returns the instances of this class.  An enumeration context provided by the Enumerate call is used as input.
func (remoteSAP RemoteSAP) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: remoteSAP.base.Pull(enumerationContext),
		},
	}
	// send the message to AMT
	err = remoteSAP.base.Execute(response.Message)
	if err != nil {
		return
	}
	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}

// Delete removes a the specified instance.
func (remoteSAP RemoteSAP) Delete(handle string) (response Response, err error) {
	selector := message.Selector{Name: "Name", Value: handle}
	response = Response{
		Message: &client.Message{
			XMLInput: remoteSAP.base.Delete(selector),
		},
	}
	// send the message to AMT
	err = remoteSAP.base.Execute(response.Message)
	if err != nil {
		return
	}
	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}
