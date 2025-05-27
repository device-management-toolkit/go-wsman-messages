/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package alarmclock facilitates communication with Intel® AMT devices to represent a single alarm clock setting.
package alarmclock

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

// NewAlarmClockOccurrence returns a new instance of the AlarmClockOccurrence struct.
func NewAlarmClockOccurrenceWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Occurrence {
	return Occurrence{
		base: message.NewBaseWithClient(wsmanMessageCreator, IPSAlarmClockOccurrence, client),
	}
}

// Get retrieves the representation of the instance.
func (occurrence Occurrence) Get(alarmName string) (response Response, err error) {
	selector := message.Selector{
		Name:  "Name",
		Value: alarmName,
	}

	response = Response{
		Message: &client.Message{
			XMLInput: occurrence.base.Get(&selector),
		},
	}

	err = occurrence.base.Execute(response.Message)
	if err != nil {
		return
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}

// Delete removes a the specified instance.
func (occurrence Occurrence) Delete(handle string) (response Response, err error) {
	selector := message.Selector{Name: "InstanceID", Value: handle}
	response = Response{
		Message: &client.Message{
			XMLInput: occurrence.base.Delete(selector),
		},
	}

	err = occurrence.base.Execute(response.Message)
	if err != nil {
		return
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}

// Enumerate returns an enumeration context which is used in a subsequent Pull call.
func (occurrence Occurrence) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: occurrence.base.Enumerate(),
		},
	}

	err = occurrence.base.Execute(response.Message)
	if err != nil {
		return
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}

// Pull returns the instances of this class.  An enumeration context provided by the Enumerate call is used as input.
func (occurrence Occurrence) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: occurrence.base.Pull(enumerationContext),
		},
	}

	err = occurrence.base.Execute(response.Message)
	if err != nil {
		return
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}
