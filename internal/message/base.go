/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package message

import (
	"fmt"

	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

func NewBase(wsmanMessageCreator *WSManMessageCreator, className string) Base {
	return Base{
		WSManMessageCreator: wsmanMessageCreator,
		ClassName:           className,
	}
}

func NewBaseWithClient(wsmanMessageCreator *WSManMessageCreator, className string, client client.WSMan) Base {
	return Base{
		WSManMessageCreator: wsmanMessageCreator,
		ClassName:           className,
		client:              client,
	}
}

// Enumerate returns an enumeration context which is used in a subsequent Pull call.
func (b *Base) Enumerate() string {
	header := b.WSManMessageCreator.CreateHeader(BaseActionsEnumerate, b.ClassName, nil, "", "")

	return b.WSManMessageCreator.CreateXML(header, EnumerateBody)
}

// Get retrieves the representation of the instance.
func (b *Base) Get(selector *Selector) string {
	selectors := []Selector{}
	if selector != nil {
		selectors = append(selectors, *selector)
	}

	header := b.WSManMessageCreator.CreateHeader(BaseActionsGet, b.ClassName, selectors, "", "")

	return b.WSManMessageCreator.CreateXML(header, GetBody)
}

// Pull returns the instances of this class.  An enumeration context provided by the Enumerate call is used as input.
func (b *Base) Pull(enumerationContext string) string {
	header := b.WSManMessageCreator.CreateHeader(BaseActionsPull, b.ClassName, nil, "", "")
	body := createCommonBodyPull(enumerationContext, 0, 0)

	return b.WSManMessageCreator.CreateXML(header, body)
}

// Delete removes a the specified instance.
func (b *Base) Delete(selector Selector) string {
	header := b.WSManMessageCreator.CreateHeader(BaseActionsDelete, b.ClassName, []Selector{selector}, "", "")

	return b.WSManMessageCreator.CreateXML(header, DeleteBody)
}

// Put will change properties of the selected instance.
func (b *Base) Put(data interface{}, useHeaderSelector bool, selectorSet []Selector) string {
	if selectorSet == nil {
		selectorSet = []Selector{{Name: "InstanceID", Value: fmt.Sprintf("%v", data)}}
	}

	var header string

	if useHeaderSelector {
		header = b.WSManMessageCreator.CreateHeader(BaseActionsPut, b.ClassName, selectorSet, "", "")
	} else {
		header = b.WSManMessageCreator.CreateHeader(BaseActionsPut, b.ClassName, nil, "", "")
	}

	body := b.WSManMessageCreator.createCommonBodyCreateOrPut(b.ClassName, data)

	return b.WSManMessageCreator.CreateXML(header, body)
}

// Creates a new instance of this class.
func (b *Base) Create(data interface{}, selectorSet []Selector) string {
	header := b.WSManMessageCreator.CreateHeader(BaseActionsCreate, b.ClassName, selectorSet, "", "")
	body := b.WSManMessageCreator.createCommonBodyCreateOrPut(b.ClassName, data)

	return b.WSManMessageCreator.CreateXML(header, body)
}

// RequestStateChange requests that the state of the element be changed to the value specified in the RequestedState parameter . . .
func (b *Base) RequestStateChange(actionName string, requestedState int) string {
	header := b.WSManMessageCreator.CreateHeader(actionName, b.ClassName, nil, "", "")
	body := createCommonBodyRequestStateChange(fmt.Sprintf("%s%s", b.WSManMessageCreator.ResourceURIBase, b.ClassName), requestedState)

	return b.WSManMessageCreator.CreateXML(header, body)
}

func (b *Base) Execute(message *client.Message) error {
	if b.client != nil {
		xmlResponse, err := b.client.Post(message.XMLInput)
		message.XMLOutput = string(xmlResponse)

		if err != nil {
			return err
		}
	}

	// potentially could return an error that says that client doesn't exist
	return nil
}
