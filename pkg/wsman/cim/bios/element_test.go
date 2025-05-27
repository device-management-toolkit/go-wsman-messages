/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package bios

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
			GetResponse: BiosElement{},
		},
	}
	expectedResult := "{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"GetResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"TargetOperatingSystem\":0,\"SoftwareElementID\":\"\",\"SoftwareElementState\":0,\"Name\":\"\",\"OperationalStatus\":null,\"ElementName\":\"\",\"Version\":\"\",\"Manufacturer\":\"\",\"PrimaryBIOS\":false,\"ReleaseDate\":{\"DateTime\":\"\"}},\"EnumerateResponse\":{\"EnumerationContext\":\"\"},\"PullResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"BiosElementItems\":null}}"
	result := response.JSON()
	assert.Equal(t, expectedResult, result)
}

func TestYaml(t *testing.T) {
	response := Response{
		Body: Body{
			GetResponse: BiosElement{},
		},
	}
	expectedResult := "xmlname:\n    space: \"\"\n    local: \"\"\ngetresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    targetoperatingsystem: 0\n    softwareelementid: \"\"\n    softwareelementstate: 0\n    name: \"\"\n    operationalstatus: []\n    elementname: \"\"\n    version: \"\"\n    manufacturer: \"\"\n    primarybios: false\n    releasedate:\n        datetime: \"\"\nenumerateresponse:\n    enumerationcontext: \"\"\npullresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    bioselementitems: []\n"
	result := response.YAML()
	assert.Equal(t, expectedResult, result)
}

func TestPositiveCIMBIOSElement(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.CIMResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/bios/element",
	}
	elementUnderTest := NewBIOSElementWithClient(wsmanMessageCreator, &client)

	t.Run("cim_BIOSElement Tests", func(t *testing.T) {
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
				"should create and parse a valid cim_BIOSElement Get call",
				CIMBIOSElement,
				wsmantesting.Get,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageGet

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetResponse: BiosElement{
						XMLName:               xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BIOSElement", Local: CIMBIOSElement},
						TargetOperatingSystem: 66,
						SoftwareElementID:     "QNCFLX70.0054.2020.0810.2227",
						SoftwareElementState:  2,
						Name:                  "Primary BIOS",
						OperationalStatus:     []OperationalStatus{0},
						ElementName:           "Primary BIOS",
						Version:               "QNCFLX70.0054.2020.0810.2227",
						Manufacturer:          "Intel Corp.",
						PrimaryBIOS:           true,
						ReleaseDate:           Time{DateTime: "2020-08-10T00:00:00Z"},
					},
				},
			},
			// ENUMERATES
			{
				"should create and parse a valid cim_BIOSElement Enumerate call",
				CIMBIOSElement,
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageEnumerate

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "C4020000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create and parse a valid cim_BIOSElement Pull call",
				CIMBIOSElement,
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
						BiosElementItems: []BiosElement{
							{
								XMLName:               xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BIOSElement", Local: CIMBIOSElement},
								TargetOperatingSystem: 66,
								SoftwareElementID:     "QNCFLX70.0054.2020.0810.2227",
								SoftwareElementState:  2,
								Name:                  "Primary BIOS",
								OperationalStatus:     []OperationalStatus{0},
								ElementName:           "Primary BIOS",
								Version:               "QNCFLX70.0054.2020.0810.2227",
								Manufacturer:          "Intel Corp.",
								PrimaryBIOS:           true,
								ReleaseDate:           Time{DateTime: "2020-08-10T00:00:00Z"},
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

func TestNegativeCIMBIOSElement(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.CIMResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/bios/element",
	}
	elementUnderTest := NewBIOSElementWithClient(wsmanMessageCreator, &client)

	t.Run("cim_BIOSElement Tests", func(t *testing.T) {
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
				"should create and parse a valid cim_BIOSElement Get call",
				CIMBIOSElement,
				wsmantesting.Get,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetResponse: BiosElement{
						XMLName:               xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BIOSElement", Local: CIMBIOSElement},
						TargetOperatingSystem: 66,
						SoftwareElementID:     "QNCFLX70.0054.2020.0810.2227",
						SoftwareElementState:  2,
						Name:                  "Primary BIOS",
						OperationalStatus:     []OperationalStatus{0},
						ElementName:           "Primary BIOS",
						Version:               "QNCFLX70.0054.2020.0810.2227",
						Manufacturer:          "Intel Corp.",
						PrimaryBIOS:           true,
						ReleaseDate:           Time{DateTime: "2020-08-10T00:00:00Z"},
					},
				},
			},
			// ENUMERATES
			{
				"should create and parse a valid cim_BIOSElement Enumerate call",
				CIMBIOSElement,
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "C4020000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create and parse a valid cim_BIOSElement Pull call",
				CIMBIOSElement,
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
						BiosElementItems: []BiosElement{
							{
								XMLName:               xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BIOSElement", Local: CIMBIOSElement},
								TargetOperatingSystem: 66,
								SoftwareElementID:     "QNCFLX70.0054.2020.0810.2227",
								SoftwareElementState:  2,
								Name:                  "Primary BIOS",
								OperationalStatus:     []OperationalStatus{0},
								ElementName:           "Primary BIOS",
								Version:               "QNCFLX70.0054.2020.0810.2227",
								Manufacturer:          "Intel Corp.",
								PrimaryBIOS:           true,
								ReleaseDate:           Time{DateTime: "2020-08-10T00:00:00Z"},
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
