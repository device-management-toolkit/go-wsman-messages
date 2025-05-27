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

func TestJson(t *testing.T) {
	response := Response{
		Body: Body{
			CredentialContextGetResponse: CredentialContextResponse{},
		},
	}
	expectedResult := "{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"ProfileGetAndPutResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"ElementName\":\"\",\"InstanceID\":\"\",\"Enabled\":false,\"ActiveInS0\":false,\"AuthenticationProtocol\":0,\"RoamingIdentity\":\"\",\"ServerCertificateName\":\"\",\"ServerCertificateNameComparison\":0,\"Username\":\"\",\"Password\":\"\",\"Domain\":\"\",\"ProtectedAccessCredential\":null,\"PACPassword\":\"\",\"ClientCertificate\":\"\",\"ServerCertificateIssue\":\"\",\"PxeTimeout\":0},\"CredentialContextGetResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"ElementInContext\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"Address\":\"\",\"ReferenceParameters\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"ResourceURI\":\"\",\"SelectorSet\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"Selectors\":null}}},\"ElementProvidingContext\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"Address\":\"\",\"ReferenceParameters\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"ResourceURI\":\"\",\"SelectorSet\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"Selectors\":null}}}},\"EnumerateResponse\":{\"EnumerationContext\":\"\"},\"PullResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"ProfileItems\":null,\"CredentialContextItems\":null}}"
	result := response.JSON()
	assert.Equal(t, expectedResult, result)
}

func TestYaml(t *testing.T) {
	response := Response{
		Body: Body{
			CredentialContextGetResponse: CredentialContextResponse{},
		},
	}
	expectedResult := "xmlname:\n    space: \"\"\n    local: \"\"\nprofilegetandputresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    elementname: \"\"\n    instanceid: \"\"\n    enabled: false\n    activeins0: false\n    authenticationprotocol: 0\n    roamingidentity: \"\"\n    servercertificatename: \"\"\n    servercertificatenamecomparison: 0\n    username: \"\"\n    password: \"\"\n    domain: \"\"\n    protectedaccesscredential: []\n    pacpassword: \"\"\n    clientcertificate: \"\"\n    servercertificateissue: \"\"\n    pxetimeout: 0\ncredentialcontextgetresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    elementincontext:\n        xmlname:\n            space: \"\"\n            local: \"\"\n        address: \"\"\n        referenceparameters:\n            xmlname:\n                space: \"\"\n                local: \"\"\n            resourceuri: \"\"\n            selectorset:\n                xmlname:\n                    space: \"\"\n                    local: \"\"\n                selectors: []\n    elementprovidingcontext:\n        xmlname:\n            space: \"\"\n            local: \"\"\n        address: \"\"\n        referenceparameters:\n            xmlname:\n                space: \"\"\n                local: \"\"\n            resourceuri: \"\"\n            selectorset:\n                xmlname:\n                    space: \"\"\n                    local: \"\"\n                selectors: []\nenumerateresponse:\n    enumerationcontext: \"\"\npullresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    profileitems: []\n    credentialcontextitems: []\n"
	result := response.YAML()
	assert.Equal(t, expectedResult, result)
}

func TestPositiveAMT_8021xCredentialContext(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.AMTResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/ieee8021x/credentialcontext",
	}
	elementUnderTest := NewIEEE8021xCredentialContextWithClient(wsmanMessageCreator, &client)

	t.Run("amt_8021xCredentialContext Tests", func(t *testing.T) {
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
				"should create a valid AMT_8021xCredentialContext Enumerate wsman message",
				AMTIEEE8021xCredentialContext,
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageEnumerate

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "02080000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create a valid AMT_8021xCredentialContext Pull wsman message",
				AMTIEEE8021xCredentialContext,
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

func TestNegativeAMT_8021xCredentialContext(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.AMTResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/ieee8021x/credentialcontext",
	}
	elementUnderTest := NewIEEE8021xCredentialContextWithClient(wsmanMessageCreator, &client)

	t.Run("amt_8021xCredentialContext Tests", func(t *testing.T) {
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
				"should create a valid AMT_8021xCredentialContext Enumerate wsman message",
				AMTIEEE8021xCredentialContext,
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "02080000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create a valid AMT_8021xCredentialContext Pull wsman message",
				AMTIEEE8021xCredentialContext,
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
