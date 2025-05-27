/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package optin

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
	expectedResult := "{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"EnumerateResponse\":{\"EnumerationContext\":\"\"},\"GetAndPutResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"Name\":\"\",\"CreationClassName\":\"\",\"SystemName\":\"\",\"SystemCreationClassName\":\"\",\"ElementName\":\"\",\"OptInCodeTimeout\":0,\"OptInRequired\":0,\"OptInState\":0,\"CanModifyOptInPolicy\":0,\"OptInDisplayTimeout\":0},\"PullResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"OptInServiceItems\":null},\"StartOptInResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"ReturnValue\":0},\"CancelOptInResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"ReturnValue\":0},\"SendOptInCodeResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"ReturnValue\":0}}"
	result := response.JSON()
	assert.Equal(t, expectedResult, result)
}

func TestYaml(t *testing.T) {
	response := Response{
		Body: Body{
			PullResponse: PullResponse{},
		},
	}
	expectedResult := "xmlname:\n    space: \"\"\n    local: \"\"\nenumerateresponse:\n    enumerationcontext: \"\"\ngetandputresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    name: \"\"\n    creationclassname: \"\"\n    systemname: \"\"\n    systemcreationclassname: \"\"\n    elementname: \"\"\n    optincodetimeout: 0\n    optinrequired: 0\n    optinstate: 0\n    canmodifyoptinpolicy: 0\n    optindisplaytimeout: 0\npullresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    optinserviceitems: []\nstartoptinresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    returnvalue: 0\ncanceloptinresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    returnvalue: 0\nsendoptincoderesponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    returnvalue: 0\n"
	result := response.YAML()
	assert.Equal(t, expectedResult, result)
}

func TestPositiveIPS_OptInService(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.IPSResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "ips/optin",
	}
	elementUnderTest := NewOptInServiceWithClient(wsmanMessageCreator, &client)

	t.Run("ips_OptInService Tests", func(t *testing.T) {
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
				"should create a valid IPS_OptInService Get wsman message",
				IPSOptInService,
				wsmantesting.Get,
				"",
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageGet

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetAndPutResponse: OptInServiceResponse{
						XMLName:                 xml.Name{Space: fmt.Sprintf("%s%s", message.IPSSchema, IPSOptInService), Local: IPSOptInService},
						CanModifyOptInPolicy:    1,
						CreationClassName:       "IPS_OptInService",
						ElementName:             "Intel(r) AMT OptIn Service",
						Name:                    "Intel(r) AMT OptIn Service",
						OptInCodeTimeout:        120,
						OptInDisplayTimeout:     300,
						OptInRequired:           0,
						OptInState:              0,
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "Intel(r) AMT",
					},
				},
			},
			// ENUMERATES
			{
				"should create a valid IPS_OptInService Enumerate wsman message",
				IPSOptInService,
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
						EnumerationContext: "9E0A0000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create a valid IPS_OptInService Pull wsman message",
				IPSOptInService,
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
						OptInServiceItems: []OptInServiceResponse{
							{
								XMLName:                 xml.Name{Space: fmt.Sprintf("%s%s", message.IPSSchema, IPSOptInService), Local: IPSOptInService},
								CanModifyOptInPolicy:    1,
								CreationClassName:       "IPS_OptInService",
								ElementName:             "Intel(r) AMT OptIn Service",
								Name:                    "Intel(r) AMT OptIn Service",
								OptInCodeTimeout:        120,
								OptInDisplayTimeout:     300,
								OptInRequired:           0,
								OptInState:              0,
								SystemCreationClassName: "CIM_ComputerSystem",
								SystemName:              "Intel(r) AMT",
							},
						},
					},
				},
			},
			// SEND_OPT_IN_CODE
			{
				"should create a valid IPS_OptInService send opt in code wsman message",
				IPSOptInService,
				wsmantesting.SendOptInCode,
				`<h:SendOptInCode_INPUT xmlns:h="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_OptInService"><h:OptInCode>1</h:OptInCode></h:SendOptInCode_INPUT>`,
				"",
				func() (Response, error) {
					client.CurrentMessage = "SendOptInCode"

					return elementUnderTest.SendOptInCode(1)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					SendOptInCodeResponse: SendOptInCode_OUTPUT{
						XMLName:     xml.Name{Space: fmt.Sprintf("%s%s", message.IPSSchema, IPSOptInService), Local: "SendOptInCode_OUTPUT"},
						ReturnValue: 2,
					},
				},
			},
			// START_OPT_IN
			{
				"should create a valid IPS_OptInService start opt in code wsman message",
				IPSOptInService,
				wsmantesting.StartOptIn,
				`<h:StartOptIn_INPUT xmlns:h="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_OptInService"></h:StartOptIn_INPUT>`,
				"",
				func() (Response, error) {
					client.CurrentMessage = "StartOptIn"

					return elementUnderTest.StartOptIn()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					StartOptInResponse: StartOptIn_OUTPUT{
						XMLName:     xml.Name{Space: fmt.Sprintf("%s%s", message.IPSSchema, IPSOptInService), Local: "StartOptIn_OUTPUT"},
						ReturnValue: 2,
					},
				},
			},
			// CANCEL_OPT_IN
			{
				"should create a valid IPS_OptInService cancel opt in code wsman message",
				IPSOptInService,
				wsmantesting.CancelOptIn,
				`<h:CancelOptIn_INPUT xmlns:h="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_OptInService"></h:CancelOptIn_INPUT>`,
				"",
				func() (Response, error) {
					client.CurrentMessage = "CancelOptIn"

					return elementUnderTest.CancelOptIn()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					CancelOptInResponse: CancelOptIn_OUTPUT{
						XMLName:     xml.Name{Space: fmt.Sprintf("%s%s", message.IPSSchema, IPSOptInService), Local: "CancelOptIn_OUTPUT"},
						ReturnValue: 2,
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

func TestNegativeIPS_OptInService(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.IPSResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "ips/optin",
	}
	elementUnderTest := NewOptInServiceWithClient(wsmanMessageCreator, &client)

	t.Run("ips_OptInService Tests", func(t *testing.T) {
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
				"should create a valid IPS_OptInService Get wsman message",
				IPSOptInService,
				wsmantesting.Get,
				"",
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetAndPutResponse: OptInServiceResponse{
						XMLName:                 xml.Name{Space: fmt.Sprintf("%s%s", message.IPSSchema, IPSOptInService), Local: IPSOptInService},
						CanModifyOptInPolicy:    1,
						CreationClassName:       "IPS_OptInService",
						ElementName:             "Intel(r) AMT OptIn Service",
						Name:                    "Intel(r) AMT OptIn Service",
						OptInCodeTimeout:        120,
						OptInDisplayTimeout:     300,
						OptInRequired:           0,
						OptInState:              0,
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "Intel(r) AMT",
					},
				},
			},
			// ENUMERATES
			{
				"should create a valid IPS_OptInService Enumerate wsman message",
				IPSOptInService,
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
						EnumerationContext: "9E0A0000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create a valid IPS_OptInService Pull wsman message",
				IPSOptInService,
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
						XMLName: xml.Name{Space: message.XMLPullResponseSpace, Local: "PullResponse"},
						OptInServiceItems: []OptInServiceResponse{
							{
								XMLName:                 xml.Name{Space: fmt.Sprintf("%s%s", message.IPSSchema, IPSOptInService), Local: IPSOptInService},
								CanModifyOptInPolicy:    1,
								CreationClassName:       "IPS_OptInService",
								ElementName:             "Intel(r) AMT OptIn Service",
								Name:                    "Intel(r) AMT OptIn Service",
								OptInCodeTimeout:        120,
								OptInDisplayTimeout:     300,
								OptInRequired:           0,
								OptInState:              0,
								SystemCreationClassName: "CIM_ComputerSystem",
								SystemName:              "Intel(r) AMT",
							},
						},
					},
				},
			},

			// SEND_OPT_IN_CODE
			{
				"should create a valid IPS_OptInService send opt in code wsman message",
				IPSOptInService,
				wsmantesting.SendOptInCode,
				`<h:SendOptInCode_INPUT xmlns:h="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_OptInService"><h:OptInCode>1</h:OptInCode></h:SendOptInCode_INPUT>`,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.SendOptInCode(1)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					SendOptInCodeResponse: SendOptInCode_OUTPUT{
						XMLName:     xml.Name{Space: fmt.Sprintf("%s%s", message.IPSSchema, IPSOptInService), Local: "SendOptInCode_OUTPUT"},
						ReturnValue: 2,
					},
				},
			},

			// START_OPT_IN
			{
				"should create a valid IPS_OptInService start opt in code wsman message",
				IPSOptInService,
				wsmantesting.StartOptIn,
				`<h:StartOptIn_INPUT xmlns:h="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_OptInService"></h:StartOptIn_INPUT>`,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.StartOptIn()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					StartOptInResponse: StartOptIn_OUTPUT{
						XMLName:     xml.Name{Space: fmt.Sprintf("%s%s", message.IPSSchema, IPSOptInService), Local: "StartOptIn_OUTPUT"},
						ReturnValue: 2,
					},
				},
			},

			// CANCEL_OPT_IN
			{
				"should create a valid IPS_OptInService cancel opt in code wsman message",
				IPSOptInService,
				wsmantesting.CancelOptIn,
				`<h:CancelOptIn_INPUT xmlns:h="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_OptInService"></h:CancelOptIn_INPUT>`,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.CancelOptIn()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					CancelOptInResponse: CancelOptIn_OUTPUT{
						XMLName:     xml.Name{Space: fmt.Sprintf("%s%s", message.IPSSchema, IPSOptInService), Local: "CancelOptIn_OUTPUT"},
						ReturnValue: 2,
					},
				},
			},
			// PUT
			{
				"should create a valid IPS_OptInService Put wsman message",
				IPSOptInService,
				wsmantesting.Put,
				`<h:IPS_OptInService xmlns:h="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_OptInService"><h:CanModifyOptInPolicy>1</h:CanModifyOptInPolicy><h:CreationClassName>IPS_OptInService</h:CreationClassName><h:ElementName>Intel(r) AMT OptIn Service</h:ElementName><h:Name>Intel(r) AMT OptIn Service</h:Name><h:OptInCodeTimeout>120</h:OptInCodeTimeout><h:OptInDisplayTimeout>300</h:OptInDisplayTimeout><h:OptInRequired>0</h:OptInRequired><h:OptInState>0</h:OptInState><h:SystemName>Intel(r) AMT</h:SystemName><h:SystemCreationClassName>CIM_ComputerSystem</h:SystemCreationClassName></h:IPS_OptInService>`,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError
					request := OptInServiceRequest{
						H:                       "http://intel.com/wbem/wscim/1/ips-schema/1//IPS_OptInService",
						CanModifyOptInPolicy:    1,
						CreationClassName:       "IPS_OptInService",
						ElementName:             "Intel(r) AMT OptIn Service",
						Name:                    "Intel(r) AMT OptIn Service",
						OptInCodeTimeout:        120,
						OptInDisplayTimeout:     300,
						OptInRequired:           0,
						OptInState:              0,
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "Intel(r) AMT",
					}

					return elementUnderTest.Put(request)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetAndPutResponse: OptInServiceResponse{
						XMLName:                 xml.Name{Space: fmt.Sprintf("%s%s", message.IPSSchema, IPSOptInService), Local: IPSOptInService},
						CanModifyOptInPolicy:    1,
						CreationClassName:       "IPS_OptInService",
						ElementName:             "Intel(r) AMT OptIn Service",
						Name:                    "Intel(r) AMT OptIn Service",
						OptInCodeTimeout:        120,
						OptInDisplayTimeout:     300,
						OptInRequired:           0,
						OptInState:              0,
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "Intel(r) AMT",
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
