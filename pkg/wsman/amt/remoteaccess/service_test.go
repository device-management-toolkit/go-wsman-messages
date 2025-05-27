/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package remoteaccess

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/amt/methods"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/wsmantesting"
)

const (
	EnvelopeResponse = `<a:Envelope xmlns:a="http://www.w3.org/2003/05/soap-envelope" x-mlns:b="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:c="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns:d="http://schemas.xmlsoap.org/ws/2005/02/trust" xmlns:e="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" xmlns:f="http://schemas.dmtf.org/wbem/wsman/1/cimbinding.xsd" xmlns:g="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService" xmlns:h="http://schemas.dmtf.org/wbem/wscim/1/common" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><a:Header><b:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:To><b:RelatesTo>0</b:RelatesTo><b:Action a:mustUnderstand="true">`
	GetBody          = `<g:AMT_AuthorizationService><g:CreationClassName>AMT_RemoteAccessService</g:CreationClassName><g:ElementName>Intel(r) AMT Remote Access Service</g:ElementName><g:Name>Intel(r) AMT Remote Access Service</g:Name><g:SystemCreationClassName>CIM_ComputerSystem</g:SystemCreationClassName><g:SystemName>ManagedSystem</g:SystemName></g:AMT Remote Access Service>`
)

func TestPositiveAMT_RemoteAccessService(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.AMTResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/remoteaccess/service",
	}
	elementUnderTest := NewRemoteAccessServiceWithClient(wsmanMessageCreator, &client)

	t.Run("amt_RemoteAccessService Tests", func(t *testing.T) {
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
				"should create a valid AMT_RemoteAccessService Get wsman message",
				AMTRemoteAccessService,
				wsmantesting.Get,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageGet

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					RemoteAccessServiceGetResponse: RemoteAccessServiceResponse{
						XMLName:                 xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessService", Local: "AMT_RemoteAccessService"},
						CreationClassName:       "AMT_RemoteAccessService",
						ElementName:             "Intel(r) AMT Remote Access Service",
						Name:                    "Intel(r) AMT Remote Access Service",
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "Intel(r) AMT",
					},
				},
			},
			// ENUMERATES
			{
				"should create a valid AMT_RemoteAccessService Enumerate wsman message",
				AMTRemoteAccessService,
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
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
						EnumerationContext: "D3000000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create a valid AMT_RemoteAccessService Pull wsman message",
				AMTRemoteAccessService,
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
						RemoteAccessItems: []RemoteAccessServiceResponse{
							{
								XMLName:                 xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessService", Local: "AMT_RemoteAccessService"},
								CreationClassName:       "AMT_RemoteAccessService",
								ElementName:             "Intel(r) AMT Remote Access Service",
								Name:                    "Intel(r) AMT Remote Access Service",
								SystemCreationClassName: "CIM_ComputerSystem",
								SystemName:              "Intel(r) AMT",
							},
						},
					},
				},
			},
			// AddMPS
			{
				"should create a valid AMT_RemoteAccessService AddMPS wsman message",
				AMTRemoteAccessService,
				methods.GenerateAction(AMTRemoteAccessService, AddMps),
				fmt.Sprintf(`<h:AddMpServer_INPUT xmlns:h="%s%s"><h:AccessInfo>%s</h:AccessInfo><h:InfoFormat>%d</h:InfoFormat><h:Port>%d</h:Port><h:AuthMethod>%d</h:AuthMethod><h:Username>%s</h:Username><h:Password>%s</h:Password><h:CN>%s</h:CN></h:AddMpServer_INPUT>`, resourceURIBase, AMTRemoteAccessService, "AccessInfo", 1, 2, 3, "Username", "Password", "CommonName"),
				func() (Response, error) {
					client.CurrentMessage = "AddMPSServer"
					mpsServer := AddMpServerRequest{
						H:          "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessService",
						AccessInfo: "AccessInfo",
						InfoFormat: 1,
						Port:       2,
						AuthMethod: 3,
						Username:   "Username",
						Password:   "Password",
						CommonName: "CommonName",
					}

					return elementUnderTest.AddMPS(mpsServer)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					AddMpServerResponse: AddMpServerResponse{
						XMLName: xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessService", Local: "AddMpServer_OUTPUT"},
						MpServer: MpServer{
							XMLName: xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessService", Local: "MpServer"},
							Address: "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous",
							ReferenceParameters: ReferenceParametersResponse{
								XMLName:     xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/08/addressing", Local: "ReferenceParameters"},
								ResourceURI: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_ManagementPresenceRemoteSAP",
								SelectorSet: SelectorSetResponse{
									XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "SelectorSet"},
									Selectors: []SelectorResponse{
										{
											XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
											Name:    "CreationClassName",
											Text:    "AMT_ManagementPresenceRemoteSAP",
										},
										{
											XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
											Name:    "Name",
											Text:    "Intel(r) AMT:Management Presence Server 0",
										},
										{
											XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
											Name:    "SystemCreationClassName",
											Text:    "CIM_ComputerSystem",
										},
										{
											XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
											Name:    "SystemName",
											Text:    "Intel(r) AMT",
										},
									},
								},
							},
						},
					},
				},
			},
			// AddRemoteAccessPolicyRule
			{
				"should create a valid AMT_RemoteAccessPolicyRule wsman message",
				AMTRemoteAccessService,
				methods.GenerateAction(AMTRemoteAccessService, AddRemoteAccessPolicyRule),
				fmt.Sprintf(`<h:AddRemoteAccessPolicyRule_INPUT xmlns:h="%s%s"><h:Trigger>%d</h:Trigger><h:TunnelLifeTime>%d</h:TunnelLifeTime><h:ExtendedData>%s</h:ExtendedData><h:MpServer><Address xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing">http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</Address><ReferenceParameters xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing"><ResourceURI xmlns="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd">%s%s</ResourceURI><SelectorSet xmlns="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd"><Selector Name="Name">true</Selector></SelectorSet></ReferenceParameters></h:MpServer></h:AddRemoteAccessPolicyRule_INPUT>`, resourceURIBase, AMTRemoteAccessService, 2, 0, "0300", "http://intel.com/wbem/wscim/1/amt-schema/1/", "AMT_ManagementPresenceRemoteSAP"),
				func() (Response, error) {
					client.CurrentMessage = "AddRemoteAccessService"
					remoteAccessPolicyRule := RemoteAccessPolicyRuleRequest{
						H:              "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessService",
						Trigger:        2,
						TunnelLifeTime: 0,
						ExtendedData:   "0300",
					}

					return elementUnderTest.AddRemoteAccessPolicyRule(remoteAccessPolicyRule, "true")
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					AddRemotePolicyRuleResponse: AddRemoteAccessPolicyRuleResponse{
						XMLName: xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessService", Local: "AddRemoteAccessPolicyRule_OUTPUT"},
						PolicyRuleResponse: PolicyRuleResponse{
							XMLName: xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessService", Local: "PolicyRule"},
							Address: "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous",
							ReferenceParameters: ReferenceParametersResponse{
								XMLName:     xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/08/addressing", Local: "ReferenceParameters"},
								ResourceURI: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyRule",
								SelectorSet: SelectorSetResponse{
									XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "SelectorSet"},
									Selectors: []SelectorResponse{
										{
											XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
											Name:    "CreationClassName",
											Text:    "AMT_RemoteAccessPolicyRule",
										},
										{
											XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
											Name:    "PolicyRuleName",
											Text:    "Periodic",
										},
										{
											XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
											Name:    "SystemCreationClassName",
											Text:    "CIM_ComputerSystem",
										},
										{
											XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
											Name:    "SystemName",
											Text:    "Intel(r) AMT",
										},
									},
								},
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
				logrus.Print(response.XMLOutput)
				assert.NoError(t, err)
				assert.Equal(t, expectedXMLInput, response.XMLInput)
				assert.Equal(t, test.expectedResponse, response.Body)
			})
		}
	})
}

func TestNegativeAMT_RemoteAccessService(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.AMTResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/remoteaccess/service",
	}
	elementUnderTest := NewRemoteAccessServiceWithClient(wsmanMessageCreator, &client)

	t.Run("amt_RemoteAccessService Tests", func(t *testing.T) {
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
				"should create a valid AMT_RemoteAccessService Get wsman message",
				AMTRemoteAccessService,
				wsmantesting.Get,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					RemoteAccessServiceGetResponse: RemoteAccessServiceResponse{
						XMLName:                 xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessService", Local: "AMT_RemoteAccessService"},
						CreationClassName:       "AMT_RemoteAccessService",
						ElementName:             "Intel(r) AMT Remote Access Service",
						Name:                    "Intel(r) AMT Remote Access Service",
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "Intel(r) AMT",
					},
				},
			},
			// ENUMERATES
			{
				"should create a valid AMT_RemoteAccessService Enumerate wsman message",
				AMTRemoteAccessService,
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
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
						EnumerationContext: "D3000000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create a valid AMT_RemoteAccessService Pull wsman message",
				AMTRemoteAccessService,
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
						RemoteAccessItems: []RemoteAccessServiceResponse{
							{
								XMLName:                 xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessService", Local: "AMT_RemoteAccessService"},
								CreationClassName:       "AMT_RemoteAccessService",
								ElementName:             "Intel(r) AMT Remote Access Service",
								Name:                    "Intel(r) AMT Remote Access Service",
								SystemCreationClassName: "CIM_ComputerSystem",
								SystemName:              "Intel(r) AMT",
							},
						},
					},
				},
			},
			// AddMPS
			{
				"should create a valid AMT_RemoteAccessService AddMPS wsman message",
				AMTRemoteAccessService,
				methods.GenerateAction(AMTRemoteAccessService, AddMps),
				fmt.Sprintf(`<h:AddMpServer_INPUT xmlns:h="%s%s"><h:AccessInfo>%s</h:AccessInfo><h:InfoFormat>%d</h:InfoFormat><h:Port>%d</h:Port><h:AuthMethod>%d</h:AuthMethod><h:Username>%s</h:Username><h:Password>%s</h:Password><h:CN>%s</h:CN></h:AddMpServer_INPUT>`, resourceURIBase, AMTRemoteAccessService, "AccessInfo", 1, 2, 3, "Username", "Password", "CommonName"),
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError
					mpsServer := AddMpServerRequest{
						H:          "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessService",
						AccessInfo: "AccessInfo",
						InfoFormat: 1,
						Port:       2,
						AuthMethod: 3,
						Username:   "Username",
						Password:   "Password",
						CommonName: "CommonName",
					}

					return elementUnderTest.AddMPS(mpsServer)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					AddMpServerResponse: AddMpServerResponse{
						XMLName: xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessService", Local: "AddMpServer_OUTPUT"},
						MpServer: MpServer{
							XMLName: xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessService", Local: "MpServer"},
							Address: "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous",
							ReferenceParameters: ReferenceParametersResponse{
								XMLName:     xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/08/addressing", Local: "ReferenceParameters"},
								ResourceURI: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_ManagementPresenceRemoteSAP",
								SelectorSet: SelectorSetResponse{
									XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "SelectorSet"},
									Selectors: []SelectorResponse{
										{
											XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
											Name:    "CreationClassName",
											Text:    "AMT_ManagementPresenceRemoteSAP",
										},
										{
											XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
											Name:    "Name",
											Text:    "Intel(r) AMT:Management Presence Server 0",
										},
										{
											XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
											Name:    "SystemCreationClassName",
											Text:    "CIM_ComputerSystem",
										},
										{
											XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
											Name:    "SystemName",
											Text:    "Intel(r) AMT",
										},
									},
								},
							},
						},
					},
				},
			},
			// AddRemoteAccessPolicyRule
			{
				"should create a valid AMT_RemoteAccessPolicyRule wsman message",
				AMTRemoteAccessService,
				methods.GenerateAction(AMTRemoteAccessService, AddRemoteAccessPolicyRule),
				fmt.Sprintf(`<h:AddRemoteAccessPolicyRule_INPUT xmlns:h="%s%s"><h:Trigger>%d</h:Trigger><h:TunnelLifeTime>%d</h:TunnelLifeTime><h:ExtendedData>%s</h:ExtendedData><h:MpServer><Address xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing">http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</Address><ReferenceParameters xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing"><ResourceURI xmlns="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd">%s%s</ResourceURI><SelectorSet xmlns="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd"><Selector Name="Name">true</Selector></SelectorSet></ReferenceParameters></h:MpServer></h:AddRemoteAccessPolicyRule_INPUT>`, resourceURIBase, AMTRemoteAccessService, 2, 0, "0300", "http://intel.com/wbem/wscim/1/amt-schema/1/", "AMT_ManagementPresenceRemoteSAP"),
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError
					remoteAccessPolicyRule := RemoteAccessPolicyRuleRequest{
						H:              "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessService",
						Trigger:        2,
						TunnelLifeTime: 0,
						ExtendedData:   "0300",
					}

					return elementUnderTest.AddRemoteAccessPolicyRule(remoteAccessPolicyRule, "true")
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					AddRemotePolicyRuleResponse: AddRemoteAccessPolicyRuleResponse{
						XMLName: xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessService", Local: "AddRemoteAccessPolicyRule_OUTPUT"},
						PolicyRuleResponse: PolicyRuleResponse{
							XMLName: xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessService", Local: "PolicyRule"},
							Address: "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous",
							ReferenceParameters: ReferenceParametersResponse{
								XMLName:     xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/08/addressing", Local: "ReferenceParameters"},
								ResourceURI: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyRule",
								SelectorSet: SelectorSetResponse{
									XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "SelectorSet"},
									Selectors: []SelectorResponse{
										{
											XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
											Name:    "CreationClassName",
											Text:    "AMT_RemoteAccessPolicyRule",
										},
										{
											XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
											Name:    "PolicyRuleName",
											Text:    "Periodic",
										},
										{
											XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
											Name:    "SystemCreationClassName",
											Text:    "CIM_ComputerSystem",
										},
										{
											XMLName: xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "Selector"},
											Name:    "SystemName",
											Text:    "Intel(r) AMT",
										},
									},
								},
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
				logrus.Print(response.XMLOutput)
				assert.Error(t, err)
				assert.Equal(t, expectedXMLInput, response.XMLInput)
				assert.NotEqual(t, test.expectedResponse, response.Body)
			})
		}
	})
}
