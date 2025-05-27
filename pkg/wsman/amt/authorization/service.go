/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package authorization facilitates communication with Intel® AMT devices to manage access control list (ACL) entries.
//
// Additional Notes:
//
// 1) Realms 'AuditLogRealm' (20) and 'ACLRealm' (21) are supported only in Intel AMT Release 4.0 and later releases.
//
// 2) Realm 'DTRealm' (23) is supported only in 'ME 5.1' and Intel AMT Release 5.1 and later releases.
//
// 3) All the methods of 'AMT_AuthorizationService' except for 'Get' are not supported in Remote Connectivity Service provisioning mode
package authorization

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/amt/methods"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

// Instantiates a new Authorization service.
func NewServiceWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Service {
	return Service{
		base: message.NewBaseWithClient(wsmanMessageCreator, AMTAuthorizationService, client),
	}
}

// Get retrieves the representation of the instance.
func (as Service) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: as.base.Get(nil),
		},
	}

	// send the message to AMT
	err = as.base.Execute(response.Message)
	if err != nil {
		return
	}

	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}

// Enumerate returns an enumeration context which is used in a subsequent Pull call.
func (as Service) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: as.base.Enumerate(),
		},
	}

	// send the message to AMT
	err = as.base.Execute(response.Message)
	if err != nil {
		return
	}

	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}

// Pull returns the instances of this class.  An enumeration context provided by the Enumerate call is used as input.
func (as Service) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: as.base.Pull(enumerationContext),
		},
	}

	// send the message to AMT
	err = as.base.Execute(response.Message)
	if err != nil {
		return
	}

	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}

// EnumerateUserACLEntries enumerates entries in the User Access Control List (ACL).
func (as Service) EnumerateUserACLEntries(startIndex int) (response Response, err error) {
	if startIndex == 0 {
		startIndex = 1
	}

	header := as.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMTAuthorizationService, EnumerateUserACLEntries), AMTAuthorizationService, nil, "", "")
	body := as.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(EnumerateUserACLEntries), AMTAuthorizationService, &EnumerateUserAclEntries_INPUT{StartIndex: startIndex})

	response = Response{
		Message: &client.Message{
			XMLInput: as.base.WSManMessageCreator.CreateXML(header, body),
		},
	}

	err = as.base.Execute(response.Message)
	if err != nil {
		return
	}

	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}

// Gets the state of a user ACL entry (enabled/disabled).
func (as Service) GetACLEnabledState(handle int) (response Response, err error) {
	header := as.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMTAuthorizationService, GetACLEnabledState), AMTAuthorizationService, nil, "", "")
	body := as.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(GetACLEnabledState), AMTAuthorizationService, &GetAclEnabledState_INPUT{Handle: handle})

	response = Response{
		Message: &client.Message{
			XMLInput: as.base.WSManMessageCreator.CreateXML(header, body),
		},
	}

	err = as.base.Execute(response.Message)
	if err != nil {
		return
	}

	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}

// Returns the username attribute of the Admin ACL.
func (as Service) GetAdminACLEntry() (response Response, err error) {
	header := as.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMTAuthorizationService, GetAdminACLEntry), AMTAuthorizationService, nil, "", "")
	body := as.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(GetAdminACLEntry), AMTAuthorizationService, nil)

	response = Response{
		Message: &client.Message{
			XMLInput: as.base.WSManMessageCreator.CreateXML(header, body),
		},
	}

	err = as.base.Execute(response.Message)
	if err != nil {
		return
	}

	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}

// Reads the Admin ACL Entry status from Intel® AMT. The return state changes as a function of the admin password.
func (as Service) GetAdminACLEntryStatus() (response Response, err error) {
	header := as.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMTAuthorizationService, GetAdminACLEntryStatus), AMTAuthorizationService, nil, "", "")
	body := as.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(GetAdminACLEntryStatus), AMTAuthorizationService, nil)

	response = Response{
		Message: &client.Message{
			XMLInput: as.base.WSManMessageCreator.CreateXML(header, body),
		},
	}

	err = as.base.Execute(response.Message)
	if err != nil {
		return
	}

	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}

// Reads the remote Admin ACL Entry status from Intel® AMT. The return state changes as a function of the remote admin password.
func (as Service) GetAdminNetACLEntryStatus() (response Response, err error) {
	header := as.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMTAuthorizationService, GetAdminNetACLEntryStatus), AMTAuthorizationService, nil, "", "")
	body := as.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(GetAdminNetACLEntryStatus), AMTAuthorizationService, nil)

	response = Response{
		Message: &client.Message{
			XMLInput: as.base.WSManMessageCreator.CreateXML(header, body),
		},
	}

	err = as.base.Execute(response.Message)
	if err != nil {
		return
	}

	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}

// Reads a user entry from the Intel® AMT device. Note: confidential information, such as password (hash) is omitted or zeroed in the response.
func (as Service) GetUserACLEntryEx(handle int) (response Response, err error) {
	header := as.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMTAuthorizationService, GetUserACLEntryEx), AMTAuthorizationService, nil, "", "")
	body := as.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(GetUserACLEntryEx), AMTAuthorizationService, &GetUserAclEntryEx_INPUT{Handle: handle})

	response = Response{
		Message: &client.Message{
			XMLInput: as.base.WSManMessageCreator.CreateXML(header, body),
		},
	}

	err = as.base.Execute(response.Message)
	if err != nil {
		return
	}

	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}

// Removes an entry from the User Access Control List (ACL), given a handle.
func (as Service) RemoveUserACLEntry(handle int) (response Response, err error) {
	header := as.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMTAuthorizationService, RemoveUserACLEntry), AMTAuthorizationService, nil, "", "")
	body := as.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(RemoveUserACLEntry), AMTAuthorizationService, &RemoveUserAclEntry_INPUT{Handle: handle})

	response = Response{
		Message: &client.Message{
			XMLInput: as.base.WSManMessageCreator.CreateXML(header, body),
		},
	}

	err = as.base.Execute(response.Message)
	if err != nil {
		return
	}

	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}

// Enables or disables a user ACL entry. Disabling ACL entries is useful when accounts that cannot be removed (system accounts - starting with $$) are required to be disabled.
func (as Service) SetACLEnabledState(handle int, enabled bool) (response Response, err error) {
	header := as.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMTAuthorizationService, SetACLEnabledState), AMTAuthorizationService, nil, "", "")
	body := as.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(SetACLEnabledState), AMTAuthorizationService, &SetAclEnabledState_INPUT{Handle: handle, Enabled: enabled})

	response = Response{
		Message: &client.Message{
			XMLInput: as.base.WSManMessageCreator.CreateXML(header, body),
		},
	}

	err = as.base.Execute(response.Message)
	if err != nil {
		return
	}

	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}

// Updates an Admin entry in the Intel® AMT device.
func (as Service) SetAdminAclEntryEx(username, digestPassword string) (response Response, err error) {
	header := as.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMTAuthorizationService, SetAdminACLEntryEx), AMTAuthorizationService, nil, "", "")
	body := as.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(SetAdminACLEntryEx), AMTAuthorizationService, &SetAdminAclEntryEx_INPUT{Username: username, DigestPassword: digestPassword})

	response = Response{
		Message: &client.Message{
			XMLInput: as.base.WSManMessageCreator.CreateXML(header, body),
		},
	}

	err = as.base.Execute(response.Message)
	if err != nil {
		return
	}

	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}
