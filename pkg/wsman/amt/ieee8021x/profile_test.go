/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package ieee8021x

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/wsmantesting"
)

func TestPositiveAMT_8021XProfile(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.AMTResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/ieee8021x/profile",
	}
	elementUnderTest := NewIEEE8021xProfileWithClient(wsmanMessageCreator, &client)

	t.Run("amt_8021XProfile Tests", func(t *testing.T) {
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
				"should create a valid AMT_8021XProfile Get wsman message",
				AMTIEEE8021xProfile,
				wsmantesting.Get,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageGet

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					ProfileGetAndPutResponse: ProfileResponse{
						XMLName:     xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_8021XProfile", Local: "AMT_8021XProfile"},
						ElementName: "Intel(r) AMT 802.1x Profile",
						InstanceID:  "Intel(r) AMT 802.1x Profile 0",
						Enabled:     false,
						ActiveInS0:  true,
						PxeTimeout:  0,
					},
				},
			},
			// ENUMERATES
			{
				"should create a valid AMT_8021XProfile Enumerate wsman message",
				AMTIEEE8021xProfile,
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageEnumerate

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "04080000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create a valid AMT_8021XProfile Pull wsman message",
				AMTIEEE8021xProfile,
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
						ProfileItems: []ProfileResponse{
							{
								XMLName:     xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_8021XProfile", Local: "AMT_8021XProfile"},
								ElementName: "Intel(r) AMT 802.1x Profile",
								InstanceID:  "Intel(r) AMT 802.1x Profile 0",
								Enabled:     false,
								ActiveInS0:  true,
								PxeTimeout:  0,
							},
						},
					},
				},
			},
			// PUTS
			{
				"should create a valid AMT_8021XProfile Put wsman message",
				AMTIEEE8021xProfile,
				wsmantesting.Put,
				"<h:AMT_8021XProfile xmlns:h=\"http://intel.com/wbem/wscim/1/amt-schema/1/AMT_8021XProfile\"><h:ElementName>Intel(r) AMT 802.1x Profile</h:ElementName><h:InstanceID>Intel(r) AMT 802.1x Profile 0</h:InstanceID><h:Enabled>false</h:Enabled><h:ActiveInS0>false</h:ActiveInS0><h:AuthenticationProtocol>0</h:AuthenticationProtocol><h:ServerCertificateNameComparison>0</h:ServerCertificateNameComparison></h:AMT_8021XProfile>",
				func() (Response, error) {
					profileRequest := ProfileRequest{
						H:           "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_8021XProfile",
						ActiveInS0:  false,
						ElementName: "Intel(r) AMT 802.1x Profile",
						Enabled:     false,
						InstanceID:  "Intel(r) AMT 802.1x Profile 0",
					}
					client.CurrentMessage = wsmantesting.CurrentMessagePut

					return elementUnderTest.Put(profileRequest)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					ProfileGetAndPutResponse: ProfileResponse{
						XMLName:     xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_8021XProfile", Local: "AMT_8021XProfile"},
						ElementName: "Intel(r) AMT 802.1x Profile",
						InstanceID:  "Intel(r) AMT 802.1x Profile 0",
						Enabled:     false,
						ActiveInS0:  false,
						PxeTimeout:  120,
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

func TestNegativeAMT_8021XProfile(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.AMTResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/ieee8021x/profile",
	}
	elementUnderTest := NewIEEE8021xProfileWithClient(wsmanMessageCreator, &client)

	t.Run("amt_8021XProfile Tests", func(t *testing.T) {
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
				"should create a valid AMT_8021XProfile Get wsman message",
				AMTIEEE8021xProfile,
				wsmantesting.Get,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					ProfileGetAndPutResponse: ProfileResponse{
						XMLName:     xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_8021XProfile", Local: "AMT_8021XProfile"},
						ElementName: "Intel(r) AMT 802.1x Profile",
						InstanceID:  "Intel(r) AMT 802.1x Profile 0",
						Enabled:     false,
						ActiveInS0:  true,
						PxeTimeout:  0,
					},
				},
			},
			// ENUMERATES
			{
				"should create a valid AMT_8021XProfile Enumerate wsman message",
				AMTIEEE8021xProfile,
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "04080000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create a valid AMT_8021XProfile Pull wsman message",
				AMTIEEE8021xProfile,
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
						ProfileItems: []ProfileResponse{
							{
								XMLName:     xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_8021XProfile", Local: "AMT_8021XProfile"},
								ElementName: "Intel(r) AMT 802.1x Profile",
								InstanceID:  "Intel(r) AMT 802.1x Profile 0",
								Enabled:     false,
								ActiveInS0:  true,
								PxeTimeout:  0,
							},
						},
					},
				},
			},
			// PUTS
			{
				"should create a valid AMT_8021XProfile Put wsman message",
				AMTIEEE8021xProfile,
				wsmantesting.Put,
				"<h:AMT_8021XProfile xmlns:h=\"http://intel.com/wbem/wscim/1/amt-schema/1/AMT_8021XProfile\"><h:ElementName>Intel(r) AMT 802.1x Profile</h:ElementName><h:InstanceID>Intel(r) AMT 802.1x Profile 0</h:InstanceID><h:Enabled>false</h:Enabled><h:ActiveInS0>false</h:ActiveInS0><h:AuthenticationProtocol>0</h:AuthenticationProtocol><h:ServerCertificateNameComparison>0</h:ServerCertificateNameComparison></h:AMT_8021XProfile>",
				func() (Response, error) {
					profileRequest := ProfileRequest{
						H:           "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_8021XProfile",
						ActiveInS0:  false,
						ElementName: "Intel(r) AMT 802.1x Profile",
						Enabled:     false,
						InstanceID:  "Intel(r) AMT 802.1x Profile 0",
					}
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Put(profileRequest)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					ProfileGetAndPutResponse: ProfileResponse{
						XMLName:     xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_8021XProfile", Local: "AMT_8021XProfile"},
						ElementName: "Intel(r) AMT 802.1x Profile",
						InstanceID:  "Intel(r) AMT 802.1x Profile 0",
						Enabled:     false,
						ActiveInS0:  false,
						PxeTimeout:  120,
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
