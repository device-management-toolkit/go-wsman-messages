/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package mediaaccess

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
	expectedResult := "{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"PullResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"MediaAccessDevices\":null},\"EnumerateResponse\":{\"EnumerationContext\":\"\"}}"
	result := response.JSON()
	assert.Equal(t, expectedResult, result)
}

func TestYaml(t *testing.T) {
	response := Response{
		Body: Body{
			PullResponse: PullResponse{},
		},
	}
	expectedResult := "xmlname:\n    space: \"\"\n    local: \"\"\npullresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    mediaaccessdevices: []\nenumerateresponse:\n    enumerationcontext: \"\"\n"
	result := response.YAML()
	assert.Equal(t, expectedResult, result)
}

func TestPositiveCIMMediaAccessDevice(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.CIMResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/mediaaccess",
	}
	elementUnderTest := NewMediaAccessDeviceWithClient(wsmanMessageCreator, &client)

	t.Run("cim_MediaAccessDevice Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			extraHeaders     string
			body             string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			// ENUMERATES
			{
				"should create and parse a valid cim_MediaAccessDevice Enumerate call",
				CIMMediaAccessDevice,
				wsmantesting.Enumerate,
				"",
				wsmantesting.EnumerateBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageEnumerate

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "CE020000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create and parse a valid cim_MediaAccessDevice Pull call",
				CIMMediaAccessDevice,
				wsmantesting.Pull,
				"",
				wsmantesting.PullBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessagePull

					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						XMLName: xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/09/enumeration", Local: "PullResponse"},
						MediaAccessDevices: []MediaAccessDevice{
							{
								Capabilities:            []Capabilities{4},
								CreationClassName:       "CIM_MediaAccessDevice",
								DeviceID:                "MEDIA DEV 0",
								ElementName:             "Managed System Media Access Device",
								EnabledDefault:          2,
								EnabledState:            0,
								MaxMediaSize:            960197124,
								OperationalStatus:       []OperationalStatus{0},
								RequestedState:          12,
								Security:                2,
								SystemCreationClassName: "CIM_ComputerSystem",
								SystemName:              "ManagedSystem",
							},
							{
								Capabilities:            []Capabilities{4},
								CreationClassName:       "CIM_MediaAccessDevice",
								DeviceID:                "MEDIA DEV 1",
								ElementName:             "Managed System Media Access Device",
								EnabledDefault:          2,
								EnabledState:            0,
								MaxMediaSize:            500107862,
								OperationalStatus:       []OperationalStatus{0},
								RequestedState:          12,
								Security:                2,
								SystemCreationClassName: "CIM_ComputerSystem",
								SystemName:              "ManagedSystem",
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

func TestNegativeCIMMediaAccessDevice(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.CIMResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/mediaaccess",
	}
	elementUnderTest := NewMediaAccessDeviceWithClient(wsmanMessageCreator, &client)

	t.Run("cim_MediaAccessDevice Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			extraHeaders     string
			body             string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			// ENUMERATES
			{
				"should handle error when cim_MediaAccessDevice Enumerate call",
				CIMMediaAccessDevice,
				wsmantesting.Enumerate,
				"",
				wsmantesting.EnumerateBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "CE020000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should handle error when cim_MediaAccessDevice Pull call",
				CIMMediaAccessDevice,
				wsmantesting.Pull,
				"",
				wsmantesting.PullBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						XMLName: xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/09/enumeration", Local: "PullResponse"},
						MediaAccessDevices: []MediaAccessDevice{
							{
								Capabilities:            []Capabilities{4},
								CreationClassName:       "CIM_MediaAccessDevice",
								DeviceID:                "MEDIA DEV 0",
								ElementName:             "Managed System Media Access Device",
								EnabledDefault:          2,
								EnabledState:            0,
								MaxMediaSize:            960197124,
								OperationalStatus:       []OperationalStatus{0},
								RequestedState:          12,
								Security:                2,
								SystemCreationClassName: "CIM_ComputerSystem",
								SystemName:              "ManagedSystem",
							},
							{
								Capabilities:            []Capabilities{4},
								CreationClassName:       "CIM_MediaAccessDevice",
								DeviceID:                "MEDIA DEV 1",
								ElementName:             "Managed System Media Access Device",
								EnabledDefault:          2,
								EnabledState:            0,
								MaxMediaSize:            500107862,
								OperationalStatus:       []OperationalStatus{0},
								RequestedState:          12,
								Security:                2,
								SystemCreationClassName: "CIM_ComputerSystem",
								SystemName:              "ManagedSystem",
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
