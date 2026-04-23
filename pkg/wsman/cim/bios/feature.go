/*********************************************************************
 * Copyright (c) Intel Corporation 2026
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package bios facilitates communication with Intel(R) AMT devices to get information about BIOS features.
package bios

import (
	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

// FWData holds captured firmware XML responses for testing/recording purposes.
type FWData struct {
	GetXML       string
	EnumerateXML string
	PullXML      string
	PutXML       string
}

type Feature struct {
	base.WSManService[Response]
}

// NewBIOSFeatureWithClient instantiates a new Feature.
func NewBIOSFeatureWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Feature {
	return Feature{
		base.NewService[Response](wsmanMessageCreator, CIMBIOSFeature, client),
	}
}

// FetchFWData captures the XML from Get, Enumerate, Pull (and optionally Put) operations.
func (f Feature) FetchFWData(request PutRequest) (FWData, error) {
	var fwData FWData

	enumerateResponse, err := f.Enumerate()
	if err != nil {
		return fwData, err
	}

	fwData.EnumerateXML = enumerateResponse.XMLOutput

	pullResponse, err := f.Pull(enumerateResponse.Body.EnumerateResponse.EnumerationContext)
	if err != nil {
		return fwData, err
	}

	fwData.PullXML = pullResponse.XMLOutput

	getResponse, err := f.Get()
	if err != nil {
		return fwData, err
	}

	fwData.GetXML = getResponse.XMLOutput

	if request.FeatureName != "" {
		feature := BIOSFeature{
			Name: request.FeatureName,
		}

		putResponse, err := f.Put(feature)
		if err != nil {
			return fwData, err
		}

		fwData.PutXML = putResponse.XMLOutput
	}

	return fwData, nil
}
