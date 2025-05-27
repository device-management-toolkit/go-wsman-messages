/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package ieee8021x

import (
	"encoding/xml"
	"fmt"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

type Profile struct {
	base message.Base
}

// NewIEEE8021xProfileWithClient instantiates a new Profile service.
func NewIEEE8021xProfileWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Profile {
	return Profile{
		base: message.NewBaseWithClient(wsmanMessageCreator, AMTIEEE8021xProfile, client),
	}
}

// Get retrieves the representation of the instance.
func (profile Profile) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: profile.base.Get(nil),
		},
	}
	// send the message to AMT
	err = profile.base.Execute(response.Message)
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
func (profile Profile) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: profile.base.Enumerate(),
		},
	}
	// send the message to AMT
	err = profile.base.Execute(response.Message)
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
func (profile Profile) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: profile.base.Pull(enumerationContext),
		},
	}
	// send the message to AMT
	err = profile.base.Execute(response.Message)
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

// Put will change properties of the selected instance.
func (profile Profile) Put(ieee8021xProfile ProfileRequest) (response Response, err error) {
	ieee8021xProfile.H = fmt.Sprintf("%s%s", message.AMTSchema, AMTIEEE8021xProfile)
	response = Response{
		Message: &client.Message{
			XMLInput: profile.base.Put(ieee8021xProfile, false, nil),
		},
	}
	// send the message to AMT
	err = profile.base.Execute(response.Message)
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
