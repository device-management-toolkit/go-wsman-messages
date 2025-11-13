/*********************************************************************
 * Copyright (c) Intel Corporation 2025
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package kvmredirection

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
		XMLName                        xml.Name `xml:"Body"`
		PullResponse                   PullResponse
		EnumerateResponse              common.EnumerateResponse
		KVMRedirectionSettingsResponse KVMRedirectionSettingsResponse
		TerminateSessionResponse       TerminateSession_OUTPUT
	}

	KVMRedirectionSettingsResponse struct {
		XMLName                        xml.Name `xml:"IPS_KVMRedirectionSettingData"`
		ElementName                    string   `xml:"ElementName"`
		InstanceID                     string   `xml:"InstanceID"`
		EnabledByMEBx                  bool     `xml:"EnabledByMEBx"`
		BackToBackFbMode               bool     `xml:"BackToBackFbMode"`
		Is5900PortEnabled              bool     `xml:"Is5900PortEnabled"`
		OptInPolicy                    bool     `xml:"OptInPolicy"`
		OptInPolicyTimeout             uint16   `xml:"OptInPolicyTimeout"`
		SessionTimeout                 uint16   `xml:"SessionTimeout"`
		RFBPassword                    string   `xml:"RFBPassword"`
		DefaultScreen                  uint8    `xml:"DefaultScreen"`
		InitialDecimationModeForLowRes uint8    `xml:"InitialDecimationModeForLowRes"`
		GreyscalePixelFormatSupported  bool     `xml:"GreyscalePixelFormatSupported"`
		ZlibControlSupported           bool     `xml:"ZlibControlSupported"`
		DoubleBufferMode               bool     `xml:"DoubleBufferMode"`
		DoubleBufferState              bool     `xml:"DoubleBufferState"`
	}
	PullResponse struct {
		XMLName                     xml.Name                         `xml:"PullResponse"`
		KVMRedirectionSettingsItems []KVMRedirectionSettingsResponse `xml:"Items>IPS_KVMRedirectionSettingData"`
	}
	TerminateSession_OUTPUT struct {
		XMLName     xml.Name `xml:"TerminateSession_OUTPUT"`
		ReturnValue ReturnValue
	}

	// ReturnValue indicates the status of the operation.
	ReturnValue int
)

// INPUT.
type (
	KVMRedirectionSettingsRequest struct {
		XMLName                        xml.Name `xml:"h:IPS_KVMRedirectionSettingData"`
		H                              string   `xml:"xmlns:h,attr"`
		ElementName                    string   `xml:"h:ElementName"`
		InstanceID                     string   `xml:"h:InstanceID"`
		EnabledByMEBx                  bool     `xml:"h:EnabledByMEBx"`
		BackToBackFbMode               bool     `xml:"h:BackToBackFbMode"`
		Is5900PortEnabled              bool     `xml:"h:Is5900PortEnabled"`
		OptInPolicy                    bool     `xml:"h:OptInPolicy"`
		SessionTimeout                 uint16   `xml:"h:SessionTimeout"`
		RFBPassword                    string   `xml:"h:RFBPassword"`
		DefaultScreen                  uint8    `xml:"h:DefaultScreen"`
		InitialDecimationModeForLowRes uint8    `xml:"h:InitialDecimationModeForLowRes"`
		GreyscalePixelFormatSupported  bool     `xml:"h:GreyscalePixelFormatSupported"`
		ZlibControlSupported           bool     `xml:"h:ZlibControlSupported"`
		DoubleBufferMode               bool     `xml:"h:DoubleBufferMode"`
		DoubleBufferState              bool     `xml:"h:DoubleBufferState"`
	}
)
