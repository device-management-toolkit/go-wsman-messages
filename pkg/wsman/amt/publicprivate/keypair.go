/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package publicprivate facilitiates communication with Intel® AMT devices to manage a public-private key in the Intel® AMT CertStore.
//
// Instances of this class can be created using the AMT_PublicKeyManagementService.AddKey method. You can't delete a key instance if it is used by some service (TLS/EAC).
package publicprivate

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

// NewPublicPrivateKeyPairWithClient instantiates a new KeyPair.
func NewPublicPrivateKeyPairWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) KeyPair {
	return KeyPair{
		base.NewService[Response](wsmanMessageCreator, AMTPublicPrivateKeyPair, client),
	}
}

// Get retrieves the representation of the instance identified by the InstanceID
// selector. Shadows the generic parameterless Get to preserve the public API.
func (keyPair KeyPair) Get(instanceID string) (response Response, err error) {
	return keyPair.GetByInstanceID(instanceID)
}

// Pull overrides the generic Pull to post-process the response into the
// RefinedPullResponse shape used by callers.
func (keyPair KeyPair) Pull(enumerationContext string) (response Response, err error) {
	var refinedOutput []RefinedPublicPrivateKeyPair

	response = Response{
		Message: &client.Message{
			XMLInput: keyPair.Base.Pull(enumerationContext),
		},
	}

	// send the message to AMT
	err = keyPair.Base.Execute(response.Message)
	if err != nil {
		return response, err
	}

	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	for _, item := range response.Body.PullResponse.PublicPrivateKeyPairItems {
		output := RefinedPublicPrivateKeyPair{
			InstanceID:  item.InstanceID,
			ElementName: item.ElementName,
			DERKey:      item.DERKey,
		}

		refinedOutput = append(refinedOutput, output)
	}

	response.Body.RefinedPullResponse.PublicPrivateKeyPairItems = refinedOutput

	return response, err
}

// Deletes an instance of a key pair.
func (keyPair KeyPair) Delete(handle string) (response Response, err error) {
	selector := message.Selector{
		Name:  "InstanceID",
		Value: handle,
	}
	response = Response{
		Message: &client.Message{
			XMLInput: keyPair.Base.Delete(selector),
		},
	}
	// send the message to AMT
	err = keyPair.Base.Execute(response.Message)
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
