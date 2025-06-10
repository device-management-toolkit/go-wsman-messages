/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package timesynchronization facilitiates communication with Intel® AMT devices to synchronize the AMT internal clock with an external clock
package timesynchronization

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

// NewTimeSynchronizationServiceWithClient instantiates a new Service.
func NewTimeSynchronizationServiceWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Service {
	return Service{
		base.NewService[Response](wsmanMessageCreator, AMTTimeSynchronizationService, client),
	}
}

// SetHighAccuracyTimeSynch is used to synchronize the Intel® AMT device's internal clock with an external clock.
//
// ta0: The time value received from invoking GetLowAccuracyTimeSynch().
//
// tm1: The remote client timestamp after getting a response from GetLowAccuracyTimeSynch().
//
// tm2: The remote client timestamp obtained immediately prior to invoking this method.
//
// ValueMap={0, 1, 36, 38}
//
// Values={PT_STATUS_SUCCESS, PT_STATUS_INTERNAL_ERROR, PT_STATUS_INVALID_PARAMETER, PT_STATUS_FLASH_WRITE_LIMIT_EXCEEDED}.
func (service Service) SetHighAccuracyTimeSynch(ta0, tm1, tm2 int64) (response Response, err error) {
	header := service.Base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMTTimeSynchronizationService, SetHighAccuracyTimeSynch), AMTTimeSynchronizationService, nil, "", "")
	body := service.Base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(SetHighAccuracyTimeSynch), AMTTimeSynchronizationService, &SetHighAccuracyTimeSynch_INPUT{
		H:   "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TimeSynchronizationService",
		Ta0: ta0,
		Tm1: tm1,
		Tm2: tm2,
	})
	response = Response{
		Message: &client.Message{
			XMLInput: service.Base.WSManMessageCreator.CreateXML(header, body),
		},
	}
	// send the message to AMT
	err = service.Base.Execute(response.Message)
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

// GetLowAccuracyTimeSynch is used for reading the Intel® AMT device's internal clock.
func (service Service) GetLowAccuracyTimeSynch() (response Response, err error) {
	header := service.Base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMTTimeSynchronizationService, GetLowAccuracyTimeSynch), AMTTimeSynchronizationService, nil, "", "")
	body := service.Base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(GetLowAccuracyTimeSynch), AMTTimeSynchronizationService, nil)
	response = Response{
		Message: &client.Message{
			XMLInput: service.Base.WSManMessageCreator.CreateXML(header, body),
		},
	}
	// send the message to AMT
	err = service.Base.Execute(response.Message)
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
