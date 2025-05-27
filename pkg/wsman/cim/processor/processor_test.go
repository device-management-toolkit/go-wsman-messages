/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package processor

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
	expectedResult := "{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"PullResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"PackageItems\":null},\"EnumerateResponse\":{\"EnumerationContext\":\"\"},\"PackageResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"DeviceID\":\"\",\"CreationClassName\":\"\",\"SystemName\":\"\",\"SystemCreationClassName\":\"\",\"ElementName\":\"\",\"OperationalStatus\":null,\"HealthState\":0,\"EnabledState\":0,\"RequestedState\":0,\"Role\":\"\",\"Family\":0,\"OtherFamilyDescription\":\"\",\"UpgradeMethod\":0,\"MaxClockSpeed\":0,\"CurrentClockSpeed\":0,\"Stepping\":\"\",\"CPUStatus\":0,\"ExternalBusClockSpeed\":0}}"
	result := response.JSON()
	assert.Equal(t, expectedResult, result)
}

func TestYaml(t *testing.T) {
	response := Response{
		Body: Body{
			PullResponse: PullResponse{},
		},
	}
	expectedResult := "xmlname:\n    space: \"\"\n    local: \"\"\npullresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    packageitems: []\nenumerateresponse:\n    enumerationcontext: \"\"\npackageresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    deviceid: \"\"\n    creationclassname: \"\"\n    systemname: \"\"\n    systemcreationclassname: \"\"\n    elementname: \"\"\n    operationalstatus: []\n    healthstate: 0\n    enabledstate: 0\n    requestedstate: 0\n    role: \"\"\n    family: 0\n    otherfamilydescription: \"\"\n    upgrademethod: 0\n    maxclockspeed: 0\n    currentclockspeed: 0\n    stepping: \"\"\n    cpustatus: 0\n    externalbusclockspeed: 0\n"
	result := response.YAML()
	assert.Equal(t, expectedResult, result)
}

func TestPositiveCIMProcessor(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.CIMResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/physical/processor",
	}
	elementUnderTest := NewProcessorWithClient(wsmanMessageCreator, &client)

	t.Run("cim_Processor Tests", func(t *testing.T) {
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
				"should create and parse a valid cim_Processor Get wsman call",
				CIMProcessor,
				wsmantesting.Get,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageGet

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PackageResponse: PackageResponse{
						XMLName:                 xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_Processor", Local: "CIM_Processor"},
						CPUStatus:               1,
						CreationClassName:       "CIM_Processor",
						CurrentClockSpeed:       2400,
						DeviceID:                "CPU 0",
						ElementName:             "Managed System CPU",
						EnabledState:            2,
						ExternalBusClockSpeed:   100,
						Family:                  198,
						HealthState:             0,
						MaxClockSpeed:           8300,
						OperationalStatus:       []OperationalStatus{0},
						OtherFamilyDescription:  "",
						RequestedState:          12,
						Role:                    "Central",
						Stepping:                "13",
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "ManagedSystem",
						UpgradeMethod:           52,
					},
				},
			},
			// ENUMERATES
			{
				"should create and parse a valid cim_Processor Enumerate wsman call",
				CIMProcessor,
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageEnumerate

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "D9020000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create and parse a valid cim_Processor Pull wsman call",
				CIMProcessor,
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
						PackageItems: []PackageResponse{
							{
								XMLName:                 xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_Processor", Local: "CIM_Processor"},
								CPUStatus:               1,
								CreationClassName:       "CIM_Processor",
								CurrentClockSpeed:       2400,
								DeviceID:                "CPU 0",
								ElementName:             "Managed System CPU",
								EnabledState:            2,
								ExternalBusClockSpeed:   100,
								Family:                  198,
								HealthState:             0,
								MaxClockSpeed:           8300,
								OperationalStatus:       []OperationalStatus{0},
								OtherFamilyDescription:  "",
								RequestedState:          12,
								Role:                    "Central",
								Stepping:                "13",
								SystemCreationClassName: "CIM_ComputerSystem",
								SystemName:              "ManagedSystem",
								UpgradeMethod:           52,
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

func TestNegativeCIMProcessor(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.CIMResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/physical/processor",
	}
	elementUnderTest := NewProcessorWithClient(wsmanMessageCreator, &client)

	t.Run("cim_Processor Tests", func(t *testing.T) {
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
				"should create and parse a valid cim_Processor Get wsman call",
				CIMProcessor,
				wsmantesting.Get,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PackageResponse: PackageResponse{
						XMLName:                 xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_Processor", Local: "CIM_Processor"},
						CPUStatus:               1,
						CreationClassName:       "CIM_Processor",
						CurrentClockSpeed:       2400,
						DeviceID:                "CPU 0",
						ElementName:             "Managed System CPU",
						EnabledState:            2,
						ExternalBusClockSpeed:   100,
						Family:                  198,
						HealthState:             0,
						MaxClockSpeed:           8300,
						OperationalStatus:       []OperationalStatus{0},
						OtherFamilyDescription:  "",
						RequestedState:          12,
						Role:                    "Central",
						Stepping:                "13",
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "ManagedSystem",
						UpgradeMethod:           52,
					},
				},
			},
			// ENUMERATES
			{
				"should create and parse a valid cim_Processor Enumerate wsman call",
				CIMProcessor,
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "D9020000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create and parse a valid cim_Processor Pull wsman call",
				CIMProcessor,
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
						PackageItems: []PackageResponse{
							{
								XMLName:                 xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_Processor", Local: "CIM_Processor"},
								CPUStatus:               1,
								CreationClassName:       "CIM_Processor",
								CurrentClockSpeed:       2400,
								DeviceID:                "CPU 0",
								ElementName:             "Managed System CPU",
								EnabledState:            2,
								ExternalBusClockSpeed:   100,
								Family:                  198,
								HealthState:             0,
								MaxClockSpeed:           8300,
								OperationalStatus:       []OperationalStatus{0},
								OtherFamilyDescription:  "",
								RequestedState:          12,
								Role:                    "Central",
								Stepping:                "13",
								SystemCreationClassName: "CIM_ComputerSystem",
								SystemName:              "ManagedSystem",
								UpgradeMethod:           52,
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
