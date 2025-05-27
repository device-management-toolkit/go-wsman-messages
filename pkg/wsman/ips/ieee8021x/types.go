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

// Package Types.
type (
	Settings struct {
		base message.Base
	}
	CredentialContext struct {
		base message.Base
	}
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
		XMLName                   xml.Name `xml:"Body"`
		PullResponse              PullResponse
		EnumerateResponse         common.EnumerateResponse
		IEEE8021xSettingsResponse IEEE8021xSettingsResponse
		SetCertificatesResponse   SetCertificates_OUTPUT
	}

	IEEE8021xSettingsResponse struct {
		XMLName       xml.Name `xml:"IPS_IEEE8021xSettings"`
		ElementName   string   `xml:"ElementName"`
		InstanceID    string   `xml:"InstanceID"`
		Enabled       Enabled  `xml:"Enabled"`
		AvailableInS0 bool     `xml:"AvailableInS0"`
		PxeTimeout    int      `xml:"PxeTimeout"`
	}

	PullResponse struct {
		XMLName                xml.Name                    `xml:"PullResponse"`
		IEEE8021xSettingsItems []IEEE8021xSettingsResponse `xml:"Items>IPS_IEEE8021xSettings"`
		CredentialContextItems []CredentialContextResponse `xml:"Items>IPS_8021xCredentialContext"`
	}
	CredentialContextResponse struct {
		XMLName                 xml.Name                        `xml:"IPS_8021xCredentialContext"`
		ElementInContext        ElementInContextResponse        `xml:"ElementInContext"`
		ElementProvidingContext ElementProvidingContextResponse `xml:"ElementProvidingContext"`
	}
	SelectorResponse struct {
		XMLName xml.Name `xml:"Selector,omitempty"`
		Name    string   `xml:"Name,attr"`
		Text    string   `xml:",chardata"`
	}
	SelectorSetResponse struct {
		XMLName   xml.Name           `xml:"SelectorSet,omitempty"`
		Selectors []SelectorResponse `xml:"Selector,omitempty"`
	}
	ReferenceParametersResponse struct {
		XMLName     xml.Name            `xml:"ReferenceParameters,omitempty"`
		ResourceURI string              `xml:"ResourceURI,omitempty"`
		SelectorSet SelectorSetResponse `xml:"SelectorSet,omitempty"`
	}
	ElementInContextResponse struct {
		XMLName             xml.Name                    `xml:"ElementInContext"`
		Address             string                      `xml:"Address,omitempty"`
		ReferenceParameters ReferenceParametersResponse `xml:"ReferenceParameters,omitempty"`
	}
	ElementProvidingContextResponse struct {
		XMLName             xml.Name                    `xml:"ElementProvidingContext"`
		Address             string                      `xml:"Address,omitempty"`
		ReferenceParameters ReferenceParametersResponse `xml:"ReferenceParameters,omitempty"`
	}
	SetCertificates_OUTPUT struct {
		XMLName     xml.Name `xml:"SetCertificates_OUTPUT"`
		ReturnValue ReturnValue
	}

	// Enabled indicates whether the IEEE 802.1x settings are enabled or disabled.
	Enabled int

	// ReturnValue indicates the status of the operation.
	ReturnValue int
)

// INPUT
// Request Types.
type (
	IEEE8021xSettingsRequest struct {
		XMLName                         xml.Name `xml:"h:IPS_IEEE8021xSettings,omitempty"`
		H                               string   `xml:"xmlns:h,attr"`
		ElementName                     string   `xml:"h:ElementName,omitempty"`
		InstanceID                      string   `xml:"h:InstanceID,omitempty"`
		AuthenticationProtocol          int      `xml:"h:AuthenticationProtocol"`
		RoamingIdentity                 string   `xml:"h:RoamingIdentity,omitempty"`
		ServerCertificateName           string   `xml:"h:ServerCertificateName,omitempty"`
		ServerCertificateNameComparison int      `xml:"h:ServerCertificateNameComparison,omitempty"`
		Username                        string   `xml:"h:Username,omitempty"`
		Password                        string   `xml:"h:Password,omitempty"`
		Domain                          string   `xml:"h:Domain,omitempty"`
		ProtectedAccessCredential       string   `xml:"h:ProtectedAccessCredential,omitempty"`
		PACPassword                     string   `xml:"h:PACPassword,omitempty"`
		PSK                             string   `xml:"h:PSK,omitempty"`
		Enabled                         int      `xml:"h:Enabled,omitempty"`
		PxeTimeout                      int      `xml:"h:PxeTimeout,omitempty"`
		AvailableInS0                   bool     `xml:"h:AvailableInS0,omitempty"`
	}
	Certificate struct {
		XMLName                 xml.Name `xml:"h:SetCertificates_INPUT"`
		H                       string   `xml:"xmlns:h,attr"`
		ServerCertificateIssuer ServerCertificateIssuer
		ClientCertificate       ClientCertificateIssuer
	}
	ServerCertificateIssuer struct {
		XMLName             xml.Name            `xml:"h:ServerCertificateIssuer"`
		Address             string              `xml:"a:Address"`
		ReferenceParameters ReferenceParameters `xml:"a:ReferenceParameters"`
	}
	ClientCertificateIssuer struct {
		XMLName             xml.Name            `xml:"h:ClientCertificate"`
		Address             string              `xml:"a:Address"`
		ReferenceParameters ReferenceParameters `xml:"a:ReferenceParameters"`
	}
	ReferenceParameters struct {
		XMLName     xml.Name    `xml:"a:ReferenceParameters"`
		ResourceURI string      `xml:"w:ResourceURI"`
		SelectorSet SelectorSet `xml:"w:SelectorSet"`
	}
	SelectorSet struct {
		XMLName  xml.Name `xml:"w:SelectorSet"`
		Selector Selector `xml:"w:Selector"`
	}
	Selector struct {
		XMLName xml.Name `xml:"w:Selector"`
		Name    string   `xml:"Name,attr"`
		Value   string   `xml:",chardata"`
	}
)
