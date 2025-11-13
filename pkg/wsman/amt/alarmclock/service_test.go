/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package alarmclock

import (
	"encoding/xml"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/amt/methods"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/models"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/wsmantesting"
)

const (
	EnvelopeResponse = `<a:Envelope xmlns:a="http://www.w3.org/2003/05/soap-envelope" xmlns:b="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:c="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns:d="http://schemas.xmlsoap.org/ws/2005/02/trust" xmlns:e="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" xmlns:f="http://schemas.dmtf.org/wbem/wsman/1/cimbinding.xsd" xmlns:g="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AlarmClockService" xmlns:h="http://schemas.dmtf.org/wbem/wscim/1/common" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><a:Header><b:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</b:To><b:RelatesTo>0</b:RelatesTo><b:Action a:mustUnderstand="true">`
	GetBody          = `<g:AMT_AlarmClockService><g:CreationClassName>AMT_AlarmClockService</g:CreationClassName><g:ElementName>Intel(r) AMT Alarm Clock Service</g:ElementName><g:Name>Intel(r) AMT Alarm Clock Service</g:Name><g:SystemCreationClassName>CIM_ComputerSystem</g:SystemCreationClassName><g:SystemName>ManagedSystem</g:SystemName></g:AMT_AlarmClockService>`
	StartTime        = "2022-12-31T23:59:00Z"
)

func TestJson(t *testing.T) {
	response := Response{
		Body: Body{
			GetResponse: AlarmClockService{},
		},
	}
	expectedResult := "{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"GetResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"Name\":\"\",\"CreationClassName\":\"\",\"SystemName\":\"\",\"SystemCreationClassName\":\"\",\"ElementName\":\"\",\"NextAMTAlarmTime\":\"\",\"AMTAlarmClockInterval\":\"\"},\"EnumerateResponse\":{\"EnumerationContext\":\"\"},\"AddAlarmOutput\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"AlarmClock\":{\"Address\":\"\",\"ReferenceParameters\":{\"ResourceURI\":\"\",\"SelectorSet\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"Selector\":null}}},\"ReturnValue\":0},\"PullResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"AlarmClockServiceItems\":null}}"
	result := response.JSON()
	assert.Equal(t, expectedResult, result)
}

func TestYaml(t *testing.T) {
	response := Response{
		Body: Body{
			GetResponse: AlarmClockService{},
		},
	}
	expectedResult := "xmlname:\n    space: \"\"\n    local: \"\"\ngetresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    name: \"\"\n    creationclassname: \"\"\n    systemname: \"\"\n    systemcreationclassname: \"\"\n    elementname: \"\"\n    nextamtalarmtime: \"\"\n    amtalarmclockinterval: \"\"\nenumerateresponse:\n    enumerationcontext: \"\"\naddalarmoutput:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    alarmclock:\n        address: \"\"\n        referenceparameters:\n            resourceuri: \"\"\n            selectorset:\n                xmlname:\n                    space: \"\"\n                    local: \"\"\n                selector: []\n    returnvalue: 0\npullresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    alarmclockserviceitems: []\n"
	result := response.YAML()
	assert.Equal(t, expectedResult, result)
}

func TestPositiveAMT_AlarmClockService(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.AMTResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/alarmclock",
	}
	elementUnderTest := NewServiceWithClient(wsmanMessageCreator, &client)

	t.Run("amt_AlarmClockService Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			body             string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			{
				"should create and parse valid AMT_AlarmClockService Get call",
				AMTAlarmClockService,
				wsmantesting.Get,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageGet

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetResponse: AlarmClockService{
						XMLName:                 xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AlarmClockService", Local: "AMT_AlarmClockService"},
						CreationClassName:       AMTAlarmClockService,
						ElementName:             "Intel(r) AMT Alarm Clock Service",
						Name:                    "Intel(r) AMT Alarm Clock Service",
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "ManagedSystem",
					},
				},
			},

			{
				"should create and parse valid AMT_AlarmClockService Enumerate call",
				AMTAlarmClockService,
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageEnumerate

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "8A000000-0000-0000-0000-000000000000",
					},
				},
			},

			{
				"should create and parse valid AMT_AlarmClockService Pull call",
				AMTAlarmClockService,
				wsmantesting.Pull,
				wsmantesting.PullBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessagePull

					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						XMLName: xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/09/enumeration", Local: "PullResponse"},
						AlarmClockServiceItems: []AlarmClockService{
							{
								XMLName:                 xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AlarmClockService", Local: "AMT_AlarmClockService"},
								Name:                    "Intel(r) AMT Alarm Clock Service",
								CreationClassName:       AMTAlarmClockService,
								SystemName:              "ManagedSystem",
								SystemCreationClassName: "CIM_ComputerSystem",
								ElementName:             "Intel(r) AMT Alarm Clock Service",
							},
						},
					},
				},
			},
			// AddAlarm
			{
				"should create and parse valid AMT_AlarmClockService AddAlarm call",
				AMTAlarmClockService,
				methods.GenerateAction(AMTAlarmClockService, AddAlarm),
				`<r:AddAlarm_INPUT xmlns:r="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AlarmClockService"><d:AlarmTemplate xmlns:d="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AlarmClockService" xmlns:s="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_AlarmClockOccurrence"><s:InstanceID>Instance</s:InstanceID><s:ElementName>Alarm instance name</s:ElementName><s:StartTime><p:Datetime xmlns:p="http://schemas.dmtf.org/wbem/wscim/1/common">2022-12-31T23:59:00Z</p:Datetime></s:StartTime><s:Interval><p:Interval xmlns:p="http://schemas.dmtf.org/wbem/wscim/1/common">P1DT23H59M</p:Interval></s:Interval><s:DeleteOnCompletion>true</s:DeleteOnCompletion></d:AlarmTemplate></r:AddAlarm_INPUT>`,
				func() (Response, error) {
					client.CurrentMessage = "AddAlarm"
					startTime := StartTime
					minutes := 59
					hours := 23
					days := 1
					interval := minutes + hours*60 + days*1440

					startTimeFormatted, err := time.Parse(time.RFC3339, startTime)
					if err != nil {
						return Response{}, err
					}

					return elementUnderTest.AddAlarm(AlarmClockOccurrence{
						InstanceID:         "Instance",
						StartTime:          startTimeFormatted,
						ElementName:        "Alarm instance name",
						Interval:           interval,
						DeleteOnCompletion: true,
					})
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					AddAlarmOutput: AddAlarmOutput{
						XMLName: xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AlarmClockService", Local: "AddAlarm_OUTPUT"},
						AlarmClock: AlarmClock{
							Address: "default",
							ReferenceParameters: models.ReferenceParameters_OUTPUT{
								ResourceURI: "",
								SelectorSet: models.SelectorSet_OUTPUT{
									XMLName: xml.Name{
										Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd",
										Local: "SelectorSet",
									},
								},
							},
						},
					},
				},
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				expectedXMLInput := wsmantesting.ExpectedResponse(messageID, resourceURIBase, test.method, test.action, "", test.body)
				messageID++
				response, err := test.responseFunc()
				assert.NoError(t, err)
				assert.Equal(t, expectedXMLInput, response.XMLInput)
				assert.Equal(t, test.expectedResponse, response.Body)
			})
		}
	})
}

func TestNegativeAMT_AlarmClockService(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.AMTResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/alarmclock",
	}
	elementUnderTest := NewServiceWithClient(wsmanMessageCreator, &client)

	t.Run("amt_AlarmClockService Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			body             string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			{
				"should create and parse valid AMT_AlarmClockService Get call",
				AMTAlarmClockService,
				wsmantesting.Get,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetResponse: AlarmClockService{
						CreationClassName:       AMTAlarmClockService,
						ElementName:             "Intel(r) AMT Alarm Clock Service",
						Name:                    "Intel(r) AMT Alarm Clock Service",
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "ManagedSystem",
					},
				},
			},

			{
				"should create and parse valid AMT_AlarmClockService Enumerate call",
				AMTAlarmClockService,
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "8A000000-0000-0000-0000-000000000000",
					},
				},
			},

			{
				"should create and parse valid AMT_AlarmClockService Pull call",
				AMTAlarmClockService,
				wsmantesting.Pull,
				wsmantesting.PullBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						XMLName: xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/09/enumeration", Local: "PullResponse"},
						AlarmClockServiceItems: []AlarmClockService{
							{
								XMLName:                 xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AlarmClockService", Local: "AMT_AlarmClockService"},
								Name:                    "Intel(r) AMT Alarm Clock Service",
								CreationClassName:       AMTAlarmClockService,
								SystemName:              "ManagedSystem",
								SystemCreationClassName: "CIM_ComputerSystem",
								ElementName:             "Intel(r) AMT Alarm Clock Service",
							},
						},
					},
				},
			},

			{
				"should create and parse valid AMT_AlarmClockService AddAlarm call",
				AMTAlarmClockService,
				methods.GenerateAction(AMTAlarmClockService, AddAlarm),
				`<r:AddAlarm_INPUT xmlns:r="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AlarmClockService"><d:AlarmTemplate xmlns:d="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AlarmClockService" xmlns:s="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_AlarmClockOccurrence"><s:InstanceID>Instance</s:InstanceID><s:ElementName>Alarm instance name</s:ElementName><s:StartTime><p:Datetime xmlns:p="http://schemas.dmtf.org/wbem/wscim/1/common">2022-12-31T23:59:00Z</p:Datetime></s:StartTime><s:Interval><p:Interval xmlns:p="http://schemas.dmtf.org/wbem/wscim/1/common">P1DT23H59M</p:Interval></s:Interval><s:DeleteOnCompletion>true</s:DeleteOnCompletion></d:AlarmTemplate></r:AddAlarm_INPUT>`,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError
					startTime := StartTime
					minutes := 59
					hours := 23
					days := 1
					interval := minutes + hours*60 + days*1440

					startTimeFormatted, err := time.Parse(time.RFC3339, startTime)
					if err != nil {
						return Response{}, err
					}

					return elementUnderTest.AddAlarm(AlarmClockOccurrence{
						InstanceID:         "Instance",
						StartTime:          startTimeFormatted,
						ElementName:        "Alarm instance name",
						Interval:           interval,
						DeleteOnCompletion: true,
					})
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					AddAlarmOutput: AddAlarmOutput{
						XMLName: xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AlarmClockService", Local: "AddAlarm_OUTPUT"},
						AlarmClock: AlarmClock{
							Address: "default",
							ReferenceParameters: models.ReferenceParameters_OUTPUT{
								ResourceURI: "",
								SelectorSet: models.SelectorSet_OUTPUT{
									XMLName: xml.Name{
										Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd",
										Local: "SelectorSet",
									},
								},
							},
						},
					},
				},
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				expectedXMLInput := wsmantesting.ExpectedResponse(messageID, resourceURIBase, test.method, test.action, "", test.body)
				messageID++
				response, err := test.responseFunc()
				assert.Error(t, err)
				assert.Equal(t, expectedXMLInput, response.XMLInput)
				assert.NotEqual(t, test.expectedResponse, response.Body)
			})
		}
	})
}
