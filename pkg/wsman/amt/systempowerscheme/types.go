/*********************************************************************
 * Copyright (c) Intel Corporation 2026
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package systempowerscheme

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
)

// OUTPUTS
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
		GetResponse       SystemPowerScheme
		EnumerateResponse common.EnumerateResponse
		PullResponse      PullResponse
	}
	PullResponse struct {
		XMLName                xml.Name            `xml:"PullResponse"`
		SystemPowerSchemeItems []SystemPowerScheme `xml:"Items>AMT_SystemPowerScheme"`
	}

	// SystemPowerScheme represents a single AMT system power scheme entry.
	SystemPowerScheme struct {
		XMLName                 xml.Name `xml:"AMT_SystemPowerScheme"`
		CreationClassName       string   `xml:"CreationClassName,omitempty"`
		Description             string   `xml:"Description,omitempty"`
		ElementName             string   `xml:"ElementName"`
		InstanceID              string   `xml:"InstanceID"`
		Name                    string   `xml:"Name,omitempty"`
		SchemeGUID              string   `xml:"SchemeGUID,omitempty"`
		SystemCreationClassName string   `xml:"SystemCreationClassName,omitempty"`
		SystemName              string   `xml:"SystemName,omitempty"`
		PolicyOwner             int      `xml:"PolicyOwner,omitempty"`
		PolicyPrecedence        int      `xml:"PolicyPrecedence,omitempty"`
	}
)
