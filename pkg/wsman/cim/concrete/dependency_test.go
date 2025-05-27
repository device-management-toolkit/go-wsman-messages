/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package concrete

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
	expectedResult := "{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"PullResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"Items\":null},\"EnumerateResponse\":{\"EnumerationContext\":\"\"}}"
	result := response.JSON()
	assert.Equal(t, expectedResult, result)
}

func TestYaml(t *testing.T) {
	response := Response{
		Body: Body{
			PullResponse: PullResponse{},
		},
	}
	expectedResult := "xmlname:\n    space: \"\"\n    local: \"\"\npullresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    items: []\nenumerateresponse:\n    enumerationcontext: \"\"\n"
	result := response.YAML()
	assert.Equal(t, expectedResult, result)
}

func TestPositiveCIMConcreteDependency(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.CIMResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/concrete/dependency",
	}
	elementUnderTest := NewDependencyWithClient(wsmanMessageCreator, &client)

	t.Run("cim_Concrete Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			extraHeader      string
			body             string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			// ENUMERATES
			{
				"should create and parse a valid cim_Concrete Enumerate call",
				CIMConcreteDependency,
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
						EnumerationContext: "AC070000-0000-0000-0000-000000000000",
					},
				},
			},

			// PULLS
			{
				"should create and parse a valid cim_Concrete Pull call",
				CIMConcreteDependency,
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
						Items: []ConcreteDependency{
							{
								Antecedent: models.AssociationReference{
									Address: "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous",
									ReferenceParameters: models.ReferenceParametersNoNamespace{
										XMLName:     xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/08/addressing", Local: "ReferenceParameters"},
										ResourceURI: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_Identity",
										SelectorSet: models.SelectorNoNamespace{
											XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "SelectorSet"},
											Selectors: []models.SelectorResponse{
												{
													XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
													Name:    "InstanceID",
													Text:    "Intel(r) AMT:$$OsAdmin",
												},
											},
										},
									},
								},
								Dependent: models.AssociationReference{
									Address: "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous",
									ReferenceParameters: models.ReferenceParametersNoNamespace{
										XMLName:     xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/08/addressing", Local: "ReferenceParameters"},
										ResourceURI: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_Role",
										SelectorSet: models.SelectorNoNamespace{
											XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "SelectorSet"},
											Selectors: []models.SelectorResponse{
												{
													XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
													Name:    "CreationClassName",
													Text:    "CIM_Role",
												},
												{
													XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
													Name:    "Name",
													Text:    "Intel(r) AMT:$$OsAdmin",
												},
											},
										},
									},
								},
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

func TestNegativeCIMConcreteDependency(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.CIMResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/concrete/dependency",
	}
	elementUnderTest := NewDependencyWithClient(wsmanMessageCreator, &client)

	t.Run("cim_Concrete Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			extraHeader      string
			body             string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			// ENUMERATES
			{
				"should handle error when cim_ConcreteDependency Enumerate call",
				CIMConcreteDependency,
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
						EnumerationContext: "AC070000-0000-0000-0000-000000000000",
					},
				},
			},

			// PULLS
			{
				"should handle error when cim_ConcreteDependency Pull call",
				CIMConcreteDependency,
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
						Items: []ConcreteDependency{
							{
								Antecedent: models.AssociationReference{
									Address: "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous",
									ReferenceParameters: models.ReferenceParametersNoNamespace{
										ResourceURI: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_Identity",
										SelectorSet: models.SelectorNoNamespace{
											Selectors: []models.SelectorResponse{
												{
													Name: "InstanceID",
													Text: "Intel(r) AMT:$$OsAdmin",
												},
											},
										},
									},
								},
								Dependent: models.AssociationReference{
									Address: "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous",
									ReferenceParameters: models.ReferenceParametersNoNamespace{
										ResourceURI: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_Role",
										SelectorSet: models.SelectorNoNamespace{
											Selectors: []models.SelectorResponse{
												{
													Name: "CreationClassName",
													Text: "CIM_Role",
												},
												{
													Name: "Name",
													Text: "Intel(r) AMT:$$OsAdmin",
												},
											},
										},
									},
								},
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
