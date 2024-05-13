/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package software facilitates communication with Intel® AMT devices to provide descriptive information about a software component for asset tracking and/or installation dependency management.
//
// When the IsEntity property has the value TRUE, the instance of SoftwareIdentity represents an individually identifiable entity similar to Physical Element. SoftwareIdentity does NOT indicate whether the software is installed, executing, etc. This extra information may be provided through specialized associations to Software Identity. For instance, both InstalledSoftwareIdentity and ElementSoftwareIdentity may be used to indicate that the software identified by this class is installed. SoftwareIdentity is used when managing the software components of a ManagedElement that is the management focus. Since software may be acquired, SoftwareIdentity can be associated with a Product using the ProductSoftwareComponent relationship. The Application Model manages the deployment and installation of software via the classes, SoftwareFeatures and SoftwareElements. SoftwareFeature and SoftwareElement are used when the software component is the management focus. The deployment/installation concepts are related to the asset/identity one. In fact, a SoftwareIdentity may correspond to a Product, or to one or more SoftwareFeatures or SoftwareElements - depending on the granularity of these classes and the deployment model. The correspondence of Software Identity to Product, SoftwareFeature or SoftwareElement is indicated using the ConcreteIdentity association. Note that there may not be sufficient detail or instrumentation to instantiate ConcreteIdentity. And, if the association is instantiated, some duplication of information may result. For example, the Vendor described in the instances of Product and SoftwareIdentity MAY be the same. However, this is not necessarily true, and it is why vendor and similar information are duplicated in this class.
//
// Note that ConcreteIdentity can also be used to describe the relationship of the software to any LogicalFiles that result from installing it. As above, there may not be sufficient detail or instrumentation to instantiate this association.
package software

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

// NewSoftwareIdentity returns a new instance of the SoftwareIdentity struct.
func NewSoftwareIdentityWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Identity {
	return Identity{
		base: message.NewBaseWithClient(wsmanMessageCreator, CIMSoftwareIdentity, client),
	}
}

// Get retrieves the representation of the instance.
func (identity Identity) Get(instanceID string) (response Response, err error) {
	selector := message.Selector{
		Name:  "InstanceID",
		Value: instanceID,
	}
	response = Response{
		Message: &client.Message{
			XMLInput: identity.base.Get(&selector),
		},
	}

	err = identity.base.Execute(response.Message)
	if err != nil {
		return response, err
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

// Enumerate returns an enumeration context which is used in a subsequent Pull call.
func (identity Identity) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: identity.base.Enumerate(),
		},
	}

	err = identity.base.Execute(response.Message)
	if err != nil {
		return response, err
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

// Pull returns the instances of this class.  An enumeration context provided by the Enumerate call is used as input.
func (identity Identity) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: identity.base.Pull(enumerationContext),
		},
	}

	err = identity.base.Execute(response.Message)
	if err != nil {
		return response, err
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	return response, nil
}
