/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package tls

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

// NewTLSProtocolEndpointCollectionWithClient instantiates a new ProtocolEndpointCollection.
func NewTLSProtocolEndpointCollectionWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) ProtocolEndpointCollection {
	return ProtocolEndpointCollection{
		base: message.NewBaseWithClient(wsmanMessageCreator, AMTTLSProtocolEndpointCollection, client),
	}
}

// Get retrieves the representation of the instance.
func (collection ProtocolEndpointCollection) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: collection.base.Get(nil),
		},
	}
	// send the message to AMT
	err = collection.base.Execute(response.Message)
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
func (collection ProtocolEndpointCollection) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: collection.base.Enumerate(),
		},
	}
	// send the message to AMT
	err = collection.base.Execute(response.Message)
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
func (collection ProtocolEndpointCollection) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: collection.base.Pull(enumerationContext),
		},
	}
	// send the message to AMT
	err = collection.base.Execute(response.Message)
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
