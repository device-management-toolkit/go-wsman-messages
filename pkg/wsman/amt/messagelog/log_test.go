/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package messagelog

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/common"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/wsmantesting"
)

func TestJson(t *testing.T) {
	response := Response{
		Body: Body{
			GetResponse: MessageLogResponse{},
		},
	}
	expectedResult := "{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"GetResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"Capabilities\":null,\"CharacterSet\":0,\"CreationClassName\":\"\",\"CurrentNumberOfRecords\":0,\"ElementName\":\"\",\"EnabledDefault\":0,\"EnabledState\":0,\"HealthState\":0,\"IsFrozen\":false,\"LastChange\":0,\"LogState\":0,\"MaxLogSize\":0,\"MaxNumberOfRecords\":0,\"MaxRecordSize\":0,\"Name\":\"\",\"OperationalStatus\":null,\"OverwritePolicy\":0,\"PercentageNearFull\":0,\"RequestedState\":0,\"SizeOfHeader\":0,\"SizeOfRecordHeader\":0,\"Status\":\"\"},\"EnumerateResponse\":{\"EnumerationContext\":\"\"},\"PullResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"MessageLogItems\":null},\"GetRecordsResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"IterationIdentifier\":0,\"NoMoreRecords\":false,\"RecordArray\":null,\"ReturnValue\":0},\"PositionToFirstRecordResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"IterationIdentifier\":0,\"ReturnValue\":0}}"
	result := response.JSON()
	assert.Equal(t, expectedResult, result)
}

func TestYaml(t *testing.T) {
	response := Response{
		Body: Body{
			GetResponse: MessageLogResponse{},
		},
	}
	expectedResult := "xmlname:\n    space: \"\"\n    local: \"\"\ngetresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    capabilities: []\n    characterset: 0\n    creationclassname: \"\"\n    currentnumberofrecords: 0\n    elementname: \"\"\n    enableddefault: 0\n    enabledstate: 0\n    healthstate: 0\n    isfrozen: false\n    lastchange: 0\n    logstate: 0\n    maxlogsize: 0\n    maxnumberofrecords: 0\n    maxrecordsize: 0\n    name: \"\"\n    operationalstatus: []\n    overwritepolicy: 0\n    percentagenearfull: 0\n    requestedstate: 0\n    sizeofheader: 0\n    sizeofrecordheader: 0\n    status: \"\"\nenumerateresponse:\n    enumerationcontext: \"\"\npullresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    messagelogitems: []\ngetrecordsresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    iterationidentifier: 0\n    nomorerecords: false\n    recordarray: []\n    returnvalue: 0\npositiontofirstrecordresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    iterationidentifier: 0\n    returnvalue: 0\n"
	result := response.YAML()
	assert.Equal(t, expectedResult, result)
}

func TestPositiveAMT_MessageLog(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/messagelog",
	}
	elementUnderTest := NewMessageLogWithClient(wsmanMessageCreator, &client)

	t.Run("amt_MessageLog Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			body             string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			//GETS
			{
				"should create a valid AMT_MessageLog Get wsman message",
				AMT_MessageLog, wsmantesting.GET,
				"",
				func() (Response, error) {
					client.CurrentMessage = "Get"
					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetResponse: MessageLogResponse{
						XMLName:                xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_MessageLog", Local: "AMT_MessageLog"},
						Capabilities:           []Capabilities{5, 6, 8, 7},
						CharacterSet:           10,
						CreationClassName:      "AMT_MessageLog",
						CurrentNumberOfRecords: 390,
						ElementName:            "Intel(r) AMT:MessageLog 1",
						EnabledDefault:         2,
						EnabledState:           2,
						HealthState:            5,
						IsFrozen:               false,
						LastChange:             0,
						LogState:               4,
						MaxLogSize:             0,
						MaxNumberOfRecords:     390,
						MaxRecordSize:          21,
						Name:                   "Intel(r) AMT:MessageLog 1",
						OperationalStatus:      []OperationalStatus{2},
						OverwritePolicy:        2,
						PercentageNearFull:     100,
						RequestedState:         12,
						SizeOfHeader:           0,
						SizeOfRecordHeader:     0,
						Status:                 "OK",
					},
				},
			},
			//ENUMERATES
			{
				"should create a valid AMT_MessageLog Enumerate wsman message",
				AMT_MessageLog,
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Enumerate"
					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "16080000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should create a valid AMT_MessageLog Pull wsman message",
				AMT_MessageLog,
				wsmantesting.PULL,
				wsmantesting.PULL_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Pull"
					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						XMLName: xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/09/enumeration", Local: "PullResponse"},
						MessageLogItems: []MessageLogResponse{
							{
								XMLName:                xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_MessageLog", Local: "AMT_MessageLog"},
								Capabilities:           []Capabilities{5, 6, 8, 7},
								CharacterSet:           10,
								CreationClassName:      "AMT_MessageLog",
								CurrentNumberOfRecords: 390,
								ElementName:            "Intel(r) AMT:MessageLog 1",
								EnabledDefault:         2,
								EnabledState:           2,
								HealthState:            5,
								IsFrozen:               false,
								LastChange:             0,
								LogState:               4,
								MaxLogSize:             0,
								MaxNumberOfRecords:     390,
								MaxRecordSize:          21,
								Name:                   "Intel(r) AMT:MessageLog 1",
								OperationalStatus:      []OperationalStatus{2},
								OverwritePolicy:        2,
								PercentageNearFull:     100,
								RequestedState:         12,
								SizeOfHeader:           0,
								SizeOfRecordHeader:     0,
								Status:                 "OK",
							},
						},
					},
				},
			},
			// POSITION TO FIRST RECORD
			{
				"should return a valid amt_MessageLog PositionToFirstRecords wsman message",
				AMT_MessageLog,
				`http://intel.com/wbem/wscim/1/amt-schema/1/AMT_MessageLog/PositionToFirstRecord`,
				`<h:PositionToFirstRecord_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_MessageLog"></h:PositionToFirstRecord_INPUT>`,
				func() (Response, error) {
					client.CurrentMessage = "PositionToFirstRecord"
					return elementUnderTest.PositionToFirstRecord()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PositionToFirstRecordResponse: PositionToFirstRecordResponse{
						XMLName:             xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_MessageLog", Local: "PositionToFirstRecord_OUTPUT"},
						IterationIdentifier: 1,
						ReturnValue:         0,
					},
				},
			},
			// GET RECORDS
			{
				"should return a valid amt_MessageLog GetRecords wsman message",
				AMT_MessageLog,
				`http://intel.com/wbem/wscim/1/amt-schema/1/AMT_MessageLog/GetRecords`,
				`<h:GetRecords_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_MessageLog"><h:IterationIdentifier>1</h:IterationIdentifier><h:MaxReadRecords>390</h:MaxReadRecords></h:GetRecords_INPUT>`,
				func() (Response, error) {
					client.CurrentMessage = "GetRecords"
					return elementUnderTest.GetRecords(1)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetRecordsResponse: GetRecordsResponse{
						XMLName:             xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_MessageLog", Local: "GetRecords_OUTPUT"},
						IterationIdentifier: 3,
						NoMoreRecords:       true,
						RecordArray:         []string{"Y8iYZf8GbwVoEP8mYaoKAAAAAAAA", "IgYBZf8PbwJoAf8iAEAHAAAAAAAA", "IgYBZf8PbwJoAf8iAEAHAAAAAAAA"},
						ReturnValue:         0,
					},
				},
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				expectedXMLInput := wsmantesting.ExpectedResponse(messageID, resourceUriBase, test.method, test.action, "", test.body)
				messageID++
				response, err := test.responseFunc()
				assert.NoError(t, err)
				assert.Equal(t, expectedXMLInput, response.XMLInput)
				assert.Equal(t, test.expectedResponse, response.Body)
			})
		}
	})
}
func TestNegativeAMT_MessageLog(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/messagelog",
	}
	elementUnderTest := NewMessageLogWithClient(wsmanMessageCreator, &client)

	t.Run("amt_MessageLog Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			body             string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			//GETS
			{
				"should create a valid AMT_MessageLog Get wsman message",
				AMT_MessageLog, wsmantesting.GET,
				"",
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetResponse: MessageLogResponse{
						XMLName:                xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_MessageLog", Local: "AMT_MessageLog"},
						Capabilities:           []Capabilities{5, 6, 8, 7},
						CharacterSet:           10,
						CreationClassName:      "AMT_MessageLog",
						CurrentNumberOfRecords: 390,
						ElementName:            "Intel(r) AMT:MessageLog 1",
						EnabledDefault:         2,
						EnabledState:           2,
						HealthState:            5,
						IsFrozen:               false,
						LastChange:             0,
						LogState:               4,
						MaxLogSize:             0,
						MaxNumberOfRecords:     390,
						MaxRecordSize:          21,
						Name:                   "Intel(r) AMT:MessageLog 1",
						OperationalStatus:      []OperationalStatus{2},
						OverwritePolicy:        2,
						PercentageNearFull:     100,
						RequestedState:         12,
						SizeOfHeader:           0,
						SizeOfRecordHeader:     0,
						Status:                 "OK",
					},
				},
			},
			//ENUMERATES
			{
				"should create a valid AMT_MessageLog Enumerate wsman message",
				AMT_MessageLog,
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "16080000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should create a valid AMT_MessageLog Pull wsman message",
				AMT_MessageLog,
				wsmantesting.PULL,
				wsmantesting.PULL_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						XMLName: xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/09/enumeration", Local: "PullResponse"},
						MessageLogItems: []MessageLogResponse{
							{
								XMLName:                xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_MessageLog", Local: "AMT_MessageLog"},
								Capabilities:           []Capabilities{5, 6, 8, 7},
								CharacterSet:           10,
								CreationClassName:      "AMT_MessageLog",
								CurrentNumberOfRecords: 390,
								ElementName:            "Intel(r) AMT:MessageLog 1",
								EnabledDefault:         2,
								EnabledState:           2,
								HealthState:            5,
								IsFrozen:               false,
								LastChange:             0,
								LogState:               4,
								MaxLogSize:             0,
								MaxNumberOfRecords:     390,
								MaxRecordSize:          21,
								Name:                   "Intel(r) AMT:MessageLog 1",
								OperationalStatus:      []OperationalStatus{2},
								OverwritePolicy:        2,
								PercentageNearFull:     100,
								RequestedState:         12,
								SizeOfHeader:           0,
								SizeOfRecordHeader:     0,
								Status:                 "OK",
							},
						},
					},
				},
			},
			// POSITION TO FIRST RECORD
			{
				"should return a valid amt_MessageLog PositionToFirstRecords wsman message",
				AMT_MessageLog,
				`http://intel.com/wbem/wscim/1/amt-schema/1/AMT_MessageLog/PositionToFirstRecord`,
				`<h:PositionToFirstRecord_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_MessageLog"></h:PositionToFirstRecord_INPUT>`,
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.PositionToFirstRecord()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PositionToFirstRecordResponse: PositionToFirstRecordResponse{
						XMLName:             xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_MessageLog", Local: "PositionToFirstRecord_OUTPUT"},
						IterationIdentifier: 1,
						ReturnValue:         0,
					},
				},
			},
			// GET RECORDS
			{
				"should return a valid amt_MessageLog GetRecords wsman message",
				AMT_MessageLog,
				`http://intel.com/wbem/wscim/1/amt-schema/1/AMT_MessageLog/GetRecords`,
				`<h:GetRecords_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_MessageLog"><h:IterationIdentifier>1</h:IterationIdentifier><h:MaxReadRecords>390</h:MaxReadRecords></h:GetRecords_INPUT>`,
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.GetRecords(1)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetRecordsResponse: GetRecordsResponse{
						XMLName:             xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_MessageLog", Local: "GetRecords_OUTPUT"},
						IterationIdentifier: 3,
						NoMoreRecords:       true,
						RecordArray:         []string{"Y8iYZf8GbwVoEP8mYaoKAAAAAAAA", "IgYBZf8PbwJoAf8iAEAHAAAAAAAA", "IgYBZf8PbwJoAf8iAEAHAAAAAAAA"},
						ReturnValue:         0,
					},
				},
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				expectedXMLInput := wsmantesting.ExpectedResponse(messageID, resourceUriBase, test.method, test.action, "", test.body)
				messageID++
				response, err := test.responseFunc()
				assert.Error(t, err)
				assert.Equal(t, expectedXMLInput, response.XMLInput)
				assert.NotEqual(t, test.expectedResponse, response.Body)
			})
		}
	})
}
