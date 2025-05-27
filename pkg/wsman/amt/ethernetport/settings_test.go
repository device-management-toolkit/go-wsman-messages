/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package ethernetport

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
			GetAndPutResponse: SettingsResponse{},
		},
	}
	expectedResult := "{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"GetAndPutResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"ElementName\":\"\",\"InstanceID\":\"\",\"VLANTag\":0,\"SharedMAC\":false,\"MACAddress\":\"\",\"LinkIsUp\":false,\"LinkPolicy\":null,\"LinkPreference\":0,\"LinkControl\":0,\"SharedStaticIp\":false,\"SharedDynamicIP\":false,\"IpSyncEnabled\":false,\"DHCPEnabled\":false,\"IPAddress\":\"\",\"SubnetMask\":\"\",\"DefaultGateway\":\"\",\"PrimaryDNS\":\"\",\"SecondaryDNS\":\"\",\"ConsoleTcpMaxRetransmissions\":0,\"WLANLinkProtectionLevel\":0,\"PhysicalConnectionType\":0,\"PhysicalNicMedium\":0},\"EnumerateResponse\":{\"EnumerationContext\":\"\"},\"PullResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"EthernetPortItems\":null}}"
	result := response.JSON()
	assert.Equal(t, expectedResult, result)
}

func TestYaml(t *testing.T) {
	response := Response{
		Body: Body{
			GetAndPutResponse: SettingsResponse{},
		},
	}
	expectedResult := "xmlname:\n    space: \"\"\n    local: \"\"\ngetandputresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    elementname: \"\"\n    instanceid: \"\"\n    vlantag: 0\n    sharedmac: false\n    macaddress: \"\"\n    linkisup: false\n    linkpolicy: []\n    linkpreference: 0\n    linkcontrol: 0\n    sharedstaticip: false\n    shareddynamicip: false\n    ipsyncenabled: false\n    dhcpenabled: false\n    ipaddress: \"\"\n    subnetmask: \"\"\n    defaultgateway: \"\"\n    primarydns: \"\"\n    secondarydns: \"\"\n    consoletcpmaxretransmissions: 0\n    wlanlinkprotectionlevel: 0\n    physicalconnectiontype: 0\n    physicalnicmedium: 0\nenumerateresponse:\n    enumerationcontext: \"\"\npullresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    ethernetportitems: []\n"
	result := response.YAML()
	assert.Equal(t, expectedResult, result)
}

func TestPositiveAMT_EthernetPortSettings(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.AMTResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/ethernetport",
	}
	elementUnderTest := NewEthernetPortSettingsWithClient(wsmanMessageCreator, &client)

	t.Run("amt_EthernetPortSettings Tests", func(t *testing.T) {
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
				"should create a valid AMT_EthernetPortSettings Get wsman message",
				AMTEthernetPortSettings,
				wsmantesting.Get,
				"",
				"<w:SelectorSet><w:Selector Name=\"InstanceID\">Intel(r) AMT Ethernet Port Settings 0</w:Selector></w:SelectorSet>",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageGet

					return elementUnderTest.Get("Intel(r) AMT Ethernet Port Settings 0")
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetAndPutResponse: SettingsResponse{
						XMLName:                xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_EthernetPortSettings", Local: "AMT_EthernetPortSettings"},
						ElementName:            "Intel(r) AMT Ethernet Port Settings",
						InstanceID:             "Intel(r) AMT Ethernet Port Settings 0",
						SharedMAC:              true,
						MACAddress:             "c8-d9-d2-7a-1e-33",
						LinkIsUp:               true,
						LinkPolicy:             []LinkPolicy{1, 14},
						SharedStaticIp:         false,
						SharedDynamicIP:        true,
						IpSyncEnabled:          true,
						DHCPEnabled:            true,
						SubnetMask:             "255.255.255.0",
						DefaultGateway:         "192.168.0.1",
						PrimaryDNS:             "68.105.28.11",
						SecondaryDNS:           "68.105.29.11",
						PhysicalConnectionType: 0,
					},
				},
			},

			// ENUMERATES
			{
				"should create a valid AMT_EthernetPortSettings Enumerate wsman message",
				AMTEthernetPortSettings,
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
						EnumerationContext: "7700000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create a valid AMT_EthernetPortSettings Pull wsman message",
				AMTEthernetPortSettings,
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
						XMLName: xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/09/enumeration", Local: "PullResponse"},
						EthernetPortItems: []SettingsResponse{
							{
								XMLName:                xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_EthernetPortSettings", Local: "AMT_EthernetPortSettings"},
								ElementName:            "Intel(r) AMT Ethernet Port Settings",
								InstanceID:             "Intel(r) AMT Ethernet Port Settings 0",
								VLANTag:                0,
								SharedMAC:              true,
								MACAddress:             "00-be-43-d8-22-a4",
								LinkIsUp:               true,
								LinkPolicy:             []LinkPolicy{1, 14},
								SharedStaticIp:         false,
								SharedDynamicIP:        true,
								IpSyncEnabled:          true,
								DHCPEnabled:            true,
								SubnetMask:             "255.255.255.0",
								DefaultGateway:         "192.168.6.1",
								PrimaryDNS:             "192.168.6.1",
								PhysicalConnectionType: 0,
							},
							{
								XMLName:                      xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_EthernetPortSettings", Local: "AMT_EthernetPortSettings"},
								ElementName:                  "Intel(r) AMT Ethernet Port Settings",
								InstanceID:                   "Intel(r) AMT Ethernet Port Settings 1",
								SharedMAC:                    true,
								MACAddress:                   "00-00-00-00-00-00",
								LinkIsUp:                     false,
								LinkPreference:               2,
								LinkControl:                  2,
								DHCPEnabled:                  true,
								ConsoleTcpMaxRetransmissions: 5,
								WLANLinkProtectionLevel:      1,
								PhysicalConnectionType:       3,
							},
						},
					},
				},
			},
			// PUTS
			{
				"should create a valid AMT_EthernetPortSettings Put wsman message",
				AMTEthernetPortSettings,
				wsmantesting.Put,
				"<h:AMT_EthernetPortSettings xmlns:h=\"http://intel.com/wbem/wscim/1/amt-schema/1/AMT_EthernetPortSettings\"><h:ElementName>Intel(r) AMT Ethernet Port Settings</h:ElementName><h:InstanceID>Intel(r) AMT Ethernet Port Settings 0</h:InstanceID><h:SharedMAC>true</h:SharedMAC><h:LinkIsUp>false</h:LinkIsUp><h:SharedStaticIp>true</h:SharedStaticIp><h:IpSyncEnabled>true</h:IpSyncEnabled><h:DHCPEnabled>true</h:DHCPEnabled></h:AMT_EthernetPortSettings>",
				"<w:SelectorSet><w:Selector Name=\"InstanceID\">Intel(r) AMT Ethernet Port Settings 0</w:Selector></w:SelectorSet>",
				func() (Response, error) {
					ethernetPortSettings := SettingsRequest{
						H:              "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_EthernetPortSettings",
						DHCPEnabled:    true,
						ElementName:    "Intel(r) AMT Ethernet Port Settings",
						InstanceID:     "Intel(r) AMT Ethernet Port Settings 0",
						IpSyncEnabled:  true,
						SharedMAC:      true,
						SharedStaticIp: true,
					}
					client.CurrentMessage = wsmantesting.CurrentMessagePut

					return elementUnderTest.Put(ethernetPortSettings.InstanceID, ethernetPortSettings)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetAndPutResponse: SettingsResponse{
						XMLName:                xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_EthernetPortSettings", Local: "AMT_EthernetPortSettings"},
						DHCPEnabled:            true,
						DefaultGateway:         "192.168.0.1",
						ElementName:            "Intel(r) AMT Ethernet Port Settings",
						IPAddress:              "192.168.0.24",
						InstanceID:             "Intel(r) AMT Ethernet Port Settings 0",
						IpSyncEnabled:          true,
						LinkIsUp:               true,
						LinkPolicy:             []LinkPolicy{1, 14, 16},
						MACAddress:             "a4-ae-11-1e-46-53",
						PhysicalConnectionType: 0,
						PrimaryDNS:             "68.105.28.11",
						SecondaryDNS:           "68.105.29.11",
						SharedDynamicIP:        true,
						SharedMAC:              true,
						SharedStaticIp:         true,
						SubnetMask:             "255.255.255.0",
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

func TestNegativeAMT_EthernetPortSettings(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.AMTResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/ethernetport",
	}
	elementUnderTest := NewEthernetPortSettingsWithClient(wsmanMessageCreator, &client)

	t.Run("amt_EthernetPortSettings Tests", func(t *testing.T) {
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
				"should create a valid AMT_EthernetPortSettings Get wsman message",
				AMTEthernetPortSettings,
				wsmantesting.Get,
				"",
				"<w:SelectorSet><w:Selector Name=\"InstanceID\">Intel(r) AMT Ethernet Port Settings 0</w:Selector></w:SelectorSet>",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Get("Intel(r) AMT Ethernet Port Settings 0")
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetAndPutResponse: SettingsResponse{
						XMLName:                xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_EthernetPortSettings", Local: "AMT_EthernetPortSettings"},
						ElementName:            "Intel(r) AMT Ethernet Port Settings",
						InstanceID:             "Intel(r) AMT Ethernet Port Settings 0",
						SharedMAC:              true,
						MACAddress:             "c8-d9-d2-7a-1e-33",
						LinkIsUp:               true,
						LinkPolicy:             []LinkPolicy{1, 14},
						SharedStaticIp:         false,
						SharedDynamicIP:        true,
						IpSyncEnabled:          true,
						DHCPEnabled:            true,
						SubnetMask:             "255.255.255.0",
						DefaultGateway:         "192.168.0.1",
						PrimaryDNS:             "68.105.28.11",
						SecondaryDNS:           "68.105.29.11",
						PhysicalConnectionType: 0,
					},
				},
			},

			// ENUMERATES
			{
				"should create a valid AMT_EthernetPortSettings Enumerate wsman message",
				AMTEthernetPortSettings,
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
						EnumerationContext: "7700000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create a valid AMT_EthernetPortSettings Pull wsman message",
				AMTEthernetPortSettings,
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
						XMLName: xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/09/enumeration", Local: "PullResponse"},
						EthernetPortItems: []SettingsResponse{
							{
								XMLName:                xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_EthernetPortSettings", Local: "AMT_EthernetPortSettings"},
								ElementName:            "Intel(r) AMT Ethernet Port Settings",
								InstanceID:             "Intel(r) AMT Ethernet Port Settings 0",
								VLANTag:                0,
								SharedMAC:              true,
								MACAddress:             "00-be-43-d8-22-a4",
								LinkIsUp:               true,
								LinkPolicy:             []LinkPolicy{1, 14},
								SharedStaticIp:         false,
								SharedDynamicIP:        true,
								IpSyncEnabled:          true,
								DHCPEnabled:            true,
								SubnetMask:             "255.255.255.0",
								DefaultGateway:         "192.168.6.1",
								PrimaryDNS:             "192.168.6.1",
								PhysicalConnectionType: 0,
							},
							{
								XMLName:                      xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_EthernetPortSettings", Local: "AMT_EthernetPortSettings"},
								ElementName:                  "Intel(r) AMT Ethernet Port Settings",
								InstanceID:                   "Intel(r) AMT Ethernet Port Settings 1",
								SharedMAC:                    true,
								MACAddress:                   "00-00-00-00-00-00",
								LinkIsUp:                     false,
								LinkPreference:               2,
								LinkControl:                  2,
								DHCPEnabled:                  true,
								ConsoleTcpMaxRetransmissions: 5,
								WLANLinkProtectionLevel:      1,
								PhysicalConnectionType:       3,
							},
						},
					},
				},
			},
			// PUTS
			{
				"should create a valid AMT_EthernetPortSettings Put wsman message",
				AMTEthernetPortSettings,
				wsmantesting.Put,
				"<h:AMT_EthernetPortSettings xmlns:h=\"http://intel.com/wbem/wscim/1/amt-schema/1/AMT_EthernetPortSettings\"><h:ElementName>Intel(r) AMT Ethernet Port Settings</h:ElementName><h:InstanceID>Intel(r) AMT Ethernet Port Settings 0</h:InstanceID><h:SharedMAC>true</h:SharedMAC><h:LinkIsUp>false</h:LinkIsUp><h:SharedStaticIp>true</h:SharedStaticIp><h:IpSyncEnabled>true</h:IpSyncEnabled><h:DHCPEnabled>true</h:DHCPEnabled></h:AMT_EthernetPortSettings>",
				"<w:SelectorSet><w:Selector Name=\"InstanceID\">Intel(r) AMT Ethernet Port Settings 0</w:Selector></w:SelectorSet>",
				func() (Response, error) {
					ethernetPortSettings := SettingsRequest{
						H:              "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_EthernetPortSettings",
						DHCPEnabled:    true,
						ElementName:    "Intel(r) AMT Ethernet Port Settings",
						InstanceID:     "Intel(r) AMT Ethernet Port Settings 0",
						IpSyncEnabled:  true,
						SharedMAC:      true,
						SharedStaticIp: true,
					}
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Put(ethernetPortSettings.InstanceID, ethernetPortSettings)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetAndPutResponse: SettingsResponse{
						XMLName:                xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_EthernetPortSettings", Local: "AMT_EthernetPortSettings"},
						DHCPEnabled:            true,
						DefaultGateway:         "192.168.0.1",
						ElementName:            "Intel(r) AMT Ethernet Port Settings",
						IPAddress:              "192.168.0.24",
						InstanceID:             "Intel(r) AMT Ethernet Port Settings 0",
						IpSyncEnabled:          true,
						LinkIsUp:               true,
						LinkPolicy:             []LinkPolicy{1, 14, 16},
						MACAddress:             "a4-ae-11-1e-46-53",
						PhysicalConnectionType: 0,
						PrimaryDNS:             "68.105.28.11",
						SecondaryDNS:           "68.105.29.11",
						SharedDynamicIP:        true,
						SharedMAC:              true,
						SharedStaticIp:         true,
						SubnetMask:             "255.255.255.0",
					},
				},
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				expectedXMLInput := wsmantesting.ExpectedResponse(messageID, resourceURIBase, test.method, test.action, test.extraHeader, test.body)
				messageID++
				response, err := test.responseFunc()
				assert.Error(t, err)
				assert.Equal(t, expectedXMLInput, response.XMLInput)
				assert.NotEqual(t, test.expectedResponse, response.Body)
			})
		}
	})
}
