/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package publickey

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/wsmantesting"
)

const X509TestCertificate = "MIIEOzCCAqOgAwIBAgIDAZMjMA0GCSqGSIb3DQEBDAUAMD0xFzAVBgNVBAMTDk1QU1Jvb3QtMGFmMWQ1MRAwDgYDVQQKEwd1bmtub3duMRAwDgYDVQQGEwd1bmtub3duMCAXDTIyMDkyNDEwNDUwOFoYDzIwNTMwOTI0MTA0NTA4WjA9MRcwFQYDVQQDEw5NUFNSb290LTBhZjFkNTEQMA4GA1UEChMHdW5rbm93bjEQMA4GA1UEBhMHdW5rbm93bjCCAaIwDQYJKoZIhvcNAQEBBQADggGPADCCAYoCggGBALz/oJNyWXlClSlteAieC8Uyd4A+tbn8b45k6LKiImhDmdz/xFo9xe0C9GNf7b42KVpg5WoH/sPhoClR9Tv5i1LnilT1SUir42fcm2NEV9dRcLsPd/RAQfz8u0D4zb3blnxE8isqzriNpG7kac35UidSr5ym8TZ3IwXx6JJuncGgfB0DFZADC/+dA74n3coykvWBYqLr6RI5pkAxvulkRlCsatJTJrvMUYJ51GI28jV56mIAc89sLrHqiSKCZBH9AcUrnZ/cB6ST/IikXpxy5wXBIvWT3VKVq75T/uIoCBEp5TLEn1EOYGqBBOCSQgmtmX7eVaB0s1+ppPW9w9a2zS45cHAtQ7tYvkkPv2dRhSzZdlk6HRXDP5wsF0aiflZCgbrjkq0SFC4e3Lo7XQX3FTNb0SOTZVTydupoMKkgJQTNlcosdu1ZzaIBl3eSkKkJZz2rUTssZC5tn9vcDd5vy3BzcGh5pvkgfAgN1sydqG7Ke1qCkNEzm11B/BsevatjjwIDAQABo0IwQDAMBgNVHRMEBTADAQH/MBEGCWCGSAGG+EIBAQQEAwIABzAdBgNVHQ4EFgQUCvHVQqerCid99eLApuLky9x6H5owDQYJKoZIhvcNAQEMBQADggGBAIzOyGV0hzsmH2biJlzwTZaHMxqS7boTFMkHw+KvzsI201tHqVmCoiQ8EHErBGLSoDOTDRgOUGOCA5XU5ie9OWupAGqKBSwIyAhmJMOzrzC4Gwpu8K1msoFJH30kx/V9purpbS3BRj0xfYXLa6IczbTg3E5IfTnZRJ9YuUtKQfI0P9c5U9CoKtddKn4+lRvOjFDoYfQGCJ7go3xjNCcGCVCjfkUhAVdbQ21DCRr6/YCZDWmjzZpL0p7UKF8roTiNuL/Z7gIXxch5HOmEWHY9uQ6K2MntuxAu0aK/mSD2kwmt/ECongdEGfUvhULLoPRQlQ2LnzcUQEgMECGQR5Yfy9jT0E8zdWDpc2tgVioNu6rEYKgp/GhG+sv7jv58pW82FRAV9xXtftW9+XDugC8tBJ6JHn0Q2v0QAflD2CEQVhWAY8bAqrbfTGUsaLfGL6kxV/qqssoMgLR8WhQ96T5le/4XGhQpbCHWIlctD6MwbrsunIAeQKp1Sc3DosY7DLq1MQ=="

func TestJson(t *testing.T) {
	response := Response{
		Body: Body{
			AddCertificate_OUTPUT: AddCertificate_OUTPUT{},
		},
	}
	expectedResult := "{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"AddTrustedRootCertificate_OUTPUT\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"CreatedCertificate\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"Address\":\"\",\"ReferenceParameters\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"ResourceURI\":\"\",\"SelectorSet\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"Selectors\":null}}},\"ReturnValue\":0},\"AddCertificate_OUTPUT\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"CreatedCertificate\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"Address\":\"\",\"ReferenceParameters\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"ResourceURI\":\"\",\"SelectorSet\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"Selectors\":null}}},\"ReturnValue\":0},\"AddKey_OUTPUT\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"CreatedKey\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"Address\":\"\",\"ReferenceParameters\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"ResourceURI\":\"\",\"SelectorSet\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"Selectors\":null}}},\"ReturnValue\":0},\"GenerateKeyPair_OUTPUT\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"KeyPair\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"Address\":\"\",\"ReferenceParameters\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"ResourceURI\":\"\",\"SelectorSet\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"Selectors\":null}}},\"ReturnValue\":0},\"GeneratePKCS10RequestEx_OUTPUT\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"SignedCertificateRequest\":\"\",\"ReturnValue\":0},\"KeyManagementGetResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"CreationClassName\":\"\",\"ElementName\":\"\",\"EnabledDefault\":0,\"EnabledState\":0,\"Name\":\"\",\"OperationalStatus\":null,\"RequestedState\":0,\"SystemCreationClassName\":\"\",\"SystemName\":\"\"},\"PublicKeyCertificateGetAndPutResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"ElementName\":\"\",\"InstanceID\":\"\",\"X509Certificate\":\"\",\"TrustedRootCertificate\":false,\"Issuer\":\"\",\"Subject\":\"\",\"ReadOnlyCertificate\":false},\"EnumerateResponse\":{\"EnumerationContext\":\"\"},\"PullResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"KeyManagementItems\":null,\"PublicKeyCertificateItems\":null},\"RefinedPullResponse\":{}}"
	result := response.JSON()
	assert.Equal(t, expectedResult, result)
}

func TestYaml(t *testing.T) {
	response := Response{
		Body: Body{
			AddCertificate_OUTPUT: AddCertificate_OUTPUT{},
		},
	}
	expectedResult := "xmlname:\n    space: \"\"\n    local: \"\"\naddtrustedrootcertificate_output:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    createdcertificate:\n        xmlname:\n            space: \"\"\n            local: \"\"\n        address: \"\"\n        referenceparameters:\n            xmlname:\n                space: \"\"\n                local: \"\"\n            resourceuri: \"\"\n            selectorset:\n                xmlname:\n                    space: \"\"\n                    local: \"\"\n                selectors: []\n    returnvalue: 0\naddcertificate_output:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    createdcertificate:\n        xmlname:\n            space: \"\"\n            local: \"\"\n        address: \"\"\n        referenceparameters:\n            xmlname:\n                space: \"\"\n                local: \"\"\n            resourceuri: \"\"\n            selectorset:\n                xmlname:\n                    space: \"\"\n                    local: \"\"\n                selectors: []\n    returnvalue: 0\naddkey_output:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    createdkey:\n        xmlname:\n            space: \"\"\n            local: \"\"\n        address: \"\"\n        referenceparameters:\n            xmlname:\n                space: \"\"\n                local: \"\"\n            resourceuri: \"\"\n            selectorset:\n                xmlname:\n                    space: \"\"\n                    local: \"\"\n                selectors: []\n    returnvalue: 0\ngeneratekeypair_output:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    keypair:\n        xmlname:\n            space: \"\"\n            local: \"\"\n        address: \"\"\n        referenceparameters:\n            xmlname:\n                space: \"\"\n                local: \"\"\n            resourceuri: \"\"\n            selectorset:\n                xmlname:\n                    space: \"\"\n                    local: \"\"\n                selectors: []\n    returnvalue: 0\ngeneratepkcs10requestex_output:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    signedcertificaterequest: \"\"\n    returnvalue: 0\nkeymanagementgetresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    creationclassname: \"\"\n    elementname: \"\"\n    enableddefault: 0\n    enabledstate: 0\n    name: \"\"\n    operationalstatus: []\n    requestedstate: 0\n    systemcreationclassname: \"\"\n    systemname: \"\"\npublickeycertificategetandputresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    elementname: \"\"\n    instanceid: \"\"\n    x509certificate: \"\"\n    trustedrootcertificate: false\n    issuer: \"\"\n    subject: \"\"\n    readonlycertificate: false\nenumerateresponse:\n    enumerationcontext: \"\"\npullresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    keymanagementitems: []\n    publickeycertificateitems: []\nrefinedpullresponse:\n    keymanagementitems: []\n    publickeycertificateitems: []\n"
	result := response.YAML()
	assert.Equal(t, expectedResult, result)
}

func TestPositiveAMT_PublicKeyCertificate(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.AMTResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/publickey/certificate",
	}
	elementUnderTest := NewPublicKeyCertificateWithClient(wsmanMessageCreator, &client)

	t.Run("amt_PublicKeyCertificate Tests", func(t *testing.T) {
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
				"should create a valid AMT_PublicKeyCertificate Get wsman message",
				AMTPublicKeyCertificate,
				wsmantesting.Get,
				"",
				"<w:SelectorSet><w:Selector Name=\"InstanceID\">Intel(r) AMT Certificate: Handle: 0</w:Selector></w:SelectorSet>",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageGet

					return elementUnderTest.Get("Intel(r) AMT Certificate: Handle: 0")
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PublicKeyCertificateGetAndPutResponse: PublicKeyCertificateResponse{
						XMLName:                xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMTPublicKeyCertificate), Local: "AMT_PublicKeyCertificate"},
						ElementName:            "Intel(r) AMT Certificate",
						InstanceID:             "Intel(r) AMT Certificate: Handle: 0",
						Issuer:                 "C=unknown,O=unknown,CN=MPSRoot-0af1d5",
						Subject:                "C=unknown,O=unknown,CN=MPSRoot-0af1d5",
						TrustedRootCertificate: true,
						X509Certificate:        X509TestCertificate,
					},
				},
			},
			// ENUMERATES
			{
				"should create a valid AMT_PublicKeyCertificate Enumerate wsman message",
				AMTPublicKeyCertificate,
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageEnumerate
					if elementUnderTest.base.WSManMessageCreator == nil {
						logrus.Print("Error")
					}

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "CB000000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create a valid AMT_PublicKeyCertificate Pull wsman message",
				AMTPublicKeyCertificate,
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
						PublicKeyCertificateItems: []PublicKeyCertificateResponse{
							{
								XMLName:                xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMTPublicKeyCertificate), Local: "AMT_PublicKeyCertificate"},
								ElementName:            "Intel(r) AMT Certificate",
								InstanceID:             "Intel(r) AMT Certificate: Handle: 0",
								Issuer:                 "C=unknown,O=unknown,CN=MPSRoot-0af1d5",
								Subject:                "C=unknown,O=unknown,CN=MPSRoot-0af1d5",
								TrustedRootCertificate: true,
								X509Certificate:        X509TestCertificate,
							},
							{
								XMLName:                xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMTPublicKeyCertificate), Local: "AMT_PublicKeyCertificate"},
								ElementName:            "Intel(r) AMT Certificate",
								InstanceID:             "Intel(r) AMT Certificate: Handle: 1",
								Issuer:                 "C=unknown,O=unknown,CN=MPSRoot-0af1d5",
								Subject:                "C=unknown,O=unknown,CN=MPSRoot-0af1d5",
								TrustedRootCertificate: false,
								X509Certificate:        "MIIEOzCCAqOgAwIBAgIDAZMjMA0GCSqGSIb3DQEBDAUAMD0xFzAVBgNVBAMTDk1QU1Jvb3QtMGFmMWQ1MRAwDgYDVQQKEwd1bmtub3duMRAwDgYDVQQGEwd1bmtub3duMCAXDTIyMDkyNDEwNDUwOFoYDzIwNTMwOTI0MTA0NTA4WjA9MRcwFQYDVQQDEw5NUFNSb290LTBhZjFkNTEQMA4GA1UEChMHdW5rbm93bjEQMA4GA1UEBhMHdW5rbm93bjCCAaIwDQYJKoZIhvcNAQEBBQADggGPADCCAYoCggGBALz/oJNyWXlClSlteAieC8Uyd4A+tbn8b45k6LKiImhDmdz/xFo9xe0C9GNf7b42KVpg5WoH/sPhoClR9Tv5i1LnilT1SUir42fcm2NEV9dRcLsPd/RAQfz8u0D4zb3blnxE8isqzriNpG7kac35UidSr5ym8TZ3IwXx6JJuncGgfB0DFZADC/+dA74n3coykvWBYqLr6RI5pkAxvulkRlCsatJTJrvMUYJ51GI28jV56mIAc89sLrHqiSKCZBH9AcUrnZ/cB6ST/IikXpxy5wXBIvWT3VKVq75T/uIoCBEp5TLEn1EOYGqBBOCSQgmtmX7eVaB0s1+ppPW9w9a2zS45cHAtQ7tYvkkPv2dRhSzZdlk6HRXDP5wsF0aiflZCgbrjkq0SFC4e3Lo7XQX3FTNb0SOTZVTydupoMKkgJQTNlcosdu1ZzaIBl3eSkKkJZz2rUTssZC5tn9vcDd5vy3BzcGh5pvkgfAgN1sydqG7Ke1qCkNEzm11B/BsevatjjwIDAQABo0IwQDAMBgNVHRMEBTADAQH/MBEGCWCGSAGG+EIBAQQEAwIABzAdBgNVHQ4EFgQUCvHVQqerCid99eLApuLky9x6H5owDQYJKoZIhvcNAQEMBQADggGBAIzOyGV0hzsmH2biJlzwTZaHMxqS7boTFMkHw+KvzsI201tHqVmCoiQ8EHErBGLSoDOTDRgOUGOCA5XU5ie9OWupAGqKBSwIyAhmJMOzrzC4Gwpu8K1msoFJH30kx/V9purpbS3BRj0xfYXLa6IczbTg3E5IfTnZRJ9YuUtKQfI0P9c5U9CoKtddKn4+lRvOjFDoYfQGCJ7go3xjNCcGCVCjfkUhAVdbQ21DCRr6/YCZDWmjzZpL0p7UKF8roTiNuL/Z7gIXxch5HOmEWHY9uQ6K2MntuxAu0aK/mSD2kwmt/ECongdEGfUvhULLoPRQlQ2LnzcUQEgMECGQR5Yfy9jT0E8zdWDpc2tgVioNu6rEYKgp/GhG+sv7jv58pW82FRAV9xXtftW9+XDugC8tBJ6JHn0Q2v0QAflD2CEQVhWAY8bAqrbfTGUsaLfGL6kxV/qqssoMgLR8Whq96T5le/4XGhQpbCHWIlctD6MwbrsunIAeQKp1Sc3DosY7DLq1MQ==",
							},
							{
								XMLName:                xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMTPublicKeyCertificate), Local: "AMT_PublicKeyCertificate"},
								ElementName:            "Intel(r) AMT Certificate",
								InstanceID:             "Intel(r) AMT Certificate: Handle: 2",
								Issuer:                 "C=unknown,O=unknown,CN=MPSRoot-0af1d5",
								Subject:                "C=unknown,O=unknown,CN=MPSRoot-0af1d5",
								TrustedRootCertificate: true,
								X509Certificate:        "MIIEOzCCAqOgAwIBAgIDAZMjMA0GCSqGSIb3DQEBDAUAMD0xFzAVBgNVBAMTDk1QU1Jvb3QtMGFmMWQ1MRAwDgYDVQQKEwd1bmtub3duMRAwDgYDVQQGEwd1bmtub3duMCAXDTIyMDkyNDEwNDUwOFoYDzIwNTMwOTI0MTA0NTA4WjA9MRcwFQYDVQQDEw5NUFNSb290LTBhZjFkNTEQMA4GA1UEChMHdW5rbm93bjEQMA4GA1UEBhMHdW5rbm93bjCCAaIwDQYJKoZIhvcNAQEBBQADggGPADCCAYoCggGBALz/oJNyWXlClSlteAieC8Uyd4A+tbn8b45k6LKiImhDmdz/xFo9xe0C9GNf7b42KVpg5WoH/sPhoClR9Tv5i1LnilT1SUir42fcm2NEV9dRcLsPd/RAQfz8u0D4zb3blnxE8isqzriNpG7kac35UidSr5ym8TZ3IwXx6JJuncGgfB0DFBADC/+dA74n3coykvWBYqLr6RI5pkAxvulkRlCsatJTJrvMUYJ51GI28jV56mIAc89sLrHqiSKCZBH9AcUrnZ/cB6ST/IikXpxy5wXBIvWT3VKVq75T/uIoCBEp5TLEn1EOYGqBBOCSQgmtmX7eVaB0s1+ppPW9w9a2zS45cHAtQ7tYvkkPv2dRhSzZdlk6HRXDP5wsF0aiflZCgbrjkq0SFC4e3Lo7XQX3FTNb0SOTZVTydupoMKkgJQTNlcosdu1ZzaIBl3eSkKkJZz2rUTssZC5tn9vcDd5vy3BzcGh5pvkgfAgN1sydqG7Ke1qCkNEzm11B/BsevatjjwIDAQABo0IwQDAMBgNVHRMEBTADAQH/MBEGCWCGSAGG+EIBAQQEAwIABzAdBgNVHQ4EFgQUCvHVQqerCid99eLApuLky9x6H5owDQYJKoZIhvcNAQEMBQADggGBAIzOyGV0hzsmH2biJlzwTZaHMxqS7boTFMkHw+KvzsI201tHqVmCoiQ8EHErBGLSoDOTDRgOUGOCA5XU5ie9OWupAGqKBSwIyAhmJMOzrzC4Gwpu8K1msoFJH30kx/V9purpbS3BRj0xfYXLa6IczbTg3E5IfTnZRJ9YuUtKQfI0P9c5U9CoKtddKn4+lRvOjFDoYfQGCJ7go3xjNCcGCVCjfkUhAVdbQ21DCRr6/YCZDWmjzZpL0p7UKF8roTiNuL/Z7gIXxch5HOmEWHY9uQ6K2MntuxAu0aK/mSD2kwmt/ECongdEGfUvhULLoPRQlQ2LnzcUQEgMECGQR5Yfy9jT0E8zdWDpc2tgVioNu6rEYKgp/GhG+sv7jv58pW82FRAV9xXtftW9+XDugC8tBJ6JHn0Q2v0QAflD2CEQVhWAY8bAqrbfTGUsaLfGL6kxV/qqssoMgLR8Whq96T5le/4XGhQpbCHWIlctD6MwbrsunIAeQKp1Sc3DosY7DLq1MQ==",
							},
						},
					},
					RefinedPullResponse: RefinedPullResponse{
						PublicKeyCertificateItems: []RefinedPublicKeyCertificateResponse{
							{
								ElementName:            "Intel(r) AMT Certificate",
								InstanceID:             "Intel(r) AMT Certificate: Handle: 0",
								Issuer:                 "C=unknown,O=unknown,CN=MPSRoot-0af1d5",
								Subject:                "C=unknown,O=unknown,CN=MPSRoot-0af1d5",
								TrustedRootCertificate: true,
								X509Certificate:        X509TestCertificate,
								PublicKeyHandle:        "",
								AssociatedProfiles:     nil,
							},
							{
								ElementName:            "Intel(r) AMT Certificate",
								InstanceID:             "Intel(r) AMT Certificate: Handle: 1",
								Issuer:                 "C=unknown,O=unknown,CN=MPSRoot-0af1d5",
								Subject:                "C=unknown,O=unknown,CN=MPSRoot-0af1d5",
								TrustedRootCertificate: false,
								X509Certificate:        "MIIEOzCCAqOgAwIBAgIDAZMjMA0GCSqGSIb3DQEBDAUAMD0xFzAVBgNVBAMTDk1QU1Jvb3QtMGFmMWQ1MRAwDgYDVQQKEwd1bmtub3duMRAwDgYDVQQGEwd1bmtub3duMCAXDTIyMDkyNDEwNDUwOFoYDzIwNTMwOTI0MTA0NTA4WjA9MRcwFQYDVQQDEw5NUFNSb290LTBhZjFkNTEQMA4GA1UEChMHdW5rbm93bjEQMA4GA1UEBhMHdW5rbm93bjCCAaIwDQYJKoZIhvcNAQEBBQADggGPADCCAYoCggGBALz/oJNyWXlClSlteAieC8Uyd4A+tbn8b45k6LKiImhDmdz/xFo9xe0C9GNf7b42KVpg5WoH/sPhoClR9Tv5i1LnilT1SUir42fcm2NEV9dRcLsPd/RAQfz8u0D4zb3blnxE8isqzriNpG7kac35UidSr5ym8TZ3IwXx6JJuncGgfB0DFZADC/+dA74n3coykvWBYqLr6RI5pkAxvulkRlCsatJTJrvMUYJ51GI28jV56mIAc89sLrHqiSKCZBH9AcUrnZ/cB6ST/IikXpxy5wXBIvWT3VKVq75T/uIoCBEp5TLEn1EOYGqBBOCSQgmtmX7eVaB0s1+ppPW9w9a2zS45cHAtQ7tYvkkPv2dRhSzZdlk6HRXDP5wsF0aiflZCgbrjkq0SFC4e3Lo7XQX3FTNb0SOTZVTydupoMKkgJQTNlcosdu1ZzaIBl3eSkKkJZz2rUTssZC5tn9vcDd5vy3BzcGh5pvkgfAgN1sydqG7Ke1qCkNEzm11B/BsevatjjwIDAQABo0IwQDAMBgNVHRMEBTADAQH/MBEGCWCGSAGG+EIBAQQEAwIABzAdBgNVHQ4EFgQUCvHVQqerCid99eLApuLky9x6H5owDQYJKoZIhvcNAQEMBQADggGBAIzOyGV0hzsmH2biJlzwTZaHMxqS7boTFMkHw+KvzsI201tHqVmCoiQ8EHErBGLSoDOTDRgOUGOCA5XU5ie9OWupAGqKBSwIyAhmJMOzrzC4Gwpu8K1msoFJH30kx/V9purpbS3BRj0xfYXLa6IczbTg3E5IfTnZRJ9YuUtKQfI0P9c5U9CoKtddKn4+lRvOjFDoYfQGCJ7go3xjNCcGCVCjfkUhAVdbQ21DCRr6/YCZDWmjzZpL0p7UKF8roTiNuL/Z7gIXxch5HOmEWHY9uQ6K2MntuxAu0aK/mSD2kwmt/ECongdEGfUvhULLoPRQlQ2LnzcUQEgMECGQR5Yfy9jT0E8zdWDpc2tgVioNu6rEYKgp/GhG+sv7jv58pW82FRAV9xXtftW9+XDugC8tBJ6JHn0Q2v0QAflD2CEQVhWAY8bAqrbfTGUsaLfGL6kxV/qqssoMgLR8Whq96T5le/4XGhQpbCHWIlctD6MwbrsunIAeQKp1Sc3DosY7DLq1MQ==",
								PublicKeyHandle:        "",
								AssociatedProfiles:     nil,
							},
							{
								ElementName:            "Intel(r) AMT Certificate",
								InstanceID:             "Intel(r) AMT Certificate: Handle: 2",
								Issuer:                 "C=unknown,O=unknown,CN=MPSRoot-0af1d5",
								Subject:                "C=unknown,O=unknown,CN=MPSRoot-0af1d5",
								TrustedRootCertificate: true,
								X509Certificate:        "MIIEOzCCAqOgAwIBAgIDAZMjMA0GCSqGSIb3DQEBDAUAMD0xFzAVBgNVBAMTDk1QU1Jvb3QtMGFmMWQ1MRAwDgYDVQQKEwd1bmtub3duMRAwDgYDVQQGEwd1bmtub3duMCAXDTIyMDkyNDEwNDUwOFoYDzIwNTMwOTI0MTA0NTA4WjA9MRcwFQYDVQQDEw5NUFNSb290LTBhZjFkNTEQMA4GA1UEChMHdW5rbm93bjEQMA4GA1UEBhMHdW5rbm93bjCCAaIwDQYJKoZIhvcNAQEBBQADggGPADCCAYoCggGBALz/oJNyWXlClSlteAieC8Uyd4A+tbn8b45k6LKiImhDmdz/xFo9xe0C9GNf7b42KVpg5WoH/sPhoClR9Tv5i1LnilT1SUir42fcm2NEV9dRcLsPd/RAQfz8u0D4zb3blnxE8isqzriNpG7kac35UidSr5ym8TZ3IwXx6JJuncGgfB0DFBADC/+dA74n3coykvWBYqLr6RI5pkAxvulkRlCsatJTJrvMUYJ51GI28jV56mIAc89sLrHqiSKCZBH9AcUrnZ/cB6ST/IikXpxy5wXBIvWT3VKVq75T/uIoCBEp5TLEn1EOYGqBBOCSQgmtmX7eVaB0s1+ppPW9w9a2zS45cHAtQ7tYvkkPv2dRhSzZdlk6HRXDP5wsF0aiflZCgbrjkq0SFC4e3Lo7XQX3FTNb0SOTZVTydupoMKkgJQTNlcosdu1ZzaIBl3eSkKkJZz2rUTssZC5tn9vcDd5vy3BzcGh5pvkgfAgN1sydqG7Ke1qCkNEzm11B/BsevatjjwIDAQABo0IwQDAMBgNVHRMEBTADAQH/MBEGCWCGSAGG+EIBAQQEAwIABzAdBgNVHQ4EFgQUCvHVQqerCid99eLApuLky9x6H5owDQYJKoZIhvcNAQEMBQADggGBAIzOyGV0hzsmH2biJlzwTZaHMxqS7boTFMkHw+KvzsI201tHqVmCoiQ8EHErBGLSoDOTDRgOUGOCA5XU5ie9OWupAGqKBSwIyAhmJMOzrzC4Gwpu8K1msoFJH30kx/V9purpbS3BRj0xfYXLa6IczbTg3E5IfTnZRJ9YuUtKQfI0P9c5U9CoKtddKn4+lRvOjFDoYfQGCJ7go3xjNCcGCVCjfkUhAVdbQ21DCRr6/YCZDWmjzZpL0p7UKF8roTiNuL/Z7gIXxch5HOmEWHY9uQ6K2MntuxAu0aK/mSD2kwmt/ECongdEGfUvhULLoPRQlQ2LnzcUQEgMECGQR5Yfy9jT0E8zdWDpc2tgVioNu6rEYKgp/GhG+sv7jv58pW82FRAV9xXtftW9+XDugC8tBJ6JHn0Q2v0QAflD2CEQVhWAY8bAqrbfTGUsaLfGL6kxV/qqssoMgLR8Whq96T5le/4XGhQpbCHWIlctD6MwbrsunIAeQKp1Sc3DosY7DLq1MQ==",
								PublicKeyHandle:        "",
								AssociatedProfiles:     nil,
							},
						},
					},
				},
			},

			// PUTS
			{
				"should create a valid AMT_PublicKeyCertificate Put wsman message",
				AMTPublicKeyCertificate,
				wsmantesting.Put,
				"<h:AMT_PublicKeyCertificate xmlns:h=\"http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyCertificate\"><h:ElementName></h:ElementName><h:InstanceID></h:InstanceID><h:X509Certificate>MIIEOzCCAqOgAwIBAgIDAZMjMA0GCSqGSIb3DQEBDAUAMD0xFzAVBgNVBAMTDk1QU1Jvb3QtMGFmMWQ1MRAwDgYDVQQKEwd1bmtub3duMRAwDgYDVQQGEwd1bmtub3duMCAXDTIyMDkyNDEwNDUwOFoYDzIwNTMwOTI0MTA0NTA4WjA9MRcwFQYDVQQDEw5NUFNSb290LTBhZjFkNTEQMA4GA1UEChMHdW5rbm93bjEQMA4GA1UEBhMHdW5rbm93bjCCAaIwDQYJKoZIhvcNAQEBBQADggGPADCCAYoCggGBALz/oJNyWXlClSlteAieC8Uyd4A+tbn8b45k6LKiImhDmdz/xFo9xe0C9GNf7b42KVpg5WoH/sPhoClR9Tv5i1LnilT1SUir42fcm2NEV9dRcLsPd/RAQfz8u0D4zb3blnxE8isqzriNpG7kac35UidSr5ym8TZ3IwXx6JJuncGgfB0DFZADC/+dA74n3coykvWBYqLr6RI5pkAxvulkRlCsatJTJrvMUYJ51GI28jV56mIAc89sLrHqiSKCZBH9AcUrnZ/cB6ST/IikXpxy5wXBIvWT3VKVq75T/uIoCBEp5TLEn1EOYGqBBOCSQgmtmX7eVaB0s1+ppPW9w9a2zS45cHAtQ7tYvkkPv2dRhSzZdlk6HRXDP5wsF0aiflZCgbrjkq0SFC4e3Lo7XQX3FTNb0SOTZVTydupoMKkgJQTNlcosdu1ZzaIBl3eSkKkJZz2rUTssZC5tn9vcDd5vy3BzcGh5pvkgfAgN1sydqG7Ke1qCkNEzm11B/BsevatjjwIDAQABo0IwQDAMBgNVHRMEBTADAQH/MBEGCWCGSAGG+EIBAQQEAwIABzAdBgNVHQ4EFgQUCvHVQqerCid99eLApuLky9x6H5owDQYJKoZIhvcNAQEMBQADggGBAIzOyGV0hzsmH2biJlzwTZaHMxqS7boTFMkHw+KvzsI201tHqVmCoiQ8EHErBGLSoDOTDRgOUGOCA5XU5ie9OWupAGqKBSwIyAhmJMOzrzC4Gwpu8K1msoFJH30kx/V9purpbS3BRj0xfYXLa6IczbTg3E5IfTnZRJ9YuUtKQfI0P9c5U9CoKtddKn4+lRvOjFDoYfQGCJ7go3xjNCcGCVCjfkUhAVdbQ21DCRr6/YCZDWmjzZpL0p7UKF8roTiNuL/Z7gIXxch5HOmEWHY9uQ6K2MntuxAu0aK/mSD2kwmt/ECongdEGfUvhULLoPRQlQ2LnzcUQEgMECGQR5Yfy9jT0E8zdWDpc2tgVioNu6rEYKgp/GhG+sv7jv58pW82FRAV9xXtftW9+XDugC8tBJ6JHn0Q2v0QAflD2CEQVhWAY8bAqrbfTGUsaLfGL6kxV/qqssoMgLR8WhQ96T5le/4XGhQpbCHWIlctD6MwbrsunIAeQKp1Sc3DosY7DLq1MQ==</h:X509Certificate><h:TrustedRootCertificate>false</h:TrustedRootCertificate><h:Issuer></h:Issuer><h:Subject></h:Subject><h:ReadOnlyCertificate>false</h:ReadOnlyCertificate></h:AMT_PublicKeyCertificate>",
				"<w:SelectorSet><w:Selector Name=\"InstanceID\">Intel(r) AMT Certificate: Handle: 0</w:Selector></w:SelectorSet>",
				func() (Response, error) {
					X509Certificate := X509TestCertificate
					client.CurrentMessage = wsmantesting.CurrentMessagePut

					return elementUnderTest.Put("Intel(r) AMT Certificate: Handle: 0", X509Certificate)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PublicKeyCertificateGetAndPutResponse: PublicKeyCertificateResponse{
						XMLName:                xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyCertificate", Local: "AMT_PublicKeyCertificate"},
						ElementName:            "Intel(r) AMT Certificate",
						InstanceID:             "Intel(r) AMT Certificate: Handle: 0",
						Issuer:                 "C=unknown,O=unknown,CN=MPSRoot-0af1d5",
						Subject:                "C=unknown,O=unknown,CN=MPSRoot-0af1d5",
						TrustedRootCertificate: true,
						X509Certificate:        X509TestCertificate,
					},
				},
			},
			// DELETE
			{
				"should create a valid AMT_PublicKeyCertificate Delete wsman message",
				AMTPublicKeyCertificate,
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

func TestNegativeAMT_PublicKeyCertificate(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.AMTResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/publickey/certificate",
	}
	elementUnderTest := NewPublicKeyCertificateWithClient(wsmanMessageCreator, &client)

	t.Run("amt_PublicKeyCertificate Tests", func(t *testing.T) {
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
				"should create a valid AMT_PublicKeyCertificate Get wsman message",
				AMTPublicKeyCertificate,
				wsmantesting.Get,
				"",
				"<w:SelectorSet><w:Selector Name=\"InstanceID\">Intel(r) AMT Certificate: Handle: 0</w:Selector></w:SelectorSet>",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Get("Intel(r) AMT Certificate: Handle: 0")
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PublicKeyCertificateGetAndPutResponse: PublicKeyCertificateResponse{
						XMLName:                xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMTPublicKeyCertificate), Local: "AMT_PublicKeyCertificate"},
						ElementName:            "Intel(r) AMT Certificate",
						InstanceID:             "Intel(r) AMT Certificate: Handle: 0",
						Issuer:                 "C=unknown,O=unknown,CN=MPSRoot-0af1d5",
						Subject:                "C=unknown,O=unknown,CN=MPSRoot-0af1d5",
						TrustedRootCertificate: true,
						X509Certificate:        X509TestCertificate,
					},
				},
			},
			// ENUMERATES
			{
				"should create a valid AMT_PublicKeyCertificate Enumerate wsman message",
				AMTPublicKeyCertificate,
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError
					if elementUnderTest.base.WSManMessageCreator == nil {
						logrus.Print("Error")
					}

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "CB000000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create a valid AMT_PublicKeyCertificate Pull wsman message",
				AMTPublicKeyCertificate,
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
						PublicKeyCertificateItems: []PublicKeyCertificateResponse{
							{
								XMLName:                xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMTPublicKeyCertificate), Local: "AMT_PublicKeyCertificate"},
								ElementName:            "Intel(r) AMT Certificate",
								InstanceID:             "Intel(r) AMT Certificate: Handle: 0",
								Issuer:                 "C=unknown,O=unknown,CN=MPSRoot-0af1d5",
								Subject:                "C=unknown,O=unknown,CN=MPSRoot-0af1d5",
								TrustedRootCertificate: true,
								X509Certificate:        X509TestCertificate,
							},
							{
								XMLName:                xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMTPublicKeyCertificate), Local: "AMT_PublicKeyCertificate"},
								ElementName:            "Intel(r) AMT Certificate",
								InstanceID:             "Intel(r) AMT Certificate: Handle: 1",
								Issuer:                 "C=unknown,O=unknown,CN=MPSRoot-0af1d5",
								Subject:                "C=unknown,O=unknown,CN=MPSRoot-0af1d5",
								TrustedRootCertificate: false,
								X509Certificate:        X509TestCertificate,
							},
							{
								XMLName:                xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMTPublicKeyCertificate), Local: "AMT_PublicKeyCertificate"},
								ElementName:            "Intel(r) AMT Certificate",
								InstanceID:             "Intel(r) AMT Certificate: Handle: 2",
								Issuer:                 "C=unknown,O=unknown,CN=MPSRoot-0af1d5",
								Subject:                "C=unknown,O=unknown,CN=MPSRoot-0af1d5",
								TrustedRootCertificate: true,
								X509Certificate:        X509TestCertificate,
							},
						},
					},
				},
			},

			// PUTS
			{
				"should create a valid AMT_PublicKeyCertificate Put wsman message",
				AMTPublicKeyCertificate,
				wsmantesting.Put,
				"<h:AMT_PublicKeyCertificate xmlns:h=\"http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyCertificate\"><h:ElementName></h:ElementName><h:InstanceID></h:InstanceID><h:X509Certificate>MIIEOzCCAqOgAwIBAgIDAZMjMA0GCSqGSIb3DQEBDAUAMD0xFzAVBgNVBAMTDk1QU1Jvb3QtMGFmMWQ1MRAwDgYDVQQKEwd1bmtub3duMRAwDgYDVQQGEwd1bmtub3duMCAXDTIyMDkyNDEwNDUwOFoYDzIwNTMwOTI0MTA0NTA4WjA9MRcwFQYDVQQDEw5NUFNSb290LTBhZjFkNTEQMA4GA1UEChMHdW5rbm93bjEQMA4GA1UEBhMHdW5rbm93bjCCAaIwDQYJKoZIhvcNAQEBBQADggGPADCCAYoCggGBALz/oJNyWXlClSlteAieC8Uyd4A+tbn8b45k6LKiImhDmdz/xFo9xe0C9GNf7b42KVpg5WoH/sPhoClR9Tv5i1LnilT1SUir42fcm2NEV9dRcLsPd/RAQfz8u0D4zb3blnxE8isqzriNpG7kac35UidSr5ym8TZ3IwXx6JJuncGgfB0DFZADC/+dA74n3coykvWBYqLr6RI5pkAxvulkRlCsatJTJrvMUYJ51GI28jV56mIAc89sLrHqiSKCZBH9AcUrnZ/cB6ST/IikXpxy5wXBIvWT3VKVq75T/uIoCBEp5TLEn1EOYGqBBOCSQgmtmX7eVaB0s1+ppPW9w9a2zS45cHAtQ7tYvkkPv2dRhSzZdlk6HRXDP5wsF0aiflZCgbrjkq0SFC4e3Lo7XQX3FTNb0SOTZVTydupoMKkgJQTNlcosdu1ZzaIBl3eSkKkJZz2rUTssZC5tn9vcDd5vy3BzcGh5pvkgfAgN1sydqG7Ke1qCkNEzm11B/BsevatjjwIDAQABo0IwQDAMBgNVHRMEBTADAQH/MBEGCWCGSAGG+EIBAQQEAwIABzAdBgNVHQ4EFgQUCvHVQqerCid99eLApuLky9x6H5owDQYJKoZIhvcNAQEMBQADggGBAIzOyGV0hzsmH2biJlzwTZaHMxqS7boTFMkHw+KvzsI201tHqVmCoiQ8EHErBGLSoDOTDRgOUGOCA5XU5ie9OWupAGqKBSwIyAhmJMOzrzC4Gwpu8K1msoFJH30kx/V9purpbS3BRj0xfYXLa6IczbTg3E5IfTnZRJ9YuUtKQfI0P9c5U9CoKtddKn4+lRvOjFDoYfQGCJ7go3xjNCcGCVCjfkUhAVdbQ21DCRr6/YCZDWmjzZpL0p7UKF8roTiNuL/Z7gIXxch5HOmEWHY9uQ6K2MntuxAu0aK/mSD2kwmt/ECongdEGfUvhULLoPRQlQ2LnzcUQEgMECGQR5Yfy9jT0E8zdWDpc2tgVioNu6rEYKgp/GhG+sv7jv58pW82FRAV9xXtftW9+XDugC8tBJ6JHn0Q2v0QAflD2CEQVhWAY8bAqrbfTGUsaLfGL6kxV/qqssoMgLR8WhQ96T5le/4XGhQpbCHWIlctD6MwbrsunIAeQKp1Sc3DosY7DLq1MQ==</h:X509Certificate><h:TrustedRootCertificate>false</h:TrustedRootCertificate><h:Issuer></h:Issuer><h:Subject></h:Subject><h:ReadOnlyCertificate>false</h:ReadOnlyCertificate></h:AMT_PublicKeyCertificate>",
				"<w:SelectorSet><w:Selector Name=\"InstanceID\">Intel(r) AMT Certificate: Handle: 0</w:Selector></w:SelectorSet>",
				func() (Response, error) {
					X509Certificate := X509TestCertificate
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Put("Intel(r) AMT Certificate: Handle: 0", X509Certificate)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PublicKeyCertificateGetAndPutResponse: PublicKeyCertificateResponse{
						XMLName:                xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyCertificate", Local: "AMT_PublicKeyCertificate"},
						ElementName:            "Intel(r) AMT Certificate",
						InstanceID:             "Intel(r) AMT Certificate: Handle: 0",
						Issuer:                 "C=unknown,O=unknown,CN=MPSRoot-0af1d5",
						Subject:                "C=unknown,O=unknown,CN=MPSRoot-0af1d5",
						TrustedRootCertificate: true,
						X509Certificate:        X509TestCertificate,
					},
				},
			},
			// DELETE
			{
				"should create a valid AMT_PublicKeyCertificate Delete wsman message",
				AMTPublicKeyCertificate,
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
