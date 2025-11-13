/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package ieee8021x

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/wsmantesting"
)

func TestPositiveIPS_IEEE8021xSettings(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.IPSResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "ips/ieee8021x/settings",
	}
	elementUnderTest := NewIEEE8021xSettingsWithClient(wsmanMessageCreator, &client)

	t.Run("ips_IEEE8021xSettings Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			body             string
			extraHeader      string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			// GETS
			{
				"should create a valid IPS_IEEE8021xSettings Get wsman message",
				"IPS_IEEE8021xSettings",
				wsmantesting.Get,
				"",
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageGet

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					IEEE8021xSettingsResponse: IEEE8021xSettingsResponse{
						XMLName:       xml.Name{Space: fmt.Sprintf("%s%s", message.IPSSchema, IPSIEEE8021xSettings), Local: IPSIEEE8021xSettings},
						ElementName:   "Intel(r) AMT: 8021X Settings",
						InstanceID:    "Intel(r) AMT: 8021X Settings",
						Enabled:       3,
						AvailableInS0: false,
						PxeTimeout:    120,
					},
				},
			},
			// ENUMERATES
			{
				"should create a valid IPS_IEEE8021xSettings Enumerate wsman message",
				"IPS_IEEE8021xSettings",
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageEnumerate

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "9C0A0000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create a valid IPS_IEEE8021xSettings Pull wsman message",
				"IPS_IEEE8021xSettings",
				wsmantesting.Pull,
				wsmantesting.PullBody,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessagePull

					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						XMLName: xml.Name{Space: message.XMLPullResponseSpace, Local: "PullResponse"},
						IEEE8021xSettingsItems: []IEEE8021xSettingsResponse{
							{
								XMLName:       xml.Name{Space: fmt.Sprintf("%s%s", message.IPSSchema, IPSIEEE8021xSettings), Local: IPSIEEE8021xSettings},
								ElementName:   "Intel(r) AMT: 8021X Settings",
								InstanceID:    "Intel(r) AMT: 8021X Settings",
								Enabled:       3,
								AvailableInS0: false,
								PxeTimeout:    120,
							},
						},
					},
				},
			},
			// SET CERTIFICATES
			{
				"should create a valid ips_IEEE8021xSettings set certificates wsman message",
				"IPS_IEEE8021xSettings",
				wsmantesting.SetCertificates,
				fmt.Sprintf(`<h:SetCertificates_INPUT xmlns:h="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_IEEE8021xSettings"><h:ServerCertificateIssuer><a:Address>default</a:Address><a:ReferenceParameters><w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyCertificate</w:ResourceURI><w:SelectorSet><w:Selector Name="InstanceID">%s</w:Selector></w:SelectorSet></a:ReferenceParameters></h:ServerCertificateIssuer><h:ClientCertificate><a:Address>default</a:Address><a:ReferenceParameters><w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyCertificate</w:ResourceURI><w:SelectorSet><w:Selector Name="InstanceID">%s</w:Selector></w:SelectorSet></a:ReferenceParameters></h:ClientCertificate></h:SetCertificates_INPUT>`, wsmantesting.ServerCertificateIssuer, wsmantesting.ClientCertificate),
				"",
				func() (Response, error) {
					client.CurrentMessage = "SetCertificates"

					return elementUnderTest.SetCertificates(wsmantesting.ServerCertificateIssuer, wsmantesting.ClientCertificate)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					SetCertificatesResponse: SetCertificates_OUTPUT{
						XMLName:     xml.Name{Space: fmt.Sprintf("%s%s", message.IPSSchema, IPSIEEE8021xSettings), Local: "SetCertificates_OUTPUT"},
						ReturnValue: 0,
					},
				},
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				expectedXMLInput := wsmantesting.ExpectedResponse(messageID, resourceURIBase, test.method, test.action, "", test.body)
				messageID++
				response, err := test.responseFunc()
				assert.NoError(t, err)
				assert.Equal(t, expectedXMLInput, response.XMLInput)
				assert.Equal(t, test.expectedResponse, response.Body)
			})
		}
	})
}

func TestNegativeIPS_IEEE8021xSettings(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.IPSResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "ips/ieee8021x/settings",
	}
	elementUnderTest := NewIEEE8021xSettingsWithClient(wsmanMessageCreator, &client)

	t.Run("ips_IEEE8021xSettings Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			body             string
			extraHeader      string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			// GETS
			{
				"should create a valid IPS_IEEE8021xSettings Get wsman message",
				"IPS_IEEE8021xSettings",
				wsmantesting.Get,
				"",
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					IEEE8021xSettingsResponse: IEEE8021xSettingsResponse{
						XMLName:       xml.Name{Space: fmt.Sprintf("%s%s", message.IPSSchema, IPSIEEE8021xSettings), Local: IPSIEEE8021xSettings},
						ElementName:   "Intel(r) AMT: 8021X Settings",
						InstanceID:    "Intel(r) AMT: 8021X Settings",
						Enabled:       3,
						AvailableInS0: false,
						PxeTimeout:    120,
					},
				},
			},
			// ENUMERATES
			{
				"should create a valid IPS_IEEE8021xSettings Enumerate wsman message",
				"IPS_IEEE8021xSettings",
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "9C0A0000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create a valid IPS_IEEE8021xSettings Pull wsman message",
				"IPS_IEEE8021xSettings",
				wsmantesting.Pull,
				wsmantesting.PullBody,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						XMLName: xml.Name{Space: message.XMLPullResponseSpace, Local: "PullResponse"},
						IEEE8021xSettingsItems: []IEEE8021xSettingsResponse{
							{
								XMLName:       xml.Name{Space: fmt.Sprintf("%s%s", message.IPSSchema, IPSIEEE8021xSettings), Local: IPSIEEE8021xSettings},
								ElementName:   "Intel(r) AMT: 8021X Settings",
								InstanceID:    "Intel(r) AMT: 8021X Settings",
								Enabled:       3,
								AvailableInS0: false,
								PxeTimeout:    120,
							},
						},
					},
				},
			},
			// SET CERTIFICATES
			{
				"should create a valid ips_IEEE8021xSettings set certificates wsman message",
				"IPS_IEEE8021xSettings",
				wsmantesting.SetCertificates,
				fmt.Sprintf(`<h:SetCertificates_INPUT xmlns:h="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_IEEE8021xSettings"><h:ServerCertificateIssuer><a:Address>default</a:Address><a:ReferenceParameters><w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyCertificate</w:ResourceURI><w:SelectorSet><w:Selector Name="InstanceID">%s</w:Selector></w:SelectorSet></a:ReferenceParameters></h:ServerCertificateIssuer><h:ClientCertificate><a:Address>default</a:Address><a:ReferenceParameters><w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyCertificate</w:ResourceURI><w:SelectorSet><w:Selector Name="InstanceID">%s</w:Selector></w:SelectorSet></a:ReferenceParameters></h:ClientCertificate></h:SetCertificates_INPUT>`, wsmantesting.ServerCertificateIssuer, wsmantesting.ClientCertificate),
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.SetCertificates(wsmantesting.ServerCertificateIssuer, wsmantesting.ClientCertificate)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					SetCertificatesResponse: SetCertificates_OUTPUT{
						XMLName:     xml.Name{Space: fmt.Sprintf("%s%s", message.IPSSchema, IPSIEEE8021xSettings), Local: "SetCertificates_OUTPUT"},
						ReturnValue: 0,
					},
				},
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				expectedXMLInput := wsmantesting.ExpectedResponse(messageID, resourceURIBase, test.method, test.action, "", test.body)
				messageID++
				response, err := test.responseFunc()
				assert.Error(t, err)
				assert.Equal(t, expectedXMLInput, response.XMLInput)
				assert.NotEqual(t, test.expectedResponse, response.Body)
			})
		}
	})
}
