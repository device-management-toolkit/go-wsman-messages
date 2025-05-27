/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package ieee8021x

import (
	"encoding/xml"
	"fmt"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/ips/methods"
)

// NewIEEE8021xSettings returns a new instance of the IEEE8021xSettings struct.
func NewIEEE8021xSettingsWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Settings {
	return Settings{
		base: message.NewBaseWithClient(wsmanMessageCreator, IPSIEEE8021xSettings, client),
	}
}

// Get retrieves the representation of the instance.
func (settings Settings) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: settings.base.Get(nil),
		},
	}

	err = settings.base.Execute(response.Message)
	if err != nil {
		return response, err
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	return
}

// Enumerate returns an enumeration context which is used in a subsequent Pull call.
func (settings Settings) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: settings.base.Enumerate(),
		},
	}

	err = settings.base.Execute(response.Message)
	if err != nil {
		return response, err
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	return
}

// Pull returns the instances of this class.  An enumeration context provided by the Enumerate call is used as input.
func (settings Settings) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: settings.base.Pull(enumerationContext),
		},
	}

	err = settings.base.Execute(response.Message)
	if err != nil {
		return response, err
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	return
}

// Put will change properties of the selected instance.
func (settings Settings) Put(ieee8021xSettings IEEE8021xSettingsRequest) (response Response, err error) {
	ieee8021xSettings.H = fmt.Sprintf("%s%s", message.IPSSchema, IPSIEEE8021xSettings)
	response = Response{
		Message: &client.Message{
			XMLInput: settings.base.Put(ieee8021xSettings, false, nil),
		},
	}

	err = settings.base.Execute(response.Message)
	if err != nil {
		return response, err
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	return
}

func (settings Settings) SetCertificates(serverCertificateIssuer, clientCertificate string) (response Response, err error) {
	header := settings.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(IPSIEEE8021xSettings, SetCertificates), IPSIEEE8021xSettings, nil, "", "")
	serverCert := ServerCertificateIssuer{
		Address: "default",
		ReferenceParameters: ReferenceParameters{
			ResourceURI: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyCertificate",
			SelectorSet: SelectorSet{
				Selector: Selector{
					Name:  "InstanceID",
					Value: serverCertificateIssuer,
				},
			},
		},
	}
	clientCert := ClientCertificateIssuer{
		Address: "default",
		ReferenceParameters: ReferenceParameters{
			ResourceURI: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyCertificate",
			SelectorSet: SelectorSet{
				Selector: Selector{
					Name:  "InstanceID",
					Value: clientCertificate,
				},
			},
		},
	}
	body := settings.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(SetCertificates), IPSIEEE8021xSettings,
		Certificate{
			H:                       "http://intel.com/wbem/wscim/1/ips-schema/1/IPS_IEEE8021xSettings",
			ServerCertificateIssuer: serverCert,
			ClientCertificate:       clientCert,
		},
	)
	response = Response{
		Message: &client.Message{
			XMLInput: settings.base.WSManMessageCreator.CreateXML(header, body),
		},
	}

	err = settings.base.Execute(response.Message)
	if err != nil {
		return response, err
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	return response, nil
}
