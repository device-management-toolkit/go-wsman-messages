/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package boot

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

// NewBootSourceSetting returns a new instance of the BootSourceSetting struct.
func NewBootSourceSettingWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) SourceSetting {
	return SourceSetting{
		base: message.NewBaseWithClient(wsmanMessageCreator, CIMBootSourceSetting, client),
	}
}

// Get retrieves the representation of the instance.
func (sourceSetting SourceSetting) Get(instanceID string) (response Response, err error) {
	selector := message.Selector{
		Name:  "InstanceID",
		Value: instanceID,
	}
	response = Response{
		Message: &client.Message{
			XMLInput: sourceSetting.base.Get(&selector),
		},
	}

	err = sourceSetting.base.Execute(response.Message)
	if err != nil {
		return
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}

// Enumerate returns an enumeration context which is used in a subsequent Pull call.
func (sourceSetting SourceSetting) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: sourceSetting.base.Enumerate(),
		},
	}

	err = sourceSetting.base.Execute(response.Message)
	if err != nil {
		return
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}

// Pull returns the instances of this class.  An enumeration context provided by the Enumerate call is used as input.
func (sourceSetting SourceSetting) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: sourceSetting.base.Pull(enumerationContext),
		},
	}

	err = sourceSetting.base.Execute(response.Message)
	if err != nil {
		return
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}
