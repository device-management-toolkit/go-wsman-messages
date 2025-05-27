/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package hostbasedsetup

import (
	"encoding/xml"
	"fmt"
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
	expectedResult := "{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"GetResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"ElementName\":\"\",\"SystemCreationClassName\":\"\",\"SystemName\":\"\",\"CreationClassName\":\"\",\"Name\":\"\",\"CurrentControlMode\":0,\"AllowedControlModes\":null,\"ConfigurationNonce\":\"\",\"CertChainStatus\":0},\"EnumerateResponse\":{\"EnumerationContext\":\"\"},\"PullResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"HostBasedSetupServiceItems\":null},\"Setup_OUTPUT\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"ReturnValue\":0},\"AdminSetup_OUTPUT\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"ReturnValue\":0},\"AddNextCertInChain_OUTPUT\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"ReturnValue\":0},\"UpgradeClientToAdmin_OUTPUT\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"ReturnValue\":0}}"
	result := response.JSON()
	assert.Equal(t, expectedResult, result)
}

func TestYaml(t *testing.T) {
	response := Response{
		Body: Body{
			PullResponse: PullResponse{},
		},
	}
	expectedResult := "xmlname:\n    space: \"\"\n    local: \"\"\ngetresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    elementname: \"\"\n    systemcreationclassname: \"\"\n    systemname: \"\"\n    creationclassname: \"\"\n    name: \"\"\n    currentcontrolmode: 0\n    allowedcontrolmodes: []\n    configurationnonce: \"\"\n    certchainstatus: 0\nenumerateresponse:\n    enumerationcontext: \"\"\npullresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    hostbasedsetupserviceitems: []\nsetup_output:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    returnvalue: 0\nadminsetup_output:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    returnvalue: 0\naddnextcertinchain_output:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    returnvalue: 0\nupgradeclienttoadmin_output:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    returnvalue: 0\n"
	result := response.YAML()
	assert.Equal(t, expectedResult, result)
}

func TestPositiveIPS_HostBasedSetupService(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.IPSResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "ips/hostbasedsetup",
	}
	elementUnderTest := NewHostBasedSetupServiceWithClient(wsmanMessageCreator, &client)

	t.Run("ips_HostBasedSetupService Tests", func(t *testing.T) {
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
				"should create a valid IPS_HostBasedSetupService Get wsman message",
				IPSHostBasedSetupService,
				wsmantesting.Get,
				"",
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageGet

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetResponse: HostBasedSetupService{
						XMLName:                 xml.Name{Space: fmt.Sprintf("%s%s", message.IPSSchema, IPSHostBasedSetupService), Local: IPSHostBasedSetupService},
						AllowedControlModes:     []AllowedControlModes{AllowedControlModesAdmin, AllowedControlModesClient},
						CertChainStatus:         CertChainStatusNotStarted,
						ConfigurationNonce:      "4P3sY7swlhjkhJNxDkEBIUcmpHE=",
						CreationClassName:       IPSHostBasedSetupService,
						CurrentControlMode:      Admin,
						ElementName:             "Intel(r) AMT Host Based Setup Service",
						Name:                    "Intel(r) AMT Host Based Setup Service",
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "Intel(r) AMT",
					},
				},
			},
			// ENUMERATES
			{
				"should create a valid IPS_HostBasedSetupService Enumerate wsman message",
				IPSHostBasedSetupService,
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
						EnumerationContext: "980A0000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create a valid IPS_HostBasedSetupService Pull wsman message",
				IPSHostBasedSetupService,
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
						HostBasedSetupServiceItems: []HostBasedSetupService{
							{
								XMLName:                 xml.Name{Space: fmt.Sprintf("%s%s", message.IPSSchema, IPSHostBasedSetupService), Local: IPSHostBasedSetupService},
								AllowedControlModes:     []AllowedControlModes{AllowedControlModesAdmin, AllowedControlModesClient},
								CertChainStatus:         CertChainStatusNotStarted,
								ConfigurationNonce:      "4P3sY7swlhjkhJNxDkEBIUcmpHE=",
								CreationClassName:       IPSHostBasedSetupService,
								CurrentControlMode:      Admin,
								ElementName:             "Intel(r) AMT Host Based Setup Service",
								Name:                    "Intel(r) AMT Host Based Setup Service",
								SystemCreationClassName: "CIM_ComputerSystem",
								SystemName:              "Intel(r) AMT",
							},
						},
					},
				},
			},

			// ADD NEXT CERT IN CHAIN
			{
				"should create a valid IPS_HostBasedSetupService AddNextCertInChain wsman message",
				"IPS_HostBasedSetupService",
				wsmantesting.AddNextCertInChain,
				fmt.Sprintf(`<h:AddNextCertInChain_INPUT xmlns:h="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HostBasedSetupService"><h:NextCertificate>%s</h:NextCertificate><h:IsLeafCertificate>true</h:IsLeafCertificate><h:IsRootCertificate>false</h:IsRootCertificate></h:AddNextCertInChain_INPUT>`, wsmantesting.ClientCertificate),
				"",
				func() (Response, error) {
					client.CurrentMessage = "AddNextCertInChain"

					return elementUnderTest.AddNextCertInChain(wsmantesting.ClientCertificate, true, false)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					AddNextCertInChain_OUTPUT: AddNextCertInChain_OUTPUT{
						XMLName:     xml.Name{Space: fmt.Sprintf("%s%s", message.IPSSchema, IPSHostBasedSetupService), Local: "AddNextCertInChain_OUTPUT"},
						ReturnValue: 0,
					},
				},
			},

			// AdminSetup
			{
				"should create a valid IPS_HostBasedSetupService AdminSetup wsman message",
				"IPS_HostBasedSetupService",
				wsmantesting.AdminSetup,
				fmt.Sprintf(`<h:AdminSetup_INPUT xmlns:h="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HostBasedSetupService"><h:NetAdminPassEncryptionType>%d</h:NetAdminPassEncryptionType><h:NetworkAdminPassword>%s</h:NetworkAdminPassword><h:McNonce>%s</h:McNonce><h:SigningAlgorithm>%d</h:SigningAlgorithm><h:DigitalSignature>%s</h:DigitalSignature></h:AdminSetup_INPUT>`, wsmantesting.AdminPassEncryptionType, "f73b2c17b1ecbd7a235ec37d66cbed71", wsmantesting.MCNonce, wsmantesting.SigningAlgorithm, wsmantesting.DigitalSignature),
				"",
				func() (Response, error) {
					client.CurrentMessage = "AdminSetup"

					return elementUnderTest.AdminSetup(wsmantesting.AdminPassEncryptionType, wsmantesting.DigestRealm, wsmantesting.AdminPassword, wsmantesting.MCNonce, wsmantesting.SigningAlgorithm, wsmantesting.DigitalSignature)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					AdminSetup_OUTPUT: AdminSetup_OUTPUT{
						XMLName:     xml.Name{Space: fmt.Sprintf("%s%s", message.IPSSchema, IPSHostBasedSetupService), Local: "AdminSetup_OUTPUT"},
						ReturnValue: 0,
					},
				},
			},

			// UpgradeToAdminSetup
			{
				"should create a valid IPS_HostBasedSetupService UpgradeToAdminSetup wsman message",
				"IPS_HostBasedSetupService",
				wsmantesting.UpgradeClientToAdmin,
				fmt.Sprintf(`<h:UpgradeClientToAdmin_INPUT xmlns:h="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HostBasedSetupService"><h:McNonce>%s</h:McNonce><h:SigningAlgorithm>%d</h:SigningAlgorithm><h:DigitalSignature>%s</h:DigitalSignature></h:UpgradeClientToAdmin_INPUT>`, wsmantesting.MCNonce, wsmantesting.SigningAlgorithm, wsmantesting.DigitalSignature),
				"",
				func() (Response, error) {
					client.CurrentMessage = "UpgradeClientToAdmin"

					return elementUnderTest.UpgradeClientToAdmin(wsmantesting.MCNonce, wsmantesting.SigningAlgorithm, wsmantesting.DigitalSignature)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					UpgradeClientToAdmin_OUTPUT: UpgradeClientToAdmin_OUTPUT{
						XMLName:     xml.Name{Space: fmt.Sprintf("%s%s", message.IPSSchema, IPSHostBasedSetupService), Local: "UpgradeClientToAdmin_OUTPUT"},
						ReturnValue: 0,
					},
				},
			},

			// Setup
			{
				"should create a valid IPS_HostBasedSetupService Setup wsman message",
				"IPS_HostBasedSetupService",
				wsmantesting.Setup,
				fmt.Sprintf(`<h:Setup_INPUT xmlns:h="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HostBasedSetupService"><h:NetAdminPassEncryptionType>%d</h:NetAdminPassEncryptionType><h:NetworkAdminPassword>%s</h:NetworkAdminPassword></h:Setup_INPUT>`, wsmantesting.AdminPassEncryptionType, "f73b2c17b1ecbd7a235ec37d66cbed71"),
				"",
				func() (Response, error) {
					client.CurrentMessage = "Setup"

					return elementUnderTest.Setup(wsmantesting.AdminPassEncryptionType, wsmantesting.DigestRealm, wsmantesting.AdminPassword)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					Setup_OUTPUT: Setup_OUTPUT{
						XMLName:     xml.Name{Space: fmt.Sprintf("%s%s", message.IPSSchema, IPSHostBasedSetupService), Local: "Setup_OUTPUT"},
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

func TestNegativeIPS_HostBasedSetupService(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.IPSResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "ips/hostbasedsetup",
	}
	elementUnderTest := NewHostBasedSetupServiceWithClient(wsmanMessageCreator, &client)

	t.Run("ips_HostBasedSetupService Tests", func(t *testing.T) {
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
				"should create a valid IPS_HostBasedSetupService Get wsman message",
				IPSHostBasedSetupService,
				wsmantesting.Get,
				"",
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetResponse: HostBasedSetupService{
						XMLName:                 xml.Name{Space: fmt.Sprintf("%s%s", message.IPSSchema, IPSHostBasedSetupService), Local: IPSHostBasedSetupService},
						AllowedControlModes:     []AllowedControlModes{AllowedControlModesAdmin, AllowedControlModesClient},
						CertChainStatus:         CertChainStatusNotStarted,
						ConfigurationNonce:      "4P3sY7swlhjkhJNxDkEBIUcmpHE=",
						CreationClassName:       IPSHostBasedSetupService,
						CurrentControlMode:      Admin,
						ElementName:             "Intel(r) AMT Host Based Setup Service",
						Name:                    "Intel(r) AMT Host Based Setup Service",
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "Intel(r) AMT",
					},
				},
			},
			// ENUMERATES
			{
				"should create a valid IPS_HostBasedSetupService Enumerate wsman message",
				IPSHostBasedSetupService,
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
						EnumerationContext: "980A0000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create a valid IPS_HostBasedSetupService Pull wsman message",
				IPSHostBasedSetupService,
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
						HostBasedSetupServiceItems: []HostBasedSetupService{
							{
								XMLName:                 xml.Name{Space: fmt.Sprintf("%s%s", message.IPSSchema, IPSHostBasedSetupService), Local: IPSHostBasedSetupService},
								AllowedControlModes:     []AllowedControlModes{AllowedControlModesAdmin, AllowedControlModesClient},
								CertChainStatus:         CertChainStatusNotStarted,
								ConfigurationNonce:      "4P3sY7swlhjkhJNxDkEBIUcmpHE=",
								CreationClassName:       IPSHostBasedSetupService,
								CurrentControlMode:      Admin,
								ElementName:             "Intel(r) AMT Host Based Setup Service",
								Name:                    "Intel(r) AMT Host Based Setup Service",
								SystemCreationClassName: "CIM_ComputerSystem",
								SystemName:              "Intel(r) AMT",
							},
						},
					},
				},
			},

			// ADD NEXT CERT IN CHAIN
			{
				"should create a valid IPS_HostBasedSetupService AddNextCertInChain wsman message",
				"IPS_HostBasedSetupService",
				wsmantesting.AddNextCertInChain,
				fmt.Sprintf(`<h:AddNextCertInChain_INPUT xmlns:h="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HostBasedSetupService"><h:NextCertificate>%s</h:NextCertificate><h:IsLeafCertificate>true</h:IsLeafCertificate><h:IsRootCertificate>false</h:IsRootCertificate></h:AddNextCertInChain_INPUT>`, wsmantesting.ClientCertificate),
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.AddNextCertInChain(wsmantesting.ClientCertificate, true, false)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					AddNextCertInChain_OUTPUT: AddNextCertInChain_OUTPUT{
						XMLName:     xml.Name{Space: fmt.Sprintf("%s%s", message.IPSSchema, IPSHostBasedSetupService), Local: "AddNextCertInChain_OUTPUT"},
						ReturnValue: 0,
					},
				},
			},

			// AdminSetup
			{
				"should create a valid IPS_HostBasedSetupService AdminSetup wsman message",
				"IPS_HostBasedSetupService",
				wsmantesting.AdminSetup,
				fmt.Sprintf(`<h:AdminSetup_INPUT xmlns:h="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HostBasedSetupService"><h:NetAdminPassEncryptionType>%d</h:NetAdminPassEncryptionType><h:NetworkAdminPassword>%s</h:NetworkAdminPassword><h:McNonce>%s</h:McNonce><h:SigningAlgorithm>%d</h:SigningAlgorithm><h:DigitalSignature>%s</h:DigitalSignature></h:AdminSetup_INPUT>`, wsmantesting.AdminPassEncryptionType, "f73b2c17b1ecbd7a235ec37d66cbed71", wsmantesting.MCNonce, wsmantesting.SigningAlgorithm, wsmantesting.DigitalSignature),
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.AdminSetup(wsmantesting.AdminPassEncryptionType, wsmantesting.DigestRealm, wsmantesting.AdminPassword, wsmantesting.MCNonce, wsmantesting.SigningAlgorithm, wsmantesting.DigitalSignature)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					AdminSetup_OUTPUT: AdminSetup_OUTPUT{
						XMLName:     xml.Name{Space: fmt.Sprintf("%s%s", message.IPSSchema, IPSHostBasedSetupService), Local: "AdminSetup_OUTPUT"},
						ReturnValue: 0,
					},
				},
			},

			// UpgradeToAdminSetup
			{
				"should create a valid IPS_HostBasedSetupService UpgradeToAdminSetup wsman message",
				"IPS_HostBasedSetupService",
				wsmantesting.UpgradeClientToAdmin,
				fmt.Sprintf(`<h:UpgradeClientToAdmin_INPUT xmlns:h="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HostBasedSetupService"><h:McNonce>%s</h:McNonce><h:SigningAlgorithm>%d</h:SigningAlgorithm><h:DigitalSignature>%s</h:DigitalSignature></h:UpgradeClientToAdmin_INPUT>`, wsmantesting.MCNonce, wsmantesting.SigningAlgorithm, wsmantesting.DigitalSignature),
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.UpgradeClientToAdmin(wsmantesting.MCNonce, wsmantesting.SigningAlgorithm, wsmantesting.DigitalSignature)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					UpgradeClientToAdmin_OUTPUT: UpgradeClientToAdmin_OUTPUT{
						XMLName:     xml.Name{Space: fmt.Sprintf("%s%s", message.IPSSchema, IPSHostBasedSetupService), Local: "UpgradeClientToAdmin_OUTPUT"},
						ReturnValue: 0,
					},
				},
			},

			// Setup
			{
				"should create a valid IPS_HostBasedSetupService Setup wsman message",
				"IPS_HostBasedSetupService",
				wsmantesting.Setup,
				fmt.Sprintf(`<h:Setup_INPUT xmlns:h="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HostBasedSetupService"><h:NetAdminPassEncryptionType>%d</h:NetAdminPassEncryptionType><h:NetworkAdminPassword>%s</h:NetworkAdminPassword></h:Setup_INPUT>`, wsmantesting.AdminPassEncryptionType, "f73b2c17b1ecbd7a235ec37d66cbed71"),
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Setup(wsmantesting.AdminPassEncryptionType, wsmantesting.DigestRealm, wsmantesting.AdminPassword)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					Setup_OUTPUT: Setup_OUTPUT{
						XMLName:     xml.Name{Space: fmt.Sprintf("%s%s", message.IPSSchema, IPSHostBasedSetupService), Local: "Setup_OUTPUT"},
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

func TestCreateMD5Hash(t *testing.T) {
	tests := []struct {
		adminPassword string
		digestRealm   string
		expected      string
	}{
		{"adminPassword1", "digestRealm1", "7eab95087308c968d56947a05e916d6b"},
		{"adminPassword2", "digestRealm2", "b404159c55fafd0b4a8e7d64833c7f26"},
	}
	for _, test := range tests {
		t.Run(test.digestRealm, func(t *testing.T) {
			result := createMD5Hash(test.adminPassword, test.digestRealm)
			assert.Equal(t, test.expected, result)
		})
	}
}
