/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package credential

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/models"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/wsmantesting"
)

func TestJson(t *testing.T) {
	response := Response{
		Body: Body{
			PullResponse: PullResponse{},
		},
	}
	expectedResult := "{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"PullResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"Items\":{\"CredentialContext\":null,\"CredentialContextTLS\":null,\"CredentialContext8021x\":null},\"EndOfSequence\":{\"Space\":\"\",\"Local\":\"\"}},\"EnumerateResponse\":{\"EnumerationContext\":\"\"}}"
	result := response.JSON()
	assert.Equal(t, expectedResult, result)
}

func TestYaml(t *testing.T) {
	response := Response{
		Body: Body{
			PullResponse: PullResponse{},
		},
	}
	expectedResult := "xmlname:\n    space: \"\"\n    local: \"\"\npullresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    items:\n        credentialcontext: []\n        credentialcontexttls: []\n        credentialcontext8021x: []\n    endofsequence:\n        space: \"\"\n        local: \"\"\nenumerateresponse:\n    enumerationcontext: \"\"\n"
	result := response.YAML()
	assert.Equal(t, expectedResult, result)
}

func TestPositiveCIMCredentialContext(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.CIMResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/credential/context",
	}
	elementUnderTest := NewContextWithClient(wsmanMessageCreator, &client)

	t.Run("cim_* Tests", func(t *testing.T) {
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
				"should create and parse a valid cim_IEEE8021xSettings Enumerate call",
				CIMCredentialContext,
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageEnumerate

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "19000000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create and parse a valid cim_IEEE8021xSettings Pull call",
				CIMCredentialContext,
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
						Items: Items{
							[]CredentialContext{
								{
									ElementInContext: models.AssociationReference{
										Address: "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous",
										ReferenceParameters: models.ReferenceParametersNoNamespace{
											XMLName:     xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/08/addressing", Local: "ReferenceParameters"},
											ResourceURI: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyCertificate",
											SelectorSet: models.SelectorNoNamespace{
												XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "SelectorSet"},
												Selectors: []models.SelectorResponse{
													{
														XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
														Name:    "InstanceID",
														Text:    "Intel(r) AMT Certificate: Handle: 0",
													},
												},
											},
										},
									},
									ElementProvidingContext: models.AssociationReference{
										Address: "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous",
										ReferenceParameters: models.ReferenceParametersNoNamespace{
											XMLName:     xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/08/addressing", Local: "ReferenceParameters"},
											ResourceURI: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_IEEE8021xSettings",
											SelectorSet: models.SelectorNoNamespace{
												XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "SelectorSet"},
												Selectors: []models.SelectorResponse{
													{
														XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
														Name:    "InstanceID",
														Text:    "Intel(r) AMT:IEEE 802.1x Settings",
													},
												},
											},
										},
									},
								},
							},
							[]CredentialContext(nil),
							[]CredentialContext(nil),
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

func TestNegativeCIMCredentialContext(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.CIMResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/credential/context",
	}
	elementUnderTest := NewContextWithClient(wsmanMessageCreator, &client)

	t.Run("cim_* Tests", func(t *testing.T) {
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
				"should handle error when cim_IEEE8021xSettings Enumerate call",
				CIMCredentialContext,
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "19000000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should handle error when cim_IEEE8021xSettings Pull call",
				CIMCredentialContext,
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
						Items: Items{
							[]CredentialContext{
								{
									ElementInContext: models.AssociationReference{
										Address: "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous",
										ReferenceParameters: models.ReferenceParametersNoNamespace{
											ResourceURI: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyCertificate",
											SelectorSet: models.SelectorNoNamespace{
												Selectors: []models.SelectorResponse{
													{
														Name: "InstanceID",
														Text: "Intel(r) AMT Certificate: Handle: 0",
													},
												},
											},
										},
									},
									ElementProvidingContext: models.AssociationReference{
										Address: "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous",
										ReferenceParameters: models.ReferenceParametersNoNamespace{
											ResourceURI: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TLSProtocolEndpointCollection",
											SelectorSet: models.SelectorNoNamespace{
												Selectors: []models.SelectorResponse{
													{
														Name: "ElementName",
														Text: "TLSProtocolEndpoint Instances Collection",
													},
												},
											},
										},
									},
								},
							},
							[]CredentialContext(nil),
							[]CredentialContext(nil),
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
