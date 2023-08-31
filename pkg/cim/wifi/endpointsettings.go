/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package wifi

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
)

type EndpointSettings struct {
	base wsman.Base
}

type EnumerationEnvelope struct {
	XMLName xml.Name        `xml:"Envelope"`
	Header  wsman.Header    `xml:"Header"`
	Body    EnumerationBody `xml:"Body"`
}

type EnumerationBody struct {
	XMLName           xml.Name          `xml:"Body"`
	EnumerateResponse EnumerateResponse `xml:"EnumerateResponse"`
}

type EnumerateResponse struct {
	XMLName            xml.Name `xml:"EnumerateResponse"`
	EnumerationContext string   `xml:"EnumerationContext"`
}

type PullEnvelope struct {
	XMLName xml.Name     `xml:"Envelope"`
	Header  wsman.Header `xml:"Header"`
	Body    PullBody     `xml:"Body"`
}

type PullBody struct {
	XMLName      xml.Name     `xml:"Body"`
	PullResponse PullResponse `xml:"PullResponse"`
}

type PullResponse struct {
	Items         PullItems `xml:"Items"`
	EndOfSequence string    `xml:"EndOfSequence"`
}

type PullItems struct {
	WifiSettings []CIMWiFiEndpointSettings `xml:"CIM_WiFiEndpointSettings"`
}

type CIMWiFiEndpointSettings struct {
	InstanceID string `xml:"InstanceID"`
}

const CIM_WiFiEndpoint = "CIM_WiFiEndpoint"
const CIM_WiFiEndpointSettings = "CIM_WiFiEndpointSettings"

// NewWiFiEndpointSettings returns a new instance of the WiFiEndpointSettings struct.
func NewWiFiEndpointSettings(wsmanMessageCreator *wsman.WSManMessageCreator) EndpointSettings {
	return EndpointSettings{
		base: wsman.NewBase(wsmanMessageCreator, string(CIM_WiFiEndpointSettings)),
	}
}

// Get retrieves the representation of the instance
func (b EndpointSettings) Get() string {
	return b.base.Get(nil)
}

// Enumerates the instances of this class
func (b EndpointSettings) Enumerate() string {
	return b.base.Enumerate()
}

// Pulls instances of this class, following an Enumerate operation
func (b EndpointSettings) Pull(enumerationContext string) string {
	return b.base.Pull(enumerationContext)
}

// Delete removes a the specified instance
func (b EndpointSettings) Delete(handle string) string {
	selector := wsman.Selector{Name: "InstanceID", Value: handle}
	return b.base.Delete(selector)
}
