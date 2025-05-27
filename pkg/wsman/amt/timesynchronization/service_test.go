/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package timesynchronization

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/amt/methods"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/wsmantesting"
)

func TestJson(t *testing.T) {
	response := Response{
		Body: Body{
			GetResponse: TimeSynchronizationServiceResponse{},
		},
	}
	expectedResult := "{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"GetResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"Name\":\"\",\"CreationClassName\":\"\",\"SystemName\":\"\",\"SystemCreationClassName\":\"\",\"ElementName\":\"\",\"EnabledState\":0,\"RequestedState\":0,\"LocalTimeSyncEnabled\":0,\"TimeSource\":0},\"EnumerateResponse\":{\"EnumerationContext\":\"\"},\"PullResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"TimeSynchronizationServiceItems\":null},\"GetLowAccuracyTimeSynchResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"Ta0\":0,\"ReturnValue\":0},\"SetHighAccuracyTimeSynchResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"ReturnValue\":0}}"
	result := response.JSON()
	assert.Equal(t, expectedResult, result)
}

func TestYaml(t *testing.T) {
	response := Response{
		Body: Body{
			GetResponse: TimeSynchronizationServiceResponse{},
		},
	}
	expectedResult := "xmlname:\n    space: \"\"\n    local: \"\"\ngetresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    name: \"\"\n    creationclassname: \"\"\n    systemname: \"\"\n    systemcreationclassname: \"\"\n    elementname: \"\"\n    enabledstate: 0\n    requestedstate: 0\n    localtimesyncenabled: 0\n    timesource: 0\nenumerateresponse:\n    enumerationcontext: \"\"\npullresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    timesynchronizationserviceitems: []\ngetlowaccuracytimesynchresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    ta0: 0\n    returnvalue: 0\nsethighaccuracytimesynchresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    returnvalue: 0\n"
	result := response.YAML()
	assert.Equal(t, expectedResult, result)
}

func TestPositiveAMT_TimeSynchronizationService(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.AMTResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/timesynchronization",
	}
	elementUnderTest := NewTimeSynchronizationServiceWithClient(wsmanMessageCreator, &client)

	t.Run("amt_TimeSynchronizationService Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			body             string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			// GETS
			{
				"should create a valid AMT_TimeSynchronizationService Get wsman message",
				AMTTimeSynchronizationService,
				wsmantesting.Get,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageGet

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetResponse: TimeSynchronizationServiceResponse{
						XMLName:                 xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TimeSynchronizationService", Local: "AMT_TimeSynchronizationService"},
						CreationClassName:       "AMT_TimeSynchronizationService",
						ElementName:             "Intel(r) AMT Time Synchronization Service",
						EnabledState:            5,
						LocalTimeSyncEnabled:    0,
						Name:                    "Intel(r) AMT Time Synchronization Service",
						RequestedState:          12,
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "Intel(r) AMT",
						TimeSource:              1,
					},
				},
			},
			// ENUMERATES
			{
				"should create a valid AMT_TimeSynchronizationService Enumerate wsman message",
				AMTTimeSynchronizationService,
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageEnumerate

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "3B080000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create a valid AMT_TimeSynchronizationService Pull wsman message",
				AMTTimeSynchronizationService,
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
						TimeSynchronizationServiceItems: []TimeSynchronizationServiceResponse{
							{
								XMLName:                 xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TimeSynchronizationService", Local: "AMT_TimeSynchronizationService"},
								CreationClassName:       "AMT_TimeSynchronizationService",
								ElementName:             "Intel(r) AMT Time Synchronization Service",
								EnabledState:            5,
								LocalTimeSyncEnabled:    0,
								Name:                    "Intel(r) AMT Time Synchronization Service",
								RequestedState:          12,
								SystemCreationClassName: "CIM_ComputerSystem",
								SystemName:              "Intel(r) AMT",
								TimeSource:              1,
							},
						},
					},
				},
			},
			{
				// GetLowAccuracyTimeSynch
				"should create a valid AMT_TimeSynchronizationService GetLowAccuracyTimeSynch wsman message",
				AMTTimeSynchronizationService,
				methods.GenerateAction(AMTTimeSynchronizationService, GetLowAccuracyTimeSynch),
				`<h:GetLowAccuracyTimeSynch_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TimeSynchronizationService"></h:GetLowAccuracyTimeSynch_INPUT>`,
				func() (Response, error) {
					client.CurrentMessage = "GetLowAccuracyTimeSynch"

					return elementUnderTest.GetLowAccuracyTimeSynch()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetLowAccuracyTimeSynchResponse: GetLowAccuracyTimeSynchResponse{
						XMLName:     xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TimeSynchronizationService", Local: "GetLowAccuracyTimeSynch_OUTPUT"},
						Ta0:         1704586865,
						ReturnValue: 0,
					},
				},
			},
			{
				// SetHighAccuracyTimeSynch
				"should create a valid AMT_TimeSynchronizationService SetHighAccuracyTimeSynch wsman message",
				AMTTimeSynchronizationService,
				methods.GenerateAction(AMTTimeSynchronizationService, SetHighAccuracyTimeSynch),
				"<h:SetHighAccuracyTimeSynch_INPUT xmlns:h=\"http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TimeSynchronizationService\"><h:Ta0>1644240911</h:Ta0><h:Tm1>1644240943</h:Tm1><h:Tm2>1644240943</h:Tm2></h:SetHighAccuracyTimeSynch_INPUT>",
				func() (Response, error) {
					client.CurrentMessage = "SetHighAccuracyTimeSynch"

					return elementUnderTest.SetHighAccuracyTimeSynch(1644240911, 1644240943, 1644240943)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					SetHighAccuracyTimeSynchResponse: SetHighAccuracyTimeSynchResponse{
						XMLName:     xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TimeSynchronizationService", Local: "SetHighAccuracyTimeSynch_OUTPUT"},
						ReturnValue: 0,
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

func TestNegativeAMT_TimeSynchronizationService(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.AMTResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/timesynchronization",
	}
	elementUnderTest := NewTimeSynchronizationServiceWithClient(wsmanMessageCreator, &client)

	t.Run("amt_TimeSynchronizationService Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			body             string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			// GETS
			{
				"should create a valid AMT_TimeSynchronizationService Get wsman message",
				AMTTimeSynchronizationService,
				wsmantesting.Get,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetResponse: TimeSynchronizationServiceResponse{
						XMLName:                 xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TimeSynchronizationService", Local: "AMT_TimeSynchronizationService"},
						CreationClassName:       "AMT_TimeSynchronizationService",
						ElementName:             "Intel(r) AMT Time Synchronization Service",
						EnabledState:            5,
						LocalTimeSyncEnabled:    0,
						Name:                    "Intel(r) AMT Time Synchronization Service",
						RequestedState:          12,
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "Intel(r) AMT",
						TimeSource:              1,
					},
				},
			},
			// ENUMERATES
			{
				"should create a valid AMT_TimeSynchronizationService Enumerate wsman message",
				AMTTimeSynchronizationService,
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "3B080000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create a valid AMT_TimeSynchronizationService Pull wsman message",
				AMTTimeSynchronizationService,
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
						TimeSynchronizationServiceItems: []TimeSynchronizationServiceResponse{
							{
								XMLName:                 xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TimeSynchronizationService", Local: "AMT_TimeSynchronizationService"},
								CreationClassName:       "AMT_TimeSynchronizationService",
								ElementName:             "Intel(r) AMT Time Synchronization Service",
								EnabledState:            5,
								LocalTimeSyncEnabled:    0,
								Name:                    "Intel(r) AMT Time Synchronization Service",
								RequestedState:          12,
								SystemCreationClassName: "CIM_ComputerSystem",
								SystemName:              "Intel(r) AMT",
								TimeSource:              1,
							},
						},
					},
				},
			},
			{
				// GetLowAccuracyTimeSynch
				"should create a valid AMT_TimeSynchronizationService GetLowAccuracyTimeSynch wsman message",
				AMTTimeSynchronizationService,
				methods.GenerateAction(AMTTimeSynchronizationService, GetLowAccuracyTimeSynch),
				`<h:GetLowAccuracyTimeSynch_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TimeSynchronizationService"></h:GetLowAccuracyTimeSynch_INPUT>`,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.GetLowAccuracyTimeSynch()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetLowAccuracyTimeSynchResponse: GetLowAccuracyTimeSynchResponse{
						XMLName:     xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TimeSynchronizationService", Local: "GetLowAccuracyTimeSynch_OUTPUT"},
						Ta0:         1704586865,
						ReturnValue: 0,
					},
				},
			},
			{
				// SetHighAccuracyTimeSynch
				"should create a valid AMT_TimeSynchronizationService SetHighAccuracyTimeSynch wsman message",
				AMTTimeSynchronizationService,
				methods.GenerateAction(AMTTimeSynchronizationService, SetHighAccuracyTimeSynch),
				"<h:SetHighAccuracyTimeSynch_INPUT xmlns:h=\"http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TimeSynchronizationService\"><h:Ta0>1644240911</h:Ta0><h:Tm1>1644240943</h:Tm1><h:Tm2>1644240943</h:Tm2></h:SetHighAccuracyTimeSynch_INPUT>",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.SetHighAccuracyTimeSynch(1644240911, 1644240943, 1644240943)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					SetHighAccuracyTimeSynchResponse: SetHighAccuracyTimeSynchResponse{
						XMLName:     xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TimeSynchronizationService", Local: "SetHighAccuracyTimeSynch_OUTPUT"},
						ReturnValue: 0,
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
