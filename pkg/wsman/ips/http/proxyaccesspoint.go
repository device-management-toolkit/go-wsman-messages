/*********************************************************************
 * Copyright (c) Intel Corporation 2025
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package http facilitates communication with Intel® AMT devices to manage HTTP proxy access point configuration.
//
// This service represents the HTTP Proxy Access Points configured for user-initiated connections through Intel AMT firmware.
package http

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
)

// ProxyAccessPointService struct represents the HTTP Proxy Access Point Service.
type ProxyAccessPointService struct {
	base.WSManService[ProxyAccessPointResponse]
}

// ProxyAccessPointResponse represents responses for IPS_HTTPProxyAccessPoint operations.
type ProxyAccessPointResponse struct {
	*client.Message
	XMLName xml.Name             `xml:"Envelope"`
	Header  message.Header       `xml:"Header"`
	Body    ProxyAccessPointBody `xml:"Body"`
}

// ProxyAccessPointBody represents the body of IPS_HTTPProxyAccessPoint responses.
type ProxyAccessPointBody struct {
	XMLName           xml.Name                     `xml:"Body"`
	EnumerateResponse common.EnumerateResponse     `xml:"EnumerateResponse"`
	PullResponse      ProxyAccessPointPullResponse `xml:"PullResponse"`
	GetAndPutResponse HTTPProxyAccessPointItem     `xml:"IPS_HTTPProxyAccessPoint"`
}

// ProxyAccessPointPullResponse represents the pull response for proxy access points.
type ProxyAccessPointPullResponse struct {
	XMLName xml.Name                   `xml:"PullResponse"`
	Items   []HTTPProxyAccessPointItem `xml:"Items>IPS_HTTPProxyAccessPoint"`
}

// HTTPProxyAccessPointItem represents an individual HTTP proxy access point configuration.
type HTTPProxyAccessPointItem struct {
	XMLName                 xml.Name `xml:"IPS_HTTPProxyAccessPoint"`
	Name                    string   `xml:"Name,omitempty"`
	CreationClassName       string   `xml:"CreationClassName,omitempty"`
	SystemName              string   `xml:"SystemName,omitempty"`
	SystemCreationClassName string   `xml:"SystemCreationClassName,omitempty"`
	ElementName             string   `xml:"ElementName,omitempty"`
	AccessInfo              string   `xml:"AccessInfo,omitempty"`       // The proxy address (IP or FQDN)
	InfoFormat              int      `xml:"InfoFormat,omitempty"`       // Format of AccessInfo: 3=IPv4, 4=IPv6, 201=FQDN
	Port                    int      `xml:"Port,omitempty"`             // Proxy port
	NetworkDnsSuffix        string   `xml:"NetworkDnsSuffix,omitempty"` // Domain suffix
}

// NewHTTPProxyAccessPointServiceWithClient returns a new instance of the ProxyAccessPointService struct.
func NewHTTPProxyAccessPointServiceWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) ProxyAccessPointService {
	return ProxyAccessPointService{
		base.NewService[ProxyAccessPointResponse](wsmanMessageCreator, IPSHTTPProxyAccessPoint, client),
	}
}

// Delete removes the specified HTTP proxy access point instance.
func (service ProxyAccessPointService) Delete(name string) (response ProxyAccessPointResponse, err error) {
	selector := message.Selector{Name: "Name", Value: name}
	response = ProxyAccessPointResponse{
		Message: &client.Message{
			XMLInput: service.Base.Delete(selector),
		},
	}

	// send the message to AMT
	err = service.Base.Execute(response.Message)
	if err != nil {
		return response, err
	}

	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	return response, nil
}
