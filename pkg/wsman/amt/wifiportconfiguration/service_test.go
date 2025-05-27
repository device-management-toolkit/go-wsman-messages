/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package wifiportconfiguration

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
			WiFiPortConfigurationService: WiFiPortConfigurationServiceResponse{},
		},
	}
	expectedResult := "{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"WiFiPortConfigurationService\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"RequestedState\":0,\"EnabledState\":0,\"HealthState\":0,\"ElementName\":\"\",\"SystemCreationClassName\":\"\",\"SystemName\":\"\",\"CreationClassName\":\"\",\"Name\":\"\",\"LocalProfileSynchronizationEnabled\":0,\"LastConnectedSsidUnderMeControl\":\"\",\"NoHostCsmeSoftwarePolicy\":0,\"UEFIWiFiProfileShareEnabled\":false},\"PullResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"WiFiPortConfigurationItems\":null},\"EnumerateResponse\":{\"EnumerationContext\":\"\"},\"AddWiFiSettingsOutput\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"ReturnValue\":0}}"
	result := response.JSON()
	assert.Equal(t, expectedResult, result)
}

func TestYaml(t *testing.T) {
	response := Response{
		Body: Body{
			WiFiPortConfigurationService: WiFiPortConfigurationServiceResponse{},
		},
	}
	expectedResult := "xmlname:\n    space: \"\"\n    local: \"\"\nwifiportconfigurationservice:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    requestedstate: 0\n    enabledstate: 0\n    healthstate: 0\n    elementname: \"\"\n    systemcreationclassname: \"\"\n    systemname: \"\"\n    creationclassname: \"\"\n    name: \"\"\n    localprofilesynchronizationenabled: 0\n    lastconnectedssidundermecontrol: \"\"\n    nohostcsmesoftwarepolicy: 0\n    uefiwifiprofileshareenabled: false\npullresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    wifiportconfigurationitems: []\nenumerateresponse:\n    enumerationcontext: \"\"\naddwifisettingsoutput:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    returnvalue: 0\n"
	result := response.YAML()
	assert.Equal(t, expectedResult, result)
}

func TestPositiveAMT_WiFiPortConfigurationService(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.AMTResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/wifiportconfiguration",
	}
	elementUnderTest := NewWiFiPortConfigurationServiceWithClient(wsmanMessageCreator, &client)

	t.Run("amt_WiFiPortConfigurationService Tests", func(t *testing.T) {
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
				"should create a valid AMT_WiFiPortConfigurationService Get wsman message",
				AMTWiFiPortConfigurationService,
				wsmantesting.Get,
				"",
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageGet

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					WiFiPortConfigurationService: WiFiPortConfigurationServiceResponse{
						XMLName:                            xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMTWiFiPortConfigurationService), Local: AMTWiFiPortConfigurationService},
						CreationClassName:                  "AMT_WiFiPortConfigurationService",
						ElementName:                        "Intel(r) AMT WiFiPort Configuration Service",
						EnabledState:                       5,
						HealthState:                        5,
						LastConnectedSsidUnderMeControl:    "",
						Name:                               "Intel(r) AMT WiFi Port Configuration Service",
						NoHostCsmeSoftwarePolicy:           0,
						RequestedState:                     12,
						SystemCreationClassName:            "CIM_ComputerSystem",
						SystemName:                         "Intel(r) AMT",
						LocalProfileSynchronizationEnabled: 1,
					},
				},
			},
			// ENUMERATES
			{
				"should create a valid AMT_WiFiPortConfigurationService Enumerate wsman message",
				AMTWiFiPortConfigurationService,
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
						EnumerationContext: "71080000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create a valid AMT_WiFiPortConfigurationService Pull wsman message",
				AMTWiFiPortConfigurationService,
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
						WiFiPortConfigurationItems: []WiFiPortConfigurationServiceResponse{
							{
								XMLName:                            xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMTWiFiPortConfigurationService), Local: AMTWiFiPortConfigurationService},
								CreationClassName:                  "AMT_WiFiPortConfigurationService",
								ElementName:                        "Intel(r) AMT WiFiPort Configuration Service",
								EnabledState:                       5,
								HealthState:                        5,
								LastConnectedSsidUnderMeControl:    "",
								Name:                               "Intel(r) AMT WiFi Port Configuration Service",
								NoHostCsmeSoftwarePolicy:           0,
								RequestedState:                     12,
								SystemCreationClassName:            "CIM_ComputerSystem",
								SystemName:                         "Intel(r) AMT",
								LocalProfileSynchronizationEnabled: 1,
							},
						},
					},
				},
			},
			// PUTS
			{
				"should create a valid AMT_WiFiPortConfigurationService Put wsman message",
				AMTWiFiPortConfigurationService,
				wsmantesting.Put,
				"<h:AMT_WiFiPortConfigurationService xmlns:h=\"http://intel.com/wbem/wscim/1/amt-schema/1/AMT_WiFiPortConfigurationService\"><h:RequestedState>12</h:RequestedState><h:EnabledState>5</h:EnabledState><h:HealthState>5</h:HealthState><h:ElementName>Intel(r) AMT WiFiPort Configuration Service</h:ElementName><h:SystemCreationClassName>CIM_ComputerSystem</h:SystemCreationClassName><h:SystemName>Intel(r) AMT</h:SystemName><h:CreationClassName>AMT_WiFiPortConfigurationService</h:CreationClassName><h:Name>Intel(r) AMT WiFi Port Configuration Service</h:Name><h:localProfileSynchronizationEnabled>1</h:localProfileSynchronizationEnabled></h:AMT_WiFiPortConfigurationService>",
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessagePut
					wifiConfiguration := WiFiPortConfigurationServiceRequest{
						RequestedState:                     12,
						EnabledState:                       5,
						HealthState:                        5,
						ElementName:                        "Intel(r) AMT WiFiPort Configuration Service",
						SystemCreationClassName:            "CIM_ComputerSystem",
						SystemName:                         "Intel(r) AMT",
						CreationClassName:                  "AMT_WiFiPortConfigurationService",
						Name:                               "Intel(r) AMT WiFi Port Configuration Service",
						LocalProfileSynchronizationEnabled: 1,
						LastConnectedSsidUnderMeControl:    "",
						NoHostCsmeSoftwarePolicy:           0,
					}

					return elementUnderTest.Put(wifiConfiguration)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					WiFiPortConfigurationService: WiFiPortConfigurationServiceResponse{
						XMLName:                            xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMTWiFiPortConfigurationService), Local: AMTWiFiPortConfigurationService},
						CreationClassName:                  "AMT_WiFiPortConfigurationService",
						ElementName:                        "Intel(r) AMT WiFiPort Configuration Service",
						EnabledState:                       5,
						HealthState:                        5,
						LastConnectedSsidUnderMeControl:    "",
						Name:                               "Intel(r) AMT WiFi Port Configuration Service",
						NoHostCsmeSoftwarePolicy:           0,
						RequestedState:                     12,
						SystemCreationClassName:            "CIM_ComputerSystem",
						SystemName:                         "Intel(r) AMT",
						LocalProfileSynchronizationEnabled: 1,
					},
				},
			},
			// WIFI PORT CONFIGURATION SERVICE
			// {
			// 	"should return a valid amt_WiFiPortConfigurationService ADD_WIFI_SETTINGS wsman message",
			// 	AMT_WiFiPortConfigurationService,
			// 	`http://intel.com/wbem/wscim/1/amt-schema/1/AMT_WiFiPortConfigurationService/AddWiFiSettings`, `<h:AddWiFiSettings_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_WiFiPortConfigurationService"><h:WiFiEndpoint><a:Address>/wsman</a:Address><a:ReferenceParameters><w:ResourceURI>http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_WiFiEndpoint</w:ResourceURI><w:SelectorSet><w:Selector Name="Name">WiFi Endpoint 0</w:Selector></w:SelectorSet></a:ReferenceParameters></h:WiFiEndpoint><h:WiFiEndpointSettingsInput xmlns:q="http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_WiFiEndpointSettings"><q:ElementName>home</q:ElementName><q:InstanceID>Intel(r) AMT:WiFi Endpoint Settings home</q:InstanceID><q:AuthenticationMethod>6</q:AuthenticationMethod><q:EncryptionMethod>4</q:EncryptionMethod><q:SSID>admin</q:SSID><q:Priority>1</q:Priority><q:PSKPassPhrase>p&#39;ass&lt;&gt;&amp;&#34;code</q:PSKPassPhrase></h:WiFiEndpointSettingsInput></h:AddWiFiSettings_INPUT>`,
			// 	"",
			// 	func() (Response, error) {
			// 		client.CurrentMessage = "AddWiFiSettings"
			// 		wifiEndpointSettings := wifi.WiFiEndpointSettings_INPUT{}
			// 		ieee8021xSettings := &models.IEEE8021xSettings{}
			// 		wifiEndpoint := "t"
			// 		clientCredential := "t"
			// 		caCredential := "t"
			// 		return elementUnderTest.AddWiFiSettings(wifiEndpointSettings, ieee8021xSettings, wifiEndpoint, clientCredential, caCredential)
			// 	},
			// 	Body{},
			// },
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
