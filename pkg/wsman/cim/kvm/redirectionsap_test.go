/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package kvm

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
	expectedResult := "{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"GetResponse\":{\"CreationClassName\":\"\",\"ElementName\":\"\",\"Name\":\"\",\"SystemCreationClassName\":\"\",\"SystemName\":\"\",\"EnabledState\":0,\"RequestedState\":0,\"KVMProtocol\":0},\"RequestStateChange_OUTPUT\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"ReturnValue\":0},\"EnumerateResponse\":{\"EnumerationContext\":\"\"},\"PullResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"Items\":null}}"
	result := response.JSON()
	assert.Equal(t, expectedResult, result)
}

func TestYaml(t *testing.T) {
	response := Response{
		Body: Body{
			PullResponse: PullResponse{},
		},
	}
	expectedResult := "xmlname:\n    space: \"\"\n    local: \"\"\ngetresponse:\n    creationclassname: \"\"\n    elementname: \"\"\n    name: \"\"\n    systemcreationclassname: \"\"\n    systemname: \"\"\n    enabledstate: 0\n    requestedstate: 0\n    kvmprotocol: 0\nrequeststatechange_output:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    returnvalue: 0\nenumerateresponse:\n    enumerationcontext: \"\"\npullresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    items: []\n"
	result := response.YAML()
	assert.Equal(t, expectedResult, result)
}

func TestPositiveCIMKVMRedirectionSAP(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.CIMResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/kvm",
	}
	elementUnderTest := NewKVMRedirectionSAPWithClient(wsmanMessageCreator, &client)

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
				"should create and parse a valid cim_KVMRedirectionSAP Get call",
				CIMKVMRedirectionSAP,
				wsmantesting.Get,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageGet

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetResponse: KVMRedirectionSAP{
						CreationClassName:       "CIM_KVMRedirectionSAP",
						ElementName:             "KVM Redirection Service Access Point",
						Name:                    "KVM Redirection Service Access Point",
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "ManagedSystem",
						KVMProtocol:             4,
						EnabledState:            6,
						RequestedState:          2,
					},
				},
			},
			// ENUMERATES
			{
				"should create and parse a valid cim_KVMRedirectionSAP Enumerate call",
				CIMKVMRedirectionSAP,
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageEnumerate

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "CB020000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create and parse a valid cim_KVMRedirectionSAP Pull call",
				CIMKVMRedirectionSAP,
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
						Items: []KVMRedirectionSAP{
							{
								CreationClassName:       "CIM_KVMRedirectionSAP",
								ElementName:             "KVM Redirection Service Access Point",
								Name:                    "KVM Redirection Service Access Point",
								SystemCreationClassName: "CIM_ComputerSystem",
								SystemName:              "ManagedSystem",
								KVMProtocol:             4,
								EnabledState:            6,
								RequestedState:          5,
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

func TestNegativeCIMKVMRedirectionSAP(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.CIMResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/kvm",
	}
	elementUnderTest := NewKVMRedirectionSAPWithClient(wsmanMessageCreator, &client)

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
				"should handle error when cim_KVMRedirectionSAP Get call",
				CIMKVMRedirectionSAP,
				wsmantesting.Get,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetResponse: KVMRedirectionSAP{
						CreationClassName:       "CIM_KVMRedirectionSAP",
						ElementName:             "KVM Redirection Service Access Point",
						Name:                    "KVM Redirection Service Access Point",
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "ManagedSystem",
						KVMProtocol:             4,
						EnabledState:            6,
						RequestedState:          2,
					},
				},
			},
			// ENUMERATES
			{
				"should handle error when cim_KVMRedirectionSAP Enumerate call",
				CIMKVMRedirectionSAP,
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "CB020000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should handle error when cim_KVMRedirectionSAP Pull call",
				CIMKVMRedirectionSAP,
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
						Items: []KVMRedirectionSAP{
							{
								CreationClassName:       "CIM_KVMRedirectionSAP",
								ElementName:             "KVM Redirection Service Access Point",
								Name:                    "KVM Redirection Service Access Point",
								SystemCreationClassName: "CIM_ComputerSystem",
								SystemName:              "ManagedSystem",
								KVMProtocol:             4,
								EnabledState:            6,
								RequestedState:          5,
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
