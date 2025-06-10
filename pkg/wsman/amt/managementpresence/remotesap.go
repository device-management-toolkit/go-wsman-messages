/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package managementpresence facilitiates communication with Intel® AMT devices to configure Management Presence Remote Service Access Points (or an MPS) to be accessed by the Intel® AMT subsystem from remote.
package managementpresence

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

type RemoteSAP struct {
	base.WSManService[Response]
}

// NewManagementPresenceRemoteSAPWithClient instantiates a new RemoteSAP.
func NewManagementPresenceRemoteSAPWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) RemoteSAP {
	return RemoteSAP{
		base.NewService[Response](wsmanMessageCreator, AMTManagementPresenceRemoteSAP, client),
	}
}

// Delete removes a the specified instance.
func (remoteSAP RemoteSAP) Delete(handle string) (response Response, err error) {
	selector := message.Selector{Name: "Name", Value: handle}
	response = Response{
		Message: &client.Message{
			XMLInput: remoteSAP.Base.Delete(selector),
		},
	}
	// send the message to AMT
	err = remoteSAP.Base.Execute(response.Message)
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
