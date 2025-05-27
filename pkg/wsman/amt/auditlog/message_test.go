/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package auditlog

import (
	"encoding/xml"
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
			GetResponse: AuditLog{},
		},
	}
	expectedResult := "{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"EnumerateResponse\":{\"EnumerationContext\":\"\"},\"GetResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"OverwritePolicy\":0,\"CurrentNumberOfRecords\":0,\"MaxNumberOfRecords\":0,\"ElementName\":\"\",\"EnabledState\":0,\"RequestedState\":0,\"PercentageFree\":0,\"Name\":\"\",\"TimeOfLastRecord\":{\"Datetime\":\"\"},\"AuditState\":0,\"MaxAllowedAuditors\":0,\"StoragePolicy\":0,\"MinDaysToKeep\":0},\"PullResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"AuditLogItems\":null},\"ReadRecordsResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"TotalRecordCount\":0,\"RecordsReturned\":0,\"EventRecords\":null,\"ReturnValue\":0},\"DecodedRecordsResponse\":null}"
	result := response.JSON()
	assert.Equal(t, expectedResult, result)
}

func TestYaml(t *testing.T) {
	response := Response{
		Body: Body{
			GetResponse: AuditLog{},
		},
	}
	expectedResult := "xmlname:\n    space: \"\"\n    local: \"\"\nenumerateresponse:\n    enumerationcontext: \"\"\ngetresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    overwritepolicy: 0\n    currentnumberofrecords: 0\n    maxnumberofrecords: 0\n    elementname: \"\"\n    enabledstate: 0\n    requestedstate: 0\n    percentagefree: 0\n    name: \"\"\n    timeoflastrecord:\n        datetime: \"\"\n    auditstate: 0\n    maxallowedauditors: 0\n    storagepolicy: 0\n    mindaystokeep: 0\npullresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    auditlogitems: []\nreadrecordsresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    totalrecordcount: 0\n    recordsreturned: 0\n    eventrecords: []\n    returnvalue: 0\ndecodedrecordsresponse: []\n"
	result := response.YAML()
	assert.Equal(t, expectedResult, result)
}

func TestPositiveAMT_AuditLog(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.AMTResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/auditlog",
	}
	elementUnderTest := NewAuditLogWithClient(wsmanMessageCreator, &client)

	t.Run("amt_AuditLog Tests", func(t *testing.T) {
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
				"should create a valid AMT_AuditLog Get wsman message",
				AMTAuditLog,
				wsmantesting.Get,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageGet

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetResponse: AuditLog{
						XMLName:                xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuditLog", Local: "AMT_AuditLog"},
						OverwritePolicy:        2,
						CurrentNumberOfRecords: 161,
						MaxNumberOfRecords:     0,
						ElementName:            "Intel(r) AMT Audit Log",
						EnabledState:           2,
						RequestedState:         2,
						PercentageFree:         92,
						Name:                   "Intel(r) AMT:Audit Log",
						TimeOfLastRecord: Datetime{
							Datetime: "2024-01-03T00:44:35Z",
						},
						AuditState:         16,
						MaxAllowedAuditors: 1,
						StoragePolicy:      1,
					},
				},
			},
			// ENUMERATES
			{
				"should create a valid AMT_AuditLog Enumerate wsman message",
				AMTAuditLog,
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageEnumerate

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "92070000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create a valid AMT_AuditLog Pull wsman message",
				AMTAuditLog,
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
						AuditLogItems: []AuditLog{
							{
								XMLName:                xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuditLog", Local: "AMT_AuditLog"},
								OverwritePolicy:        2,
								CurrentNumberOfRecords: 162,
								MaxNumberOfRecords:     0,
								ElementName:            "Intel(r) AMT Audit Log",
								EnabledState:           2,
								RequestedState:         2,
								PercentageFree:         92,
								Name:                   "Intel(r) AMT:Audit Log",
								TimeOfLastRecord: Datetime{
									Datetime: "2024-01-03T00:45:41Z",
								},
								AuditState:         16,
								MaxAllowedAuditors: 1,
								StoragePolicy:      1,
							},
						},
					},
				},
			},
			// READ RECORDS
			{
				"should create a valid AMT_AuditLog Read Records wsman message",
				AMTAuditLog,
				`http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuditLog/ReadRecords`,
				`<h:ReadRecords_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuditLog"><h:StartIndex>1</h:StartIndex></h:ReadRecords_INPUT>`,
				func() (Response, error) {
					client.CurrentMessage = "ReadRecords"

					return elementUnderTest.ReadRecords(1)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					ReadRecordsResponse: ReadRecords_OUTPUT{
						XMLName:          xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuditLog", Local: "ReadRecords_OUTPUT"},
						TotalRecordCount: 2,
						RecordsReturned:  2,
						EventRecords:     []string{"ABMAAAI/9M1uAgAQAAwAAAAoBZkADAAAACQFlA==", "ABMAAAI/9M1vAgAQAAwAAAAoBZkADAAAACQFlA=="},
						ReturnValue:      0,
					},
					DecodedRecordsResponse: []AuditLogRecord{
						{
							AuditAppID:     19,
							EventID:        0,
							InitiatorType:  0x2,
							AuditApp:       "Firmware Update Manager Events",
							Event:          "Firmware Update Started",
							Initiator:      "Local",
							Time:           time.Unix(int64(1073007983), 0),
							MCLocationType: 0x2,
							NetAddress:     "",
							Ex:             "\x00\f\x00\x00\x00(\x05\x99\x00\f\x00\x00\x00$\x05\x94",
							ExStr:          "Firmware update was started.\nOld version: 3072.0.10240.39173\nNew version: 3072.0.9216.37893",
						},
						{
							AuditAppID:     19,
							EventID:        0,
							InitiatorType:  0x2,
							AuditApp:       "Firmware Update Manager Events",
							Event:          "Firmware Update Started",
							Initiator:      "Local",
							Time:           time.Unix(int64(1073007982), 0),
							MCLocationType: 0x2,
							NetAddress:     "",
							Ex:             "\x00\f\x00\x00\x00(\x05\x99\x00\f\x00\x00\x00$\x05\x94",
							ExStr:          "Firmware update was started.\nOld version: 3072.0.10240.39173\nNew version: 3072.0.9216.37893",
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

func TestNegativeAMT_AuditLog(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.AMTResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/auditlog",
	}
	elementUnderTest := NewAuditLogWithClient(wsmanMessageCreator, &client)

	t.Run("amt_AuditLog Tests", func(t *testing.T) {
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
				"should create a valid AMT_AuditLog Get wsman message",
				AMTAuditLog,
				wsmantesting.Get,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetResponse: AuditLog{
						XMLName:                xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuditLog", Local: "AMT_AuditLog"},
						OverwritePolicy:        2,
						CurrentNumberOfRecords: 161,
						MaxNumberOfRecords:     0,
						ElementName:            "Intel(r) AMT Audit Log",
						EnabledState:           2,
						RequestedState:         2,
						PercentageFree:         92,
						Name:                   "Intel(r) AMT:Audit Log",
						TimeOfLastRecord: Datetime{
							Datetime: "2024-01-03T00:44:35Z",
						},
						AuditState:         16,
						MaxAllowedAuditors: 1,
						StoragePolicy:      1,
					},
				},
			},
			// ENUMERATES
			{
				"should create a valid AMT_AuditLog Enumerate wsman message",
				AMTAuditLog,
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "92070000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create a valid AMT_AuditLog Pull wsman message",
				AMTAuditLog,
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
						AuditLogItems: []AuditLog{
							{
								XMLName:                xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuditLog", Local: "AMT_AuditLog"},
								OverwritePolicy:        2,
								CurrentNumberOfRecords: 162,
								MaxNumberOfRecords:     0,
								ElementName:            "Intel(r) AMT Audit Log",
								EnabledState:           2,
								RequestedState:         2,
								PercentageFree:         92,
								Name:                   "Intel(r) AMT:Audit Log",
								TimeOfLastRecord: Datetime{
									Datetime: "2024-01-03T00:45:41Z",
								},
								AuditState:         16,
								MaxAllowedAuditors: 1,
								StoragePolicy:      1,
							},
						},
					},
				},
			},
			// READ RECORDS
			{
				"should create a valid AMT_AuditLog Read Records wsman message",
				AMTAuditLog,
				`http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuditLog/ReadRecords`,
				`<h:ReadRecords_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuditLog"><h:StartIndex>1</h:StartIndex></h:ReadRecords_INPUT>`,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.ReadRecords(1)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					ReadRecordsResponse: ReadRecords_OUTPUT{
						XMLName:          xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_AuditLog", Local: "ReadRecords_OUTPUT"},
						TotalRecordCount: 2,
						RecordsReturned:  2,
						EventRecords:     []string{"ABMAAAI/9M1uAgAQAAwAAAAoBZkADAAAACQFlA==", "ABMAAAI/9M1vAgAQAAwAAAAoBZkADAAAACQFlA=="},
						ReturnValue:      0,
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
