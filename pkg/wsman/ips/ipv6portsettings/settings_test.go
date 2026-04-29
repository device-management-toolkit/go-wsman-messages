/*********************************************************************
 * Copyright (c) Intel Corporation 2026
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package ipv6portsettings

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
	expectedResult := "{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"GetResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"CurrentDefaultRouter\":\"\",\"CurrentPrimaryDNS\":\"\",\"CurrentSecondaryDNS\":\"\",\"DefaultRouter\":\"\",\"ElementName\":\"\",\"IPv6Address\":\"\",\"InstanceID\":\"\",\"InterfaceIDType\":0,\"ManualInterfaceID\":\"\",\"PrimaryDNS\":\"\",\"SecondaryDNS\":\"\"},\"PullResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"IPv6PortSettingsItems\":null},\"EnumerateResponse\":{\"EnumerationContext\":\"\"}}"
	result := response.JSON()
	assert.Equal(t, expectedResult, result)
}

func TestYaml(t *testing.T) {
	response := Response{
		Body: Body{
			PullResponse: PullResponse{},
		},
	}
	expectedResult := "xmlname:\n    space: \"\"\n    local: \"\"\ngetresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    currentdefaultrouter: \"\"\n    currentprimarydns: \"\"\n    currentsecondarydns: \"\"\n    defaultrouter: \"\"\n    elementname: \"\"\n    ipv6address: \"\"\n    instanceid: \"\"\n    interfaceidtype: 0\n    manualinterfaceid: \"\"\n    primarydns: \"\"\n    secondarydns: \"\"\npullresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    ipv6portsettingsitems: []\nenumerateresponse:\n    enumerationcontext: \"\"\n"
	result := response.YAML()
	assert.Equal(t, expectedResult, result)
}

func TestPositiveIPS_IPv6PortSettings(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.IPSResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{PackageUnderTest: "ips/ipv6portsettings"}
	elementUnderTest := NewIPv6PortSettingsWithClient(wsmanMessageCreator, &client)

	t.Run("ips_IPv6PortSettings Tests", func(t *testing.T) {
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
				"should create a valid IPS_IPv6PortSettings Get wsman message",
				IPSIPv6PortSettings,
				wsmantesting.Get,
				"",
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageGet

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetResponse: IPv6PortSettings{
						XMLName:              xml.Name{Space: fmt.Sprintf("%s%s", message.IPSSchema, IPSIPv6PortSettings), Local: IPSIPv6PortSettings},
						CurrentDefaultRouter: "::",
						CurrentPrimaryDNS:    "::",
						CurrentSecondaryDNS:  "::",
						DefaultRouter:        "::",
						ElementName:          "Intel(r) IPS IPv6 Settings 0",
						IPv6Address:          "::",
						InstanceID:           "Intel(r) IPS IPv6 Settings 0",
						InterfaceIDType:      0,
						ManualInterfaceID:    "0000000000000000",
						PrimaryDNS:           "::",
						SecondaryDNS:         "::",
					},
				},
			},
			{
				"should create a valid IPS_IPv6PortSettings Enumerate wsman message",
				IPSIPv6PortSettings,
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageEnumerate

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName:           xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{EnumerationContext: "9C000000-0000-0000-0000-000000000000"},
				},
			},
			{
				"should create a valid IPS_IPv6PortSettings Pull wsman message",
				IPSIPv6PortSettings,
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
						IPv6PortSettingsItems: []IPv6PortSettings{{
							XMLName:              xml.Name{Space: fmt.Sprintf("%s%s", message.IPSSchema, IPSIPv6PortSettings), Local: IPSIPv6PortSettings},
							CurrentDefaultRouter: "::",
							CurrentPrimaryDNS:    "::",
							CurrentSecondaryDNS:  "::",
							DefaultRouter:        "::",
							ElementName:          "Intel(r) IPS IPv6 Settings 0",
							IPv6Address:          "::",
							InstanceID:           "Intel(r) IPS IPv6 Settings 0",
							InterfaceIDType:      0,
							ManualInterfaceID:    "0000000000000000",
							PrimaryDNS:           "::",
							SecondaryDNS:         "::",
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
