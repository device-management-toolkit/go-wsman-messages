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
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/methods"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/wsmantesting"
)

func TestPositiveService(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.CIMResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/boot/service",
	}
	elementUnderTest := NewBootServiceWithClient(wsmanMessageCreator, &client)

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
				"should create and parse a valid cim_BootService Get call",
				CIMBootService,
				wsmantesting.Get,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageGet

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					ServiceGetResponse: BootService{
						XMLName:                 xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootService", Local: CIMBootService},
						Name:                    "Intel(r) AMT Boot Service",
						CreationClassName:       "CIM_BootService",
						SystemName:              "Intel(r) AMT",
						SystemCreationClassName: "CIM_ComputerSystem",
						ElementName:             "Intel(r) AMT Boot Service",
						OperationalStatus:       []OperationalStatus{0},
						EnabledState:            2,
						RequestedState:          12,
					},
				},
			},
			// ENUMERATES
			{
				"should create and parse a valid cim_BootService Enumerate call",
				CIMBootService,
				wsmantesting.Enumerate,
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
				"should create and parse a valid cim_BootService Pull call",
				CIMBootService,
				wsmantesting.Pull,
				wsmantesting.PullBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessagePull

					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						BootServiceItems: []BootService{
							{
								XMLName:                 xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootService", Local: CIMBootService},
								Name:                    "Intel(r) AMT Boot Service",
								CreationClassName:       "CIM_BootService",
								SystemName:              "Intel(r) AMT",
								SystemCreationClassName: "CIM_ComputerSystem",
								ElementName:             "Intel(r) AMT Boot Service",
								OperationalStatus:       []OperationalStatus{0},
								EnabledState:            2,
								RequestedState:          12,
							},
						},
					},
				},
			},
			// SetBootConfigRole
			{
				"should handle error when making cim_BootService SetBootConfigRole wsman message",
				CIMBootService,
				wsmantesting.SetBootConfigRole,
				`<h:SetBootConfigRole_INPUT xmlns:h="http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootService"><h:BootConfigSetting><Address xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing">http://schemas.xmlsoap.org/ws/2004/08/addressing</Address><ReferenceParameters xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing"><ResourceURI xmlns="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd">http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootConfigSetting</ResourceURI><SelectorSet xmlns="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd"><Selector Name="InstanceID">InstanceID</Selector></SelectorSet></ReferenceParameters></h:BootConfigSetting><h:Role>0</h:Role></h:SetBootConfigRole_INPUT>`,
				func() (Response, error) {
					client.CurrentMessage = "SetBootConfigRole"

					return elementUnderTest.SetBootConfigRole("InstanceID", 0)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					SetBootConfigRole_OUTPUT: SetBootConfigRole_OUTPUT{
						ReturnValue: 0,
					},
				},
			},
			// Request State Change
			{
				"should create and parse a valid cim_BootService Request State Change call",
				CIMBootService,
				methods.RequestStateChange(CIMBootService),
				"<h:RequestStateChange_INPUT xmlns:h=\"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootService\"><h:RequestedState>3</h:RequestedState></h:RequestStateChange_INPUT>",
				func() (Response, error) {
					client.CurrentMessage = "RequestStateChange"

					return elementUnderTest.RequestStateChange(3)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					RequestStateChange_OUTPUT: common.ReturnValue{
						XMLName:     xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootService", Local: "RequestStateChange_OUTPUT"},
						ReturnValue: 0,
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

func TestNegativeService(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.CIMResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/boot/service",
	}
	elementUnderTest := NewBootServiceWithClient(wsmanMessageCreator, &client)

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
				"should handle error when making cim_BootService Get call",
				CIMBootService,
				wsmantesting.Get,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					ServiceGetResponse: BootService{
						XMLName:                 xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootService", Local: CIMBootService},
						Name:                    "Intel(r) AMT Boot Service",
						CreationClassName:       "CIM_BootService",
						SystemName:              "Intel(r) AMT",
						SystemCreationClassName: "CIM_ComputerSystem",
						ElementName:             "Intel(r) AMT Boot Service",
						OperationalStatus:       []OperationalStatus{0},
						EnabledState:            2,
						RequestedState:          12,
					},
				},
			},
			// ENUMERATES
			{
				"should handle error when making cim_BootService Enumerate call",
				CIMBootService,
				wsmantesting.Enumerate,
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
				"should handle error when making cim_BootService Pull wsman message",
				CIMBootService,
				wsmantesting.Pull,
				wsmantesting.PullBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						BootServiceItems: []BootService{
							{
								XMLName:                 xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootService", Local: CIMBootService},
								Name:                    "Intel(r) AMT Boot Service",
								CreationClassName:       "CIM_BootService",
								SystemName:              "Intel(r) AMT",
								SystemCreationClassName: "CIM_ComputerSystem",
								ElementName:             "Intel(r) AMT Boot Service",
								OperationalStatus:       []OperationalStatus{0},
								EnabledState:            2,
								RequestedState:          12,
							},
						},
					},
				},
			},
			// SetBootConfigRole
			{
				"should handle error when making cim_BootService SetBootConfigRole wsman message",
				CIMBootService,
				wsmantesting.SetBootConfigRole,
				`<h:SetBootConfigRole_INPUT xmlns:h="http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootService"><h:BootConfigSetting><Address xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing">http://schemas.xmlsoap.org/ws/2004/08/addressing</Address><ReferenceParameters xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing"><ResourceURI xmlns="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd">http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootConfigSetting</ResourceURI><SelectorSet xmlns="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd"><Selector Name="InstanceID">InstanceID</Selector></SelectorSet></ReferenceParameters></h:BootConfigSetting><h:Role>0</h:Role></h:SetBootConfigRole_INPUT>`,
				func() (Response, error) {
					client.CurrentMessage = "Error"

					return elementUnderTest.SetBootConfigRole("InstanceID", 0)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					SetBootConfigRole_OUTPUT: SetBootConfigRole_OUTPUT{
						ReturnValue: 0,
					},
				},
			},
			// Request State Change
			{
				"should handle error when making cim_BootService requestStateChange wsman message",
				CIMBootService,
				methods.RequestStateChange(CIMBootService),
				"<h:RequestStateChange_INPUT xmlns:h=\"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootService\"><h:RequestedState>3</h:RequestedState></h:RequestStateChange_INPUT>",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.RequestStateChange(3)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					RequestStateChange_OUTPUT: common.ReturnValue{
						XMLName:     xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootService", Local: "RequestStateChange_OUTPUT"},
						ReturnValue: 0,
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
