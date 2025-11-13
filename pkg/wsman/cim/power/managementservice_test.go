/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package power

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/methods"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/wsmantesting"
)

const (
	RequestPowerStateChangeBODY = "<h:RequestPowerStateChange_INPUT xmlns:h=\"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_PowerManagementService\"><h:PowerState>8</h:PowerState><h:ManagedElement><Address xmlns=\"http://schemas.xmlsoap.org/ws/2004/08/addressing\">http://schemas.xmlsoap.org/ws/2004/08/addressing</Address><ReferenceParameters xmlns=\"http://schemas.xmlsoap.org/ws/2004/08/addressing\"><ResourceURI xmlns=\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\">http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ComputerSystem</ResourceURI><SelectorSet xmlns=\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\"><Selector Name=\"CreationClassName\">CIM_ComputerSystem</Selector><Selector Name=\"Name\">ManagedSystem</Selector></SelectorSet></ReferenceParameters></h:ManagedElement></h:RequestPowerStateChange_INPUT>"
)

func TestJson(t *testing.T) {
	response := Response{
		Body: Body{
			PullResponse: PullResponse{},
		},
	}
	expectedResult := "{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"RequestPowerStateChangeResponse\":{\"ReturnValue\":0},\"GetResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"CreationClassName\":\"\",\"ElementName\":\"\",\"EnabledState\":0,\"Name\":\"\",\"RequestedState\":0,\"SystemCreationClassName\":\"\",\"SystemName\":\"\"},\"EnumerateResponse\":{\"EnumerationContext\":\"\"},\"PullResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"PowerManagementServiceItems\":null}}"
	result := response.JSON()
	assert.Equal(t, expectedResult, result)
}

func TestYaml(t *testing.T) {
	response := Response{
		Body: Body{
			PullResponse: PullResponse{},
		},
	}
	expectedResult := "xmlname:\n    space: \"\"\n    local: \"\"\nrequestpowerstatechangeresponse:\n    returnvalue: 0\ngetresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    creationclassname: \"\"\n    elementname: \"\"\n    enabledstate: 0\n    name: \"\"\n    requestedstate: 0\n    systemcreationclassname: \"\"\n    systemname: \"\"\nenumerateresponse:\n    enumerationcontext: \"\"\npullresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    powermanagementserviceitems: []\n"
	result := response.YAML()
	assert.Equal(t, expectedResult, result)
}

func TestPositiveCIMPowerManagementService(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.CIMResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/power/managementservice",
	}
	elementUnderTest := NewPowerManagementServiceWithClient(wsmanMessageCreator, &client)

	t.Run("cim_PowerManagementService Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			body             string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			{
				"Should issue a valid cim_PowerManagementService RequestPowerStateChange call",
				CIMPowerManagementService,
				methods.GenerateAction(CIMPowerManagementService, RequestPowerStateChange),
				RequestPowerStateChangeBODY,
				func() (Response, error) {
					client.CurrentMessage = "RequestPowerStateChange"
					powerState := PowerOffHard

					return elementUnderTest.RequestPowerStateChange(powerState)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					RequestPowerStateChangeResponse: PowerActionResponse{
						ReturnValue: 0,
					},
				},
			},
			{
				"Should issue a valid cim_PowerManagementService Get call",
				CIMPowerManagementService,
				wsmantesting.Get,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageGet

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetResponse: PowerManagementService{
						XMLName:                 xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_PowerManagementService", Local: "CIM_PowerManagementService"},
						CreationClassName:       "CIM_PowerManagementService",
						ElementName:             "Intel(r) AMT Power Management Service",
						EnabledState:            5,
						Name:                    "Intel(r) AMT Power Management Service",
						RequestedState:          12,
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "Intel(r) AMT",
					},
				},
			},
			{
				"Should issue a valid cim_PowerManagementService Enumerate call",
				CIMPowerManagementService,
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageEnumerate

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "DB020000-0000-0000-0000-000000000000",
					},
				},
			},
			{
				"Should issue a valid cim_PowerManagementService Pull call",
				CIMPowerManagementService,
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
						PowerManagementServiceItems: []PowerManagementService{
							{
								XMLName:                 xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_PowerManagementService", Local: "CIM_PowerManagementService"},
								CreationClassName:       "CIM_PowerManagementService",
								ElementName:             "Intel(r) AMT Power Management Service",
								EnabledState:            5,
								Name:                    "Intel(r) AMT Power Management Service",
								RequestedState:          12,
								SystemCreationClassName: "CIM_ComputerSystem",
								SystemName:              "Intel(r) AMT",
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

func TestNegativeCIMPowerManagementService(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.CIMResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/power/managementservice",
	}
	elementUnderTest := NewPowerManagementServiceWithClient(wsmanMessageCreator, &client)

	t.Run("cim_PowerManagementService Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			body             string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			{
				"Should issue a valid cim_PowerManagementService RequestPowerStateChange call",
				CIMPowerManagementService,
				methods.GenerateAction(CIMPowerManagementService, RequestPowerStateChange),
				RequestPowerStateChangeBODY,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError
					powerState := PowerOffHard

					return elementUnderTest.RequestPowerStateChange(powerState)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					RequestPowerStateChangeResponse: PowerActionResponse{
						ReturnValue: 0,
					},
				},
			},
			{
				"Should issue a valid cim_PowerManagementService Get call",
				CIMPowerManagementService,
				wsmantesting.Get,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetResponse: PowerManagementService{
						XMLName:                 xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_PowerManagementService", Local: "CIM_PowerManagementService"},
						CreationClassName:       "CIM_PowerManagementService",
						ElementName:             "Intel(r) AMT Power Management Service",
						EnabledState:            5,
						Name:                    "Intel(r) AMT Power Management Service",
						RequestedState:          12,
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "Intel(r) AMT",
					},
				},
			},
			{
				"Should issue a valid cim_PowerManagementService Enumerate call",
				CIMPowerManagementService,
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "DB020000-0000-0000-0000-000000000000",
					},
				},
			},
			{
				"Should issue a valid cim_PowerManagementService Pull call",
				CIMPowerManagementService,
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
						PowerManagementServiceItems: []PowerManagementService{
							{
								XMLName:                 xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_PowerManagementService", Local: "CIM_PowerManagementService"},
								CreationClassName:       "CIM_PowerManagementService",
								ElementName:             "Intel(r) AMT Power Management Service",
								EnabledState:            5,
								Name:                    "Intel(r) AMT Power Management Service",
								RequestedState:          12,
								SystemCreationClassName: "CIM_ComputerSystem",
								SystemName:              "Intel(r) AMT",
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
