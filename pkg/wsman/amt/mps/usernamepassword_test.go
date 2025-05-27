/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package mps

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
			GetResponse: MPSUsernamePasswordResponse{},
		},
	}
	expectedResult := "{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"GetResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"InstanceID\":\"\",\"RemoteID\":\"\",\"Secret\":\"\",\"Algorithm\":\"\",\"Protocol\":\"\"},\"EnumerateResponse\":{\"EnumerationContext\":\"\"},\"PullResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"MPSUsernamePasswordItems\":null}}"
	result := response.JSON()
	assert.Equal(t, expectedResult, result)
}

func TestYaml(t *testing.T) {
	response := Response{
		Body: Body{
			GetResponse: MPSUsernamePasswordResponse{},
		},
	}
	expectedResult := "xmlname:\n    space: \"\"\n    local: \"\"\ngetresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    instanceid: \"\"\n    remoteid: \"\"\n    secret: \"\"\n    algorithm: \"\"\n    protocol: \"\"\nenumerateresponse:\n    enumerationcontext: \"\"\npullresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    mpsusernamepassworditems: []\n"
	result := response.YAML()
	assert.Equal(t, expectedResult, result)
}

func TestPositiveAMT_MPSUsernamePassword(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.AMTResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/mps",
	}
	elementUnderTest := NewMPSUsernamePasswordWithClient(wsmanMessageCreator, &client)

	t.Run("amt_MPSUsernamePassword Tests", func(t *testing.T) {
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
				"should create a valid AMT_MPSUsernamePassword Get wsman message",
				AMTMPSUsernamePassword,
				wsmantesting.Get,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageGet

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetResponse: MPSUsernamePasswordResponse{
						XMLName:    xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_MPSUsernamePassword", Local: "AMT_MPSUsernamePassword"},
						InstanceID: "Intel(r) AMT:MPS Username Password 0",
						RemoteID:   "test",
					},
				},
			},
			// ENUMERATES
			{
				"should create a valid AMT_MPSUsernamePassword Enumerate wsman message",
				AMTMPSUsernamePassword,
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageEnumerate

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "19080000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create a valid AMT_MPSUsernamePassword Pull wsman message",
				AMTMPSUsernamePassword,
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
						MPSUsernamePasswordItems: []MPSUsernamePasswordResponse{
							{
								XMLName:    xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_MPSUsernamePassword", Local: "AMT_MPSUsernamePassword"},
								InstanceID: "Intel(r) AMT:MPS Username Password 0",
								RemoteID:   "test",
							},
						},
					},
				},
			},
			{
				"should create a valid AMT_MPSUsernamePassword Put wsman message",
				AMTMPSUsernamePassword,
				wsmantesting.Put,
				"<h:AMT_MPSUsernamePassword xmlns:h=\"http://intel.com/wbem/wscim/1/amt-schema/1/AMT_MPSUsernamePassword\"><h:InstanceID>Intel(r) AMT:MPS Username Password 0</h:InstanceID><h:RemoteID>test</h:RemoteID><h:Secret>P@ssw0rd</h:Secret></h:AMT_MPSUsernamePassword>",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessagePut
					mpsUsernamePassword := MPSUsernamePasswordRequest{
						H:          "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_MPSUsernamePassword",
						InstanceID: "Intel(r) AMT:MPS Username Password 0",
						Secret:     "P@ssw0rd",
						RemoteID:   "test",
					}

					return elementUnderTest.Put(mpsUsernamePassword)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
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

func TestNegativeAMT_MPSUsernamePassword(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.AMTResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/mps",
	}
	elementUnderTest := NewMPSUsernamePasswordWithClient(wsmanMessageCreator, &client)

	t.Run("amt_MPSUsernamePassword Tests", func(t *testing.T) {
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
				"should create a valid AMT_MPSUsernamePassword Get wsman message",
				AMTMPSUsernamePassword,
				wsmantesting.Get,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetResponse: MPSUsernamePasswordResponse{
						XMLName:    xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_MPSUsernamePassword", Local: "AMT_MPSUsernamePassword"},
						InstanceID: "Intel(r) AMT:MPS Username Password 0",
						RemoteID:   "test",
					},
				},
			},
			// ENUMERATES
			{
				"should create a valid AMT_MPSUsernamePassword Enumerate wsman message",
				AMTMPSUsernamePassword,
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "19080000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create a valid AMT_MPSUsernamePassword Pull wsman message",
				AMTMPSUsernamePassword,
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
						MPSUsernamePasswordItems: []MPSUsernamePasswordResponse{
							{
								XMLName:    xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_MPSUsernamePassword", Local: "AMT_MPSUsernamePassword"},
								InstanceID: "Intel(r) AMT:MPS Username Password 0",
								RemoteID:   "test",
							},
						},
					},
				},
			},
			{
				"should create a valid AMT_MPSUsernamePassword Put wsman message",
				AMTMPSUsernamePassword,
				wsmantesting.Put,
				"<h:AMT_MPSUsernamePassword xmlns:h=\"http://intel.com/wbem/wscim/1/amt-schema/1/AMT_MPSUsernamePassword\"><h:InstanceID>Intel(r) AMT:MPS Username Password 0</h:InstanceID><h:RemoteID>test</h:RemoteID><h:Secret>P@ssw0rd</h:Secret></h:AMT_MPSUsernamePassword>",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError
					mpsUsernamePassword := MPSUsernamePasswordRequest{
						H:          "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_MPSUsernamePassword",
						InstanceID: "Intel(r) AMT:MPS Username Password 0",
						Secret:     "P@ssw0rd",
						RemoteID:   "test",
					}

					return elementUnderTest.Put(mpsUsernamePassword)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
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
