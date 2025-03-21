/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package setupandconfiguration facilitiates communication with Intel® AMT devices to access and interact with the Setup and Configuration Service, which is the logic in Intel® AMT that responds to Setup and Configuration requests.
package setupandconfiguration

import (
	"encoding/base64"
	"encoding/xml"
	"errors"
	"fmt"
	"strconv"

	"github.com/google/uuid"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/amt/methods"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

// DecodeUUID formats the returned AMT base64 encoded UUID into a human readable UUID.
func (r *Response) DecodeUUID() (amtUUID string, err error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(r.Body.GetUuid_OUTPUT.UUID)
	if err != nil {
		return "", err
	}

	rearrangeBytes := []byte{
		decodedBytes[3], decodedBytes[2], decodedBytes[1], decodedBytes[0],
		decodedBytes[5], decodedBytes[4],
		decodedBytes[7], decodedBytes[6],
		decodedBytes[8], decodedBytes[9],
		decodedBytes[10], decodedBytes[11], decodedBytes[12], decodedBytes[13], decodedBytes[14], decodedBytes[15],
	}

	uuidSlice, err := uuid.FromBytes(rearrangeBytes)
	if err != nil {
		return "", err
	}

	amtUUID = uuidSlice.String()

	return amtUUID, err
}

// NewSetupAndConfigurationServiceWithClient instantiates a new Service.
func NewSetupAndConfigurationServiceWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Service {
	return Service{
		base: message.NewBaseWithClient(wsmanMessageCreator, AMTSetupAndConfigurationService, client),
	}
}

// Gets the representation of the instance.
func (s Service) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: s.base.Get(nil),
		},
	}

	// send the message to AMT
	err = s.base.Execute(response.Message)
	if err != nil {
		return response, err
	}

	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	return response, err
}

// Enumerate returns an enumeration context which is used in a subsequent Pull call.
func (s Service) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: s.base.Enumerate(),
		},
	}

	// send the message to AMT
	err = s.base.Execute(response.Message)
	if err != nil {
		return response, err
	}

	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	return response, err
}

// Pull returns the instances of this class.  An enumeration context provided by the Enumerate call is used as input.
func (s Service) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: s.base.Pull(enumerationContext),
		},
	}

	// send the message to AMT
	err = s.base.Execute(response.Message)
	if err != nil {
		return response, err
	}

	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	return response, err
}

// Put will change properties of the selected instance.
func (s Service) Put(setupAndConfigurationService SetupAndConfigurationServiceRequest) (response Response, err error) {
	setupAndConfigurationService.H = fmt.Sprintf("%s%s", message.AMTSchema, AMTSetupAndConfigurationService)

	response = Response{
		Message: &client.Message{
			XMLInput: s.base.Put(setupAndConfigurationService, false, nil),
		},
	}

	// send the message to AMT
	err = s.base.Execute(response.Message)
	if err != nil {
		return response, err
	}

	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	return response, err
}

// CommitChanges saves pending configuration commands made to the Intel® AMT device.
// Completes configuration when in "IN-provisioning" state.
// This routine commits pending configuration commands which are dependent on an internal restart sequence or a cumulative validity check.
//
// Failure to execute this command prevents the pending configurations (which are not stored in flash memory) to take effect.
// Operations (or situations such as a power loss) that immediately change flash memory depend on a call to CommitChanges()to refresh the internal Firmware state.
//
// Note:
//
// 1. If TLS is enabled, RSA Key and Certificate must be configured in order to work properly with the changes being committed.
//
// 2. If DHCP is enabled, host-name must be set.
//
// 3. If mutual authentication is configured, then at least one trusted root certificate must exist.
//
// 4. When using TLS mutual authentication, the user must first configure the Intel AMT system time.
//
// 5. If in EnterpriseMode Provisioning, then caller must update the internal clock and change the PRNG.
//
// Since committing changes may cause an internal restart sequence, remote applications should allow sufficient time for Intel AMT to reload before issuing the next command.
//
// ValueMap={0, 1, 38, 2057}
//
// Values={PT_STATUS_SUCCESS, PT_STATUS_INTERNAL_ERROR, PT_STATUS_FLASH_WRITE_LIMIT_EXCEEDED, PT_STATUS_DATA_MISSING}.
func (s Service) CommitChanges() (response Response, err error) {
	header := s.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMTSetupAndConfigurationService, CommitChanges), AMTSetupAndConfigurationService, nil, "", "")
	body := s.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(CommitChanges), AMTSetupAndConfigurationService, nil)

	response = Response{
		Message: &client.Message{
			XMLInput: s.base.WSManMessageCreator.CreateXML(header, body),
		},
	}

	// send the message to AMT
	err = s.base.Execute(response.Message)
	if err != nil {
		return response, err
	}

	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	err = checkReturnValue(int(response.Body.SetMEBxPassword_OUTPUT.ReturnValue), "Commit Changes")

	return response, err
}

// GetUUID gets the AMT UUID from the device.
//
// The returned value is in base64 format.  DecodeUUID can be used to format this value into a human readable UUID
//
// ValueMap={0, 1}
//
// Values={PT_STATUS_SUCCESS, PT_STATUS_INTERNAL_ERROR}.
func (s Service) GetUUID() (response Response, err error) {
	header := s.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMTSetupAndConfigurationService, GetUUID), AMTSetupAndConfigurationService, nil, "", "")
	body := s.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(GetUUID), AMTSetupAndConfigurationService, nil)

	response = Response{
		Message: &client.Message{
			XMLInput: s.base.WSManMessageCreator.CreateXML(header, body),
		},
	}

	// send the message to AMT
	err = s.base.Execute(response.Message)
	if err != nil {
		return response, err
	}

	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	return response, err
}

// SetMEBXPassword changes the ME Bios extension password.
// It allows a remote caller to change the ME access password for the BIOS extension screen.
// This call succeeds depending on the password policy rule defined in MEBx (BIOS extension):
//
// "Default Password Only" - Method succeeds only when the current password is still the default value and only in PKI provisioning.
//
// "During Setup and Configuration" - Method succeeds only during provisioning, regardless of provisioning method or previous password value.
//
// "ANYTIME" - Method will always succeed. (i.e. even when configured).
//
// Note: API is blocked in client control mode
//
// ValueMap={0, 1, 16, 2054}
//
// Values={PT_STATUS_SUCCESS, PT_STATUS_INTERNAL_ERROR, PT_STATUS_NOT_PERMITTED, PT_STATUS_INVALID_PASSWORD}.
func (s Service) SetMEBXPassword(password string) (response Response, err error) {
	header := s.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMTSetupAndConfigurationService, SetMEBxPassword), AMTSetupAndConfigurationService, nil, "", "")

	mebxPassword := MEBXPassword{
		Password: password,
	}
	body := s.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(SetMEBxPassword), AMTSetupAndConfigurationService, &mebxPassword)

	response = Response{
		Message: &client.Message{
			XMLInput: s.base.WSManMessageCreator.CreateXML(header, body),
		},
	}

	// send the message to AMT
	err = s.base.Execute(response.Message)
	if err != nil {
		return response, err
	}

	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	err = checkReturnValue(int(response.Body.SetMEBxPassword_OUTPUT.ReturnValue), "MEBx Password")

	return response, err
}

// PartialUnprovision Transfers Intel(R) AMT into a partially-unprovisioned state: Except for configuration settings required for the next provisioning: Admin ACL settings, TLS-PSK keys, Host & Domain name, and provisioning server IP and port number, settings will be restored to factory defaults. The device will need to be re-provisioned after this command.
//
// Product Specific Usage:
// This command puts Intel AMT into a Partial Unprovision state.
// The effect of this command is similar to calling the Unprovision() command (with ProvisioningMode set to "Enterprise").
// The only difference is that the following settings remain unchanged:
// - Admin ACL settings (name and password)
// - TLS-PSK keys
// - Host name
// - Provisioning server IP and port number
// - Domain name
//
// # In Client Control Mode, call will succeed even if auditor is blocking the operation
//
// Qualifiers:
// -------------
// ValueMap={0, 1, 16, 2076}
// Values={PT_STATUS_SUCCESS, PT_STATUS_INTERNAL_ERROR, PT_STATUS_NOT_PERMITTED, PT_STATUS_BLOCKING_COMPONENT}.
func (s Service) PartialUnprovision() (response Response, err error) {
	header := s.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMTSetupAndConfigurationService, PartialUnprovision), AMTSetupAndConfigurationService, nil, "", "")
	body := s.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(PartialUnprovision), AMTSetupAndConfigurationService, nil)

	response = Response{
		Message: &client.Message{
			XMLInput: s.base.WSManMessageCreator.CreateXML(header, body),
		},
	}

	// send the message to AMT
	err = s.base.Execute(response.Message)
	if err != nil {
		return response, err
	}

	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	if response.Body.PartialUnprovision_OUTPUT.ReturnValue != 0 {
		err = errors.New("Status: Failed to partially deactivate. ReturnValue: " + fmt.Sprintf("%d", response.Body.PartialUnprovision_OUTPUT.ReturnValue))
	}

	return response, err
}

// Unprovision unconfigures and deactivates the Intel® AMT device. The device will need to be re-provisioned after this command before being able to use AMT features.
//
// In Client Control Mode, call will succeed even if auditor is blocking the operation.
//
// ValueMap={0, 1, 16, 36, 2076}
//
// Values={PT_STATUS_SUCCESS, PT_STATUS_INTERNAL_ERROR, PT_STATUS_NOT_PERMITTED, PT_STATUS_INVALID_PARAMETER, PT_STATUS_BLOCKING_COMPONENT}.
func (s Service) Unprovision(provisioningMode ProvisioningModeValue) (response Response, err error) {
	if provisioningMode == 0 {
		provisioningMode = 1
	}

	pMode := ProvisioningMode{
		ProvisioningMode: provisioningMode,
	}

	header := s.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMTSetupAndConfigurationService, Unprovision), AMTSetupAndConfigurationService, nil, "", "")
	body := s.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(Unprovision), AMTSetupAndConfigurationService, &pMode)

	response = Response{
		Message: &client.Message{
			XMLInput: s.base.WSManMessageCreator.CreateXML(header, body),
		},
	}

	// send the message to AMT
	err = s.base.Execute(response.Message)
	if err != nil {
		return response, err
	}

	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	if response.Body.Unprovision_OUTPUT.ReturnValue != 0 {
		err = errors.New("Status: Failed to deactivate. ReturnValue: " + fmt.Sprintf("%d", response.Body.Unprovision_OUTPUT.ReturnValue))
	}

	return response, err
}

func checkReturnValue(rc int, item string) (err error) {
	if rc != int(ReturnValueSuccess) {
		if rc == int(ReturnValueNotPermitted) {
			return errors.New(item + " is not permitted")
		} else if rc == int(ReturnValueInvalidPassword) {
			return errors.New(item + " is invalid")
		}

		return errors.New(item + " non-zero return code: " + strconv.Itoa(rc))
	}

	return nil
}
