/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package general

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
			GetResponse: GeneralSettingsResponse{},
		},
	}
	expectedResult := "{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"GetResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"ElementName\":\"\",\"InstanceID\":\"\",\"NetworkInterfaceEnabled\":false,\"DigestRealm\":\"\",\"IdleWakeTimeout\":0,\"HostName\":\"\",\"DomainName\":\"\",\"PingResponseEnabled\":false,\"WsmanOnlyMode\":false,\"PreferredAddressFamily\":0,\"DHCPv6ConfigurationTimeout\":0,\"DDNSUpdateEnabled\":false,\"DDNSUpdateByDHCPServerEnabled\":false,\"SharedFQDN\":false,\"HostOSFQDN\":\"\",\"DDNSTTL\":0,\"AMTNetworkEnabled\":0,\"RmcpPingResponseEnabled\":false,\"DDNSPeriodicUpdateInterval\":0,\"PresenceNotificationInterval\":0,\"PrivacyLevel\":0,\"PowerSource\":0,\"ThunderboltDockEnabled\":0,\"OemID\":0,\"DHCPSyncRequiresHostname\":0},\"EnumerateResponse\":{\"EnumerationContext\":\"\"},\"PullResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"GeneralSettingsItems\":null},\"PutResponse\":{}}"
	result := response.JSON()
	assert.Equal(t, expectedResult, result)
}

func TestYaml(t *testing.T) {
	response := Response{
		Body: Body{
			GetResponse: GeneralSettingsResponse{},
		},
	}
	expectedResult := "xmlname:\n    space: \"\"\n    local: \"\"\ngetresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    elementname: \"\"\n    instanceid: \"\"\n    networkinterfaceenabled: false\n    digestrealm: \"\"\n    idlewaketimeout: 0\n    hostname: \"\"\n    domainname: \"\"\n    pingresponseenabled: false\n    wsmanonlymode: false\n    preferredaddressfamily: 0\n    dhcpv6configurationtimeout: 0\n    ddnsupdateenabled: false\n    ddnsupdatebydhcpserverenabled: false\n    sharedfqdn: false\n    hostosfqdn: \"\"\n    ddnsttl: 0\n    amtnetworkenabled: 0\n    rmcppingresponseenabled: false\n    ddnsperiodicupdateinterval: 0\n    presencenotificationinterval: 0\n    privacylevel: 0\n    powersource: 0\n    thunderboltdockenabled: 0\n    oemid: 0\n    dhcpsyncrequireshostname: 0\nenumerateresponse:\n    enumerationcontext: \"\"\npullresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    generalsettingsitems: []\nputresponse: {}\n"
	result := response.YAML()
	assert.Equal(t, expectedResult, result)
}

func TestPositiveAMT_GeneralSettings(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.AMTResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)

	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/general",
	}

	elementUnderTest := NewGeneralSettingsWithClient(wsmanMessageCreator, &client)

	t.Run("amt_* Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			body             string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			// GETS
			{
				"should create a valid AMT_GeneralSettings Get wsman message",
				AMTGeneralSettings,
				wsmantesting.Get,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageGet

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetResponse: GeneralSettingsResponse{
						XMLName:                       xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_GeneralSettings", Local: AMTGeneralSettings},
						ElementName:                   "Intel(r) AMT: General Settings",
						InstanceID:                    "Intel(r) AMT: General Settings",
						AMTNetworkEnabled:             1,
						DDNSPeriodicUpdateInterval:    1440,
						DDNSTTL:                       900,
						DDNSUpdateByDHCPServerEnabled: true,
						DDNSUpdateEnabled:             false,
						DHCPSyncRequiresHostname:      1, // Intel SDK documentation missing
						DHCPv6ConfigurationTimeout:    0,
						DigestRealm:                   "Digest:F3EB554784E729164447A89F60B641C5",
						DomainName:                    "Test Domain Name",
						HostName:                      "Test Host Name",
						HostOSFQDN:                    "Test Host OS FQDN",
						IdleWakeTimeout:               1,
						NetworkInterfaceEnabled:       true,
						PingResponseEnabled:           true,
						PowerSource:                   0,
						PreferredAddressFamily:        0,
						PresenceNotificationInterval:  0,
						PrivacyLevel:                  0,
						RmcpPingResponseEnabled:       true,
						SharedFQDN:                    true,
						ThunderboltDockEnabled:        0,
						WsmanOnlyMode:                 false,
					},
				},
			},
			// ENUMERATES
			{
				"should create a valid AMT_GeneralSettings Enumerate wsman message",
				AMTGeneralSettings,
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageEnumerate

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "14000000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create a valid AMT_GeneralSettings Pull wsman message",
				AMTGeneralSettings,
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
						GeneralSettingsItems: []GeneralSettingsResponse{
							{
								XMLName:                       xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_GeneralSettings", Local: AMTGeneralSettings},
								ElementName:                   "Intel(r) AMT: General Settings",
								InstanceID:                    "Intel(r) AMT: General Settings",
								AMTNetworkEnabled:             1,
								DDNSPeriodicUpdateInterval:    1440,
								DDNSTTL:                       900,
								DDNSUpdateByDHCPServerEnabled: true,
								DDNSUpdateEnabled:             false,
								DHCPv6ConfigurationTimeout:    0,
								DigestRealm:                   "Digest:6EE8C61BA74893E059F032EA919D699E",
								DomainName:                    "",
								HostOSFQDN:                    "",
								IdleWakeTimeout:               65535,
								NetworkInterfaceEnabled:       true,
								PingResponseEnabled:           true,
								PowerSource:                   0,
								PreferredAddressFamily:        0,
								PresenceNotificationInterval:  0,
								PrivacyLevel:                  0,
								RmcpPingResponseEnabled:       true,
								SharedFQDN:                    true,
								WsmanOnlyMode:                 false,
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

func TestNegativeAMT_GeneralSettings(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.AMTResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)

	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/general",
	}
	elementUnderTest := NewGeneralSettingsWithClient(wsmanMessageCreator, &client)

	t.Run("amt_* Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			body             string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			// GETS
			{
				"should create a valid AMT_GeneralSettings Get wsman message",
				AMTGeneralSettings,
				wsmantesting.Get,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetResponse: GeneralSettingsResponse{
						XMLName:                       xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_GeneralSettings", Local: AMTGeneralSettings},
						ElementName:                   "Intel(r) AMT: General Settings",
						InstanceID:                    "Intel(r) AMT: General Settings",
						AMTNetworkEnabled:             1,
						DDNSPeriodicUpdateInterval:    1440,
						DDNSTTL:                       900,
						DDNSUpdateByDHCPServerEnabled: true,
						DDNSUpdateEnabled:             false,
						DHCPSyncRequiresHostname:      1, // Intel SDK documentation missing
						DHCPv6ConfigurationTimeout:    0,
						DigestRealm:                   "Digest:F3EB554784E729164447A89F60B641C5",
						DomainName:                    "Test Domain Name",
						HostName:                      "Test Host Name",
						HostOSFQDN:                    "Test Host OS FQDN",
						IdleWakeTimeout:               1,
						NetworkInterfaceEnabled:       true,
						PingResponseEnabled:           true,
						PowerSource:                   0,
						PreferredAddressFamily:        0,
						PresenceNotificationInterval:  0,
						PrivacyLevel:                  0,
						RmcpPingResponseEnabled:       true,
						SharedFQDN:                    true,
						ThunderboltDockEnabled:        0,
						WsmanOnlyMode:                 false,
					},
				},
			},
			// ENUMERATES
			{
				"should create a valid AMT_GeneralSettings Enumerate wsman message",
				AMTGeneralSettings,
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "14000000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create a valid AMT_GeneralSettings Pull wsman message",
				AMTGeneralSettings,
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
						GeneralSettingsItems: []GeneralSettingsResponse{
							{
								XMLName:                       xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_GeneralSettings", Local: AMTGeneralSettings},
								ElementName:                   "Intel(r) AMT: General Settings",
								InstanceID:                    "Intel(r) AMT: General Settings",
								AMTNetworkEnabled:             1,
								DDNSPeriodicUpdateInterval:    1440,
								DDNSTTL:                       900,
								DDNSUpdateByDHCPServerEnabled: true,
								DDNSUpdateEnabled:             false,
								DHCPv6ConfigurationTimeout:    0,
								DigestRealm:                   "Digest:6EE8C61BA74893E059F032EA919D699E",
								DomainName:                    "",
								HostOSFQDN:                    "",
								IdleWakeTimeout:               65535,
								NetworkInterfaceEnabled:       true,
								PingResponseEnabled:           true,
								PowerSource:                   0,
								PreferredAddressFamily:        0,
								PresenceNotificationInterval:  0,
								PrivacyLevel:                  0,
								RmcpPingResponseEnabled:       true,
								SharedFQDN:                    true,
								WsmanOnlyMode:                 false,
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
