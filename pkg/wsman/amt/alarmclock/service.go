/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package alarmclock facilitates communication with Intel® AMT devices to set an alarm time to turn the host computer system on. Setting an alarm time is done by calling "AddAlarm" method.
package alarmclock

import (
	"encoding/xml"
	"strconv"
	"strings"
	"time"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/amt/methods"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

// NewServiceWithClient instantiates a new Alarm Clock service
func NewServiceWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Service {
	return Service{
		base: message.NewBaseWithClient(wsmanMessageCreator, AMT_AlarmClockService, client),
	}
}

// Get retrieves the representation of the instance
func (acs Service) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: acs.base.Get(nil),
		},
	}

	// send the message to AMT
	err = acs.base.Execute(response.Message)
	if err != nil {
		return
	}

	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}

// Enumerate returns an enumeration context which is used in a subsequent Pull call
func (acs Service) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: acs.base.Enumerate(),
		},
	}
	// send the message to AMT
	err = acs.base.Execute(response.Message)
	if err != nil {
		return
	}

	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}

// Pull returns the instances of this class.  An enumeration context provided by the Enumerate call is used as input.
func (acs Service) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: acs.base.Pull(enumerationContext),
		},
	}
	// send the message to AMT
	err = acs.base.Execute(response.Message)
	if err != nil {
		return
	}

	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}

// AddAlarm creates an alarm that would wake the system at a given time. The method receives as input an embedded instance of type IPS_AlarmClockOccurrence, with the following fields set: StartTime, Interval, InstanceID, DeleteOnCompletion. Upon success, the method creates an instance of IPS_AlarmClockOccurrence which is associated with AlarmClockService. The method would fail if 5 instances or more of IPS_AlarmClockOccurrence already exist in the system.
func (acs Service) AddAlarm(alarmClockOccurrence AlarmClockOccurrence) (response Response, err error) {
	header := acs.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMT_AlarmClockService, AddAlarm), AMT_AlarmClockService, nil, "", "")
	startTime := alarmClockOccurrence.StartTime.UTC().Format(time.RFC3339Nano)
	startTime = strings.Split(startTime, ".")[0]

	var body strings.Builder
	body.WriteString(`<Body><p:AddAlarm_INPUT xmlns:p="`)
	body.WriteString(acs.base.WSManMessageCreator.ResourceURIBase)
	body.WriteString(`AMT_AlarmClockService"><p:AlarmTemplate><s:InstanceID xmlns:s="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_AlarmClockOccurrence">`)
	body.WriteString(alarmClockOccurrence.InstanceID)
	body.WriteString(`</s:InstanceID>`)

	if alarmClockOccurrence.ElementName != "" {
		body.WriteString(`<s:ElementName xmlns:s="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_AlarmClockOccurrence">`)
		body.WriteString(alarmClockOccurrence.ElementName)
		body.WriteString(`</s:ElementName>`)
	}

	body.WriteString(`<s:StartTime xmlns:s="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_AlarmClockOccurrence"><p:Datetime xmlns:p="http://schemas.dmtf.org/wbem/wscim/1/common">`)
	body.WriteString(startTime)
	body.WriteString(`</p:Datetime></s:StartTime>`)

	minutes := alarmClockOccurrence.Interval % 60
	hours := (alarmClockOccurrence.Interval / 60) % 24
	days := alarmClockOccurrence.Interval / 1440

	body.WriteString(`<s:Interval xmlns:s="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_AlarmClockOccurrence"><p:Interval xmlns:p="http://schemas.dmtf.org/wbem/wscim/1/common">P`)
	body.WriteString(strconv.Itoa(days))
	body.WriteString("DT")
	body.WriteString(strconv.Itoa(hours))
	body.WriteString("H")
	body.WriteString(strconv.Itoa(minutes))
	body.WriteString(`M</p:Interval></s:Interval>`)

	body.WriteString(`<s:DeleteOnCompletion xmlns:s="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_AlarmClockOccurrence">`)
	body.WriteString(strconv.FormatBool(alarmClockOccurrence.DeleteOnCompletion))
	body.WriteString(`</s:DeleteOnCompletion></p:AlarmTemplate></p:AddAlarm_INPUT></Body>`)

	response = Response{
		Message: &client.Message{
			XMLInput: acs.base.WSManMessageCreator.CreateXML(header, body.String()),
		},
	}
	// send the message to AMT
	err = acs.base.Execute(response.Message)
	if err != nil {
		return
	}

	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}
