/*********************************************************************
 * Copyright (c) Intel Corporation 2026
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package hostipsettings

import (
	"encoding/xml"
	"fmt"
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
	expectedResult := "{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"GetResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"DHCPEnabled\":false,\"ElementName\":\"\",\"InstanceID\":\"\"},\"PullResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"HostIPSettingsItems\":null},\"EnumerateResponse\":{\"EnumerationContext\":\"\"}}"
	result := response.JSON()
	assert.Equal(t, expectedResult, result)
}

func TestYaml(t *testing.T) {
	response := Response{
		Body: Body{
			PullResponse: PullResponse{},
		},
	}
	expectedResult := "xmlname:\n    space: \"\"\n    local: \"\"\ngetresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    dhcpenabled: false\n    elementname: \"\"\n    instanceid: \"\"\npullresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    hostipsettingsitems: []\nenumerateresponse:\n    enumerationcontext: \"\"\n"
	result := response.YAML()
	assert.Equal(t, expectedResult, result)
}

func TestPositiveIPS_HostIPSettings(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.IPSResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{PackageUnderTest: "ips/hostipsettings"}
	elementUnderTest := NewHostIPSettingsWithClient(wsmanMessageCreator, &client)

	t.Run("ips_HostIPSettings Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			body             string
			extraHeader      string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			{
				"should create a valid IPS_HostIPSettings Get wsman message",
				IPSHostIPSettings,
				wsmantesting.Get,
				"",
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageGet

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetResponse: HostIPSettings{
						XMLName:     xml.Name{Space: fmt.Sprintf("%s%s", message.IPSSchema, IPSHostIPSettings), Local: IPSHostIPSettings},
						DHCPEnabled: false,
						ElementName: "Intel(r) AMT: Host IP Settings",
						InstanceID:  "Intel(r) AMT: Host IP Settings",
					},
				},
			},
			{
				"should create a valid IPS_HostIPSettings Enumerate wsman message",
				IPSHostIPSettings,
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageEnumerate

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName:           xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{EnumerationContext: "9A000000-0000-0000-0000-000000000000"},
				},
			},
			{
				"should create a valid IPS_HostIPSettings Pull wsman message",
				IPSHostIPSettings,
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
						HostIPSettingsItems: []HostIPSettings{{
							XMLName:     xml.Name{Space: fmt.Sprintf("%s%s", message.IPSSchema, IPSHostIPSettings), Local: IPSHostIPSettings},
							DHCPEnabled: false,
							ElementName: "Intel(r) AMT: Host IP Settings",
							InstanceID:  "Intel(r) AMT: Host IP Settings",
						}},
					},
				},
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				expectedXMLInput := wsmantesting.ExpectedResponse(messageID, resourceURIBase, test.method, test.action, test.extraHeader, test.body)
				messageID++
				response, err := test.responseFunc()
				assert.NoError(t, err)
				assert.Equal(t, expectedXMLInput, response.XMLInput)
				assert.Equal(t, test.expectedResponse, response.Body)
			})
		}
	})
}
