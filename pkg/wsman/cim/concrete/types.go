/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package concrete

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/models"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
)

// Response Types.
type (
	Response struct {
		*client.Message
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}
	Body struct {
		XMLName           xml.Name `xml:"Body"`
		PullResponse      PullResponse
		EnumerateResponse common.EnumerateResponse
	}
	PullResponse struct {
		XMLName xml.Name             `xml:"PullResponse"`
		Items   []ConcreteDependency `xml:"Items>CIM_ConcreteDependency"`
	}
	ConcreteDependency struct {
		Antecedent models.AssociationReference `xml:"Antecedent"` // Antecedent represents the independent object in this association.
		Dependent  models.AssociationReference `xml:"Dependent"`  // Dependent represents the object that is dependent on the Antecedent.
	}
)
