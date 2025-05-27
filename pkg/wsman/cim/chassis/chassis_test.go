/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package chassis

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
	expectedResult := "{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"PullResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"PackageItems\":null},\"EnumerateResponse\":{\"EnumerationContext\":\"\"},\"PackageResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"Version\":\"\",\"SerialNumber\":\"\",\"Model\":\"\",\"Manufacturer\":\"\",\"ElementName\":\"\",\"CreationClassName\":\"\",\"Tag\":\"\",\"OperationalStatus\":null,\"PackageType\":0,\"ChassisPackageType\":0}}"
	result := response.JSON()
	assert.Equal(t, expectedResult, result)
}

func TestYaml(t *testing.T) {
	response := Response{
		Body: Body{
			PullResponse: PullResponse{},
		},
	}
	expectedResult := "xmlname:\n    space: \"\"\n    local: \"\"\npullresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    packageitems: []\nenumerateresponse:\n    enumerationcontext: \"\"\npackageresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    version: \"\"\n    serialnumber: \"\"\n    model: \"\"\n    manufacturer: \"\"\n    elementname: \"\"\n    creationclassname: \"\"\n    tag: \"\"\n    operationalstatus: []\n    packagetype: 0\n    chassispackagetype: 0\n"
	result := response.YAML()
	assert.Equal(t, expectedResult, result)
}

func TestPositiveCIMChassis(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.CIMResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/chassis",
	}
	elementUnderTest := NewChassisWithClient(wsmanMessageCreator, &client)

	t.Run("cim_* Tests", func(t *testing.T) {
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
				"should create and parse a valid cim_Chassis Get call",
				CIMChassis,
				wsmantesting.Get,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageGet

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PackageResponse: PackageResponse{
						XMLName:            xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_Chassis", Local: "CIM_Chassis"},
						ChassisPackageType: 0,
						CreationClassName:  "CIM_Chassis",
						ElementName:        "Managed System Chassis",
						Manufacturer:       "Intel(R) Client Systems",
						Model:              "NUC9V7QNX",
						OperationalStatus:  []OperationalStatus{0},
						PackageType:        3,
						SerialNumber:       "JRQN0243007J",
						Tag:                "CIM_Chassis",
						Version:            "K47174-402",
					},
				},
			},
			// ENUMERATES
			{
				"should create and parse a valid cim_Chassis Enumerate call",
				CIMChassis,
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageEnumerate

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "D1020000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create and parse a valid cim_Chassis Pull call",
				CIMChassis,
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
								XMLName:            xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_Chassis", Local: "CIM_Chassis"},
								ChassisPackageType: 0,
								CreationClassName:  "CIM_Chassis",
								ElementName:        "Managed System Chassis",
								Manufacturer:       "Intel(R) Client Systems",
								Model:              "NUC9V7QNX",
								OperationalStatus:  []OperationalStatus{0},
								PackageType:        3,
								SerialNumber:       "JRQN0243007J",
								Tag:                "CIM_Chassis",
								Version:            "K47174-402",
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

func TestNegativeCIMChassis(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.CIMResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/chassis",
	}
	elementUnderTest := NewChassisWithClient(wsmanMessageCreator, &client)

	t.Run("cim_Chassis Tests", func(t *testing.T) {
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
				"should create and parse a valid cim_Chassis Get call",
				CIMChassis,
				wsmantesting.Get,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PackageResponse: PackageResponse{
						XMLName:            xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_Chassis", Local: "CIM_Chassis"},
						ChassisPackageType: 0,
						CreationClassName:  "CIM_Chassis",
						ElementName:        "Managed System Chassis",
						Manufacturer:       "Intel(R) Client Systems",
						Model:              "NUC9V7QNX",
						OperationalStatus:  []OperationalStatus{0},
						PackageType:        3,
						SerialNumber:       "JRQN0243007J",
						Tag:                "CIM_Chassis",
						Version:            "K47174-402",
					},
				},
			},
			// ENUMERATES
			{
				"should create and parse a valid cim_Chassis Enumerate call",
				CIMChassis,
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "D1020000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create and parse a valid cim_Chassis Pull call",
				CIMChassis,
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
								XMLName:            xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_Chassis", Local: "CIM_Chassis"},
								ChassisPackageType: 0,
								CreationClassName:  "CIM_Chassis",
								ElementName:        "Managed System Chassis",
								Manufacturer:       "Intel(R) Client Systems",
								Model:              "NUC9V7QNX",
								OperationalStatus:  []OperationalStatus{0},
								PackageType:        3,
								SerialNumber:       "JRQN0243007J",
								Tag:                "CIM_Chassis",
								Version:            "K47174-402",
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
