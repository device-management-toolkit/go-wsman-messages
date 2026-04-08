/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package asset

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

// NewServiceWithClient instantiates a new Asset Table Service.
func NewServiceWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Service {
	return Service{
		base.NewService[Response](wsmanMessageCreator, AMTAssetTableService, client),
	}
}

// GetAssetTableData retrieves asset table data for a specified table ID.
func (s Service) GetAssetTableData(tableID int) (response Response, err error) {
	header := s.Base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMTAssetTableService, GetAssetTableData), AMTAssetTableService, nil, "", "")
	body := s.Base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(GetAssetTableData), AMTAssetTableService, &GetAssetTableData_INPUT{TableID: tableID})

	response = Response{
		Message: &client.Message{
			XMLInput: s.Base.WSManMessageCreator.CreateXML(header, body),
		},
	}

	err = s.Base.Execute(response.Message)
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

// GetAssetTableSize retrieves the size of a specified asset table.
func (s Service) GetAssetTableSize(tableID int) (response Response, err error) {
	header := s.Base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMTAssetTableService, GetAssetTableSize), AMTAssetTableService, nil, "", "")
	body := s.Base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(GetAssetTableSize), AMTAssetTableService, &GetAssetTableSize_INPUT{TableID: tableID})

	response = Response{
		Message: &client.Message{
			XMLInput: s.Base.WSManMessageCreator.CreateXML(header, body),
		},
	}

	err = s.Base.Execute(response.Message)
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
