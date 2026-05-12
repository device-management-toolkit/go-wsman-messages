/*********************************************************************
 * Copyright (c) Intel Corporation 2026
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package battery

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
	expectedResult := "{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"PullResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"BatteryItems\":null,\"EndOfSequence\":{\"Space\":\"\",\"Local\":\"\"}},\"EnumerateResponse\":{\"EnumerationContext\":\"\"},\"PackageResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"DeviceID\":\"\",\"CreationClassName\":\"\",\"SystemName\":\"\",\"SystemCreationClassName\":\"\",\"ElementName\":\"\",\"OperationalStatus\":null,\"HealthState\":0,\"EnabledState\":0,\"RequestedState\":0,\"BatteryStatus\":0,\"EstimatedChargeRemaining\":0,\"Chemistry\":0,\"DesignCapacity\":0,\"FullChargeCapacity\":0,\"DesignVoltage\":0,\"ChargingStatus\":0}}"
	result := response.JSON()
	assert.Equal(t, expectedResult, result)
}

func TestYaml(t *testing.T) {
	response := Response{
		Body: Body{
			PullResponse: PullResponse{},
		},
	}
	expectedResult := "xmlname:\n    space: \"\"\n    local: \"\"\npullresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    batteryitems: []\n    endofsequence:\n        space: \"\"\n        local: \"\"\nenumerateresponse:\n    enumerationcontext: \"\"\npackageresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    deviceid: \"\"\n    creationclassname: \"\"\n    systemname: \"\"\n    systemcreationclassname: \"\"\n    elementname: \"\"\n    operationalstatus: []\n    healthstate: 0\n    enabledstate: 0\n    requestedstate: 0\n    batterystatus: 0\n    estimatedchargeremaining: 0\n    chemistry: 0\n    designcapacity: 0\n    fullchargecapacity: 0\n    designvoltage: 0\n    chargingstatus: 0\n"
	result := response.YAML()
	assert.Equal(t, expectedResult, result)
}

func TestPositiveCIMBattery(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.CIMResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/battery",
	}
	elementUnderTest := NewBatteryWithClient(wsmanMessageCreator, &client)

	expectedBattery := PackageResponse{
		XMLName:                 xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_Battery", Local: "CIM_Battery"},
		BatteryStatus:           BatteryStatusUnknown,
		ChargingStatus:          ChargingStatusUnknown,
		Chemistry:               ChemistryUnknown,
		CreationClassName:       "CIM_Battery",
		DesignCapacity:          5601,
		DesignVoltage:           11550,
		DeviceID:                "CC03056XL",
		EnabledState:            EnabledStateEnabled,
		HealthState:             HealthStateUnknown,
		RequestedState:          RequestedStateEnabled,
		SystemCreationClassName: "CIM_ComputerSystem",
		SystemName:              "ManagedSystem",
	}

	t.Run("cim_Battery Tests", func(t *testing.T) {
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
				"should create and parse a valid cim_Battery Get wsman call",
				CIMBattery,
				wsmantesting.Get,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageGet

					return elementUnderTest.Get()
				},
				Body{
					XMLName:         xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PackageResponse: expectedBattery,
				},
			},
			// ENUMERATES
			{
				"should create and parse a valid cim_Battery Enumerate wsman call",
				CIMBattery,
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageEnumerate

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "B5000000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create and parse a valid cim_Battery Pull wsman call",
				CIMBattery,
				wsmantesting.Pull,
				wsmantesting.PullBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessagePull

					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						XMLName:       xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/09/enumeration", Local: "PullResponse"},
						BatteryItems:  []PackageResponse{expectedBattery},
						EndOfSequence: xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/09/enumeration", Local: "EndOfSequence"},
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

func TestNegativeCIMBattery(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.CIMResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/battery",
	}
	elementUnderTest := NewBatteryWithClient(wsmanMessageCreator, &client)

	expectedBattery := PackageResponse{
		XMLName:                 xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_Battery", Local: "CIM_Battery"},
		BatteryStatus:           BatteryStatusUnknown,
		ChargingStatus:          ChargingStatusUnknown,
		Chemistry:               ChemistryUnknown,
		CreationClassName:       "CIM_Battery",
		DesignCapacity:          5601,
		DesignVoltage:           11550,
		DeviceID:                "CC03056XL",
		EnabledState:            EnabledStateEnabled,
		HealthState:             HealthStateUnknown,
		RequestedState:          RequestedStateEnabled,
		SystemCreationClassName: "CIM_ComputerSystem",
		SystemName:              "ManagedSystem",
	}

	t.Run("cim_Battery Tests", func(t *testing.T) {
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
				"should create and parse a valid cim_Battery Get wsman call",
				CIMBattery,
				wsmantesting.Get,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Get()
				},
				Body{
					XMLName:         xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PackageResponse: expectedBattery,
				},
			},
			// ENUMERATES
			{
				"should create and parse a valid cim_Battery Enumerate wsman call",
				CIMBattery,
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "B5000000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create and parse a valid cim_Battery Pull wsman call",
				CIMBattery,
				wsmantesting.Pull,
				wsmantesting.PullBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						XMLName:       xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/09/enumeration", Local: "PullResponse"},
						BatteryItems:  []PackageResponse{expectedBattery},
						EndOfSequence: xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/09/enumeration", Local: "EndOfSequence"},
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
