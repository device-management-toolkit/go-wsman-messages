/*********************************************************************
 * Copyright (c) Intel Corporation 2025
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package screensetting

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
)

// OUTPUT.
type (
	Response struct {
		*client.Message
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}

	Body struct {
		XMLName                   xml.Name `xml:"Body"`
		PullResponse              PullResponse
		EnumerateResponse         common.EnumerateResponse
		ScreenSettingDataResponse ScreenSettingDataResponse
	}

	ScreenSettingDataResponse struct {
		XMLName        xml.Name `xml:"IPS_ScreenSettingData"`
		ElementName    string   `xml:"ElementName"`
		InstanceID     string   `xml:"InstanceID"`
		PrimaryIndex   int      `xml:"PrimaryIndex"`
		SecondaryIndex int      `xml:"SecondaryIndex"`
		TertiaryIndex  int      `xml:"TertiaryIndex"`
		QuadraryIndex  int      `xml:"QuadraryIndex"`
		IsActive       []bool   `xml:"IsActive"`
		UpperLeftX     []int    `xml:"UpperLeftX"`
		UpperLeftY     []int    `xml:"UpperLeftY"`
		ResolutionX    []int    `xml:"ResolutionX"`
		ResolutionY    []int    `xml:"ResolutionY"`
	}

	PullResponse struct {
		XMLName                xml.Name                    `xml:"PullResponse"`
		ScreenSettingDataItems []ScreenSettingDataResponse `xml:"Items>IPS_ScreenSettingData"`
	}
)

// INPUT.
type (
	ScreenSettingDataRequest struct {
		XMLName        xml.Name `xml:"h:IPS_ScreenSettingData,omitempty"`
		H              string   `xml:"xmlns:h,attr"`
		ElementName    string   `xml:"h:ElementName,omitempty"`
		InstanceID     string   `xml:"h:InstanceID,omitempty"`
		PrimaryIndex   uint8    `xml:"h:PrimaryIndex,omitempty"`
		SecondaryIndex uint8    `xml:"h:SecondaryIndex,omitempty"`
		TertiaryIndex  uint8    `xml:"h:TertiaryIndex,omitempty"`
		QuadraryIndex  uint8    `xml:"h:QuadraryIndex,omitempty"`
		IsActive       []bool   `xml:"h:IsActive,omitempty"`
		UpperLeftX     []int32  `xml:"h:UpperLeftX,omitempty"`
		UpperLeftY     []int32  `xml:"h:UpperLeftY,omitempty"`
		ResolutionX    []uint32 `xml:"h:ResolutionX,omitempty"`
		ResolutionY    []uint32 `xml:"h:ResolutionY,omitempty"`
	}
)
