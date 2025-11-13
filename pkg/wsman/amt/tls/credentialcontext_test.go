/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package tls

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/wsmantesting"
)

var createCredentialContextResponse = CredentialContextCreateResponse{
	XMLName: xml.Name{
		Space: "http://schemas.xmlsoap.org/ws/2004/09/transfer",
		Local: "ResourceCreated",
	},
	Address: "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous",
	ReferenceParameters: ReferenceParametersResponse{
		XMLName: xml.Name{
			Space: "http://schemas.xmlsoap.org/ws/2004/08/addressing",
			Local: "ReferenceParameters",
		},
		ResourceURI: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TLSCredentialContext",
		SelectorSet: SelectorSetResponse{
			XMLName: xml.Name{
				Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd",
				Local: "SelectorSet",
			},
			Selectors: []SelectorResponse{
				{
					XMLName: xml.Name{
						Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd",
						Local: "Selector",
					},
					Name: "ElementInContext",
					EndpointReference: EndpointReferenceResponse{
						XMLName: xml.Name{
							Space: "http://schemas.xmlsoap.org/ws/2004/08/addressing",
							Local: "EndpointReference",
						},
						Address: "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous",
						ReferenceParameters: ReferenceParametersResponse{
							XMLName: xml.Name{
								Space: "http://schemas.xmlsoap.org/ws/2004/08/addressing",
								Local: "ReferenceParameters",
							},
							ResourceURI: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyCertificate",
							SelectorSet: SelectorSetResponse{
								XMLName: xml.Name{
									Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd",
									Local: "SelectorSet",
								},
								Selectors: []SelectorResponse{
									{
										XMLName: xml.Name{
											Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd",
											Local: "Selector",
										},
										Name: "InstanceID",
										Text: "testCertificate",
									},
								},
							},
						},
					},
				},
				{
					XMLName: xml.Name{
						Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd",
						Local: "Selector",
					},
					Name: "ElementProvidingContext",
					EndpointReference: EndpointReferenceResponse{
						XMLName: xml.Name{
							Space: "http://schemas.xmlsoap.org/ws/2004/08/addressing",
							Local: "EndpointReference",
						},
						Address: "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous",
						ReferenceParameters: ReferenceParametersResponse{
							XMLName: xml.Name{
								Space: "http://schemas.xmlsoap.org/ws/2004/08/addressing",
								Local: "ReferenceParameters",
							},
							ResourceURI: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TLSProtocolEndpointCollection",
							SelectorSet: SelectorSetResponse{
								XMLName: xml.Name{
									Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd",
									Local: "SelectorSet",
								},
								Selectors: []SelectorResponse{
									{
										XMLName: xml.Name{
											Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd",
											Local: "Selector",
										},
										Name: "ElementName",
										Text: "TLSProtocolEndpoint Instances Collection",
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

var putCredentialContextResponse = CredentialContextResponse{
	XMLName: xml.Name{
		Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TLSCredentialContext",
		Local: "AMT_TLSCredentialContext",
	},
	ElementInContext: ElementInContextResponse{
		XMLName: xml.Name{
			Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TLSCredentialContext",
			Local: "ElementInContext",
		},
		Address: "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous",
		ReferenceParameters: ReferenceParametersResponse{
			XMLName: xml.Name{
				Space: "http://schemas.xmlsoap.org/ws/2004/08/addressing",
				Local: "ReferenceParameters",
			},
			ResourceURI: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyCertificate",
			SelectorSet: SelectorSetResponse{
				XMLName: xml.Name{
					Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd",
					Local: "SelectorSet",
				},
				Selectors: []SelectorResponse{
					{
						XMLName: xml.Name{
							Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd",
							Local: "Selector",
						},
						Name: "InstanceID",
						Text: "Intel(r) AMT Certificate: Handle: 4",
					},
				},
			},
		},
	},
	ElementProvidingContext: ElementProvidingContextResponse{
		XMLName: xml.Name{
			Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TLSCredentialContext",
			Local: "ElementProvidingContext",
		},
		Address: "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous",
		ReferenceParameters: ReferenceParametersResponse{
			XMLName: xml.Name{
				Space: "http://schemas.xmlsoap.org/ws/2004/08/addressing",
				Local: "ReferenceParameters",
			},
			ResourceURI: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TLSProtocolEndpointCollection",
			SelectorSet: SelectorSetResponse{
				XMLName: xml.Name{
					Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd",
					Local: "SelectorSet",
				},
				Selectors: []SelectorResponse{
					{
						XMLName: xml.Name{
							Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd",
							Local: "Selector",
						},
						Name: "ElementName",
						Text: "TLSProtocolEndpoint Instances Collection",
					},
				},
			},
		},
	},
}

func TestJson(t *testing.T) {
	response := Response{
		Body: Body{
			CredentialContextGetResponse:    CredentialContextResponse{},
			CredentialContextCreateResponse: createCredentialContextResponse,
		},
	}
	expectedResult := "{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"SettingDataGetAndPutResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"ElementName\":\"\",\"InstanceID\":\"\",\"MutualAuthentication\":false,\"Enabled\":false,\"TrustedCN\":null,\"AcceptNonSecureConnections\":false,\"NonSecureConnectionsSupported\":null},\"CredentialContextGetResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"ElementInContext\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"Address\":\"\",\"ReferenceParameters\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"ResourceURI\":\"\",\"SelectorSet\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"Selectors\":null}}},\"ElementProvidingContext\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"Address\":\"\",\"ReferenceParameters\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"ResourceURI\":\"\",\"SelectorSet\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"Selectors\":null}}}},\"CredentialContextCreateResponse\":{\"XMLName\":{\"Space\":\"http://schemas.xmlsoap.org/ws/2004/09/transfer\",\"Local\":\"ResourceCreated\"},\"Address\":\"http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous\",\"ReferenceParameters\":{\"XMLName\":{\"Space\":\"http://schemas.xmlsoap.org/ws/2004/08/addressing\",\"Local\":\"ReferenceParameters\"},\"ResourceURI\":\"http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TLSCredentialContext\",\"SelectorSet\":{\"XMLName\":{\"Space\":\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\",\"Local\":\"SelectorSet\"},\"Selectors\":[{\"XMLName\":{\"Space\":\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\",\"Local\":\"Selector\"},\"Name\":\"ElementInContext\",\"Text\":\"\",\"EndpointReference\":{\"XMLName\":{\"Space\":\"http://schemas.xmlsoap.org/ws/2004/08/addressing\",\"Local\":\"EndpointReference\"},\"Address\":\"http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous\",\"ReferenceParameters\":{\"XMLName\":{\"Space\":\"http://schemas.xmlsoap.org/ws/2004/08/addressing\",\"Local\":\"ReferenceParameters\"},\"ResourceURI\":\"http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyCertificate\",\"SelectorSet\":{\"XMLName\":{\"Space\":\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\",\"Local\":\"SelectorSet\"},\"Selectors\":[{\"XMLName\":{\"Space\":\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\",\"Local\":\"Selector\"},\"Name\":\"InstanceID\",\"Text\":\"testCertificate\",\"EndpointReference\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"Address\":\"\",\"ReferenceParameters\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"ResourceURI\":\"\",\"SelectorSet\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"Selectors\":null}}}}]}}}},{\"XMLName\":{\"Space\":\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\",\"Local\":\"Selector\"},\"Name\":\"ElementProvidingContext\",\"Text\":\"\",\"EndpointReference\":{\"XMLName\":{\"Space\":\"http://schemas.xmlsoap.org/ws/2004/08/addressing\",\"Local\":\"EndpointReference\"},\"Address\":\"http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous\",\"ReferenceParameters\":{\"XMLName\":{\"Space\":\"http://schemas.xmlsoap.org/ws/2004/08/addressing\",\"Local\":\"ReferenceParameters\"},\"ResourceURI\":\"http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TLSProtocolEndpointCollection\",\"SelectorSet\":{\"XMLName\":{\"Space\":\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\",\"Local\":\"SelectorSet\"},\"Selectors\":[{\"XMLName\":{\"Space\":\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\",\"Local\":\"Selector\"},\"Name\":\"ElementName\",\"Text\":\"TLSProtocolEndpoint Instances Collection\",\"EndpointReference\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"Address\":\"\",\"ReferenceParameters\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"ResourceURI\":\"\",\"SelectorSet\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"Selectors\":null}}}}]}}}}]}}},\"ProtocolEndpointCollectionGetResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"ElementName\":\"\"},\"EnumerateResponse\":{\"EnumerationContext\":\"\"},\"PullResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"SettingDataItems\":null,\"ProtocolEndpointCollectionItems\":null,\"CredentialContextItems\":null}}"
	result := response.JSON()
	assert.Equal(t, expectedResult, result)
}

func TestYaml(t *testing.T) {
	response := Response{
		Body: Body{
			CredentialContextGetResponse:    CredentialContextResponse{},
			CredentialContextCreateResponse: createCredentialContextResponse,
		},
	}
	expectedResult := "xmlname:\n    space: \"\"\n    local: \"\"\nsettingdatagetandputresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    elementname: \"\"\n    instanceid: \"\"\n    mutualauthentication: false\n    enabled: false\n    trustedcn: []\n    acceptnonsecureconnections: false\n    nonsecureconnectionssupported: null\ncredentialcontextgetresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    elementincontext:\n        xmlname:\n            space: \"\"\n            local: \"\"\n        address: \"\"\n        referenceparameters:\n            xmlname:\n                space: \"\"\n                local: \"\"\n            resourceuri: \"\"\n            selectorset:\n                xmlname:\n                    space: \"\"\n                    local: \"\"\n                selectors: []\n    elementprovidingcontext:\n        xmlname:\n            space: \"\"\n            local: \"\"\n        address: \"\"\n        referenceparameters:\n            xmlname:\n                space: \"\"\n                local: \"\"\n            resourceuri: \"\"\n            selectorset:\n                xmlname:\n                    space: \"\"\n                    local: \"\"\n                selectors: []\ncredentialcontextcreateresponse:\n    xmlname:\n        space: http://schemas.xmlsoap.org/ws/2004/09/transfer\n        local: ResourceCreated\n    address: http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous\n    referenceparameters:\n        xmlname:\n            space: http://schemas.xmlsoap.org/ws/2004/08/addressing\n            local: ReferenceParameters\n        resourceuri: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TLSCredentialContext\n        selectorset:\n            xmlname:\n                space: http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\n                local: SelectorSet\n            selectors:\n                - xmlname:\n                    space: http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\n                    local: Selector\n                  name: ElementInContext\n                  text: \"\"\n                  endpointreference:\n                    xmlname:\n                        space: http://schemas.xmlsoap.org/ws/2004/08/addressing\n                        local: EndpointReference\n                    address: http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous\n                    referenceparameters:\n                        xmlname:\n                            space: http://schemas.xmlsoap.org/ws/2004/08/addressing\n                            local: ReferenceParameters\n                        resourceuri: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyCertificate\n                        selectorset:\n                            xmlname:\n                                space: http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\n                                local: SelectorSet\n                            selectors:\n                                - xmlname:\n                                    space: http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\n                                    local: Selector\n                                  name: InstanceID\n                                  text: testCertificate\n                                  endpointreference:\n                                    xmlname:\n                                        space: \"\"\n                                        local: \"\"\n                                    address: \"\"\n                                    referenceparameters:\n                                        xmlname:\n                                            space: \"\"\n                                            local: \"\"\n                                        resourceuri: \"\"\n                                        selectorset:\n                                            xmlname:\n                                                space: \"\"\n                                                local: \"\"\n                                            selectors: []\n                - xmlname:\n                    space: http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\n                    local: Selector\n                  name: ElementProvidingContext\n                  text: \"\"\n                  endpointreference:\n                    xmlname:\n                        space: http://schemas.xmlsoap.org/ws/2004/08/addressing\n                        local: EndpointReference\n                    address: http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous\n                    referenceparameters:\n                        xmlname:\n                            space: http://schemas.xmlsoap.org/ws/2004/08/addressing\n                            local: ReferenceParameters\n                        resourceuri: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TLSProtocolEndpointCollection\n                        selectorset:\n                            xmlname:\n                                space: http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\n                                local: SelectorSet\n                            selectors:\n                                - xmlname:\n                                    space: http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\n                                    local: Selector\n                                  name: ElementName\n                                  text: TLSProtocolEndpoint Instances Collection\n                                  endpointreference:\n                                    xmlname:\n                                        space: \"\"\n                                        local: \"\"\n                                    address: \"\"\n                                    referenceparameters:\n                                        xmlname:\n                                            space: \"\"\n                                            local: \"\"\n                                        resourceuri: \"\"\n                                        selectorset:\n                                            xmlname:\n                                                space: \"\"\n                                                local: \"\"\n                                            selectors: []\nprotocolendpointcollectiongetresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    elementname: \"\"\nenumerateresponse:\n    enumerationcontext: \"\"\npullresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    settingdataitems: []\n    protocolendpointcollectionitems: []\n    credentialcontextitems: []\n"
	result := response.YAML()
	assert.Equal(t, expectedResult, result)
}

func TestPositiveAMT_TLSCredentialContext(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.AMTResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/tls/credentialcontext",
	}
	elementUnderTest := NewTLSCredentialContextWithClient(wsmanMessageCreator, &client)

	t.Run("amt_TLSCredentialContext Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			body             string
			extraHeader      string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			// GETS
			// {
			// 	"should create a valid AMT_TLSCredentialContext Get wsman message",
			// 	AMT_TLSCredentialContext,
			// 	wsmantesting.GET,
			// 	"",
			// 	"",
			// 	func() (Response, error) {
			// 		 client.CurrentMessage = wsmantesting.CurrentMessageGet
			// 		return elementUnderTest.Get()
			// 	},
			// 	Body{},
			// },
			// ENUMERATES
			{
				"should create a valid AMT_TLSCredentialContext Enumerate wsman message",
				AMTTLSCredentialContext,
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageEnumerate

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "6B080000-0000-0000-0000-000000000000",
					},
				},
			},

			// PULLS
			{
				"should create a valid AMT_TLSCredentialContext Pull wsman message",
				AMTTLSCredentialContext,
				wsmantesting.Pull,
				wsmantesting.PullBody,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessagePull

					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						XMLName: xml.Name{Space: message.XMLPullResponseSpace, Local: "PullResponse"},
						CredentialContextItems: []CredentialContextResponse{
							{
								XMLName: xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TLSCredentialContext", Local: "AMT_TLSCredentialContext"},
								ElementInContext: ElementInContextResponse{
									XMLName: xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TLSCredentialContext", Local: "ElementInContext"},
									Address: "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous",
									ReferenceParameters: ReferenceParametersResponse{
										XMLName:     xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/08/addressing", Local: "ReferenceParameters"},
										ResourceURI: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyCertificate",
										SelectorSet: SelectorSetResponse{
											XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "SelectorSet"},
											Selectors: []SelectorResponse{
												{
													XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
													Name:    "InstanceID",
													Text:    "Intel(r) AMT Certificate: Handle: 3",
												},
											},
										},
									},
								},
								ElementProvidingContext: ElementProvidingContextResponse{
									XMLName: xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TLSCredentialContext", Local: "ElementProvidingContext"},
									Address: "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous",
									ReferenceParameters: ReferenceParametersResponse{
										XMLName:     xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/08/addressing", Local: "ReferenceParameters"},
										ResourceURI: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TLSProtocolEndpointCollection",
										SelectorSet: SelectorSetResponse{
											XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "SelectorSet"},
											Selectors: []SelectorResponse{
												{
													XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
													Name:    "ElementName",
													Text:    "TLSProtocolEndpoint Instances Collection",
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
			// DELETE
			// {
			// 	"should create a valid AMT_TLSCredentialContext Delete wsman message",
			// 	AMT_TLSCredentialContext,
			// 	wsmantesting.DELETE,
			// 	"",
			// 	"<w:SelectorSet><w:Selector Name=\"Name\">instanceID123</w:Selector></w:SelectorSet>",
			// 	func() (Response, error) {
			// 		 client.CurrentMessage = wsmantesting.CurrentMessageDelete
			// 		return elementUnderTest.Delete("instanceID123")
			// 	},
			// 	Body{},
			// },
			// Create
			{
				"should create a valid AMT_TLSCredentialContext Create wsman message",
				AMTTLSCredentialContext,
				wsmantesting.Create,
				"<h:AMT_TLSCredentialContext xmlns:h=\"http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TLSCredentialContext\"><h:ElementInContext><a:Address>/wsman</a:Address><a:ReferenceParameters><w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyCertificate</w:ResourceURI><w:SelectorSet><w:Selector Name=\"InstanceID\">testCertificate</w:Selector></w:SelectorSet></a:ReferenceParameters></h:ElementInContext><h:ElementProvidingContext><a:Address>/wsman</a:Address><a:ReferenceParameters><w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TLSProtocolEndpointCollection</w:ResourceURI><w:SelectorSet><w:Selector Name=\"ElementName\">TLSProtocolEndpointInstances Collection</w:Selector></w:SelectorSet></a:ReferenceParameters></h:ElementProvidingContext></h:AMT_TLSCredentialContext>",
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageCreate

					return elementUnderTest.Create("testCertificate")
				},
				Body{
					XMLName: xml.Name{
						Space: "http://www.w3.org/2003/05/soap-envelope",
						Local: "Body",
					},
					CredentialContextCreateResponse: createCredentialContextResponse,
				},
			},
			// Put
			{
				"should create a valid AMT_TLSCredentialContext Put wsman message",
				AMTTLSCredentialContext,
				wsmantesting.Put,
				fmt.Sprintf(`<h:AMT_TLSCredentialContext xmlns:h="%sAMT_TLSCredentialContext"><h:ElementInContext><a:Address>/wsman</a:Address><a:ReferenceParameters><w:ResourceURI>%sAMT_PublicKeyCertificate</w:ResourceURI><w:SelectorSet><w:Selector Name="InstanceID">%s</w:Selector></w:SelectorSet></a:ReferenceParameters></h:ElementInContext><h:ElementProvidingContext><a:Address>/wsman</a:Address><a:ReferenceParameters><w:ResourceURI>%sAMT_TLSProtocolEndpointCollection</w:ResourceURI><w:SelectorSet><w:Selector Name="ElementName">TLSProtocolEndpointInstances Collection</w:Selector></w:SelectorSet></a:ReferenceParameters></h:ElementProvidingContext></h:AMT_TLSCredentialContext>`, "http://intel.com/wbem/wscim/1/amt-schema/1/", "http://intel.com/wbem/wscim/1/amt-schema/1/", "testCertificate", "http://intel.com/wbem/wscim/1/amt-schema/1/"),
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessagePut

					return elementUnderTest.Put("testCertificate")
				},
				Body{
					XMLName: xml.Name{
						Space: "http://www.w3.org/2003/05/soap-envelope",
						Local: "Body",
					},
					CredentialContextGetResponse: putCredentialContextResponse,
				},
			},
		}
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				expectedXMLInput := wsmantesting.ExpectedResponse(messageID, resourceURIBase, test.method, test.action, test.extraHeader, test.body)
				messageID++
				response, err := test.responseFunc()
				assert.NoError(t, err)
				assert.Equal(t, expectedXMLInput, response.XMLInput)
				assert.Equal(t, test.expectedResponse, response.Body)
			})
		}
	})
}
