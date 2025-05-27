/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package boot facilitates communication with Intel® AMT devices to access the boot capabilities and boot setting data.
//
// Capabilities reports what boot options that the Intel® AMT device supports.
//
// SettingData provides configuration-related and operational parameters for the boot service in the Intel® AMT device.  In order to activate these settings use [pkg/github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/power] RequestPowerStateChange().  Notice that you can't set certain values while others are enabled (for example: You can't set UseIDER or UseSOL if a CIM_BootSourceSetting is chosen).
package boot

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

// NewBootCapabilitiesWithClient instantiates a new Boot Capabilities service.
func NewBootCapabilitiesWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Capabilities {
	return Capabilities{
		base: message.NewBaseWithClient(wsmanMessageCreator, AMTBootCapabilities, client),
	}
}

// Get retrieves the representation of the instance.
func (bootCapabilities Capabilities) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: bootCapabilities.base.Get(nil),
		},
	}
	// send the message to AMT
	err = bootCapabilities.base.Execute(response.Message)
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
func (bootCapabilities Capabilities) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: bootCapabilities.base.Enumerate(),
		},
	}
	// send the message to AMT
	err = bootCapabilities.base.Execute(response.Message)
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
func (bootCapabilities Capabilities) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: bootCapabilities.base.Pull(enumerationContext),
		},
	}
	// send the message to AMT
	err = bootCapabilities.base.Execute(response.Message)
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
