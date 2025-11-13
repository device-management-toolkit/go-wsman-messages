/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package wifi

import (
	"encoding/xml"
	"errors"
	"strconv"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/methods"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

type Port struct {
	base.WSManService[Response]
}

// NewWiFiPort returns a new instance of the WiFiPort struct.
func NewWiFiPortWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Port {
	return Port{
		base.NewService[Response](wsmanMessageCreator, CIMWiFiPort, client),
	}
}

// RequestStateChange requests that the state of the element be changed to the value specified in the RequestedState parameter . . .
func (port Port) RequestStateChange(requestedState int) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: port.Base.RequestStateChange(methods.GenerateAction(CIMWiFiPort, "RequestStateChange"), requestedState),
		},
	}

	err = port.Base.Execute(response.Message)
	if err != nil {
		return response, err
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	if response.Body.RequestStateChange_OUTPUT.ReturnValue != 0 {
		err = errors.New("RequestStateChange failed with return code " + strconv.Itoa(response.Body.RequestStateChange_OUTPUT.ReturnValue))
	}

	return response, err
}
