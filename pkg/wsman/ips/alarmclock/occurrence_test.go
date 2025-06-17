/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package alarmclock

import (
	"encoding/xml"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/wsmantesting"
)

func TestJson(t *testing.T) {
	response := Response{
		Body: Body{
			PullResponse: PullResponse{},
		},
	}
	expectedResult := "{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"PullResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"Items\":null},\"EnumerateResponse\":{\"EnumerationContext\":\"\"},\"GetResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"ElementName\":\"\",\"InstanceID\":\"\",\"StartTime\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"Datetime\":\"0001-01-01T00:00:00Z\"},\"Interval\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"Interval\":\"\"},\"DeleteOnCompletion\":false}}"
	result := response.JSON()
	assert.Equal(t, expectedResult, result)
}

func TestYaml(t *testing.T) {
	response := Response{
		Body: Body{
			PullResponse: PullResponse{},
		},
	}
	expectedResult := "xmlname:\n    space: \"\"\n    local: \"\"\npullresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    items: []\nenumerateresponse:\n    enumerationcontext: \"\"\ngetresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    elementname: \"\"\n    instanceid: \"\"\n    starttime:\n        xmlname:\n            space: \"\"\n            local: \"\"\n        datetime: 0001-01-01T00:00:00Z\n    interval:\n        xmlname:\n            space: \"\"\n            local: \"\"\n        interval: \"\"\n    deleteoncompletion: false\n"
	result := response.YAML()
	assert.Equal(t, expectedResult, result)
}

func TestPositiveIPS_AlarmClockOccurrence(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.IPSResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "ips/alarmclock",
	}
	elementUnderTest := NewAlarmClockOccurrenceWithClient(wsmanMessageCreator, &client)

	t.Run("ips_AlarmClockOccurrence Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			body             string
			extraHeader      string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			// GETS
			{
				"should create a valid ips_AlarmClockOccurrence Get wsman message",
				"IPS_AlarmClockOccurrence",
				wsmantesting.Get,
				"",
				"<w:SelectorSet><w:Selector Name=\"Name\">testalarm</w:Selector></w:SelectorSet>",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageGet

					return elementUnderTest.GetByName("testalarm")
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetResponse: AlarmClockOccurrence{
						XMLName:     xml.Name{Space: fmt.Sprintf("%s%s", message.IPSSchema, IPSAlarmClockOccurrence), Local: IPSAlarmClockOccurrence},
						ElementName: "testalarm",
						InstanceID:  "testalarm",
						StartTime: StartTime{
							XMLName:  xml.Name{Space: "http://intel.com/wbem/wscim/1/ips-schema/1/IPS_AlarmClockOccurrence", Local: "StartTime"},
							Datetime: time.Time{},
						},
						Interval: Interval{
							XMLName:  xml.Name{Space: "http://intel.com/wbem/wscim/1/ips-schema/1/IPS_AlarmClockOccurrence", Local: "Interval"},
							Interval: "",
						},
						DeleteOnCompletion: true,
					},
				},
			},
			// ENUMERATES
			{
				"should create a valid IPS_AlarmClockOccurrence Enumerate wsman message",
				"IPS_AlarmClockOccurrence",
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageEnumerate

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "9C0A0000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create a valid IPS_AlarmClockOccurrence Pull wsman message",
				"IPS_AlarmClockOccurrence",
				wsmantesting.Pull,
				wsmantesting.PullBody,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessagePull

					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						XMLName: xml.Name{Space: message.XMLPullResponseSpace, Local: "PullResponse"},
						Items: []AlarmClockOccurrence{
							{
								XMLName:     xml.Name{Space: fmt.Sprintf("%s%s", message.IPSSchema, IPSAlarmClockOccurrence), Local: IPSAlarmClockOccurrence},
								ElementName: "testalarm",
								InstanceID:  "testalarm",
								StartTime: StartTime{
									XMLName:  xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/09/enumeration", Local: "StartTime"},
									Datetime: time.Time{},
								},
								Interval: Interval{
									XMLName:  xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/09/enumeration", Local: "Interval"},
									Interval: "",
								},
								DeleteOnCompletion: true,
							},
						},
					},
				},
			},
			// DELETE
			{
				"should create a valid ips_AlarmClockOccurrence Delete wsman message",
				"IPS_AlarmClockOccurrence",
				wsmantesting.Delete,
				"",
				"<w:SelectorSet><w:Selector Name=\"InstanceID\">testalarm</w:Selector></w:SelectorSet>",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageDelete

					return elementUnderTest.Delete("testalarm")
				},
				Body{XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"}},
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				expectedXMLInput := wsmantesting.ExpectedResponse(messageID, resourceURIBase, test.method, test.action, test.extraHeader, test.body)
				messageID++
				response, err := test.responseFunc()
				assert.NoError(t, err)
				assert.Equal(t, expectedXMLInput, response.XMLInput)
				assert.Equal(t, test.expectedResponse, response.Body)
			})
		}
	})
}

func TestNegativeIPS_AlarmClockOccurrence(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.IPSResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "ips/alarmclock",
	}
	elementUnderTest := NewAlarmClockOccurrenceWithClient(wsmanMessageCreator, &client)

	t.Run("ips_AlarmClockOccurrence Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			body             string
			extraHeader      string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			// GETS
			{
				"should create a valid ips_AlarmClockOccurrence Get wsman message",
				"IPS_AlarmClockOccurrence",
				wsmantesting.Get,
				"",
				"<w:SelectorSet><w:Selector Name=\"Name\">testalarm</w:Selector></w:SelectorSet>",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.GetByName("testalarm")
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetResponse: AlarmClockOccurrence{
						XMLName:     xml.Name{Space: fmt.Sprintf("%s%s", message.IPSSchema, IPSAlarmClockOccurrence), Local: IPSAlarmClockOccurrence},
						ElementName: "testalarm",
						InstanceID:  "testalarm",
						StartTime: StartTime{
							XMLName:  xml.Name{Space: "http://intel.com/wbem/wscim/1/ips-schema/1/IPS_AlarmClockOccurrence", Local: "StartTime"},
							Datetime: time.Time{},
						},
						Interval: Interval{
							XMLName:  xml.Name{Space: "http://intel.com/wbem/wscim/1/ips-schema/1/IPS_AlarmClockOccurrence", Local: "Interval"},
							Interval: "",
						},
						DeleteOnCompletion: true,
					},
				},
			},
			// ENUMERATES
			{
				"should create a valid IPS_AlarmClockOccurrence Enumerate wsman message",
				"IPS_AlarmClockOccurrence",
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "9C0A0000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create a valid IPS_AlarmClockOccurrence Pull wsman message",
				"IPS_AlarmClockOccurrence",
				wsmantesting.Pull,
				wsmantesting.PullBody,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						XMLName: xml.Name{Space: message.XMLPullResponseSpace, Local: "PullResponse"},
						Items: []AlarmClockOccurrence{
							{
								XMLName:            xml.Name{Space: fmt.Sprintf("%s%s", message.IPSSchema, IPSAlarmClockOccurrence), Local: IPSAlarmClockOccurrence},
								ElementName:        "testalarm",
								InstanceID:         "testalarm",
								StartTime:          StartTime{Datetime: time.Time{}},
								Interval:           Interval{Interval: "0"},
								DeleteOnCompletion: true,
							},
						},
					},
				},
			},
			// DELETE
			{
				"should create a valid ips_AlarmClockOccurrence Delete wsman message",
				"IPS_AlarmClockOccurrence",
				wsmantesting.Delete,
				"",
				"<w:SelectorSet><w:Selector Name=\"InstanceID\">testalarm</w:Selector></w:SelectorSet>",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Delete("testalarm")
				},
				Body{XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"}},
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				expectedXMLInput := wsmantesting.ExpectedResponse(messageID, resourceURIBase, test.method, test.action, test.extraHeader, test.body)
				messageID++
				response, err := test.responseFunc()
				assert.Error(t, err)
				assert.Equal(t, expectedXMLInput, response.XMLInput)
				assert.NotEqual(t, test.expectedResponse, response.Body)
			})
		}
	})
}
