/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package redirection

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/wsmantesting"
)

func TestJson(t *testing.T) {
	response := Response{
		Body: Body{
			GetAndPutResponse: RedirectionResponse{},
		},
	}
	expectedResult := "{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"GetAndPutResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"CreationClassName\":\"\",\"ElementName\":\"\",\"EnabledState\":0,\"ListenerEnabled\":false,\"Name\":\"\",\"SystemCreationClassName\":\"\",\"SystemName\":\"\",\"AccessLog\":null},\"RequestStateChange_OUTPUT\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"ReturnValue\":0},\"EnumerateResponse\":{\"EnumerationContext\":\"\"},\"PullResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"RedirectionItems\":null}}"
	result := response.JSON()
	assert.Equal(t, expectedResult, result)
}

func TestYaml(t *testing.T) {
	response := Response{
		Body: Body{
			GetAndPutResponse: RedirectionResponse{},
		},
	}
	expectedResult := "xmlname:\n    space: \"\"\n    local: \"\"\ngetandputresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    creationclassname: \"\"\n    elementname: \"\"\n    enabledstate: 0\n    listenerenabled: false\n    name: \"\"\n    systemcreationclassname: \"\"\n    systemname: \"\"\n    accesslog: []\nrequeststatechange_output:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    returnvalue: 0\nenumerateresponse:\n    enumerationcontext: \"\"\npullresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    redirectionitems: []\n"
	result := response.YAML()
	assert.Equal(t, expectedResult, result)
}

func TestPositiveAMT_RedirectionService(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.AMTResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/redirectionservice",
	}
	elementUnderTest := NewRedirectionServiceWithClient(wsmanMessageCreator, &client)

	t.Run("amt_RedirectionService Tests", func(t *testing.T) {
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
				"should create a valid AMT_RedirectionService Get wsman message",
				AMTRedirectionService,
				wsmantesting.Get,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageGet

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetAndPutResponse: RedirectionResponse{
						XMLName:                 xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMTRedirectionService), Local: AMTRedirectionService},
						CreationClassName:       AMTRedirectionService,
						ElementName:             "Intel(r) AMT Redirection Service",
						EnabledState:            32771,
						ListenerEnabled:         true,
						Name:                    "Intel(r) AMT Redirection Service",
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "Intel(r) AMT",
					},
				},
			},
			// ENUMERATES
			{
				"should create a valid AMT_RedirectionService Enumerate wsman message",
				AMTRedirectionService,
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageEnumerate
					if elementUnderTest.Base.WSManMessageCreator == nil {
						logrus.Println("Error")
					}

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
				"should create a valid AMT_RedirectionService Pull wsman message",
				AMTRedirectionService,
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
						RedirectionItems: []RedirectionResponse{
							{
								XMLName:                 xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMTRedirectionService), Local: AMTRedirectionService},
								CreationClassName:       AMTRedirectionService,
								ElementName:             "Intel(r) AMT Redirection Service",
								EnabledState:            32771,
								ListenerEnabled:         true,
								Name:                    "Intel(r) AMT Redirection Service",
								SystemCreationClassName: "CIM_ComputerSystem",
								SystemName:              "Intel(r) AMT",
							},
						},
					},
				},
			},
			// PUTS
			{
				"should create a valid AMT_RedirectionService Put wsman message",
				AMTRedirectionService,
				wsmantesting.Put,
				"<h:AMT_RedirectionService xmlns:h=\"http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RedirectionService\"><h:CreationClassName>AMT_RedirectionService</h:CreationClassName><h:ElementName>Intel(r) AMT Redirection Service</h:ElementName><h:EnabledState>32771</h:EnabledState><h:ListenerEnabled>true</h:ListenerEnabled><h:Name>Intel(r) AMT Redirection Service</h:Name><h:SystemCreationClassName>CIM_ComputerSystem</h:SystemCreationClassName><h:SystemName>Intel(r) AMT</h:SystemName></h:AMT_RedirectionService>",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessagePut
					redirectionRequest := &RedirectionRequest{
						CreationClassName:       AMTRedirectionService,
						ElementName:             "Intel(r) AMT Redirection Service",
						EnabledState:            32771,
						ListenerEnabled:         true,
						Name:                    "Intel(r) AMT Redirection Service",
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "Intel(r) AMT",
					}

					return elementUnderTest.Put(redirectionRequest)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetAndPutResponse: RedirectionResponse{
						XMLName:                 xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMTRedirectionService), Local: AMTRedirectionService},
						CreationClassName:       AMTRedirectionService,
						ElementName:             "Intel(r) AMT Redirection Service",
						EnabledState:            32771,
						ListenerEnabled:         true,
						Name:                    "Intel(r) AMT Redirection Service",
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "Intel(r) AMT",
					},
				},
			},
			// REQUEST STATE CHANGE
			{
				"should create a valid AMT_RedirectionService Request State Change wsman message",
				AMTRedirectionService,
				fmt.Sprintf("%s%s/%s", message.AMTSchema, AMTRedirectionService, "RequestStateChange"),
				"<h:RequestStateChange_INPUT xmlns:h=\"http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RedirectionService\"><h:RequestedState>32771</h:RequestedState></h:RequestStateChange_INPUT>",
				func() (Response, error) {
					client.CurrentMessage = "RequestStateChange"

					return elementUnderTest.RequestStateChange(EnableIDERAndSOL)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					RequestStateChange_OUTPUT: RequestStateChange_OUTPUT{
						XMLName:     xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMTRedirectionService), Local: "RequestStateChange_OUTPUT"},
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

func TestNegativeAMT_RedirectionService(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.AMTResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/redirectionservice",
	}
	elementUnderTest := NewRedirectionServiceWithClient(wsmanMessageCreator, &client)

	t.Run("amt_RedirectionService Tests", func(t *testing.T) {
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
				"should create a valid AMT_RedirectionService Get wsman message",
				AMTRedirectionService,
				wsmantesting.Get,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetAndPutResponse: RedirectionResponse{
						XMLName:                 xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMTRedirectionService), Local: AMTRedirectionService},
						CreationClassName:       AMTRedirectionService,
						ElementName:             "Intel(r) AMT Redirection Service",
						EnabledState:            32771,
						ListenerEnabled:         true,
						Name:                    "Intel(r) AMT Redirection Service",
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "Intel(r) AMT",
					},
				},
			},
			// ENUMERATES
			{
				"should create a valid AMT_RedirectionService Enumerate wsman message",
				AMTRedirectionService,
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError
					if elementUnderTest.Base.WSManMessageCreator == nil {
						logrus.Println("Error")
					}

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
				"should create a valid AMT_RedirectionService Pull wsman message",
				AMTRedirectionService,
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
						RedirectionItems: []RedirectionResponse{
							{
								XMLName:                 xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMTRedirectionService), Local: AMTRedirectionService},
								CreationClassName:       AMTRedirectionService,
								ElementName:             "Intel(r) AMT Redirection Service",
								EnabledState:            32771,
								ListenerEnabled:         true,
								Name:                    "Intel(r) AMT Redirection Service",
								SystemCreationClassName: "CIM_ComputerSystem",
								SystemName:              "Intel(r) AMT",
							},
						},
					},
				},
			},
			// PUTS
			{
				"should create a valid AMT_RedirectionService Put wsman message",
				AMTRedirectionService,
				wsmantesting.Put,
				"<h:AMT_RedirectionService xmlns:h=\"http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RedirectionService\"><h:CreationClassName>AMT_RedirectionService</h:CreationClassName><h:ElementName>Intel(r) AMT Redirection Service</h:ElementName><h:EnabledState>32771</h:EnabledState><h:ListenerEnabled>true</h:ListenerEnabled><h:Name>Intel(r) AMT Redirection Service</h:Name><h:SystemCreationClassName>CIM_ComputerSystem</h:SystemCreationClassName><h:SystemName>Intel(r) AMT</h:SystemName></h:AMT_RedirectionService>",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError
					redirectionRequest := &RedirectionRequest{
						CreationClassName:       AMTRedirectionService,
						ElementName:             "Intel(r) AMT Redirection Service",
						EnabledState:            32771,
						ListenerEnabled:         true,
						Name:                    "Intel(r) AMT Redirection Service",
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "Intel(r) AMT",
					}

					return elementUnderTest.Put(redirectionRequest)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetAndPutResponse: RedirectionResponse{
						XMLName:                 xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMTRedirectionService), Local: AMTRedirectionService},
						CreationClassName:       AMTRedirectionService,
						ElementName:             "Intel(r) AMT Redirection Service",
						EnabledState:            32771,
						ListenerEnabled:         true,
						Name:                    "Intel(r) AMT Redirection Service",
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "Intel(r) AMT",
					},
				},
			},
			// REQUEST STATE CHANGE
			{
				"should create a valid AMT_RedirectionService Request State Change wsman message",
				AMTRedirectionService,
				fmt.Sprintf("%s%s/%s", message.AMTSchema, AMTRedirectionService, "RequestStateChange"),
				"<h:RequestStateChange_INPUT xmlns:h=\"http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RedirectionService\"><h:RequestedState>32771</h:RequestedState></h:RequestStateChange_INPUT>",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.RequestStateChange(EnableIDERAndSOL)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					RequestStateChange_OUTPUT: RequestStateChange_OUTPUT{
						XMLName:     xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMTRedirectionService), Local: "RequestStateChange_OUTPUT"},
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
