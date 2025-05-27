/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package ieee8021x

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/wsmantesting"
)

func TestJson(t *testing.T) {
	response := Response{
		Body: Body{
			PullResponse: PullResponse{},
		},
	}
	expectedResult := "{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"EnumerateResponse\":{\"EnumerationContext\":\"\"},\"PullResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"IEEE8021xSettingsItems\":null},\"GetResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"ElementName\":\"\",\"InstanceID\":\"\",\"AuthenticationProtocol\":0,\"RoamingIdentity\":\"\",\"ServerCertificateName\":\"\",\"ServerCertificateNameComparison\":0,\"Username\":\"\",\"Password\":\"\",\"Domain\":\"\",\"ProtectedAccessCredential\":\"\",\"PACPassword\":\"\",\"PSK\":\"\"}}"
	result := response.JSON()
	assert.Equal(t, expectedResult, result)
}

func TestYaml(t *testing.T) {
	response := Response{
		Body: Body{
			PullResponse: PullResponse{},
		},
	}
	expectedResult := "xmlname:\n    space: \"\"\n    local: \"\"\nenumerateresponse:\n    enumerationcontext: \"\"\npullresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    ieee8021xsettingsitems: []\ngetresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    elementname: \"\"\n    instanceid: \"\"\n    authenticationprotocol: 0\n    roamingidentity: \"\"\n    servercertificatename: \"\"\n    servercertificatenamecomparison: 0\n    username: \"\"\n    password: \"\"\n    domain: \"\"\n    protectedaccesscredential: \"\"\n    pacpassword: \"\"\n    psk: \"\"\n"
	result := response.YAML()
	assert.Equal(t, expectedResult, result)
}

func TestPositiveCIMIEEE8021xSettings(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.CIMResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/ieee8021x/settings",
	}
	elementUnderTest := NewIEEE8021xSettingsWithClient(wsmanMessageCreator, &client)

	t.Run("cim_* Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			body             string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			// ENUMERATES
			{
				"should create and parse a valid cim_IEEE8021xSettings Enumerate call",
				CIMIEEE8021xSettings,
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageEnumerate

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "1A000000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create and parse a valid cim_IEEE8021xSettings Pull call",
				CIMIEEE8021xSettings,
				wsmantesting.Pull,
				wsmantesting.PullBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessagePull

					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						XMLName: xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/09/enumeration", Local: "PullResponse"},
						IEEE8021xSettingsItems: []IEEE8021xSettingsResponse{
							{
								XMLName:                         xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_IEEE8021xSettings", Local: "CIM_IEEE8021xSettings"},
								ElementName:                     "Wifi8021xTLS",
								InstanceID:                      "Intel(r) AMT:IEEE 802.1x Settings Wifi8021xTLS",
								AuthenticationProtocol:          0,
								RoamingIdentity:                 "",
								ServerCertificateName:           "",
								ServerCertificateNameComparison: 0,
								Username:                        "iME$amt16$",
								Password:                        "",
								Domain:                          "",
								ProtectedAccessCredential:       "",
								PACPassword:                     "",
								PSK:                             "",
							},
						},
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

func TestNegativeCIMIEEE8021xSettings(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.CIMResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/ieee8021x/settings",
	}
	elementUnderTest := NewIEEE8021xSettingsWithClient(wsmanMessageCreator, &client)

	t.Run("cim_* Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			body             string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			// ENUMERATES
			{
				"should handle error when cim_IEEE8021xSettings Enumerate wsman message",
				CIMIEEE8021xSettings,
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "1A000000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should handle error when cim_IEEE8021xSettings Pull wsman message",
				CIMIEEE8021xSettings,
				wsmantesting.Pull,
				wsmantesting.PullBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						XMLName: xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/09/enumeration", Local: "PullResponse"},
						IEEE8021xSettingsItems: []IEEE8021xSettingsResponse{
							{
								XMLName:                         xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_IEEE8021xSettings", Local: "CIM_IEEE8021xSettings"},
								ElementName:                     "Wifi8021xTLS",
								InstanceID:                      "Intel(r) AMT:IEEE 802.1x Settings Wifi8021xTLS",
								AuthenticationProtocol:          0,
								RoamingIdentity:                 "",
								ServerCertificateName:           "",
								ServerCertificateNameComparison: 0,
								Username:                        "iME$amt16$",
								Password:                        "",
								Domain:                          "",
								ProtectedAccessCredential:       "",
								PACPassword:                     "",
								PSK:                             "",
							},
						},
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
