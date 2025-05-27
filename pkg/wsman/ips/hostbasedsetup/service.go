/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package hostbasedsetup facilitates communication with Intel® AMT devices to describe the Host Based Setup Service, which is the logic in Intel(R) AMT that responds to setup requests initiated from the host using OS Administrator credentials.
//
// Also provides a method to upgrade to Admin Control mode that can be initiated remotely.
package hostbasedsetup

import (
	"crypto/md5"
	"encoding/xml"
	"errors"
	"fmt"
	"io"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/ips/methods"
)

// NewHostBasedSetupService returns a new instance of the HostBasedSetupService struct.
func NewHostBasedSetupServiceWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Service {
	return Service{
		base: message.NewBaseWithClient(wsmanMessageCreator, IPSHostBasedSetupService, client),
	}
}

// Get retrieves the representation of the instance.
func (service Service) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: service.base.Get(nil),
		},
	}

	err = service.base.Execute(response.Message)
	if err != nil {
		return response, err
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

// Enumerate returns an enumeration context which is used in a subsequent Pull call.
func (service Service) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: service.base.Enumerate(),
		},
	}

	err = service.base.Execute(response.Message)
	if err != nil {
		return response, err
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

// Pull returns the instances of this class.  An enumeration context provided by the Enumerate call is used as input.
func (service Service) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: service.base.Pull(enumerationContext),
		},
	}

	err = service.base.Execute(response.Message)
	if err != nil {
		return response, err
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

// Add a certificate to the provisioning certificate chain, to be used by AdminSetup or UpgradeClientToAdmin methods.
func (service Service) AddNextCertInChain(cert string, isLeaf, isRoot bool) (response Response, err error) {
	header := service.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(IPSHostBasedSetupService, AddNextCertInChain), IPSHostBasedSetupService, nil, "", "")
	body := service.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(AddNextCertInChain), IPSHostBasedSetupService, AddNextCertInChainInput{
		H:                 "http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HostBasedSetupService",
		NextCertificate:   cert,
		IsLeafCertificate: isLeaf,
		IsRootCertificate: isRoot,
	})
	response = Response{
		Message: &client.Message{
			XMLInput: service.base.WSManMessageCreator.CreateXML(header, body),
		},
	}

	err = service.base.Execute(response.Message)
	if err != nil {
		return response, err
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	if response.Body.AddNextCertInChain_OUTPUT.ReturnValue != 0 {
		err = generateErrorMessage("addnextcertinchain", response.Body.AdminSetup_OUTPUT.ReturnValue)
	}

	return response, err
}

// Setup Intel® AMT from the local host, resulting in Admin Setup Mode. Requires OS administrator rights, and moves Intel® AMT from "Pre Provisioned" state to "Post Provisioned" state. The control mode after this method is run will be "Admin".
func (service Service) AdminSetup(adminPassEncryptionType AdminPassEncryptionType, digestRealm, adminPassword, mcNonce string, signingAlgorithm SigningAlgorithm, digitalSignature string) (response Response, err error) {
	hashInHex := createMD5Hash(adminPassword, digestRealm)
	header := service.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(IPSHostBasedSetupService, AdminSetup), IPSHostBasedSetupService, nil, "", "")
	body := service.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(AdminSetup), IPSHostBasedSetupService, AdminSetupInput{
		H:                          "http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HostBasedSetupService",
		NetAdminPassEncryptionType: int(adminPassEncryptionType),
		NetworkAdminPassword:       hashInHex,
		McNonce:                    mcNonce,
		SigningAlgorithm:           int(signingAlgorithm),
		DigitalSignature:           digitalSignature,
	})
	response = Response{
		Message: &client.Message{
			XMLInput: service.base.WSManMessageCreator.CreateXML(header, body),
		},
	}

	err = service.base.Execute(response.Message)
	if err != nil {
		return response, err
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	if response.Body.AdminSetup_OUTPUT.ReturnValue != 0 {
		err = generateErrorMessage("adminsetup", response.Body.AdminSetup_OUTPUT.ReturnValue)
	}

	return response, err
}

func (service Service) Setup(adminPassEncryptionType AdminPassEncryptionType, digestRealm, adminPassword string) (response Response, err error) {
	hashInHex := createMD5Hash(adminPassword, digestRealm)
	header := service.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(IPSHostBasedSetupService, Setup), IPSHostBasedSetupService, nil, "", "")
	body := service.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(Setup), IPSHostBasedSetupService, SetupInput{
		H:                          "http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HostBasedSetupService",
		NetAdminPassEncryptionType: int(adminPassEncryptionType),
		NetworkAdminPassword:       hashInHex,
	})
	response = Response{
		Message: &client.Message{
			XMLInput: service.base.WSManMessageCreator.CreateXML(header, body),
		},
	}

	err = service.base.Execute(response.Message)
	if err != nil {
		return response, err
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	if response.Body.Setup_OUTPUT.ReturnValue != 0 {
		err = generateErrorMessage("setup", response.Body.Setup_OUTPUT.ReturnValue)
	}

	return response, err
}

func createMD5Hash(adminPassword, digestRealm string) string {
	// Create an md5 hash.
	setupPassword := "admin:" + digestRealm + ":" + adminPassword
	hash := md5.New()

	_, err := io.WriteString(hash, setupPassword)
	if err != nil {
		return ""
	}

	hashInHex := fmt.Sprintf("%x", hash.Sum(nil))

	return hashInHex
}

// Upgrade Intel® AMT from Client to Admin Control Mode.
func (service Service) UpgradeClientToAdmin(mcNonce string, signingAlgorithm SigningAlgorithm, digitalSignature string) (response Response, err error) {
	header := service.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(IPSHostBasedSetupService, UpgradeClientToAdmin), IPSHostBasedSetupService, nil, "", "")
	body := service.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(UpgradeClientToAdmin), IPSHostBasedSetupService, UpgradeClientToAdminInput{
		H:                "http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HostBasedSetupService",
		McNonce:          mcNonce,
		SigningAlgorithm: int(signingAlgorithm),
		DigitalSignature: digitalSignature,
	})
	response = Response{
		Message: &client.Message{
			XMLInput: service.base.WSManMessageCreator.CreateXML(header, body),
		},
	}

	err = service.base.Execute(response.Message)
	if err != nil {
		return response, err
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

// generateErrorMessage returns an error message based on the return value.
func generateErrorMessage(call string, returnValue SetupReturnValue) error {
	ErrSetupFailed := errors.New(call + " failed")

	return fmt.Errorf("%w: returned %d", ErrSetupFailed, returnValue)
}
