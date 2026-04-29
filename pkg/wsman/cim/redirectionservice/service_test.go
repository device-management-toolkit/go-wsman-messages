/*********************************************************************
 * Copyright (c) Intel Corporation 2026
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package redirectionservice

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/wsmantesting"
)

func TestPositiveCIMRedirectionService(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.CIMResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/redirectionservice",
	}
	elementUnderTest := NewRedirectionServiceWithClient(wsmanMessageCreator, &client)

	t.Run("cim_RedirectionService Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			body             string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			{
				"should create and parse a valid CIM_RedirectionService Get call",
				CIMRedirectionService,
				wsmantesting.Get,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageGet

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetResponse: RedirectionService{
						XMLName:                 xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_RedirectionService", Local: CIMRedirectionService},
						CreationClassName:       "CIM_RedirectionService",
						ElementName:             "Intel(r) ME KVM Redirection Service",
						EnabledState:            3,
						Name:                    "Intel(r) AMT KVM Redirection Service",
						RequestedState:          12,
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "Intel(r) AMT",
						MaxCurrentEnabledSAPs:   1,
						RedirectionServiceType:  []int{3},
						SharingMode:             3,
					},
				},
			},
			{
				"should create a valid CIM_RedirectionService Enumerate call",
				CIMRedirectionService,
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageEnumerate

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "30000000-0000-0000-0000-000000000000",
					},
				},
			},
			{
				"should create a valid CIM_RedirectionService Pull call",
				CIMRedirectionService,
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
						RedirectionServiceItems: []RedirectionService{
							{
								XMLName:                 xml.Name{Space: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_RedirectionService", Local: CIMRedirectionService},
								CreationClassName:       "CIM_RedirectionService",
								ElementName:             "Intel(r) ME KVM Redirection Service",
								EnabledState:            3,
								Name:                    "Intel(r) AMT KVM Redirection Service",
								RequestedState:          12,
								SystemCreationClassName: "CIM_ComputerSystem",
								SystemName:              "Intel(r) AMT",
								MaxCurrentEnabledSAPs:   1,
								RedirectionServiceType:  []int{3},
								SharingMode:             3,
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

func TestNegativeCIMRedirectionService(t *testing.T) {
	resourceURIBase := wsmantesting.CIMResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/redirectionservice",
	}
	elementUnderTest := NewRedirectionServiceWithClient(wsmanMessageCreator, &client)

	t.Run("cim_RedirectionService Error Handling Tests", func(t *testing.T) {
		tests := []struct {
			name         string
			responseFunc func() (Response, error)
		}{
			{
				"should handle error when Get fails",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Get()
				},
			},
			{
				"should handle error when Enumerate fails",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Enumerate()
				},
			},
			{
				"should handle error when Pull fails",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				_, err := test.responseFunc()
				assert.Error(t, err)
			})
		}
	})
}
