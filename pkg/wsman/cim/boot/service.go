/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package boot

import (
	"encoding/xml"
	"errors"
	"strconv"
	"strings"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/methods"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

type Service struct {
	base.WSManService[Response]
}

// NewBootService returns a new instance of the BootService struct.
func NewBootServiceWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Service {
	return Service{
		base.NewService[Response](wsmanMessageCreator, CIMBootService, client),
	}
}

func (service Service) SetBootConfigRole(instanceID string, role int) (response Response, err error) {
	header := service.Base.WSManMessageCreator.CreateHeader(methods.GenerateAction(CIMBootService, SetBootConfigRole), CIMBootService, nil, "", "")

	var body strings.Builder

	body.WriteString(`<Body><h:SetBootConfigRole_INPUT xmlns:h="`)
	body.WriteString(service.Base.WSManMessageCreator.ResourceURIBase)
	body.WriteString(`CIM_BootService"><h:BootConfigSetting><Address xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing">http://schemas.xmlsoap.org/ws/2004/08/addressing</Address><ReferenceParameters xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing"><ResourceURI xmlns="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd">http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootConfigSetting</ResourceURI><SelectorSet xmlns="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd">`)
	body.WriteString(`<Selector Name="InstanceID">`)
	body.WriteString(instanceID)
	body.WriteString(`</Selector></SelectorSet></ReferenceParameters></h:BootConfigSetting>`)
	body.WriteString(`<h:Role>`)
	body.WriteString(strconv.Itoa(role))
	body.WriteString(`</h:Role></h:SetBootConfigRole_INPUT></Body>`)

	response = Response{
		Message: &client.Message{
			XMLInput: service.Base.WSManMessageCreator.CreateXML(header, body.String()),
		},
	}

	err = service.Base.Execute(response.Message)
	if err != nil {
		return response, err
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

// RequestStateChange requests that the state of the element be changed to the value specified in the RequestedState parameter . . .
func (service Service) RequestStateChange(requestedState int) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: service.Base.RequestStateChange(methods.GenerateAction(CIMBootService, "RequestStateChange"), requestedState),
		},
	}

	err = service.Base.Execute(response.Message)
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
