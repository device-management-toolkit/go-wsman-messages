/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package http

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/models"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/wsmantesting"
)

func TestJson(t *testing.T) {
	response := Response{
		Body: Body{
			PullResponse: PullResponse{},
		},
	}
	expectedResult := "{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"EnumerateResponse\":{\"EnumerationContext\":\"\"},\"GetAndPutResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"Name\":\"\",\"CreationClassName\":\"\",\"SystemName\":\"\",\"SystemCreationClassName\":\"\",\"ElementName\":\"\",\"SyncEnabled\":false},\"PullResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"Items\":null},\"AddProxyAccessPointResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"ProxyAccessPoint\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"Address\":\"\",\"ReferenceParameters\":{\"ResourceURI\":\"\",\"SelectorSet\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"Selector\":null}}},\"ReturnValue\":0}}"
	result := response.JSON()
	assert.Equal(t, expectedResult, result)
}

func TestYaml(t *testing.T) {
	response := Response{
		Body: Body{
			PullResponse: PullResponse{},
		},
	}
	expectedResult := "xmlname:\n    space: \"\"\n    local: \"\"\nenumerateresponse:\n    enumerationcontext: \"\"\ngetandputresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    name: \"\"\n    creationclassname: \"\"\n    systemname: \"\"\n    systemcreationclassname: \"\"\n    elementname: \"\"\n    syncenabled: false\npullresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    items: []\naddproxyaccesspointresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    proxyaccesspoint:\n        xmlname:\n            space: \"\"\n            local: \"\"\n        address: \"\"\n        referenceparameters:\n            resourceuri: \"\"\n            selectorset:\n                xmlname:\n                    space: \"\"\n                    local: \"\"\n                selector: []\n    returnvalue: 0\n"
	result := response.YAML()
	assert.Equal(t, expectedResult, result)
}

func TestPositiveIPS_HTTPProxyService(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.IPSResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "ips/http",
	}
	elementUnderTest := NewHTTPProxyServiceWithClient(wsmanMessageCreator, &client)

	t.Run("ips_HTTPProxyService Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			body             string
			extraHeader      string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			// AddProxyAccessPoint
			{
				"should create a valid AddProxyAccessPoint wsman message",
				IPSHTTPProxyService,
				wsmantesting.AddProxyAccessPoint,
				`<h:AddProxyAccessPoint_INPUT xmlns:h="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HTTPProxyService"><h:AccessInfo>proxy.example.com</h:AccessInfo><h:InfoFormat>201</h:InfoFormat><h:Port>8080</h:Port><h:NetworkDnsSuffix>example.com</h:NetworkDnsSuffix></h:AddProxyAccessPoint_INPUT>`,
				"",
				func() (Response, error) {
					client.CurrentMessage = "AddProxyAccessPoint"

					return elementUnderTest.AddProxyAccessPoint("proxy.example.com", InfoFormatFQDN, 8080, "example.com")
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					AddProxyAccessPointResponse: AddProxyAccessPoint_OUTPUT{
						XMLName: xml.Name{Space: fmt.Sprintf("%s%s", message.IPSSchema, IPSHTTPProxyService), Local: "AddProxyAccessPoint_OUTPUT"},
						ProxyAccessPoint: ProxyAccessPoint{
							XMLName: xml.Name{Space: fmt.Sprintf("%s%s", message.IPSSchema, IPSHTTPProxyService), Local: "ProxyAccessPoint"},
							Address: "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous",
							ReferenceParameters: models.ReferenceParameters_OUTPUT{
								ResourceURI: "http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HTTPProxyAccessPoint",
								SelectorSet: models.SelectorSet_OUTPUT{
									XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "SelectorSet"},
									Selector: []message.Selector_OUTPUT{
										{XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"}, Name: "CreationClassName", Value: "IPS_HTTPProxyAccessPoint"},
										{XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"}, Name: "Name", Value: "Intel(r) ME:HTTP Proxy Access Point 3"},
										{XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"}, Name: "SystemCreationClassName", Value: "CIM_ComputerSystem"},
										{XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"}, Name: "SystemName", Value: "Intel(r) AMT"},
									},
								},
							},
						},
						ReturnValue: 0,
					},
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
