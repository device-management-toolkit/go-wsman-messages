/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package authorization

import (
	"encoding/xml"
	"testing"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/common"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/wsmantesting"
	"github.com/stretchr/testify/assert"
)

func TestPositiveAMT_AuthorizationService(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/authorization",
	}
	elementUnderTest := NewServiceWithClient(wsmanMessageCreator, &client)
	t.Run("amt_AuthorizationService Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			body             string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			//GETS
			{"should create a valid AMT_AuthorizationService Get wsman message",
				AMT_AuthorizationService,
				wsmantesting.GET,
				"",
				func() (Response, error) {
					client.CurrentMessage = "Get"
					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetResponse: AuthorizationOccurrence{
						XMLName:                 xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService", Local: "AMT_AuthorizationService"},
						AllowHttpQopAuthOnly:    1,
						CreationClassName:       AMT_AuthorizationService,
						ElementName:             "Intel(r) AMT Authorization Service",
						EnabledState:            5,
						Name:                    "Intel(r) AMT Authorization Service",
						RequestedState:          12,
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "Intel(r) AMT",
					},
				},
			},
			//ENUMERATES
			{"should create a valid AMT_AuthorizationService Enumerate wsman message",
				AMT_AuthorizationService,
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Enumerate"
					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "5C000000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should create a valid AMT_AuthorizationService Pull wsman message",
				AMT_AuthorizationService,
				wsmantesting.PULL,
				wsmantesting.PULL_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Pull"
					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						XMLName: xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/09/enumeration", Local: "PullResponse"},
						AuthorizationOccurrenceItems: []AuthorizationOccurrence{
							{
								XMLName:                 xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService", Local: "AMT_AuthorizationService"},
								AllowHttpQopAuthOnly:    1,
								CreationClassName:       AMT_AuthorizationService,
								ElementName:             "Intel(r) AMT Authorization Service",
								EnabledState:            5,
								Name:                    "Intel(r) AMT Authorization Service",
								RequestedState:          12,
								SystemCreationClassName: "CIM_ComputerSystem",
								SystemName:              "Intel(r) AMT",
							},
						},
					},
				},
			},
			// // AUTHORIZATION SERVICE

			// // ADD USER ACL ENTRY EX
			// // Verify with Matt - Typescript is referring to wrong realm values
			// // {"should return a valid amt_AuthorizationService ADD_USER_ACL_ENTRY_EX wsman message using digest", AMT_AuthorizationService, ADD_USER_ACL_ENTRY_EX, fmt.Sprintf(`<h:AddUserAclEntryEx_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService"><h:DigestUsername>%s</h:DigestUsername><h:DigestPassword>%s</h:DigestPassword><h:AccessPermission>%d</h:AccessPermission><h:Realms>%d</h:Realms></h:AddUserAclEntryEx_INPUT>`, "test", "P@ssw0rd", 2, 3), func() string {
			// // 	return elementUnderTest.AddUserAclEntryEx(authorization.AccessPermissionLocalAndNetworkAccess, []authorization.RealmValues{authorization.RedirectionRealm}, "test", "P@ssw0rd", "")
			// // }},
			// // {"should return a valid amt_AuthorizationService ADD_USER_ACL_ENTRY_EX wsman message using kerberos", AMT_AuthorizationService, ADD_USER_ACL_ENTRY_EX, fmt.Sprintf(`<h:AddUserAclEntryEx_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService"><h:KerberosUserSid>%d</h:KerberosUserSid><h:AccessPermission>%d</h:AccessPermission><h:Realms>%d3</h:Realms></h:AddUserAclEntryEx_INPUT>`, 64, 2, 3), func() string {
			// // 	return elementUnderTest.AddUserAclEntryEx(authorization.AccessPermissionLocalAndNetworkAccess, []authorization.RealmValues{authorization.RedirectionRealm}, "", "", "64")
			// // }},
			// // // Check how to verify for exceptions
			// // // {"should throw an error if the digestUsername is longer than 16 when calling AddUserAclEntryEx", "", "", "", func() string {
			// // // 	return elementUnderTest.AddUserAclEntryEx(2, []models.RealmValues{models.RedirectionRealm}, "thisusernameistoolong", "test", "")
			// // // }},
			// // ENUMERATE USER ACL ENTRIES
			// {"should return a valid amt_AuthorizationService EnumerateUserAclEntries wsman message when startIndex is undefined", AMT_AuthorizationService, `http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService/EnumerateUserAclEntries`, fmt.Sprintf(`<h:EnumerateUserAclEntries_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService"><h:StartIndex>%d</h:StartIndex></h:EnumerateUserAclEntries_INPUT>`, 1), func() string {
			// 	var index int
			// 	return elementUnderTest.EnumerateUserAclEntries(index)
			// }},
			// {"should return a valid amt_AuthorizationService EnumerateUserAclEntries wsman message when startIndex is not 1", AMT_AuthorizationService, `http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService/EnumerateUserAclEntries`, fmt.Sprintf(`<h:EnumerateUserAclEntries_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService"><h:StartIndex>%d</h:StartIndex></h:EnumerateUserAclEntries_INPUT>`, 50), func() string {
			// 	return elementUnderTest.EnumerateUserAclEntries(50)
			// }},
			// // GET USER ACL ENTRY EX
			// {"should return a valid amt_AuthorizationService GetUserAclEntryEx wsman message", AMT_AuthorizationService, `http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService/GetUserAclEntryEx`, `<h:GetUserAclEntryEx_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService"><h:Handle>1</h:Handle></h:GetUserAclEntryEx_INPUT>`, func() string {
			// 	return elementUnderTest.GetUserAclEntryEx(1)
			// }},
			// // UPDATE USER ACL ENTRY EX
			// // {"should return a valid amt_AuthorizationService UpdateUserAclEntryEx wsman message using digest", AMT_AuthorizationService, `http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService/UpdateUserAclEntryEx`, `<h:GetUserAclEntryEx_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService"><h:Handle>1</h:Handle></h:GetUserAclEntryEx_INPUT>`, func() string {
			// // 	return elementUnderTest.UpdateUserAclEntryEx(1, 2, []authorization.RealmValues{authorization.RedirectionRealm}, "test", "test123!", "")
			// // }},
			// // {"should return a valid amt_AuthorizationService UpdateUserAclEntryEx wsman message using kerberos", AMT_AuthorizationService, `http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService/UpdateUserAclEntryEx`, `<h:UpdateUserAclEntryEx_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService"><h:Handle>1</h:Handle><h:KerberosUserSid>64</h:KerberosUserSid><h:AccessPermission>2</h:AccessPermission><h:Realms>3</h:Realms></h:UpdateUserAclEntryEx_INPUT>`, func() string {
			// // 	return elementUnderTest.UpdateUserAclEntryEx(1, 2, []authorization.RealmValues{authorization.RedirectionRealm}, "", "", "64")
			// // }},
			// // // should throw an error if digest or kerberos credentials are not provided to UpdateUserAclEntryEx
			// // // should throw an error if the digestUsername is longer than 16 when calling UpdateUserAclEntryEx

			// // REMOVE USER ACL ENTRY
			// {"should return a valid amt_AuthorizationService RemoveUserAclEntry wsman message", AMT_AuthorizationService, `http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService/RemoveUserAclEntry`, `<h:RemoveUserAclEntry_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService"><h:Handle>1</h:Handle></h:RemoveUserAclEntry_INPUT>`, func() string {
			// 	return elementUnderTest.RemoveUserAclEntry(1)
			// }},

			// // GET ADMIN ACL ENTRY
			// {"should return a valid amt_AuthorizationService GetAdminAclEntry wsman message", AMT_AuthorizationService, `http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService/GetAdminAclEntry`, `<h:GetAdminAclEntry_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService"></h:GetAdminAclEntry_INPUT>`, func() string {
			// 	return elementUnderTest.GetAdminAclEntry()
			// }},

			// // GET ADMIN ACL ENTRY STATUS
			// {"should return a valid amt_AuthorizationService GetAdminAclEntry wsman message", AMT_AuthorizationService, `http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService/GetAdminAclEntryStatus`, `<h:GetAdminAclEntryStatus_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService"></h:GetAdminAclEntryStatus_INPUT>`, func() string {
			// 	return elementUnderTest.GetAdminAclEntryStatus()
			// }},

			// // GET ADMIN NET ACL ENTRY STATUS
			// {"should return a valid amt_AuthorizationService GetAdminNetAclEntryStatus wsman message", AMT_AuthorizationService, `http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService/GetAdminNetAclEntryStatus`, `<h:GetAdminNetAclEntryStatus_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService"></h:GetAdminNetAclEntryStatus_INPUT>`, func() string {
			// 	return elementUnderTest.GetAdminNetAclEntryStatus()
			// }},

			// // GET ACL ENABLED STATE
			// {"should return a valid amt_AuthorizationService GetAclEnabledState wsman message", AMT_AuthorizationService, `http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService/GetAclEnabledState`, `<h:GetAclEnabledState_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService"><h:Handle>1</h:Handle></h:GetAclEnabledState_INPUT>`, func() string {
			// 	return elementUnderTest.GetAclEnabledState(1)
			// }},
		}
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				expectedXMLInput := wsmantesting.ExpectedResponse(messageID, resourceUriBase, test.method, test.action, "", test.body)
				messageID++
				response, err := test.responseFunc()
				assert.NoError(t, err)
				assert.Equal(t, expectedXMLInput, response.XMLInput)
				assert.Equal(t, test.expectedResponse, response.Body)
			})
		}
	})
}
func TestNegativeAMT_AuthorizationService(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/authorization",
	}
	elementUnderTest := NewServiceWithClient(wsmanMessageCreator, &client)
	t.Run("amt_AuthorizationService Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			body             string
			extraHeader      string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			{
				"should create an invalid AMT_EthernetPortSettings Get wsman message",
				"AMT_EthernetPortSettings",
				wsmantesting.GET,
				"",
				"",
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetResponse: AuthorizationOccurrence{
						XMLName:                 xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService", Local: "AMT_AuthorizationService"},
						AllowHttpQopAuthOnly:    1,
						CreationClassName:       AMT_AuthorizationService,
						ElementName:             "Intel(r) AMT Authorization Service",
						EnabledState:            5,
						Name:                    "Intel(r) AMT Authorization Service",
						RequestedState:          12,
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "Intel(r) AMT",
					},
				},
			},
			{
				"should create an invalid AMT_EthernetPortSettings Pull wsman message",
				"AMT_EthernetPortSettings",
				wsmantesting.PULL,
				wsmantesting.PULL_BODY,
				"",
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Pull("")
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						XMLName: xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/09/enumeration", Local: "PullResponse"},
						AuthorizationOccurrenceItems: []AuthorizationOccurrence{
							{
								XMLName:                 xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuthorizationService", Local: "AMT_AuthorizationService"},
								AllowHttpQopAuthOnly:    1,
								CreationClassName:       AMT_AuthorizationService,
								ElementName:             "Intel(r) AMT Authorization Service",
								EnabledState:            5,
								Name:                    "Intel(r) AMT Authorization Service",
								RequestedState:          12,
								SystemCreationClassName: "CIM_ComputerSystem",
								SystemName:              "Intel(r) AMT",
							},
						},
					},
				},
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				expectedXMLInput := wsmantesting.ExpectedResponse(messageID, resourceUriBase, test.method, test.action, test.extraHeader, test.body)
				messageID++
				response, err := test.responseFunc()
				assert.Error(t, err)
				assert.NotEqual(t, expectedXMLInput, response.XMLInput)
				assert.NotEqual(t, test.expectedResponse, response.Body)
			})
		}
	})
}
