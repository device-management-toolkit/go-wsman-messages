/*********************************************************************
 * Copyright (c) Intel Corporation 2026
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package fan

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
	expectedResult := "{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"PullResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"FanItems\":null},\"EnumerateResponse\":{\"EnumerationContext\":\"\"},\"FanResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"ActiveCooling\":false,\"CreationClassName\":\"\",\"DesiredSpeed\":0,\"DeviceID\":\"\",\"ElementName\":\"\",\"EnabledState\":0,\"HealthState\":0,\"OperationalStatus\":null,\"RequestedState\":0,\"SystemCreationClassName\":\"\",\"SystemName\":\"\",\"VariableSpeed\":false}}"
	result := response.JSON()
	assert.Equal(t, expectedResult, result)
}

func TestYaml(t *testing.T) {
	response := Response{
		Body: Body{
			PullResponse: PullResponse{},
		},
	}
	expectedResult := "xmlname:\n    space: \"\"\n    local: \"\"\npullresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    fanitems: []\nenumerateresponse:\n    enumerationcontext: \"\"\nfanresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    activecooling: false\n    creationclassname: \"\"\n    desiredspeed: 0\n    deviceid: \"\"\n    elementname: \"\"\n    enabledstate: 0\n    healthstate: 0\n    operationalstatus: []\n    requestedstate: 0\n    systemcreationclassname: \"\"\n    systemname: \"\"\n    variablespeed: false\n"
	result := response.YAML()
	assert.Equal(t, expectedResult, result)
}

func TestPositiveCIMFan(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.CIMResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/fan",
	}
	elementUnderTest := NewFanWithClient(wsmanMessageCreator, &client)

	t.Run("cim_Fan Tests", func(t *testing.T) {
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
				"should create and parse a valid cim_Fan Get call",
				CIMFan, wsmantesting.Get,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageGet

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					FanResponse: FanResponse{
						XMLName:                 xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_Fan", Local: "CIM_Fan"},
						ActiveCooling:           true,
						CreationClassName:       "CIM_Fan",
						DesiredSpeed:            0,
						DeviceID:                "Fan 0",
						ElementName:             "Fan",
						EnabledState:            EnabledStateNotApplicable,
						HealthState:             HealthStateOK,
						OperationalStatus:       []OperationalStatus{OperationalStatusOK},
						RequestedState:          RequestedStateNotApplicable,
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "ManagedSystem",
						VariableSpeed:           false,
					},
				},
			},
			// ENUMERATES
			{
				"should create and parse a valid cim_Fan Enumerate call",
				CIMFan, wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageEnumerate

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "D0510500-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create and parse a valid cim_Fan Pull call",
				CIMFan, wsmantesting.Pull,
				wsmantesting.PullBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessagePull

					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						XMLName: xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/09/enumeration", Local: "PullResponse"},
						FanItems: []FanResponse{
							{
								XMLName:                 xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_Fan", Local: "CIM_Fan"},
								ActiveCooling:           true,
								CreationClassName:       "CIM_Fan",
								DesiredSpeed:            0,
								DeviceID:                "Fan 0",
								ElementName:             "Fan",
								EnabledState:            EnabledStateNotApplicable,
								HealthState:             HealthStateOK,
								OperationalStatus:       []OperationalStatus{OperationalStatusOK},
								RequestedState:          RequestedStateNotApplicable,
								SystemCreationClassName: "CIM_ComputerSystem",
								SystemName:              "ManagedSystem",
								VariableSpeed:           false,
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

func TestNegativeCIMFan(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.CIMResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/fan",
	}
	elementUnderTest := NewFanWithClient(wsmanMessageCreator, &client)

	t.Run("cim_Fan Tests", func(t *testing.T) {
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
				"should handle an error response on cim_Fan Get call",
				CIMFan, wsmantesting.Get,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					FanResponse: FanResponse{
						XMLName:                 xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_Fan", Local: "CIM_Fan"},
						ActiveCooling:           true,
						CreationClassName:       "CIM_Fan",
						DesiredSpeed:            0,
						DeviceID:                "Fan 0",
						ElementName:             "Fan",
						EnabledState:            EnabledStateNotApplicable,
						HealthState:             HealthStateOK,
						OperationalStatus:       []OperationalStatus{OperationalStatusOK},
						RequestedState:          RequestedStateNotApplicable,
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "ManagedSystem",
						VariableSpeed:           false,
					},
				},
			},
			// ENUMERATES
			{
				"should handle an error response on cim_Fan Enumerate call",
				CIMFan, wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "D0510500-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should handle an error response on cim_Fan Pull call",
				CIMFan, wsmantesting.Pull,
				wsmantesting.PullBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						XMLName: xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/09/enumeration", Local: "PullResponse"},
						FanItems: []FanResponse{
							{
								XMLName:                 xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_Fan", Local: "CIM_Fan"},
								ActiveCooling:           true,
								CreationClassName:       "CIM_Fan",
								DesiredSpeed:            0,
								DeviceID:                "Fan 0",
								ElementName:             "Fan",
								EnabledState:            EnabledStateNotApplicable,
								HealthState:             HealthStateOK,
								OperationalStatus:       []OperationalStatus{OperationalStatusOK},
								RequestedState:          RequestedStateNotApplicable,
								SystemCreationClassName: "CIM_ComputerSystem",
								SystemName:              "ManagedSystem",
								VariableSpeed:           false,
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
