/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package wifi

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
	expectedResult := "{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"WiFiPortGetResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"LinkTechnology\":0,\"DeviceID\":\"\",\"CreationClassName\":\"\",\"SystemName\":\"\",\"SystemCreationClassName\":\"\",\"ElementName\":\"\",\"HealthState\":0,\"EnabledState\":0,\"RequestedState\":0,\"PortType\":0,\"PermanentAddress\":\"\"},\"EnumerateResponse\":{\"EnumerationContext\":\"\"},\"PullResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"EndpointSettingsItems\":null,\"WiFiPortItems\":null},\"RequestStateChange_OUTPUT\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"ReturnValue\":0,\"ReturnValueStr\":\"\"}}"
	result := response.JSON()
	assert.Equal(t, expectedResult, result)
}

func TestYaml(t *testing.T) {
	response := Response{
		Body: Body{
			PullResponse: PullResponse{},
		},
	}
	expectedResult := "xmlname:\n    space: \"\"\n    local: \"\"\nwifiportgetresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    linktechnology: 0\n    deviceid: \"\"\n    creationclassname: \"\"\n    systemname: \"\"\n    systemcreationclassname: \"\"\n    elementname: \"\"\n    healthstate: 0\n    enabledstate: 0\n    requestedstate: 0\n    porttype: 0\n    permanentaddress: \"\"\nenumerateresponse:\n    enumerationcontext: \"\"\npullresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    endpointsettingsitems: []\n    wifiportitems: []\nrequeststatechange_output:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    returnvalue: 0\n    returnvaluestr: \"\"\n"
	result := response.YAML()
	assert.Equal(t, expectedResult, result)
}

func TestPositiveCIMWifiEndpointSettings(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.CIMResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/wifi/endpointsettings",
	}
	elementUnderTest := NewWiFiEndpointSettingsWithClient(wsmanMessageCreator, &client)

	t.Run("cim_* Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			extraHeader      string
			body             string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			// ENUMERATES
			{
				"should create and parse a valid cim_WiFiEndpointSettings Enumerate call",
				CIMWiFiEndpointSettings,
				wsmantesting.Enumerate,
				"",
				wsmantesting.EnumerateBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageEnumerate

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "95040000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create and parse a valid cim_WiFiEndpointSettings Pull call",
				CIMWiFiEndpointSettings,
				wsmantesting.Pull,
				"",
				wsmantesting.PullBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessagePull

					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						XMLName: xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/09/enumeration", Local: "PullResponse"},
						EndpointSettingsItems: []WiFiEndpointSettingsResponse{
							{
								XMLName:              xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_WiFiEndpointSettings", Local: "CIM_WiFiEndpointSettings"},
								AuthenticationMethod: 2,
								BSSType:              3,
								ElementName:          "Test",
								EncryptionMethod:     2,
								InstanceID:           "Test",
								Priority:             1,
								SSID:                 "testSSID",
							},
						},
					},
				},
			},
			// DELETE
			{
				"should create and parse a valid cim_WiFiEndpointSettings Delete call",
				CIMWiFiEndpointSettings,
				wsmantesting.Delete,
				"<w:SelectorSet><w:Selector Name=\"InstanceID\">instanceID123</w:Selector></w:SelectorSet>",
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageDelete

					return elementUnderTest.Delete("instanceID123")
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					RequestStateChange_OUTPUT: common.ReturnValue{
						ReturnValue: 0,
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

func TestNegativeCIMWifiEndpointSettings(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.CIMResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/wifi/endpointsettings",
	}
	elementUnderTest := NewWiFiEndpointSettingsWithClient(wsmanMessageCreator, &client)

	t.Run("cim_* Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			extraHeader      string
			body             string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			// ENUMERATES
			{
				"should handle error when cim_WiFiEndpointSettings Enumerate call",
				CIMWiFiEndpointSettings,
				wsmantesting.Enumerate,
				"",
				wsmantesting.EnumerateBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "95040000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should handle error when cim_WiFiEndpointSettings Pull call",
				CIMWiFiEndpointSettings,
				wsmantesting.Pull,
				"",
				wsmantesting.PullBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						XMLName: xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/09/enumeration", Local: "PullResponse"},
						EndpointSettingsItems: []WiFiEndpointSettingsResponse{
							{
								XMLName:              xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_WiFiEndpointSettings", Local: "CIM_WiFiEndpointSettings"},
								AuthenticationMethod: 2,
								BSSType:              3,
								ElementName:          "Test",
								EncryptionMethod:     2,
								InstanceID:           "Test",
								Priority:             1,
								SSID:                 "testSSID",
							},
						},
					},
				},
			},
			// DELETE
			{
				"should handle error when cim_WiFiEndpointSettings Delete call",
				CIMWiFiEndpointSettings,
				wsmantesting.Delete,
				"<w:SelectorSet><w:Selector Name=\"InstanceID\">instanceID123</w:Selector></w:SelectorSet>",
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Delete("instanceID123")
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					RequestStateChange_OUTPUT: common.ReturnValue{
						ReturnValue: 0,
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
