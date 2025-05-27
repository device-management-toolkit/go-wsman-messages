/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package tls

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/wsmantesting"
)

const (
	EnvelopeResponse = `<a:Envelope xmlns:a="http://www.w3.org/2003/05/soap-envelope" x-mlns:b="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:c="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns:d="http://schemas.xmlsoap.org/ws/2005/02/trust" xmlns:e="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" xmlns:f="http://schemas.dmtf.org/wbem/wsman/1/cimbinding.xsd" xmlns:g="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TLSSettingData" xmlns:h="http://schemas.dmtf.org/wbem/wscim/1/common" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><a:Header><b:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:To><b:RelatesTo>0</b:RelatesTo><b:Action a:mustUnderstand="true">`
	GetBody          = `<g:AMT_TLSSettingData><g:CreationClassName>AMT_TLSSettingData</g:CreationClassName><g:ElementName>Intel(r) TLS Setting Data</g:ElementName><g:Name>Intel(r) AMT TLS Setting Data</g:Name><g:SystemCreationClassName>CIM_ComputerSystem</g:SystemCreationClassName><g:SystemName>ManagedSystem</g:SystemName></g:AMT_TLSSettingData>`
)

func TestPositiveAMT_TLSSettingData(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.AMTResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/tls/settingdata",
	}
	elementUnderTest := NewTLSSettingDataWithClient(wsmanMessageCreator, &client)

	t.Run("amt_TLSSettingData Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			extraHeader      string
			body             string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			// GETS
			{
				"should create a valid AMT_TLSSettingData Get wsman message",
				AMTTLSSettingData,
				wsmantesting.Get,
				"<w:SelectorSet><w:Selector Name=\"InstanceID\">Intel(r) AMT 802.3 TLS Settings</w:Selector></w:SelectorSet>",
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageGet

					return elementUnderTest.Get("Intel(r) AMT 802.3 TLS Settings")
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					SettingDataGetAndPutResponse: SettingDataResponse{
						XMLName:                    xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMTTLSSettingData), Local: AMTTLSSettingData},
						AcceptNonSecureConnections: false,
						ElementName:                "Intel(r) AMT 802.3 TLS Settings",
						Enabled:                    false,
						InstanceID:                 "Intel(r) AMT 802.3 TLS Settings",
						MutualAuthentication:       false,
					},
				},
			},

			// ENUMERATES
			{
				"should create a valid AMT_TLSSettingData Enumerate wsman message",
				AMTTLSSettingData,
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
						EnumerationContext: "CA000000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create a valid AMT_TLSSettingData Pull wsman message",
				AMTTLSSettingData,
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
						XMLName: xml.Name{Space: message.XMLPullResponseSpace, Local: "PullResponse"},
						SettingDataItems: []SettingDataResponse{
							{
								XMLName:                    xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMTTLSSettingData), Local: AMTTLSSettingData},
								AcceptNonSecureConnections: false,
								ElementName:                "Intel(r) AMT 802.3 TLS Settings",
								Enabled:                    false,
								InstanceID:                 "Intel(r) AMT 802.3 TLS Settings",
								MutualAuthentication:       false,
							},
							{
								XMLName:                    xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMTTLSSettingData), Local: AMTTLSSettingData},
								AcceptNonSecureConnections: true,
								ElementName:                "Intel(r) AMT LMS TLS Settings",
								Enabled:                    false,
								InstanceID:                 "Intel(r) AMT LMS TLS Settings",
								MutualAuthentication:       false,
							},
						},
					},
				},
			},

			// PUTS
			{
				"should create a valid AMT_TLSSettingData Put wsman message",
				AMTTLSSettingData,
				wsmantesting.Put,
				"<w:SelectorSet><w:Selector Name=\"InstanceID\">Intel(r) AMT 802.3 TLS Settings</w:Selector></w:SelectorSet>",
				"<h:AMT_TLSSettingData xmlns:h=\"http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TLSSettingData\"><h:ElementName>Intel(r) AMT 802.3 TLS Settings</h:ElementName><h:InstanceID>Intel(r) AMT 802.3 TLS Settings</h:InstanceID><h:MutualAuthentication>false</h:MutualAuthentication><h:Enabled>true</h:Enabled><h:AcceptNonSecureConnections>false</h:AcceptNonSecureConnections><h:NonSecureConnectionsSupported>false</h:NonSecureConnectionsSupported></h:AMT_TLSSettingData>",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessagePut
					tlsSettingData := SettingDataRequest{
						ElementName: "Intel(r) AMT 802.3 TLS Settings",
						InstanceID:  "Intel(r) AMT 802.3 TLS Settings",
						Enabled:     true,
					}

					return elementUnderTest.Put("Intel(r) AMT 802.3 TLS Settings", tlsSettingData)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					SettingDataGetAndPutResponse: SettingDataResponse{
						XMLName:                    xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMTTLSSettingData), Local: AMTTLSSettingData},
						AcceptNonSecureConnections: false,
						ElementName:                "Intel(r) AMT 802.3 TLS Settings",
						Enabled:                    false,
						InstanceID:                 "Intel(r) AMT 802.3 TLS Settings",
						MutualAuthentication:       false,
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

func TestNegativeAMT_TLSSettingData(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.AMTResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/tls/settingdata",
	}
	elementUnderTest := NewTLSSettingDataWithClient(wsmanMessageCreator, &client)

	t.Run("amt_TLSSettingData Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			extraHeader      string
			body             string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			// GETS
			{
				"should create a valid AMT_TLSSettingData Get wsman message",
				AMTTLSSettingData,
				wsmantesting.Get,
				"<w:SelectorSet><w:Selector Name=\"InstanceID\">Intel(r) AMT 802.3 TLS Settings</w:Selector></w:SelectorSet>",
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Get("Intel(r) AMT 802.3 TLS Settings")
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					SettingDataGetAndPutResponse: SettingDataResponse{
						XMLName:                    xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMTTLSSettingData), Local: AMTTLSSettingData},
						AcceptNonSecureConnections: false,
						ElementName:                "Intel(r) AMT 802.3 TLS Settings",
						Enabled:                    false,
						InstanceID:                 "Intel(r) AMT 802.3 TLS Settings",
						MutualAuthentication:       false,
					},
				},
			},

			// ENUMERATES
			{
				"should create a valid AMT_TLSSettingData Enumerate wsman message",
				AMTTLSSettingData,
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
						EnumerationContext: "CA000000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create a valid AMT_TLSSettingData Pull wsman message",
				AMTTLSSettingData,
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
						XMLName: xml.Name{Space: message.XMLPullResponseSpace, Local: "PullResponse"},
						SettingDataItems: []SettingDataResponse{
							{
								XMLName:                    xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMTTLSSettingData), Local: AMTTLSSettingData},
								AcceptNonSecureConnections: false,
								ElementName:                "Intel(r) AMT 802.3 TLS Settings",
								Enabled:                    false,
								InstanceID:                 "Intel(r) AMT 802.3 TLS Settings",
								MutualAuthentication:       false,
							},
							{
								XMLName:                    xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMTTLSSettingData), Local: AMTTLSSettingData},
								AcceptNonSecureConnections: true,
								ElementName:                "Intel(r) AMT LMS TLS Settings",
								Enabled:                    false,
								InstanceID:                 "Intel(r) AMT LMS TLS Settings",
								MutualAuthentication:       false,
							},
						},
					},
				},
			},

			// PUTS
			{
				"should create a valid AMT_TLSSettingData Put wsman message",
				AMTTLSSettingData,
				wsmantesting.Put,
				"<w:SelectorSet><w:Selector Name=\"InstanceID\">Intel(r) AMT 802.3 TLS Settings</w:Selector></w:SelectorSet>",
				"<h:AMT_TLSSettingData xmlns:h=\"http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TLSSettingData\"><h:ElementName>Intel(r) AMT 802.3 TLS Settings</h:ElementName><h:InstanceID>Intel(r) AMT 802.3 TLS Settings</h:InstanceID><h:MutualAuthentication>false</h:MutualAuthentication><h:Enabled>true</h:Enabled><h:AcceptNonSecureConnections>false</h:AcceptNonSecureConnections><h:NonSecureConnectionsSupported>false</h:NonSecureConnectionsSupported></h:AMT_TLSSettingData>",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError
					tlsSettingData := SettingDataRequest{
						ElementName: "Intel(r) AMT 802.3 TLS Settings",
						InstanceID:  "Intel(r) AMT 802.3 TLS Settings",
						Enabled:     true,
					}

					return elementUnderTest.Put("Intel(r) AMT 802.3 TLS Settings", tlsSettingData)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					SettingDataGetAndPutResponse: SettingDataResponse{
						XMLName:                    xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMTTLSSettingData), Local: AMTTLSSettingData},
						AcceptNonSecureConnections: false,
						ElementName:                "Intel(r) AMT 802.3 TLS Settings",
						Enabled:                    false,
						InstanceID:                 "Intel(r) AMT 802.3 TLS Settings",
						MutualAuthentication:       false,
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
