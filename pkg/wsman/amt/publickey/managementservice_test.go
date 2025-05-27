/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package publickey

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/wsmantesting"
)

const (
	EnvelopeResponseService = `<a:Envelope xmlns:a="http://www.w3.org/2003/05/soap-envelope" x-mlns:b="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:c="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns:d="http://schemas.xmlsoap.org/ws/2005/02/trust" xmlns:e="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" xmlns:f="http://schemas.dmtf.org/wbem/wsman/1/cimbinding.xsd" xmlns:g="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService" xmlns:h="http://schemas.dmtf.org/wbem/wscim/1/common" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><a:Header><b:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:To><b:RelatesTo>0</b:RelatesTo><b:Action a:mustUnderstand="true">`
	GetBodyService          = `<g:AMT_PublicKeyManagementService><g:CreationClassName>AMT_PublicKeyManagementService</g:CreationClassName><g:ElementName>Intel(r) AMT Public Key Management Service</g:ElementName><g:Name>Intel(r) AMT Public Key Management Service</g:Name><g:SystemCreationClassName>CIM_ComputerSystem</g:SystemCreationClassName><g:SystemName>ManagedSystem</g:SystemName></g:AMT_PublicKeyManagementService>`
	PrivateKey              = "privatekey"
)

func TestPositiveAMT_PublicKeyManagementService(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.AMTResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/publickey/management",
	}
	elementUnderTest := NewPublicKeyManagementServiceWithClient(wsmanMessageCreator, &client)

	t.Run("amt_PublicKeyManagementService Tests", func(t *testing.T) {
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
			{
				"should create a valid AMT_PublicKeyManagementService Get wsman message",
				AMTPublicKeyManagementService,
				wsmantesting.Get,
				"",
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageGet

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					KeyManagementGetResponse: KeyManagementResponse{
						XMLName:                 xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMTPublicKeyManagementService), Local: AMTPublicKeyManagementService},
						CreationClassName:       "AMT_PublicKeyManagementService",
						ElementName:             "Intel(r) AMT Certificate Store Service",
						EnabledDefault:          5,
						EnabledState:            5,
						Name:                    "Intel(r) AMT Public Key Management Service",
						RequestedState:          12,
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "Intel(r) AMT",
					},
				},
			},
			// ENUMERATES
			{
				"should create a valid AMT_PublicKeyManagementService Enumerate wsman message",
				AMTPublicKeyManagementService,
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
						EnumerationContext: "7E000000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create a valid AMT_PublicKeyManagementService Pull wsman message",
				AMTPublicKeyManagementService,
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
						KeyManagementItems: []KeyManagementResponse{
							{
								XMLName:                 xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMTPublicKeyManagementService), Local: AMTPublicKeyManagementService},
								CreationClassName:       "AMT_PublicKeyManagementService",
								ElementName:             "Intel(r) AMT Certificate Store Service",
								EnabledDefault:          5,
								EnabledState:            5,
								Name:                    "Intel(r) AMT Public Key Management Service",
								RequestedState:          12,
								SystemCreationClassName: "CIM_ComputerSystem",
								SystemName:              "Intel(r) AMT",
							},
						},
					},
				},
			},

			// AddTrustedRootCertificate
			{
				"should return a valid amt_PublicKeyManagementService AddTrustedRootCertificate wsman message",
				AMTPublicKeyManagementService,
				`http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService/AddTrustedRootCertificate`,
				fmt.Sprintf(`<h:AddTrustedRootCertificate_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService"><h:CertificateBlob>%s</h:CertificateBlob></h:AddTrustedRootCertificate_INPUT>`, wsmantesting.TrustedRootCert),
				"",
				func() (Response, error) {
					client.CurrentMessage = "AddTrustedRootCertificate"

					return elementUnderTest.AddTrustedRootCertificate(wsmantesting.TrustedRootCert)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					AddTrustedRootCertificate_OUTPUT: AddTrustedRootCertificate_OUTPUT{
						XMLName: xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMTPublicKeyManagementService), Local: "AddTrustedRootCertificate_OUTPUT"},
						CreatedCertificate: CreatedCertificateResponse{
							XMLName: xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService", Local: "CreatedCertificate"},
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
											Text:    "Intel(r) AMT Certificate: Handle: 2",
										},
									},
								},
							},
						},
						ReturnValue: 0,
					},
				},
			},

			// GenerateKeyPair
			{
				"should return a valid amt_PublicKeyManagementService GenerateKeyPair wsman message",
				AMTPublicKeyManagementService,
				`http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService/GenerateKeyPair`,
				`<h:GenerateKeyPair_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService"><h:KeyAlgorithm>0</h:KeyAlgorithm><h:KeyLength>2048</h:KeyLength></h:GenerateKeyPair_INPUT>`,
				"",
				func() (Response, error) {
					client.CurrentMessage = "GenerateKeyPair"

					return elementUnderTest.GenerateKeyPair(RSA, KeyLength2048)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GenerateKeyPair_OUTPUT: GenerateKeyPair_OUTPUT{
						XMLName: xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMTPublicKeyManagementService), Local: "GenerateKeyPair_OUTPUT"},
						KeyPair: KeyPairResponse{
							XMLName: xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService", Local: "KeyPair"},
							Address: "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous",
							ReferenceParameters: ReferenceParametersResponse{
								XMLName:     xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/08/addressing", Local: "ReferenceParameters"},
								ResourceURI: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicPrivateKeyPair",
								SelectorSet: SelectorSetResponse{
									XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "SelectorSet"},
									Selectors: []SelectorResponse{
										{
											XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
											Name:    "InstanceID",
											Text:    "Intel(r) AMT Key: Handle: 0",
										},
									},
								},
							},
						},
						ReturnValue: 0,
					},
				},
			},
			// AddCertificate
			{
				"should return a valid amt_PublicKeyManagementService AddCertificate wsman message",
				AMTPublicKeyManagementService,
				`http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService/AddCertificate`,
				fmt.Sprintf(`<h:AddCertificate_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService"><h:CertificateBlob>%s</h:CertificateBlob></h:AddCertificate_INPUT>`, wsmantesting.TrustedRootCert),
				"",
				func() (Response, error) {
					client.CurrentMessage = "AddCertificate"

					return elementUnderTest.AddCertificate(wsmantesting.TrustedRootCert)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					AddCertificate_OUTPUT: AddCertificate_OUTPUT{
						XMLName: xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMTPublicKeyManagementService), Local: "AddCertificate_OUTPUT"},
						CreatedCertificate: CreatedCertificateResponse{
							XMLName: xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService", Local: "CreatedCertificate"},
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
											Text:    "Intel(r) AMT Certificate: Handle: 1",
										},
									},
								},
							},
						},
						ReturnValue: 0,
					},
				},
			},
			// PKCS10RequestEx
			{
				"should return a valid amt_PublicKeyManagementService GeneratePKCS10RequestEx wsman message",
				AMTPublicKeyManagementService,
				"http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService/GeneratePKCS10RequestEx",
				fmt.Sprintf("<h:GeneratePKCS10RequestEx_INPUT xmlns:h=\"http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService\"><h:KeyPair><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address><a:ReferenceParameters><w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicPrivateKeyPair</w:ResourceURI><w:SelectorSet><w:Selector Name=\"InstanceID\">%s</w:Selector></w:SelectorSet></a:ReferenceParameters></h:KeyPair><h:SigningAlgorithm>1</h:SigningAlgorithm><h:NullSignedCertificateRequest>reallylongcertificateteststring</h:NullSignedCertificateRequest></h:GeneratePKCS10RequestEx_INPUT>", "test"),
				"",
				func() (Response, error) {
					client.CurrentMessage = "GeneratePKCS10RequestEx"

					return elementUnderTest.GeneratePKCS10RequestEx("test", "reallylongcertificateteststring", 1)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GeneratePKCS10RequestEx_OUTPUT: GeneratePKCS10RequestEx_OUTPUT{
						XMLName:                  xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMTPublicKeyManagementService), Local: "GeneratePKCS10RequestEx_OUTPUT"},
						SignedCertificateRequest: "test?",
						ReturnValue:              0,
					},
				},
			},
			// AddKey
			{
				"should return a valid amt_PublicKeyManagementService AddKey wsman message",
				AMTPublicKeyManagementService,
				"http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService/AddKey",
				`<h:AddKey_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService"><h:KeyBlob>privatekey</h:KeyBlob></h:AddKey_INPUT>`,
				"",
				func() (Response, error) {
					client.CurrentMessage = "AddKey"
					cert := PrivateKey

					return elementUnderTest.AddKey(cert)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					AddKey_OUTPUT: AddKey_OUTPUT{
						XMLName: xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMTPublicKeyManagementService), Local: "AddKey_OUTPUT"},
						CreatedKey: CreatedKeyResponse{
							XMLName: xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService", Local: "CreatedKey"},
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
											Text:    "Intel(r) AMT Key: Handle: 1",
										},
									},
								},
							},
						},
						ReturnValue: 0,
					},
				},
			},
			// DELETE
			{
				"should create a valid amt_PublicKeyManagementService Delete wsman message",
				AMTPublicKeyManagementService,
				wsmantesting.Delete,
				"",
				"<w:SelectorSet><w:Selector Name=\"InstanceID\">instanceID123</w:Selector></w:SelectorSet>",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageDelete

					return elementUnderTest.Delete("instanceID123")
				},
				Body{XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"}},
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

func TestNegativeAMT_PublicKeyManagementService(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.AMTResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/publickey/management",
	}
	elementUnderTest := NewPublicKeyManagementServiceWithClient(wsmanMessageCreator, &client)

	t.Run("amt_PublicKeyManagementService Tests", func(t *testing.T) {
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
			{
				"should create a invalid AMT_PublicKeyManagementService Get wsman message",
				AMTPublicKeyManagementService,
				wsmantesting.Get,
				"",
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					KeyManagementGetResponse: KeyManagementResponse{
						XMLName:                 xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMTPublicKeyManagementService), Local: AMTPublicKeyManagementService},
						CreationClassName:       "AMT_PublicKeyManagementService",
						ElementName:             "Intel(r) AMT Certificate Store Service",
						EnabledDefault:          5,
						EnabledState:            5,
						Name:                    "Intel(r) AMT Public Key Management Service",
						RequestedState:          12,
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "Intel(r) AMT",
					},
				},
			},
			// ENUMERATES
			{
				"should create a invalid AMT_PublicKeyManagementService Enumerate wsman message",
				AMTPublicKeyManagementService,
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "7E000000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create a invalid AMT_PublicKeyManagementService Pull wsman message",
				AMTPublicKeyManagementService,
				wsmantesting.Pull,
				wsmantesting.PullBody,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						XMLName: xml.Name{Space: message.XMLPullResponseSpace, Local: "PullResponse"},
						KeyManagementItems: []KeyManagementResponse{
							{
								XMLName:                 xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMTPublicKeyManagementService), Local: AMTPublicKeyManagementService},
								CreationClassName:       "AMT_PublicKeyManagementService",
								ElementName:             "Intel(r) AMT Certificate Store Service",
								EnabledDefault:          5,
								EnabledState:            5,
								Name:                    "Intel(r) AMT Public Key Management Service",
								RequestedState:          12,
								SystemCreationClassName: "CIM_ComputerSystem",
								SystemName:              "Intel(r) AMT",
							},
						},
					},
				},
			},

			// AddTrustedRootCertificate
			{
				"should return a invalid amt_PublicKeyManagementService AddTrustedRootCertificate wsman message",
				AMTPublicKeyManagementService,
				`http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService/AddTrustedRootCertificate`,
				fmt.Sprintf(`<h:AddTrustedRootCertificate_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService"><h:CertificateBlob>%s</h:CertificateBlob></h:AddTrustedRootCertificate_INPUT>`, wsmantesting.TrustedRootCert),
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.AddTrustedRootCertificate(wsmantesting.TrustedRootCert)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					AddTrustedRootCertificate_OUTPUT: AddTrustedRootCertificate_OUTPUT{
						XMLName: xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMTPublicKeyManagementService), Local: "AddTrustedRootCertificate_OUTPUT"},
						CreatedCertificate: CreatedCertificateResponse{
							XMLName: xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService", Local: "CreatedCertificate"},
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
											Text:    "Intel(r) AMT Certificate: Handle: 2",
										},
									},
								},
							},
						},
						ReturnValue: 0,
					},
				},
			},

			// GenerateKeyPair
			{
				"should return a invalid amt_PublicKeyManagementService GenerateKeyPair wsman message",
				AMTPublicKeyManagementService,
				`http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService/GenerateKeyPair`,
				`<h:GenerateKeyPair_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService"><h:KeyAlgorithm>0</h:KeyAlgorithm><h:KeyLength>2048</h:KeyLength></h:GenerateKeyPair_INPUT>`,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.GenerateKeyPair(RSA, KeyLength2048)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GenerateKeyPair_OUTPUT: GenerateKeyPair_OUTPUT{
						XMLName: xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMTPublicKeyManagementService), Local: "GenerateKeyPair_OUTPUT"},
						KeyPair: KeyPairResponse{
							XMLName: xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService", Local: "KeyPair"},
							Address: "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous",
							ReferenceParameters: ReferenceParametersResponse{
								XMLName:     xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/08/addressing", Local: "ReferenceParameters"},
								ResourceURI: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicPrivateKeyPair",
								SelectorSet: SelectorSetResponse{
									XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "SelectorSet"},
									Selectors: []SelectorResponse{
										{
											XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
											Name:    "InstanceID",
											Text:    "Intel(r) AMT Key: Handle: 0",
										},
									},
								},
							},
						},
						ReturnValue: 0,
					},
				},
			},
			// AddCertificate
			{
				"should return a invalid amt_PublicKeyManagementService AddCertificate wsman message",
				AMTPublicKeyManagementService,
				`http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService/AddCertificate`,
				fmt.Sprintf(`<h:AddCertificate_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService"><h:CertificateBlob>%s</h:CertificateBlob></h:AddCertificate_INPUT>`, wsmantesting.TrustedRootCert),
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.AddCertificate(wsmantesting.TrustedRootCert)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					AddCertificate_OUTPUT: AddCertificate_OUTPUT{
						XMLName: xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMTPublicKeyManagementService), Local: "AddCertificate_OUTPUT"},
						CreatedCertificate: CreatedCertificateResponse{
							XMLName: xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService", Local: "CreatedCertificate"},
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
											Text:    "Intel(r) AMT Certificate: Handle: 1",
										},
									},
								},
							},
						},
						ReturnValue: 0,
					},
				},
			},

			// PKCS10RequestEx
			{
				"should return a invalid amt_PublicKeyManagementService GeneratePKCS10RequestEx wsman message",
				AMTPublicKeyManagementService,
				"http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService/GeneratePKCS10RequestEx",
				fmt.Sprintf("<h:GeneratePKCS10RequestEx_INPUT xmlns:h=\"http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService\"><h:KeyPair><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address><a:ReferenceParameters><w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicPrivateKeyPair</w:ResourceURI><w:SelectorSet><w:Selector Name=\"InstanceID\">%s</w:Selector></w:SelectorSet></a:ReferenceParameters></h:KeyPair><h:SigningAlgorithm>1</h:SigningAlgorithm><h:NullSignedCertificateRequest>reallylongcertificateteststring</h:NullSignedCertificateRequest></h:GeneratePKCS10RequestEx_INPUT>", "test"),
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.GeneratePKCS10RequestEx("test", "reallylongcertificateteststring", 1)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GeneratePKCS10RequestEx_OUTPUT: GeneratePKCS10RequestEx_OUTPUT{
						XMLName:                  xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMTPublicKeyManagementService), Local: "GeneratePKCS10RequestEx_OUTPUT"},
						SignedCertificateRequest: "test?",
						ReturnValue:              0,
					},
				},
			},

			// AddKey
			{
				"should return a invalid amt_PublicKeyManagementService AddKey wsman message",
				AMTPublicKeyManagementService,
				"http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService/AddKey",
				`<h:AddKey_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService"><h:KeyBlob>privatekey</h:KeyBlob></h:AddKey_INPUT>`,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError
					cert := PrivateKey

					return elementUnderTest.AddKey(cert)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					AddKey_OUTPUT: AddKey_OUTPUT{
						XMLName: xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMTPublicKeyManagementService), Local: "AddKey_OUTPUT"},
						CreatedKey: CreatedKeyResponse{
							XMLName: xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService", Local: "CreatedKey"},
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
											Text:    "Intel(r) AMT Key: Handle: 1",
										},
									},
								},
							},
						},
						ReturnValue: 0,
					},
				},
			},
			// DELETE
			{
				"should create a invalid amt_PublicKeyManagementService Delete wsman message",
				AMTPublicKeyManagementService,
				wsmantesting.Delete,
				"",
				"<w:SelectorSet><w:Selector Name=\"InstanceID\">instanceID123</w:Selector></w:SelectorSet>",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Delete("instanceID123")
				},
				Body{XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"}},
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				expectedXMLInput := wsmantesting.ExpectedResponse(messageID, resourceURIBase, test.method, test.action, test.extraHeader, test.body)
				messageID++
				response, err := test.responseFunc()
				assert.Error(t, err)
				assert.Equal(t, expectedXMLInput, response.XMLInput)
				assert.NotEqual(t, test.expectedResponse, response.Body)
			})
		}
	})
}
