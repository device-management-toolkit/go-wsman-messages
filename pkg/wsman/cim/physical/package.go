/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package physical

import (
	"encoding/xml"
	"errors"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

// NewPhysicalPackage returns a new instance of the PhysicalPackage struct.
func NewPhysicalPackageWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Package {
	return Package{
		base.NewService[Response](wsmanMessageCreator, CIMPhysicalPackage, client),
	}
}

// Pull returns the instances of this class.  An enumeration context provided by the Enumerate call is used as input.
func (physicalPackage Package) Pull(enumerationContext string) (response Response, err error) {
	loopMax := 3
	loopCnt := 0

	response = Response{
		Message: &client.Message{
			XMLInput: physicalPackage.Base.Pull(enumerationContext),
		},
	}

	for {
		err = physicalPackage.Base.Execute(response.Message)
		if err != nil {
			return response, err
		}

		err = xml.Unmarshal([]byte(response.XMLOutput), &response)
		if err != nil {
			return response, err
		}

		if response.Body.PullResponse.EndOfSequence.Local != "" {
			break
		}

		loopCnt++
		if loopCnt == loopMax {
			err = errors.New("CIM_PhysicalPackage.Pull() - maximum pull attempts exceeded")

			break
		}
	}

	return response, err
}
