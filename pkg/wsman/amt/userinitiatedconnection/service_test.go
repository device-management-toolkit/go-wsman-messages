/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package userinitiatedconnection

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
			GetResponse: UserResponse{},
		},
	}
	expectedResult := "{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"RequestStateChange_OUTPUT\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"ReturnValue\":0},\"GetResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"CreationClassName\":\"\",\"ElementName\":\"\",\"EnabledState\":0,\"Name\":\"\",\"SystemCreationClassName\":\"\",\"SystemName\":\"\"},\"EnumerateResponse\":{\"EnumerationContext\":\"\"},\"PullResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"UserItems\":null}}"
	result := response.JSON()
	assert.Equal(t, expectedResult, result)
}

func TestYaml(t *testing.T) {
	response := Response{
		Body: Body{
			GetResponse: UserResponse{},
		},
	}
	expectedResult := "xmlname:\n    space: \"\"\n    local: \"\"\nrequeststatechange_output:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    returnvalue: 0\ngetresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    creationclassname: \"\"\n    elementname: \"\"\n    enabledstate: 0\n    name: \"\"\n    systemcreationclassname: \"\"\n    systemname: \"\"\nenumerateresponse:\n    enumerationcontext: \"\"\npullresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    useritems: []\n"
	result := response.YAML()
	assert.Equal(t, expectedResult, result)
}

func TestPositiveAMT_UserInitiatedConnectionService(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.AMTResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/userinitiatedconnection",
	}
	elementUnderTest := NewUserInitiatedConnectionServiceWithClient(wsmanMessageCreator, &client)

	t.Run("amt_UserInitiatedConnectionService Tests", func(t *testing.T) {
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
				"should create a valid AMT_UserInitiatedConnectionService Get wsman message",
				AMTUserInitiatedConnectionService,
				wsmantesting.Get,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageGet

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetResponse: UserResponse{
						XMLName:                 xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMTUserInitiatedConnectionService), Local: AMTUserInitiatedConnectionService},
						CreationClassName:       AMTUserInitiatedConnectionService,
						ElementName:             "Intel(r) AMT User Initiated Connection Service",
						EnabledState:            32771,
						Name:                    "Intel(r) AMT User Initiated Connection Service",
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "Intel(r) AMT",
					},
				},
			},
			// ENUMERATES
			{
				"should create a valid AMT_UserInitiatedConnectionService Enumerate wsman message",
				AMTUserInitiatedConnectionService,
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageEnumerate

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "D3000000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create a valid AMT_UserInitiatedConnectionService Pull wsman message",
				AMTUserInitiatedConnectionService,
				wsmantesting.Pull,
				wsmantesting.PullBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessagePull

					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						XMLName: xml.Name{Space: message.XMLPullResponseSpace, Local: "PullResponse"},
						UserItems: []UserResponse{
							{
								XMLName:                 xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMTUserInitiatedConnectionService), Local: AMTUserInitiatedConnectionService},
								CreationClassName:       AMTUserInitiatedConnectionService,
								ElementName:             "Intel(r) AMT User Initiated Connection Service",
								EnabledState:            32771,
								Name:                    "Intel(r) AMT User Initiated Connection Service",
								SystemCreationClassName: "CIM_ComputerSystem",
								SystemName:              "Intel(r) AMT",
							},
						},
					},
				},
			},
			// REQUEST STATE CHANGE
			{
				"should create a valid AMT_UserInitiatedConnectionService RequestStateChange wsman message",
				AMTUserInitiatedConnectionService,
				fmt.Sprintf("%s%s/%s", message.AMTSchema, AMTUserInitiatedConnectionService, "RequestStateChange"),
				"<h:RequestStateChange_INPUT xmlns:h=\"http://intel.com/wbem/wscim/1/amt-schema/1/AMT_UserInitiatedConnectionService\"><h:RequestedState>32771</h:RequestedState></h:RequestStateChange_INPUT>",
				func() (Response, error) {
					client.CurrentMessage = "Request"

					return elementUnderTest.RequestStateChange(32771)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					RequestStateChange_OUTPUT: RequestStateChange_OUTPUT{
						XMLName:     xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMTUserInitiatedConnectionService), Local: "RequestStateChange_OUTPUT"},
						ReturnValue: 0,
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

func TestNegativeAMT_UserInitiatedConnectionService(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.AMTResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/userinitiatedconnection",
	}
	elementUnderTest := NewUserInitiatedConnectionServiceWithClient(wsmanMessageCreator, &client)

	t.Run("amt_UserInitiatedConnectionService Tests", func(t *testing.T) {
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
				"should create a valid AMT_UserInitiatedConnectionService Get wsman message",
				AMTUserInitiatedConnectionService,
				wsmantesting.Get,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetResponse: UserResponse{
						XMLName:                 xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMTUserInitiatedConnectionService), Local: AMTUserInitiatedConnectionService},
						CreationClassName:       AMTUserInitiatedConnectionService,
						ElementName:             "Intel(r) AMT User Initiated Connection Service",
						EnabledState:            32771,
						Name:                    "Intel(r) AMT User Initiated Connection Service",
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "Intel(r) AMT",
					},
				},
			},
			// ENUMERATES
			{
				"should create a valid AMT_UserInitiatedConnectionService Enumerate wsman message",
				AMTUserInitiatedConnectionService,
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "D3000000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create a valid AMT_UserInitiatedConnectionService Pull wsman message",
				AMTUserInitiatedConnectionService,
				wsmantesting.Pull,
				wsmantesting.PullBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						XMLName: xml.Name{Space: message.XMLPullResponseSpace, Local: "PullResponse"},
						UserItems: []UserResponse{
							{
								XMLName:                 xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMTUserInitiatedConnectionService), Local: AMTUserInitiatedConnectionService},
								CreationClassName:       AMTUserInitiatedConnectionService,
								ElementName:             "Intel(r) AMT User Initiated Connection Service",
								EnabledState:            32771,
								Name:                    "Intel(r) AMT User Initiated Connection Service",
								SystemCreationClassName: "CIM_ComputerSystem",
								SystemName:              "Intel(r) AMT",
							},
						},
					},
				},
			},
			// REQUEST STATE CHANGE
			{
				"should create a valid AMT_UserInitiatedConnectionService RequestStateChange wsman message",
				AMTUserInitiatedConnectionService,
				fmt.Sprintf("%s%s/%s", message.AMTSchema, AMTUserInitiatedConnectionService, "RequestStateChange"),
				"<h:RequestStateChange_INPUT xmlns:h=\"http://intel.com/wbem/wscim/1/amt-schema/1/AMT_UserInitiatedConnectionService\"><h:RequestedState>32771</h:RequestedState></h:RequestStateChange_INPUT>",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.RequestStateChange(32771)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					RequestStateChange_OUTPUT: RequestStateChange_OUTPUT{
						XMLName:     xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMTUserInitiatedConnectionService), Local: "RequestStateChange_OUTPUT"},
						ReturnValue: 0,
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
