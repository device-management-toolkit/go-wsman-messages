/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package concrete facilitates communication with Intel® AMT devices and is a generic association used to establish dependency relationships between ManagedElements.
//
// It is defined as a concrete subclass of the abstract CIM_Dependency class, to be used in place of many specific subclasses of Dependency that add no semantics, that is subclasses that do not clarify the type of dependency, update cardinalities, or add or remove qualifiers. Note that when you define additional semantics for Dependency, this class must not be subclassed. Specific semantics continue to be defined as subclasses of the abstract CIM_Dependency. ConcreteDependency is limited in its use as a concrete form of a general dependency.
//
// It was deemed more prudent to create this concrete subclass than to change Dependency from an abstract to a concrete class. Dependency already had multiple abstract subclasses in the CIM Schema, and wider industry usage and impact could not be anticipated.
package concrete

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

// NewDependency returns a new instance of the NewDependency struct.
// should be NewDependency() because concrete is scoped already as package name.
func NewDependencyWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Dependency {
	return Dependency{
		base: message.NewBaseWithClient(wsmanMessageCreator, CIMConcreteDependency, client),
	}
}

// TODO: Figure out how to call GET requiring resourceURIs and Selectors

// Enumerate the instances of this class.
func (dependency Dependency) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: dependency.base.Enumerate(),
		},
	}

	err = dependency.base.Execute(response.Message)
	if err != nil {
		return
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}

// Pull instances of this class, following an Enumerate operation.
func (dependency Dependency) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: dependency.base.Pull(enumerationContext),
		},
	}

	err = dependency.base.Execute(response.Message)
	if err != nil {
		return
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}
