/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package concrete

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/cim/models"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/common"
)

type Dependency struct {
	base   message.Base
	client client.WSMan
}

// Response Types
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
		Antecedent models.AssociationReference `xml:"Antecedent"`
		Dependent  models.AssociationReference `xml:"Dependent"`
	}
)
