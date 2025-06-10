/*********************************************************************
 * Copyright (c) Intel Corporation 2025
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package power facilitates communication with Intel® AMT devices where a class derived from Service describes power management functionality, hosted on a System.
//
// Whether this service might be used to affect the power state of a particular element is defined by the CIM_ServiceAvailable ToElement association.
package power

import (
	"encoding/xml"
	"fmt"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/ips/methods"
)

type ManagementService struct {
	base.WSManService[Response]
}

// NewPowerManagementService returns a new instance of the PowerManagementService struct.
func NewPowerManagementServiceWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) ManagementService {
	return ManagementService{
		base.NewService[Response](wsmanMessageCreator, IPSPowerManagementService, client),
	}
}

// RequestOSPowerSavingStateChange defines the desired OS powersaving state of the managed element, and when the element should be put into that state.
func (managementService ManagementService) RequestOSPowerSavingStateChange(osPowerSavingState OSPowerSavingState) (response Response, err error) {
	header := managementService.Base.WSManMessageCreator.CreateHeader(methods.GenerateAction(IPSPowerManagementService, RequestOSPowerSavingStateChange), IPSPowerManagementService, nil, "", "")

	body := fmt.Sprintf(`<Body><h:RequestOSPowerSavingStateChange_INPUT xmlns:h="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_PowerManagementService"><h:OSPowerSavingState>%d</h:OSPowerSavingState><h:ManagedElement><Address xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing">http://schemas.xmlsoap.org/ws/2004/08/addressing</Address><ReferenceParameters xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing"><ResourceURI xmlns="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd">http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ComputerSystem</ResourceURI><SelectorSet xmlns="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd"><Selector Name="CreationClassName">CIM_ComputerSystem</Selector><Selector Name="Name">ManagedSystem</Selector></SelectorSet></ReferenceParameters></h:ManagedElement></h:RequestOSPowerSavingStateChange_INPUT></Body>`, osPowerSavingState)
	response = Response{
		Message: &client.Message{
			XMLInput: managementService.Base.WSManMessageCreator.CreateXML(header, body),
		},
	}

	// send the message to AMT
	err = managementService.Base.Execute(response.Message)
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
