/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package physical

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
	expectedResult := "{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"PullResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"EnumerationContext\":\"\",\"MemoryItems\":null,\"Card\":null,\"PhysicalPackage\":null,\"Chassis\":null,\"EndOfSequence\":{\"Space\":\"\",\"Local\":\"\"}},\"EnumerateResponse\":{\"EnumerationContext\":\"\"},\"MemoryResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"PartNumber\":\"\",\"SerialNumber\":\"\",\"Manufacturer\":\"\",\"ElementName\":\"\",\"CreationClassName\":\"\",\"Tag\":\"\",\"OperationalStatus\":null,\"FormFactor\":0,\"MemoryType\":0,\"Speed\":0,\"Capacity\":0,\"BankLabel\":\"\",\"ConfiguredMemoryClockSpeed\":0,\"IsSpeedInMhz\":false,\"MaxMemorySpeed\":0}}"
	result := response.JSON()
	assert.Equal(t, expectedResult, result)
}

func TestYaml(t *testing.T) {
	response := Response{
		Body: Body{
			PullResponse: PullResponse{},
		},
	}
	expectedResult := "xmlname:\n    space: \"\"\n    local: \"\"\npullresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    enumerateresponse:\n        enumerationcontext: \"\"\n    memoryitems: []\n    card: []\n    physicalpackage: []\n    chassis: []\n    endofsequence:\n        space: \"\"\n        local: \"\"\nenumerateresponse:\n    enumerationcontext: \"\"\nmemoryresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    partnumber: \"\"\n    serialnumber: \"\"\n    manufacturer: \"\"\n    elementname: \"\"\n    creationclassname: \"\"\n    tag: \"\"\n    operationalstatus: []\n    formfactor: 0\n    memorytype: 0\n    speed: 0\n    capacity: 0\n    banklabel: \"\"\n    configuredmemoryclockspeed: 0\n    isspeedinmhz: false\n    maxmemoryspeed: 0\n"
	result := response.YAML()
	assert.Equal(t, expectedResult, result)
}

func TestPositiveCIMMemory(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.CIMResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/physical/memory",
	}
	elementUnderTest := NewPhysicalMemoryWithClient(wsmanMessageCreator, &client)

	t.Run("cim_Memory Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			body             string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			// ENUMERATES
			{
				"should create and parse a valid cim_PhysicalMemory Enumerate call",
				"CIM_PhysicalMemory",
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageEnumerate

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "D6020000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create and parse a valid cim_PhysicalMemory Pull call",
				"CIM_PhysicalMemory",
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
						MemoryItems: []PhysicalMemory{
							{
								XMLName:                    xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_PhysicalMemory", Local: "CIM_PhysicalMemory"},
								BankLabel:                  "BANK 0",
								Capacity:                   17179869184,
								ConfiguredMemoryClockSpeed: 2400,
								CreationClassName:          "CIM_PhysicalMemory",
								ElementName:                "Managed System Memory Chip",
								FormFactor:                 13,
								IsSpeedInMhz:               true,
								Manufacturer:               "859B",
								MaxMemorySpeed:             2400,
								MemoryType:                 26,
								PartNumber:                 "CT16G4SFD824A.M16FE ",
								SerialNumber:               "E0E8D670",
								Speed:                      0,
								Tag:                        "9876543210",
							},
							{
								XMLName:                    xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_PhysicalMemory", Local: "CIM_PhysicalMemory"},
								BankLabel:                  "BANK 2",
								Capacity:                   17179869184,
								ConfiguredMemoryClockSpeed: 2400,
								CreationClassName:          "CIM_PhysicalMemory",
								ElementName:                "Managed System Memory Chip",
								FormFactor:                 13,
								IsSpeedInMhz:               true,
								Manufacturer:               "859B",
								MaxMemorySpeed:             2400,
								MemoryType:                 26,
								PartNumber:                 "CT16G4SFD824A.M16FE ",
								SerialNumber:               "E0E8D070",
								Speed:                      0,
								Tag:                        "9876543210 (#1)",
							},
						},
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

func TestNegativeCIMMemory(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.CIMResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/physical/memory",
	}
	elementUnderTest := NewPhysicalMemoryWithClient(wsmanMessageCreator, &client)

	t.Run("cim_Memory Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			body             string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			// ENUMERATES
			{
				"should create and parse a valid cim_PhysicalMemory Enumerate call",
				"CIM_PhysicalMemory",
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "D6020000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create and parse a valid cim_PhysicalMemory Pull call",
				"CIM_PhysicalMemory",
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
						MemoryItems: []PhysicalMemory{
							{
								XMLName:                    xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_PhysicalMemory", Local: "CIM_PhysicalMemory"},
								BankLabel:                  "BANK 0",
								Capacity:                   17179869184,
								ConfiguredMemoryClockSpeed: 2400,
								CreationClassName:          "CIM_PhysicalMemory",
								ElementName:                "Managed System Memory Chip",
								FormFactor:                 13,
								IsSpeedInMhz:               true,
								Manufacturer:               "859B",
								MaxMemorySpeed:             2400,
								MemoryType:                 26,
								PartNumber:                 "CT16G4SFD824A.M16FE ",
								SerialNumber:               "E0E8D670",
								Speed:                      0,
								Tag:                        "9876543210",
							},
							{
								XMLName:                    xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_PhysicalMemory", Local: "CIM_PhysicalMemory"},
								BankLabel:                  "BANK 2",
								Capacity:                   17179869184,
								ConfiguredMemoryClockSpeed: 2400,
								CreationClassName:          "CIM_PhysicalMemory",
								ElementName:                "Managed System Memory Chip",
								FormFactor:                 13,
								IsSpeedInMhz:               true,
								Manufacturer:               "859B",
								MaxMemorySpeed:             2400,
								MemoryType:                 26,
								PartNumber:                 "CT16G4SFD824A.M16FE ",
								SerialNumber:               "E0E8D070",
								Speed:                      0,
								Tag:                        "9876543210 (#1)",
							},
						},
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
