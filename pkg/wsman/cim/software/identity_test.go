/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package software

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
	expectedResult := "{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"PullResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"SoftwareIdentityItems\":null},\"EnumerateResponse\":{\"EnumerationContext\":\"\"},\"SoftwareIdentityResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"InstanceID\":\"\",\"VersionString\":\"\",\"IsEntity\":false}}"
	result := response.JSON()
	assert.Equal(t, expectedResult, result)
}

func TestYaml(t *testing.T) {
	response := Response{
		Body: Body{
			PullResponse: PullResponse{},
		},
	}
	expectedResult := "xmlname:\n    space: \"\"\n    local: \"\"\npullresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    softwareidentityitems: []\nenumerateresponse:\n    enumerationcontext: \"\"\nsoftwareidentityresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    instanceid: \"\"\n    versionstring: \"\"\n    isentity: false\n"
	result := response.YAML()
	assert.Equal(t, expectedResult, result)
}

func TestPositiveCIMSoftwareIdentity(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.CIMResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/software/identity",
	}
	elementUnderTest := NewSoftwareIdentityWithClient(wsmanMessageCreator, &client)

	t.Run("cim_SoftwareIdentity Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			extraHeaders     string
			body             string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			// GETS
			{
				"should create and parse a valid cim_SoftwareIdentity Get call",
				CIMSoftwareIdentity,
				wsmantesting.Get,
				"<w:SelectorSet><w:Selector Name=\"InstanceID\">AMTApps</w:Selector></w:SelectorSet>",
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageGet

					return elementUnderTest.GetByInstanceID("AMTApps")
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					SoftwareIdentityResponse: SoftwareIdentity{
						XMLName:       xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_SoftwareIdentity", Local: "CIM_SoftwareIdentity"},
						InstanceID:    "AMTApps",
						IsEntity:      true,
						VersionString: "12.0.67",
					},
				},
			},
			// ENUMERATES
			{
				"should create and parse a valid cim_SoftwareIdentity Enumerate call",
				CIMSoftwareIdentity,
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
						EnumerationContext: "E2020000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create and parse a valid cim_SoftwareIdentity Pull call",
				CIMSoftwareIdentity,
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
						SoftwareIdentityItems: []SoftwareIdentity{
							{
								XMLName:       xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_SoftwareIdentity", Local: "CIM_SoftwareIdentity"},
								InstanceID:    "Flash",
								IsEntity:      true,
								VersionString: "12.0.67",
							},
							{
								XMLName:       xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_SoftwareIdentity", Local: "CIM_SoftwareIdentity"},
								InstanceID:    "Netstack",
								IsEntity:      true,
								VersionString: "12.0.67",
							},
							{
								XMLName:       xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_SoftwareIdentity", Local: "CIM_SoftwareIdentity"},
								InstanceID:    "AMTApps",
								IsEntity:      true,
								VersionString: "12.0.67",
							},
							{
								XMLName:       xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_SoftwareIdentity", Local: "CIM_SoftwareIdentity"},
								InstanceID:    "AMT",
								IsEntity:      true,
								VersionString: "12.0.67",
							},
							{
								XMLName:       xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_SoftwareIdentity", Local: "CIM_SoftwareIdentity"},
								InstanceID:    "Sku",
								IsEntity:      true,
								VersionString: "16392",
							},
							{
								XMLName:       xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_SoftwareIdentity", Local: "CIM_SoftwareIdentity"},
								InstanceID:    "VendorID",
								IsEntity:      true,
								VersionString: "8086",
							},
							{
								XMLName:       xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_SoftwareIdentity", Local: "CIM_SoftwareIdentity"},
								InstanceID:    "Build Number",
								IsEntity:      true,
								VersionString: "1579",
							},
							{
								XMLName:       xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_SoftwareIdentity", Local: "CIM_SoftwareIdentity"},
								InstanceID:    "Recovery Version",
								IsEntity:      true,
								VersionString: "12.0.67",
							},
							{
								XMLName:       xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_SoftwareIdentity", Local: "CIM_SoftwareIdentity"},
								InstanceID:    "Recovery Build Num",
								IsEntity:      true,
								VersionString: "1579",
							},
							{
								XMLName:       xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_SoftwareIdentity", Local: "CIM_SoftwareIdentity"},
								InstanceID:    "Legacy Mode",
								IsEntity:      true,
								VersionString: "False",
							},
							{
								XMLName:       xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_SoftwareIdentity", Local: "CIM_SoftwareIdentity"},
								InstanceID:    "AMT FW Core Version",
								IsEntity:      true,
								VersionString: "12.0.67",
							},
						},
					},
				},
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				expectedXMLInput := wsmantesting.ExpectedResponse(messageID, resourceURIBase, test.method, test.action, test.extraHeaders, test.body)
				messageID++
				response, err := test.responseFunc()
				assert.NoError(t, err)
				assert.Equal(t, expectedXMLInput, response.XMLInput)
				assert.Equal(t, test.expectedResponse, response.Body)
			})
		}
	})
}

func TestNegativeCIMSoftwareIdentity(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.CIMResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/software/identity",
	}
	elementUnderTest := NewSoftwareIdentityWithClient(wsmanMessageCreator, &client)

	t.Run("cim_SoftwareIdentity Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			extraHeaders     string
			body             string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			// GETS
			{
				"should create and parse a valid cim_SoftwareIdentity Get call",
				CIMSoftwareIdentity,
				wsmantesting.Get,
				"<w:SelectorSet><w:Selector Name=\"InstanceID\">AMTApps</w:Selector></w:SelectorSet>",
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.GetByInstanceID("AMTApps")
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					SoftwareIdentityResponse: SoftwareIdentity{
						XMLName:       xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_SoftwareIdentity", Local: "CIM_SoftwareIdentity"},
						InstanceID:    "AMTApps",
						IsEntity:      true,
						VersionString: "12.0.67",
					},
				},
			},
			// ENUMERATES
			{
				"should create and parse a valid cim_SoftwareIdentity Enumerate call",
				CIMSoftwareIdentity,
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
						EnumerationContext: "E2020000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create and parse a valid cim_SoftwareIdentity Pull call",
				CIMSoftwareIdentity,
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
						SoftwareIdentityItems: []SoftwareIdentity{
							{
								XMLName:       xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_SoftwareIdentity", Local: "CIM_SoftwareIdentity"},
								InstanceID:    "Flash",
								IsEntity:      true,
								VersionString: "12.0.67",
							},
							{
								XMLName:       xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_SoftwareIdentity", Local: "CIM_SoftwareIdentity"},
								InstanceID:    "Netstack",
								IsEntity:      true,
								VersionString: "12.0.67",
							},
							{
								XMLName:       xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_SoftwareIdentity", Local: "CIM_SoftwareIdentity"},
								InstanceID:    "AMTApps",
								IsEntity:      true,
								VersionString: "12.0.67",
							},
							{
								XMLName:       xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_SoftwareIdentity", Local: "CIM_SoftwareIdentity"},
								InstanceID:    "AMT",
								IsEntity:      true,
								VersionString: "12.0.67",
							},
							{
								XMLName:       xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_SoftwareIdentity", Local: "CIM_SoftwareIdentity"},
								InstanceID:    "Sku",
								IsEntity:      true,
								VersionString: "16392",
							},
							{
								XMLName:       xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_SoftwareIdentity", Local: "CIM_SoftwareIdentity"},
								InstanceID:    "VendorID",
								IsEntity:      true,
								VersionString: "8086",
							},
							{
								XMLName:       xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_SoftwareIdentity", Local: "CIM_SoftwareIdentity"},
								InstanceID:    "Build Number",
								IsEntity:      true,
								VersionString: "1579",
							},
							{
								XMLName:       xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_SoftwareIdentity", Local: "CIM_SoftwareIdentity"},
								InstanceID:    "Recovery Version",
								IsEntity:      true,
								VersionString: "12.0.67",
							},
							{
								XMLName:       xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_SoftwareIdentity", Local: "CIM_SoftwareIdentity"},
								InstanceID:    "Recovery Build Num",
								IsEntity:      true,
								VersionString: "1579",
							},
							{
								XMLName:       xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_SoftwareIdentity", Local: "CIM_SoftwareIdentity"},
								InstanceID:    "Legacy Mode",
								IsEntity:      true,
								VersionString: "False",
							},
							{
								XMLName:       xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_SoftwareIdentity", Local: "CIM_SoftwareIdentity"},
								InstanceID:    "AMT FW Core Version",
								IsEntity:      true,
								VersionString: "12.0.67",
							},
						},
					},
				},
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				expectedXMLInput := wsmantesting.ExpectedResponse(messageID, resourceURIBase, test.method, test.action, test.extraHeaders, test.body)
				messageID++
				response, err := test.responseFunc()
				assert.Error(t, err)
				assert.Equal(t, expectedXMLInput, response.XMLInput)
				assert.NotEqual(t, test.expectedResponse, response.Body)
			})
		}
	})
}
