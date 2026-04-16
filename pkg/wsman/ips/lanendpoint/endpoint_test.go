/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package lanendpoint

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
	expectedResult := "{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"GetResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"CreationClassName\":\"\",\"EnabledState\":0,\"LANType\":0,\"MACAddress\":\"\",\"Name\":\"\",\"ProtocolType\":0,\"RequestedState\":0,\"SystemCreationClassName\":\"\",\"SystemName\":\"\"},\"PullResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"LANEndpointItems\":null},\"EnumerateResponse\":{\"EnumerationContext\":\"\"}}"
	result := response.JSON()
	assert.Equal(t, expectedResult, result)
}

func TestYaml(t *testing.T) {
	response := Response{
		Body: Body{
			PullResponse: PullResponse{},
		},
	}
	expectedResult := "xmlname:\n    space: \"\"\n    local: \"\"\ngetresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    creationclassname: \"\"\n    enabledstate: 0\n    lantype: 0\n    macaddress: \"\"\n    name: \"\"\n    protocoltype: 0\n    requestedstate: 0\n    systemcreationclassname: \"\"\n    systemname: \"\"\npullresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    lanendpointitems: []\nenumerateresponse:\n    enumerationcontext: \"\"\n"
	result := response.YAML()
	assert.Equal(t, expectedResult, result)
}

func TestPositiveIPS_LANEndpoint(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.IPSResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{PackageUnderTest: "ips/lanendpoint"}
	elementUnderTest := NewLANEndpointWithClient(wsmanMessageCreator, &client)

	t.Run("ips_LANEndpoint Tests", func(t *testing.T) {
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
				"should create a valid IPS_LANEndpoint Get wsman message",
				IPSLANEndpoint,
				wsmantesting.Get,
				"",
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageGet

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetResponse: LANEndpoint{
						XMLName:                 xml.Name{Space: fmt.Sprintf("%s%s", message.IPSSchema, IPSLANEndpoint), Local: IPSLANEndpoint},
						CreationClassName:       "IPS_LANEndpoint",
						EnabledState:            2,
						LANType:                 2,
						MACAddress:              "48210b50d8c9",
						Name:                    "Intel(r) AMT LAN Endpoint",
						ProtocolType:            0,
						RequestedState:          12,
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "Intel(r) AMT",
					},
				},
			},
			{
				"should create a valid IPS_LANEndpoint Enumerate wsman message",
				IPSLANEndpoint,
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageEnumerate

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName:           xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{EnumerationContext: "9D000000-0000-0000-0000-000000000000"},
				},
			},
			{
				"should create a valid IPS_LANEndpoint Pull wsman message",
				IPSLANEndpoint,
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
						LANEndpointItems: []LANEndpoint{{
							XMLName:                 xml.Name{Space: fmt.Sprintf("%s%s", message.IPSSchema, IPSLANEndpoint), Local: IPSLANEndpoint},
							CreationClassName:       "IPS_LANEndpoint",
							EnabledState:            2,
							LANType:                 2,
							MACAddress:              "48210b50d8c9",
							Name:                    "Intel(r) AMT LAN Endpoint",
							ProtocolType:            0,
							RequestedState:          12,
							SystemCreationClassName: "CIM_ComputerSystem",
							SystemName:              "Intel(r) AMT",
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
