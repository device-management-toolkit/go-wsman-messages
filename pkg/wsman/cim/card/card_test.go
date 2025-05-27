/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package card

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
	expectedResult := "{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"PullResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"CardItems\":null},\"EnumerateResponse\":{\"EnumerationContext\":\"\"},\"PackageResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"CanBeFRUed\":false,\"CreationClassName\":\"\",\"ElementName\":\"\",\"Manufacturer\":\"\",\"Model\":\"\",\"OperationalStatus\":null,\"PackageType\":0,\"SerialNumber\":\"\",\"Tag\":\"\",\"Version\":\"\"}}"
	result := response.JSON()
	assert.Equal(t, expectedResult, result)
}

func TestYaml(t *testing.T) {
	response := Response{
		Body: Body{
			PullResponse: PullResponse{},
		},
	}
	expectedResult := "xmlname:\n    space: \"\"\n    local: \"\"\npullresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    carditems: []\nenumerateresponse:\n    enumerationcontext: \"\"\npackageresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    canbefrued: false\n    creationclassname: \"\"\n    elementname: \"\"\n    manufacturer: \"\"\n    model: \"\"\n    operationalstatus: []\n    packagetype: 0\n    serialnumber: \"\"\n    tag: \"\"\n    version: \"\"\n"
	result := response.YAML()
	assert.Equal(t, expectedResult, result)
}

func TestPositiveCIMCard(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.CIMResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/card",
	}
	elementUnderTest := NewCardWithClient(wsmanMessageCreator, &client)

	t.Run("cim_Card Tests", func(t *testing.T) {
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
				"should create a valid cim_Card Get wsman message",
				CIMCard,
				wsmantesting.Get,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageGet

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PackageResponse: PackageResponse{
						XMLName:           xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_Card", Local: "CIM_Card"},
						CanBeFRUed:        true,
						CreationClassName: "CIM_Card",
						ElementName:       "Managed System Base Board",
						Manufacturer:      "Intel Corporation",
						Model:             "NUC9V7QNB",
						OperationalStatus: []OperationalStatus{0},
						PackageType:       9,
						SerialNumber:      "KNQN0221020W",
						Tag:               "CIM_Card",
						Version:           "K47180-402",
					},
				},
			},
			// ENUMERATES
			{
				"should create a valid cim_Card Enumerate wsman message",
				CIMCard,
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageEnumerate

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "CF020000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create a valid cim_Card Pull wsman message",
				CIMCard,
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
						CardItems: []PackageResponse{
							{
								XMLName:           xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_Card", Local: "CIM_Card"},
								CanBeFRUed:        true,
								CreationClassName: "CIM_Card",
								ElementName:       "Managed System Base Board",
								Manufacturer:      "Intel Corporation",
								Model:             "NUC9V7QNB",
								OperationalStatus: []OperationalStatus{0},
								PackageType:       9,
								SerialNumber:      "KNQN0221020W",
								Tag:               "CIM_Card",
								Version:           "K47180-402",
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

func TestNegativeCIMCard(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.CIMResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/card",
	}
	elementUnderTest := NewCardWithClient(wsmanMessageCreator, &client)

	t.Run("cim_Card* Tests", func(t *testing.T) {
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
				"should create a valid cim_Card Get wsman message",
				CIMCard,
				wsmantesting.Get,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PackageResponse: PackageResponse{
						XMLName:           xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_Card", Local: "CIM_Card"},
						CanBeFRUed:        true,
						CreationClassName: "CIM_Card",
						ElementName:       "Managed System Base Board",
						Manufacturer:      "Intel Corporation",
						Model:             "NUC9V7QNB",
						OperationalStatus: []OperationalStatus{0},
						PackageType:       9,
						SerialNumber:      "KNQN0221020W",
						Tag:               "CIM_Card",
						Version:           "K47180-402",
					},
				},
			},
			// ENUMERATES
			{
				"should create a valid cim_Card Enumerate wsman message",
				CIMCard,
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "CF020000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create a valid cim_Card Pull wsman message",
				CIMCard,
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
						CardItems: []PackageResponse{
							{
								XMLName:           xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_Card", Local: "CIM_Card"},
								CanBeFRUed:        true,
								CreationClassName: "CIM_Card",
								ElementName:       "Managed System Base Board",
								Manufacturer:      "Intel Corporation",
								Model:             "NUC9V7QNB",
								OperationalStatus: []OperationalStatus{0},
								PackageType:       9,
								SerialNumber:      "KNQN0221020W",
								Tag:               "CIM_Card",
								Version:           "K47180-402",
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
