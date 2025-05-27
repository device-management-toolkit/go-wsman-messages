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
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/card"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/wsmantesting"
)

func TestPositiveCIMPackage(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.CIMResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/physical/package",
	}
	elementUnderTest := NewPhysicalPackageWithClient(wsmanMessageCreator, &client)

	t.Run("cim_Package Tests", func(t *testing.T) {
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
				"should create a valid cim_PhysicalPackage Enumerate wsman message",
				CIMPhysicalPackage,
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageEnumerate

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "D8020000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create a valid cim_PhysicalPackage Pull wsman message",
				CIMPhysicalPackage,
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
						EnumerateResponse: common.EnumerateResponse{
							EnumerationContext: "D7020000-0000-0000-0000-000000000000",
						},
						Card: []card.PackageResponse{
							{
								XMLName:           xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_Card", Local: "CIM_Card"},
								CanBeFRUed:        true,
								CreationClassName: "CIM_Card",
								ElementName:       "Managed System Base Board",
								Manufacturer:      "Intel Corporation",
								Model:             "NUC9V7QNB",
								OperationalStatus: []card.OperationalStatus{0},
								PackageType:       9,
								SerialNumber:      "KNQN0221020W",
								Tag:               "CIM_Card",
								Version:           "K47180-402",
							},
						},
						PhysicalPackage: nil,
						Chassis:         nil,
						EndOfSequence:   xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/09/enumeration", Local: "EndOfSequence"},
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

func TestNegativeCIMPackage(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.CIMResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/physical/package",
	}
	elementUnderTest := NewPhysicalPackageWithClient(wsmanMessageCreator, &client)

	t.Run("cim_Package Tests", func(t *testing.T) {
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
				"should create a valid cim_PhysicalPackage Enumerate wsman message",
				CIMPhysicalPackage,
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "D8020000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create a valid cim_PhysicalPackage Pull wsman message",
				CIMPhysicalPackage,
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
						EnumerateResponse: common.EnumerateResponse{
							EnumerationContext: "D7020000-0000-0000-0000-000000000000",
						},
						Card: []card.PackageResponse{
							{
								XMLName:           xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_Card", Local: "CIM_Card"},
								CanBeFRUed:        true,
								CreationClassName: "CIM_Card",
								ElementName:       "Managed System Base Board",
								Manufacturer:      "Intel Corporation",
								Model:             "NUC9V7QNB",
								OperationalStatus: []card.OperationalStatus{0},
								PackageType:       9,
								SerialNumber:      "KNQN0221020W",
								Tag:               "CIM_Card",
								Version:           "K47180-402",
							},
						},
						PhysicalPackage: nil,
						Chassis:         nil,
						EndOfSequence:   xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/09/enumeration", Local: "EndOfSequence"},
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
