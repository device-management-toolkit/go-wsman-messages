/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package boot facilitates communication with Intel® AMT devices to access and configure Boot Config Setting, Boot Service, and Boot Source Setting features of AMT.
//
// ConfigSetting:
// A class derived from SettingData that provides the container to arrange all the BootSourceSetting instances in an ordered sequence.
// There can be one or more of the BootConfigSetting instances associated to a ComputerSystem.
// For example, one such BootConfigSetting could be a default boot configuration supplied by the manufacturer, a second one could be a configuration recommended by the IT Administrator.
// A third one could be the one actually to be used on next system boot.
//
// Service:
// A class derived from Service that provides the controls to manage the boot configuration of a managed computer system or device.
// This includes changing the order of the boot devices and affecting settings on managed elements during the boot process.
// This service can also affect the load of a specific operating system on the computer system through a BootSourceSetting that points to a specific operating system image.
//
// SourceSetting:
// A class derived from SettingData that provides the information necessary to describe a boot source.
// This may be optionally associated to a bootable logical device, such as a hard disk partition, or a network device.
// The information from this class instance is used by the boot manager, such as BIOS/EFI or OS Loader to initiate the boot process, when this instance appears in a BootConfigSetting collection.
package boot

import (
	"encoding/xml"
	"errors"
	"fmt"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/methods"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

type ConfigSetting struct {
	base.WSManService[Response]
}

// NewBootConfigSettingWithClient instantiates a new ConfigSetting.
func NewBootConfigSettingWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) ConfigSetting {
	return ConfigSetting{
		base.NewService[Response](wsmanMessageCreator, CIMBootConfigSetting, client),
	}
}

// ChangeBootOrder sets the boot order within a boot configuration.
//
// An ordered array of BootSourceSetting instances is passed to this method.
// Each BootSourceSetting instance MUST already be associated with this BootConfigSetting instance via an instance of OrderedComponent.
// If not, the implementation MUST return a value of "Invalid Parameter" Upon execution of this method,
// the value of the AssignedSequence property on each instance of OrderedComponent will be updated such that the values are monotonically increasing in correlation with the position of the referenced BootSourceSetting instance in the source input parameter.
// That is, the first position in the array will have the lowest value for AssignedSequence.
// The second position will have the second lowest value, and so on.
// For BootSourceSetting instances which are associated with the BootConfigSetting instance via OrderedComponent and not present in the input array, the AssignedSequence property on the OrderedComponent association will be assigned a value of 0.
//
// Additional Notes:
//
// 1) A boot source cannot be set if some special boot options were set in AMT_BootSettingData (such as UseSOL, UseIDER, ReflashBIOS, BIOSPause, BIOSSetup)
//
// 2) Parameter 'Source' changed in capitalization. Intel AMT Release 5.0 and earlier releases use 2.13.0 MOF version and therefor expect 'Source' parameter as 'source'.
//
// 3) Intel AMT Release 7.0: Returns WSMAN Fault = “access denied” if user consent is required but IPS_OptInService.OptInState value is not 'Received' or 'In Session'. An exception to this rule is when the Source parameter is an empty array.
func (configSetting ConfigSetting) ChangeBootOrder(source Source) (response Response, err error) {
	header := configSetting.Base.WSManMessageCreator.CreateHeader(methods.GenerateAction(CIMBootConfigSetting, ChangeBootOrder), CIMBootConfigSetting, nil, "", "")

	var body string

	if source == "" {
		body = `<Body><h:ChangeBootOrder_INPUT xmlns:h="http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootConfigSetting"></h:ChangeBootOrder_INPUT></Body>`
	} else {
		body = fmt.Sprintf(`<Body><h:ChangeBootOrder_INPUT xmlns:h="http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootConfigSetting"><h:Source><Address xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing">http://schemas.xmlsoap.org/ws/2004/08/addressing</Address><ReferenceParameters xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing"><ResourceURI xmlns="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd">http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting</ResourceURI><SelectorSet xmlns="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd"><Selector Name="InstanceID">%s</Selector></SelectorSet></ReferenceParameters></h:Source></h:ChangeBootOrder_INPUT></Body>`, source)
	}

	response = Response{
		Message: &client.Message{
			XMLInput: configSetting.Base.WSManMessageCreator.CreateXML(header, body),
		},
	}

	err = configSetting.Base.Execute(response.Message)
	if err != nil {
		return response, err
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	if response.Body.ChangeBootOrder_OUTPUT.ReturnValue != 0 {
		err = generateErrorMessage("changebootorder", response.Body.ChangeBootOrder_OUTPUT.ReturnValue)
	}

	return response, err
}

// generateErrorMessage returns an error message based on the return value.
func generateErrorMessage(call string, returnValue ReturnValue) error {
	ErrSetupFailed := errors.New(call + " failed")

	return fmt.Errorf("%w: returned %d", ErrSetupFailed, returnValue)
}
