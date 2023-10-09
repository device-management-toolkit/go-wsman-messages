/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package publickey

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
	
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/common"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsmantesting"
)

func TestAMT_PublicKeyManagementService(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	//client := MockClient{} // wsman.NewClient("http://localhost:16992/wsman", "admin", "P@ssw0rd", true)
	//elementUnderTest := NewServiceWithClient(wsmanMessageCreator, &client)
	// enumerationId := ""
	client := wsman.NewClient("http://localhost:16992/wsman", "admin", "Intel123!", true)
	elementUnderTest := NewPublicKeyManagementServiceWithClient(wsmanMessageCreator, client)

	t.Run("amt_* Tests", func(t *testing.T) {
		tests := []struct {
			name         string
			method       string
			action       string
			body         string
			extraHeader  string
			responseFunc 	 	func() (Response, error)
			expectedResponse 	interface{}
		}{
			//GETS
			{
				"should create a valid AMT_PublicKeyManagementService Get wsman message", 
				"AMT_PublicKeyManagementService", 
				wsmantesting.GET, 
				"", 
				"", 
				func() (Response, error) {
					//client.CurrentMessage = "Get"
					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: "http://www.w3.org/2003/05/soap-envelope", Local: "Body"},
					KeyManagement: KeyManagement{
						CreationClassName: "",
						ElementName: "",
						EnabledDefault: 0, 
						EnabledState: 0,
						Name: "",
						RequestedState: 0,
						SystemCreationClassName: "", 
						SystemName: "",
					},
				},
			},
			//ENUMERATES
			{
				"should create a valid AMT_PublicKeyManagementService Enumerate wsman message", 
				"AMT_PublicKeyManagementService", 
				wsmantesting.ENUMERATE, 
				wsmantesting.ENUMERATE_BODY, 
				"", 
				func() (Response, error) {
					//client.CurrentMessage = "Enumerate"
					return elementUnderTest.Enumerate()
				}, 
				Body{
					XMLName: xml.Name{Space: "http://www.w3.org/2003/05/soap-envelope", Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "80000000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			//{"should create a valid AMT_PublicKeyManagementService Pull wsman message", "AMT_PublicKeyManagementService", wsmantesting.PULL, wsmantesting.PULL_BODY, "", func() string { return elementUnderTest.Pull(wsmantesting.EnumerationContext) }},

			// // PUBLIC KEY MANAGEMENT SERVICE
			// {"should return a valid amt_PublicKeyManagementService AddTrustedRootCertificate wsman message", "AMT_PublicKeyManagementService", `http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService/AddTrustedRootCertificate`, fmt.Sprintf(`<h:AddTrustedRootCertificate_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService"><h:CertificateBlob>%s</h:CertificateBlob></h:AddTrustedRootCertificate_INPUT>`, wsmantesting.TrustedRootCert), "", func() string {
			// 	return elementUnderTest.AddTrustedRootCertificate(wsmantesting.TrustedRootCert)
			// }},

			// {"should return a valid amt_PublicKeyManagementService GenerateKeyPair wsman message", "AMT_PublicKeyManagementService", `http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService/GenerateKeyPair`, `<h:GenerateKeyPair_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService"><h:KeyAlgorithm>0</h:KeyAlgorithm><h:KeyLength>2048</h:KeyLength></h:GenerateKeyPair_INPUT>`, "", func() string {
			// 	params := GenerateKeyPair_INPUT{
			// 		KeyAlgorithm: 0,
			// 		KeyLength:    2048,
			// 	}
			// 	return elementUnderTest.GenerateKeyPair(params)
			// }},

			// {"should return a valid amt_PublicKeyManagementService AddCertificate wsman message", "AMT_PublicKeyManagementService", `http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService/AddCertificate`, fmt.Sprintf(`<h:AddCertificate_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService"><h:CertificateBlob>%s</h:CertificateBlob></h:AddCertificate_INPUT>`, wsmantesting.TrustedRootCert), "", func() string {
			// 	return elementUnderTest.AddCertificate(wsmantesting.TrustedRootCert)
			// }},

			// {"should return a valid amt_PublicKeyManagementService GeneratePKCS10RequestEx wsman message", "AMT_PublicKeyManagementService", "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService/GeneratePKCS10RequestEx", `<h:GeneratePKCS10RequestEx_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService"><h:KeyPair>test</h:KeyPair><h:NullSignedCertificateRequest>reallylongcertificateteststring</h:NullSignedCertificateRequest><h:SigningAlgorithm>1</h:SigningAlgorithm></h:GeneratePKCS10RequestEx_INPUT>`, "", func() string {
			// 	pkcs10Request := PKCS10Request{
			// 		KeyPair:                      "test",
			// 		NullSignedCertificateRequest: "reallylongcertificateteststring",
			// 		SigningAlgorithm:             1,
			// 	}
			// 	return elementUnderTest.GeneratePKCS10RequestEx(pkcs10Request)
			// }},
			// {"should return a valid amt_PublicKeyManagementService AddKey wsman message", "AMT_PublicKeyManagementService", "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService/AddKey", `<h:AddKey_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService"><h:KeyBlob>privatekey</h:KeyBlob></h:AddKey_INPUT>`, "", func() string {
			// 	cert := []byte("privatekey")
			// 	return elementUnderTest.AddKey(cert)
			// }},
			//DELETE
			//{"should create a valid amt_PublicKeyManagementService Delete wsman message", "AMT_PublicKeyManagementService", wsmantesting.DELETE, "", "<w:SelectorSet><w:Selector Name=\"InstanceID\">instanceID123</w:Selector></w:SelectorSet>", func() string { return elementUnderTest.Delete("instanceID123") }},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				expectedXMLInput := wsmantesting.ExpectedResponse(messageID, resourceUriBase, test.method, test.action, "", test.body)
				messageID++
				response, err := test.responseFunc()
				println(response.XMLOutput)
				assert.NoError(t, err)
				assert.Equal(t, expectedXMLInput, response.XMLInput)
				assert.Equal(t, test.expectedResponse, response.Body)
			})
		}
	})
}
