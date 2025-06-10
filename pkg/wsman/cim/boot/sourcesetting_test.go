/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package boot

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/wsmantesting"
)

func TestPositiveSourceSetting(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.CIMResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/boot/sourcesetting",
	}
	elementUnderTest := NewBootSourceSettingWithClient(wsmanMessageCreator, &client)

	t.Run("cim_* Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			extraHeader      string
			body             string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			// GETS
			{
				"should create and parse a valid cim_BootSourceSetting Get call",
				CIMBootSourceSetting,
				wsmantesting.Get,
				"<w:SelectorSet><w:Selector Name=\"InstanceID\">Intel(r) AMT: Force Hard-drive Boot</w:Selector></w:SelectorSet>",
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageGet

					return elementUnderTest.GetByInstanceID("Intel(r) AMT: Force Hard-drive Boot")
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					SourceSettingGetResponse: BootSourceSetting{
						XMLName:              xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting", Local: CIMBootSourceSetting},
						ElementName:          "Intel(r) AMT: Boot Source",
						InstanceID:           "Intel(r) AMT: Force Hard-drive Boot",
						StructuredBootString: "CIM:Hard-Disk:1",
						FailThroughSupported: 2,
					},
				},
			},
			// ENUMERATES
			{
				"should create and parse a valid cim_BootSourceSetting Enumerate call",
				CIMBootSourceSetting,
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
						EnumerationContext: "14000000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create and parse a valid cim_BootSourceSetting Pull call",
				CIMBootSourceSetting,
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
						BootSourceSettingItems: []BootSourceSetting{
							{
								XMLName:              xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting", Local: CIMBootSourceSetting},
								InstanceID:           "Intel(r) AMT: Force Hard-drive Boot",
								ElementName:          "Intel(r) AMT: Boot Source",
								StructuredBootString: "CIM:Hard-Disk:1",
								FailThroughSupported: 2,
							},
							{
								XMLName:              xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting", Local: CIMBootSourceSetting},
								InstanceID:           "Intel(r) AMT: Force PXE Boot",
								ElementName:          "Intel(r) AMT: Boot Source",
								StructuredBootString: "CIM:Network:1",
								FailThroughSupported: 2,
							},
							{
								XMLName:              xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting", Local: CIMBootSourceSetting},
								InstanceID:           "Intel(r) AMT: Force CD/DVD Boot",
								ElementName:          "Intel(r) AMT: Boot Source",
								StructuredBootString: "CIM:CD/DVD:1",
								FailThroughSupported: 2,
							},
						},
					},
				},
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				expectedResponse := wsmantesting.ExpectedResponse(messageID, resourceURIBase, test.method, test.action, test.extraHeader, test.body)
				messageID++
				response, err := test.responseFunc()
				assert.NoError(t, err)
				assert.Equal(t, expectedResponse, response.XMLInput)
				assert.Equal(t, test.expectedResponse, response.Body)
			})
		}
	})
}

func TestNegativeSourceSetting(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.CIMResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/boot/sourcesetting",
	}
	elementUnderTest := NewBootSourceSettingWithClient(wsmanMessageCreator, &client)

	t.Run("cim_* Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			extraHeader      string
			body             string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			// GETS
			{
				"should handle error when cim_BootSourceSetting Get call",
				CIMBootSourceSetting,
				wsmantesting.Get,
				"<w:SelectorSet><w:Selector Name=\"InstanceID\">Intel(r) AMT: Force Hard-drive Boot</w:Selector></w:SelectorSet>",
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.GetByInstanceID("Intel(r) AMT: Force Hard-drive Boot")
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					SourceSettingGetResponse: BootSourceSetting{
						XMLName:              xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting", Local: CIMBootSourceSetting},
						ElementName:          "Intel(r) AMT: Boot Source",
						InstanceID:           "Intel(r) AMT: Force Hard-drive Boot",
						StructuredBootString: "CIM:Hard-Disk:1",
						FailThroughSupported: 2,
					},
				},
			},
			// ENUMERATES
			{
				"should handle error when cim_BootSourceSetting Enumerate call",
				CIMBootSourceSetting,
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
						EnumerationContext: "14000000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should handle error when cim_BootSourceSetting Pull call",
				CIMBootSourceSetting,
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
						BootSourceSettingItems: []BootSourceSetting{
							{
								XMLName:              xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting", Local: CIMBootSourceSetting},
								InstanceID:           "Intel(r) AMT: Force Hard-drive Boot",
								ElementName:          "Intel(r) AMT: Boot Source",
								StructuredBootString: "CIM:Hard-Disk:1",
								FailThroughSupported: 2,
							},
							{
								XMLName:              xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting", Local: CIMBootSourceSetting},
								InstanceID:           "Intel(r) AMT: Force PXE Boot",
								ElementName:          "Intel(r) AMT: Boot Source",
								StructuredBootString: "CIM:Network:1",
								FailThroughSupported: 2,
							},
							{
								XMLName:              xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting", Local: CIMBootSourceSetting},
								InstanceID:           "Intel(r) AMT: Force CD/DVD Boot",
								ElementName:          "Intel(r) AMT: Boot Source",
								StructuredBootString: "CIM:CD/DVD:1",
								FailThroughSupported: 2,
							},
						},
					},
				},
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				expectedResponse := wsmantesting.ExpectedResponse(messageID, resourceURIBase, test.method, test.action, test.extraHeader, test.body)
				messageID++
				response, err := test.responseFunc()
				assert.Error(t, err)
				assert.Equal(t, expectedResponse, response.XMLInput)
				assert.NotEqual(t, test.expectedResponse, response.Body)
			})
		}
	})
}
