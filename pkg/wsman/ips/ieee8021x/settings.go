/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package ieee8021x

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/ips/methods"
)

type Settings struct {
	base.WSManService[Response]
}

// NewIEEE8021xSettings returns a new instance of the IEEE8021xSettings struct.
func NewIEEE8021xSettingsWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Settings {
	return Settings{
		base.NewService[Response](wsmanMessageCreator, IPSIEEE8021xSettings, client),
	}
}

func (settings Settings) SetCertificates(serverCertificateIssuer, clientCertificate string) (response Response, err error) {
	header := settings.Base.WSManMessageCreator.CreateHeader(methods.GenerateAction(IPSIEEE8021xSettings, SetCertificates), IPSIEEE8021xSettings, nil, "", "")
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
	body := settings.Base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(SetCertificates), IPSIEEE8021xSettings,
		Certificate{
			H:                       "http://intel.com/wbem/wscim/1/ips-schema/1/IPS_IEEE8021xSettings",
			ServerCertificateIssuer: serverCert,
			ClientCertificate:       clientCert,
		},
	)
	response = Response{
		Message: &client.Message{
			XMLInput: settings.Base.WSManMessageCreator.CreateXML(header, body),
		},
	}

	err = settings.Base.Execute(response.Message)
	if err != nil {
		return response, err
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	return response, nil
}
