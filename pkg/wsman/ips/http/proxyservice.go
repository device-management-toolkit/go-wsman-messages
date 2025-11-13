/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package http facilitates communication with Intel® AMT devices to manage HTTP proxy service configuration.
//
// This service represents the HTTP Proxy Service used for managing proxy settings and access points
// for user-initiated connections through Intel AMT firmware.
package http

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/models"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/ips/methods"
)

// ProxyService struct represents the HTTP Proxy Service.
type ProxyService struct {
	base.WSManService[Response]
}

// OUTPUT
// Response Types.
type (
	Response struct {
		*client.Message
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}

	Body struct {
		XMLName                     xml.Name `xml:"Body"`
		EnumerateResponse           common.EnumerateResponse
		GetAndPutResponse           HTTPProxyServiceResponse   `xml:"IPS_HTTPProxyService"`
		PullResponse                PullResponse               `xml:"PullResponse"`
		AddProxyAccessPointResponse AddProxyAccessPoint_OUTPUT `xml:"AddProxyAccessPoint_OUTPUT"`
	}

	HTTPProxyServiceResponse struct {
		XMLName                 xml.Name `xml:"IPS_HTTPProxyService"`
		Name                    string   `xml:"Name,omitempty"`
		CreationClassName       string   `xml:"CreationClassName,omitempty"`
		SystemName              string   `xml:"SystemName,omitempty"`
		SystemCreationClassName string   `xml:"SystemCreationClassName,omitempty"`
		ElementName             string   `xml:"ElementName,omitempty"`
		SyncEnabled             bool     `xml:"SyncEnabled,omitempty"` // Defines whether HTTP proxy sync (from local) is allowed
	}

	PullResponse struct {
		XMLName xml.Name                   `xml:"PullResponse"`
		Items   []HTTPProxyServiceResponse `xml:"Items>IPS_HTTPProxyService"`
	}

	// AddProxyAccessPoint_OUTPUT represents the output of AddProxyAccessPoint method.
	AddProxyAccessPoint_OUTPUT struct {
		XMLName          xml.Name         `xml:"AddProxyAccessPoint_OUTPUT"`
		ProxyAccessPoint ProxyAccessPoint `xml:"ProxyAccessPoint,omitempty"` // Reference to the created Proxy Access Point
		ReturnValue      int              `xml:"ReturnValue,omitempty"`      // Return value of the operation
	}

	// ProxyAccessPoint represents an endpoint reference to a proxy access point.
	ProxyAccessPoint struct {
		XMLName             xml.Name                          `xml:"ProxyAccessPoint,omitempty"`
		Address             string                            `xml:"Address,omitempty"`
		ReferenceParameters models.ReferenceParameters_OUTPUT `xml:"ReferenceParameters,omitempty"`
	}
)

// INPUT
// Request Types.
type (
	// AddProxyAccessPoint_INPUT represents the input parameters for AddProxyAccessPoint method.
	AddProxyAccessPoint_INPUT struct {
		XMLName          xml.Name `xml:"h:AddProxyAccessPoint_INPUT"`
		H                string   `xml:"xmlns:h,attr"`
		AccessInfo       string   `xml:"h:AccessInfo"`       // IP address or FQDN of the server (max 256 chars)
		InfoFormat       int      `xml:"h:InfoFormat"`       // Format of AccessInfo: 3=IPv4, 4=IPv6, 201=FQDN
		Port             int      `xml:"h:Port"`             // Port number
		NetworkDnsSuffix string   `xml:"h:NetworkDnsSuffix"` // Domain name of the network (max 192 chars)
	}
)

// InfoFormat enumeration constants.
const (
	InfoFormatIPv4 InfoFormat = 3   // IPv4 Address
	InfoFormatIPv6 InfoFormat = 4   // IPv6 Address
	InfoFormatFQDN InfoFormat = 201 // FQDN
)

// InfoFormat represents the format and interpretation of the AccessInfo property.
type InfoFormat int

// Return value constants for AddProxyAccessPoint method.
const (
	PTStatusSuccess          = 0    // PT_STATUS_SUCCESS
	PTStatusInternalError    = 1    // PT_STATUS_INTERNAL_ERROR
	PTStatusNotPermitted     = 16   // PT_STATUS_NOT_PERMITTED
	PTStatusMaxLimitReached  = 23   // PT_STATUS_MAX_LIMIT_REACHED
	PTStatusInvalidParameter = 36   // PT_STATUS_INVALID_PARAMETER
	PTStatusDuplicate        = 2058 // PT_STATUS_DUPLICATE
)

// NewHTTPProxyServiceWithClient returns a new instance of the HTTPProxyService struct.
func NewHTTPProxyServiceWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) ProxyService {
	return ProxyService{
		base.NewService[Response](wsmanMessageCreator, IPSHTTPProxyService, client),
	}
}

// AddProxyAccessPoint adds a Proxy access point that will be used when the Intel AMT firmware
// needs to open a user-initiated connection.
func (service ProxyService) AddProxyAccessPoint(accessInfo string, infoFormat InfoFormat, port int, networkDnsSuffix string) (response Response, err error) {
	header := service.Base.WSManMessageCreator.CreateHeader(methods.GenerateAction(IPSHTTPProxyService, AddProxyAccessPoint), IPSHTTPProxyService, nil, "", "")
	body := service.Base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod("AddProxyAccessPoint"), IPSHTTPProxyService, AddProxyAccessPoint_INPUT{
		H:                "http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HTTPProxyService",
		AccessInfo:       accessInfo,
		InfoFormat:       int(infoFormat),
		Port:             port,
		NetworkDnsSuffix: networkDnsSuffix,
	})

	response = Response{
		Message: &client.Message{
			XMLInput: service.Base.WSManMessageCreator.CreateXML(header, body),
		},
	}

	err = service.Base.Execute(response.Message)
	if err != nil {
		return response, err
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	return response, err
}
