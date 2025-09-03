/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package userinitiatedconnection facilitiates communication with Intel® AMT devices to access and change the state of the user initiated connection feature of AMT.
package userinitiatedconnection

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/amt/methods"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

type Service struct {
	base.WSManService[Response]
}

func NewUserInitiatedConnectionServiceWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Service {
	return Service{
		base.NewService[Response](wsmanMessageCreator, AMTUserInitiatedConnectionService, client),
	}
}

// Requests that the state of the element be changed to the value specified in the RequestedState parameter.
// When the requested state change takes place, the EnabledState and RequestedState of the element will be the same.
// Invoking the RequestStateChange method multiple times could result in earlier requests being overwritten or lost.
// If 0 is returned, then the task completed successfully and the use of ConcreteJob was not required.
// If 4096 (0x1000) is returned, then the task will take some time to complete, ConcreteJob will be created, and its reference returned in the output parameter Job.
// Any other return code indicates an error condition.
//
// Additional Notes:
//
// 1) In Intel AMT Release 5.0 and earlier releases 'datetime' format is simple string. In Intel AMT Release 5.1 and later releases 'datetime' format is as defined in DSP0230 'DMTF WS-CIM Mapping Specification'.
//
// 2) AMT doesn't support the TimeoutPeriod parameter (only value 0 is valid).
//
// 3) The supported values in RequestedState are 32768-32771.
//
// ValueMap={0, 1, 2, 3, 4, 5, 6, .., 4096, 4097, 4098, 4099, 4100..32767, 32768..65535}
//
// Values={Completed with No Error, Not Supported, Unknown or Unspecified Error, Cannot complete within Timeout Period, Failed, Invalid Parameter, In Use, DMTF Reserved, Method Parameters Checked - Job Started, Invalid State Transition, Use of Timeout Parameter Not Supported, Busy, Method Reserved, Vendor Specific}.
func (service Service) RequestStateChange(requestedState RequestedState) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: service.Base.RequestStateChange(methods.RequestStateChange(AMTUserInitiatedConnectionService), int(requestedState)),
		},
	}
	// send the message to AMT
	err = service.Base.Execute(response.Message)
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
