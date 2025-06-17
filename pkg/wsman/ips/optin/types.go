/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package optin

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
)

// OUTPUT
// Response Types.
type (
	Response struct {
		*client.Message
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}

	Body struct {
		XMLName               xml.Name `xml:"Body"`
		EnumerateResponse     common.EnumerateResponse
		GetAndPutResponse     OptInServiceResponse `xml:"IPS_OptInService"`
		PullResponse          PullResponse         `xml:"PullResponse"`
		StartOptInResponse    StartOptIn_OUTPUT    `xml:"StartOptIn_OUTPUT"`
		CancelOptInResponse   CancelOptIn_OUTPUT   `xml:"CancelOptIn_OUTPUT"`
		SendOptInCodeResponse SendOptInCode_OUTPUT `xml:"SendOptInCode_OUTPUT"`
	}

	OptInServiceResponse struct {
		XMLName                 xml.Name `xml:"IPS_OptInService"`
		Name                    string   `xml:"Name,omitempty"`
		CreationClassName       string   `xml:"CreationClassName,omitempty"`
		SystemName              string   `xml:"SystemName,omitempty"`
		SystemCreationClassName string   `xml:"SystemCreationClassName,omitempty"`
		ElementName             string   `xml:"ElementName,omitempty"`
		OptInCodeTimeout        int      `xml:"OptInCodeTimeout,omitempty"`
		OptInRequired           uint32   `xml:"OptInRequired"`
		OptInState              int      `xml:"OptInState"`
		CanModifyOptInPolicy    int      `xml:"CanModifyOptInPolicy,omitempty"`
		OptInDisplayTimeout     int      `xml:"OptInDisplayTimeout,omitempty"`
	}

	PullResponse struct {
		XMLName           xml.Name               `xml:"PullResponse"`
		OptInServiceItems []OptInServiceResponse `xml:"Items>IPS_OptInService"`
	}
	StartOptIn_OUTPUT struct {
		XMLName     xml.Name `xml:"StartOptIn_OUTPUT"`
		ReturnValue int
	}

	CancelOptIn_OUTPUT struct {
		XMLName     xml.Name `xml:"CancelOptIn_OUTPUT"`
		ReturnValue int
	}

	SendOptInCode_OUTPUT struct {
		XMLName     xml.Name `xml:"SendOptInCode_OUTPUT"`
		ReturnValue int
	}

	// OptInRequired indicates the OptIn (user consent) policy for redirection operations.
	OptInRequired uint32
	// OptInState indicates the current state of the OptIn (user consent) policy.
	OptInState int
	// CanModifyOptInPolicy indicates whether the OptIn policy can be modified.
	CanModifyOptInPolicy int
	// ReturnValue indicates the status of the operation.
	ReturnValue int
)

// INPUT
// Request Types.
type (
	OptInCode struct {
		XMLName   xml.Name `xml:"h:SendOptInCode_INPUT"`
		H         string   `xml:"xmlns:h,attr"`
		OptInCode int      `xml:"h:OptInCode"`
	}
	OptInServiceRequest struct {
		XMLName                 xml.Name `xml:"h:IPS_OptInService"`
		H                       string   `xml:"xmlns:h,attr"`
		CanModifyOptInPolicy    int      `xml:"h:CanModifyOptInPolicy,omitempty"`
		CreationClassName       string   `xml:"h:CreationClassName,omitempty"`
		ElementName             string   `xml:"h:ElementName,omitempty"`
		Name                    string   `xml:"h:Name,omitempty"`
		OptInCodeTimeout        int      `xml:"h:OptInCodeTimeout,omitempty"`
		OptInDisplayTimeout     int      `xml:"h:OptInDisplayTimeout,omitempty"`
		OptInRequired           int      `xml:"h:OptInRequired"`
		OptInState              int      `xml:"h:OptInState"`
		SystemName              string   `xml:"h:SystemName,omitempty"`
		SystemCreationClassName string   `xml:"h:SystemCreationClassName,omitempty"`
	}
)
