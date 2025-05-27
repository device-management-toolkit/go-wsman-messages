/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package ieee8021x

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
)

type Settings struct {
	base message.Base
}

// Response Types.
type (
	Response struct {
		*client.Message
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}

	Body struct {
		XMLName           xml.Name `xml:"Body"`
		EnumerateResponse common.EnumerateResponse
		PullResponse      PullResponse `xml:"PullResponse"`
		GetResponse       IEEE8021xSettingsResponse
	}

	IEEE8021xSettingsResponse struct {
		XMLName                         xml.Name `xml:"CIM_IEEE8021xSettings"`
		ElementName                     string   `xml:"ElementName"`
		InstanceID                      string   `xml:"InstanceID"`
		AuthenticationProtocol          int      `xml:"AuthenticationProtocol"`
		RoamingIdentity                 string   `XML:"RoamingIdentity"`
		ServerCertificateName           string   `xml:"ServerCertificateName"`
		ServerCertificateNameComparison int      `xml:"ServerCertificateNameComparison"`
		Username                        string   `xml:"Username"`
		Password                        string   `xml:"Password"`
		Domain                          string   `xml:"Domain"`
		ProtectedAccessCredential       string   `xml:"ProtectedAccessCredential"`
		PACPassword                     string   `xml:"PACPassword"`
		PSK                             string   `xml:"PSK"`
	}

	Time struct {
		DateTime string `xml:"Datetime"`
	}

	PullResponse struct {
		XMLName                xml.Name                    `xml:"PullResponse"`
		IEEE8021xSettingsItems []IEEE8021xSettingsResponse `xml:"Items>CIM_IEEE8021xSettings"`
	}
)

type (
	IEEE8021xSettingsRequest struct {
		H             string   `xml:"xmlns:w,attr"`
		XMLName       xml.Name `xml:"CIM_IEEE8021xSettings"`
		ElementName   string   `xml:"w:ElementName"`   // The user-friendly name for this instance of SettingData.
		InstanceID    string   `xml:"w:InstanceID"`    // Within the scope of the instantiating Namespace, InstanceID opaquely and uniquely identifies an instance of this class.
		Enabled       int      `xml:"w:Enabled"`       // Indicates whether the 802.1x profile is enabled.
		AvailableInS0 bool     `xml:"w:AvailableInS0"` // Indicates the activity setting of the 802.1X module in S0 state.
		PxeTimeout    int      `xml:"w:PxeTimeout"`    // Timeout in seconds, in which the Intel(R) AMT will hold an authenticated 802.1X session.
	}
)
