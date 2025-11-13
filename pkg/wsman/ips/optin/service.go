/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package optin facilitates communication with Intel® AMT devices to describe the user consent service.
//
// This service manages user opt-in options and sends a user consent code for KVM, redirection, and set boot options.
package optin

import (
	"encoding/xml"
	"fmt"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/ips/actions"
)

type Service struct {
	base.WSManService[Response]
}

// NewOptInService returns a new instance of the OptInService struct.
func NewOptInServiceWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Service {
	return Service{
		base.NewService[Response](wsmanMessageCreator, IPSOptInService, client),
	}
}

// Send the opt-in code to Intel® AMT.
func (service Service) SendOptInCode(optInCode int) (response Response, err error) {
	header := service.Base.WSManMessageCreator.CreateHeader(string(actions.SendOptInCode), IPSOptInService, nil, "", "")
	body := service.Base.WSManMessageCreator.CreateBody("SendOptInCode_INPUT", IPSOptInService, OptInCode{
		H:         "http://intel.com/wbem/wscim/1/ips-schema/1/IPS_OptInService",
		OptInCode: optInCode,
	})
	response = Response{
		Message: &client.Message{
			XMLInput: service.Base.WSManMessageCreator.CreateXML(header, body),
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

	return response, err
}

// Request an opt-in code.
func (service Service) StartOptIn() (response Response, err error) {
	header := service.Base.WSManMessageCreator.CreateHeader(string(actions.StartOptIn), IPSOptInService, nil, "", "")
	body := service.Base.WSManMessageCreator.CreateBody("StartOptIn_INPUT", IPSOptInService, nil)
	response = Response{
		Message: &client.Message{
			XMLInput: service.Base.WSManMessageCreator.CreateXML(header, body),
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

	return response, err
}

// Cancel a previous opt-in code request.
func (service Service) CancelOptIn() (response Response, err error) {
	header := service.Base.WSManMessageCreator.CreateHeader(string(actions.CancelOptIn), IPSOptInService, nil, "", "")
	body := service.Base.WSManMessageCreator.CreateBody("CancelOptIn_INPUT", IPSOptInService, nil)
	response = Response{
		Message: &client.Message{
			XMLInput: service.Base.WSManMessageCreator.CreateXML(header, body),
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

	return response, err
}

// Put will change properties of the selected instance.
func (service Service) Put(request OptInServiceRequest) (response Response, err error) {
	request.H = fmt.Sprintf("%s%s", message.IPSSchema, IPSOptInService)
	response = Response{
		Message: &client.Message{
			XMLInput: service.Base.Put(request, false, nil),
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

	return response, err
}
