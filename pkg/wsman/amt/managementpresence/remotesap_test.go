/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package managementpresence

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
			GetResponse: ManagementRemoteResponse{},
		},
	}
	expectedResult := "{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"GetResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"AccessInfo\":\"\",\"CN\":\"\",\"CreationClassName\":\"\",\"ElementName\":\"\",\"InfoFormat\":0,\"Name\":\"\",\"Port\":0,\"SystemCreationClassName\":\"\",\"SystemName\":\"\"},\"EnumerateResponse\":{\"EnumerationContext\":\"\"},\"PullResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"ManagementRemoteItems\":null}}"
	result := response.JSON()
	assert.Equal(t, expectedResult, result)
}

func TestYaml(t *testing.T) {
	response := Response{
		Body: Body{
			GetResponse: ManagementRemoteResponse{},
		},
	}
	expectedResult := "xmlname:\n    space: \"\"\n    local: \"\"\ngetresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    accessinfo: \"\"\n    cn: \"\"\n    creationclassname: \"\"\n    elementname: \"\"\n    infoformat: 0\n    name: \"\"\n    port: 0\n    systemcreationclassname: \"\"\n    systemname: \"\"\nenumerateresponse:\n    enumerationcontext: \"\"\npullresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    managementremoteitems: []\n"
	result := response.YAML()
	assert.Equal(t, expectedResult, result)
}

func TestPositiveAMT_ManagementPresenceRemoteSAP(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.AMTResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/managementpresence",
	}
	elementUnderTest := NewManagementPresenceRemoteSAPWithClient(wsmanMessageCreator, &client)

	t.Run("amt_ManagementPresenceRemoteSAP Tests", func(t *testing.T) {
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
				"should create a valid AMT_ManagementPresenceRemoteSAP Get wsman message",
				AMTManagementPresenceRemoteSAP,
				wsmantesting.Get,
				"",
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageGet

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetResponse: ManagementRemoteResponse{
						XMLName:                 xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_ManagementPresenceRemoteSAP", Local: "AMT_ManagementPresenceRemoteSAP"},
						AccessInfo:              "192.168.0.208",
						CN:                      "192.168.0.208",
						CreationClassName:       "AMT_ManagementPresenceRemoteSAP",
						ElementName:             "Intel(r) AMT:Management Presence Server",
						InfoFormat:              3,
						Name:                    "Intel(r) AMT:Management Presence Server 0",
						Port:                    4433,
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "Intel(r) AMT",
					},
				},
			},
			// ENUMERATES
			{
				"should create a valid AMT_ManagementPresenceRemoteSAP Enumerate wsman message",
				AMTManagementPresenceRemoteSAP,
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
						EnumerationContext: "C9000000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create a valid AMT_ManagementPresenceRemoteSAP Pull wsman message",
				AMTManagementPresenceRemoteSAP,
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
						ManagementRemoteItems: []ManagementRemoteResponse{
							{
								XMLName:                 xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_ManagementPresenceRemoteSAP", Local: "AMT_ManagementPresenceRemoteSAP"},
								AccessInfo:              "192.168.10.196",
								CN:                      "192.168.10.196",
								CreationClassName:       "AMT_ManagementPresenceRemoteSAP",
								ElementName:             "Intel(r) AMT:Management Presence Server",
								InfoFormat:              3,
								Name:                    "Intel(r) AMT:Management Presence Server 0",
								Port:                    4433,
								SystemCreationClassName: "CIM_ComputerSystem",
								SystemName:              "Intel(r) AMT",
							},
						},
					},
				},
			},
			// DELETE
			{
				"should create a valid AMT_ManagementPresenceRemoteSAP Delete wsman message",
				AMTManagementPresenceRemoteSAP,
				wsmantesting.Delete,
				"",
				"<w:SelectorSet><w:Selector Name=\"Name\">instanceID123</w:Selector></w:SelectorSet>",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageDelete

					return elementUnderTest.Delete("instanceID123")
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
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

func TestNegativeAMT_ManagementPresenceRemoteSAP(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.AMTResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/managementpresence",
	}
	elementUnderTest := NewManagementPresenceRemoteSAPWithClient(wsmanMessageCreator, &client)

	t.Run("amt_ManagementPresenceRemoteSAP Tests", func(t *testing.T) {
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
				"should create a valid AMT_ManagementPresenceRemoteSAP Get wsman message",
				AMTManagementPresenceRemoteSAP,
				wsmantesting.Get,
				"",
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetResponse: ManagementRemoteResponse{
						XMLName:                 xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_ManagementPresenceRemoteSAP", Local: "AMT_ManagementPresenceRemoteSAP"},
						AccessInfo:              "192.168.0.208",
						CN:                      "192.168.0.208",
						CreationClassName:       "AMT_ManagementPresenceRemoteSAP",
						ElementName:             "Intel(r) AMT:Management Presence Server",
						InfoFormat:              3,
						Name:                    "Intel(r) AMT:Management Presence Server 0",
						Port:                    4433,
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "Intel(r) AMT",
					},
				},
			},
			// ENUMERATES
			{
				"should create a valid AMT_ManagementPresenceRemoteSAP Enumerate wsman message",
				AMTManagementPresenceRemoteSAP,
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
						EnumerationContext: "C9000000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create a valid AMT_ManagementPresenceRemoteSAP Pull wsman message",
				AMTManagementPresenceRemoteSAP,
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
						ManagementRemoteItems: []ManagementRemoteResponse{
							{
								XMLName:                 xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_ManagementPresenceRemoteSAP", Local: "AMT_ManagementPresenceRemoteSAP"},
								AccessInfo:              "192.168.10.196",
								CN:                      "192.168.10.196",
								CreationClassName:       "AMT_ManagementPresenceRemoteSAP",
								ElementName:             "Intel(r) AMT:Management Presence Server",
								InfoFormat:              3,
								Name:                    "Intel(r) AMT:Management Presence Server 0",
								Port:                    4433,
								SystemCreationClassName: "CIM_ComputerSystem",
								SystemName:              "Intel(r) AMT",
							},
						},
					},
				},
			},
			// DELETE
			{
				"should create a valid AMT_ManagementPresenceRemoteSAP Delete wsman message",
				AMTManagementPresenceRemoteSAP,
				wsmantesting.Delete,
				"",
				"<w:SelectorSet><w:Selector Name=\"Name\">instanceID123</w:Selector></w:SelectorSet>",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Delete("instanceID123")
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
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
