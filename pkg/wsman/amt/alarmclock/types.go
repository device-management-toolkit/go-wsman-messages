/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package alarmclock

import (
	"encoding/xml"
	"time"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/models"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
)

type Service struct {
	base message.Base
}

// INPUTS
// AlarmClockOccurrence represents a single alarm clock setting.
type AlarmClockOccurrence struct {
	ElementName        string    `json:"ElementName"`        // Elementname is a user-friendly name for the object
	InstanceID         string    `json:"InstanceID"`         // InstanceID is the instance key, set by the caller of AMT_AlarmClockService.AddAlarm.
	StartTime          time.Time `json:"StartTime"`          // StartTime is the next time when the alarm is scheduled to be set.
	Interval           int       `json:"Interval"`           // Interval between occurrences of the alarm (0 if the alarm is scheduled to run once).
	DeleteOnCompletion bool      `json:"DeleteOnCompletion"` // DeleteOnComplete if set to TRUE, the instance will be deleted by the FW when the alarm is completed
}

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
		XMLName           xml.Name          `xml:"Body"`
		GetResponse       AlarmClockService `xml:"AMT_AlarmClockService"`
		EnumerateResponse common.EnumerateResponse
		AddAlarmOutput    AddAlarmOutput `xml:"AddAlarm_OUTPUT"`
		PullResponse      PullResponse
	}
	PullResponse struct {
		XMLName                xml.Name            `xml:"PullResponse"`
		AlarmClockServiceItems []AlarmClockService `xml:"Items>AMT_AlarmClockService"`
	}
	AlarmClockService struct {
		XMLName                 xml.Name `xml:"AMT_AlarmClockService"`
		Name                    string   // The Name property uniquely identifies the Service and provides an indication of the functionality that is managed
		CreationClassName       string   // CreationClassName indicates the name of the class or the subclass that is used in the creation of an instance
		SystemName              string   // The Name of the scoping System.
		SystemCreationClassName string   // The CreationClassName of the scoping System.
		ElementName             string   // A user-friendly name for the object
		NextAMTAlarmTime        string   // Specifies the next AMT alarm time
		AMTAlarmClockInterval   string   // Specifies the alarm time interval
	}
	AddAlarmOutput struct {
		XMLName     xml.Name    `xml:"AddAlarm_OUTPUT"`
		AlarmClock  AlarmClock  // A reference to the created instance of IPS_AlarmClockOccurrence.
		ReturnValue ReturnValue // Return code. 0 indicates success
	}
	AlarmClock struct {
		// Reference address to the created instance of IPS_AlarmClockOccurrence.
		Address             string
		ReferenceParameters models.ReferenceParameters_OUTPUT
	}

	// ReturnValue is a return code. 0 indicates success.
	ReturnValue int
)
