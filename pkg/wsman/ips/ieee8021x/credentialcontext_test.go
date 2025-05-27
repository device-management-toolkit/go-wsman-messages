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
	expectedResult := "{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"PullResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"IEEE8021xSettingsItems\":null,\"CredentialContextItems\":null},\"EnumerateResponse\":{\"EnumerationContext\":\"\"},\"IEEE8021xSettingsResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"ElementName\":\"\",\"InstanceID\":\"\",\"Enabled\":0,\"AvailableInS0\":false,\"PxeTimeout\":0},\"SetCertificatesResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"ReturnValue\":0}}"
	result := response.JSON()
	assert.Equal(t, expectedResult, result)
}

func TestYaml(t *testing.T) {
	response := Response{
		Body: Body{
			PullResponse: PullResponse{},
		},
	}
	expectedResult := "xmlname:\n    space: \"\"\n    local: \"\"\npullresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    ieee8021xsettingsitems: []\n    credentialcontextitems: []\nenumerateresponse:\n    enumerationcontext: \"\"\nieee8021xsettingsresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    elementname: \"\"\n    instanceid: \"\"\n    enabled: 0\n    availableins0: false\n    pxetimeout: 0\nsetcertificatesresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    returnvalue: 0\n"
	result := response.YAML()
	assert.Equal(t, expectedResult, result)
}

func TestPositiveIPS_8021xCredentialContext(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.IPSResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "ips/ieee8021x/credentialcontext",
	}
	elementUnderTest := NewIEEE8021xCredentialContextWithClient(wsmanMessageCreator, &client)

	t.Run("ips_8021xCredentialContext Tests", func(t *testing.T) {
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
			// {
			// 	"should create a valid IPS_8021xCredentialContext Get wsman message",
			// 	"IPS_8021xCredentialContext",
			// 	wsmantesting.GET,
			// 	"",
			// 	"",
			// 	func() (Response, error) {
			// 		 client.CurrentMessage = wsmantesting.CurrentMessageGet

			// 		return elementUnderTest.Get()
			// 	},
			// 	Body{},
			// },
			// ENUMERATES
			{
				"should create a valid IPS_8021xCredentialContext Enumerate wsman message",
				"IPS_8021xCredentialContext",
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
						EnumerationContext: "9A0A0000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			// {
			// 	"should create a valid IPS_8021xCredentialContext Pull wsman message",
			// 	"IPS_8021xCredentialContext",
			// 	wsmantesting.PULL,
			// 	wsmantesting.PULL_BODY,
			// 	"",
			// 	func() (Response, error) {
			// 		 client.CurrentMessage = wsmantesting.CurrentMessagePull

			// 		return elementUnderTest.Pull(wsmantesting.EnumerationContext)
			// 	},
			// 	Body{},
			// },
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
