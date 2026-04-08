/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package eventlogentry

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
)

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
		XMLName           xml.Name `xml:"Body"`
		GetResponse       EventLogEntry
		EnumerateResponse common.EnumerateResponse
		PullResponse      PullResponse
	}
	PullResponse struct {
		XMLName            xml.Name        `xml:"PullResponse"`
		EventLogEntryItems []EventLogEntry `xml:"Items>AMT_EventLogEntry"`
	}

	// EventLogEntry represents a single AMT event log record.
	EventLogEntry struct {
		XMLName           xml.Name `xml:"AMT_EventLogEntry"`
		CreationClassName string   `xml:"CreationClassName"`
		DeviceAddress     int      `xml:"DeviceAddress"`
		ElementName       string   `xml:"ElementName"`
		EventData         string   `xml:"EventData"`
		EventSensorType   int      `xml:"EventSensorType"`
		EventType         int      `xml:"EventType"`
		RecordID          int      `xml:"RecordID"`
		TimeStamp         int64    `xml:"TimeStamp"`
	}
)
