/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package provisioningrecordlog

import (
	"encoding/xml"
	"fmt"
	"testing"

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
	expectedResult := "{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"GetResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"CurrentNumberOfRecords\":0,\"ElementName\":\"\",\"EnabledState\":0,\"HealthState\":0,\"InstanceID\":\"\",\"LogState\":0,\"MaxNumberOfRecords\":0,\"Name\":\"\",\"OverwritePolicy\":0,\"RequestedState\":0},\"PullResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"ProvisioningRecordLogItems\":null},\"EnumerateResponse\":{\"EnumerationContext\":\"\"}}"
	result := response.JSON()
	assert.Equal(t, expectedResult, result)
}

func TestYaml(t *testing.T) {
	response := Response{
		Body: Body{
			PullResponse: PullResponse{},
		},
	}
	expectedResult := "xmlname:\n    space: \"\"\n    local: \"\"\ngetresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    currentnumberofrecords: 0\n    elementname: \"\"\n    enabledstate: 0\n    healthstate: 0\n    instanceid: \"\"\n    logstate: 0\n    maxnumberofrecords: 0\n    name: \"\"\n    overwritepolicy: 0\n    requestedstate: 0\npullresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    provisioningrecordlogitems: []\nenumerateresponse:\n    enumerationcontext: \"\"\n"
	result := response.YAML()
	assert.Equal(t, expectedResult, result)
}

func TestPositiveIPS_ProvisioningRecordLog(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.IPSResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{PackageUnderTest: "ips/provisioningrecordlog"}
	elementUnderTest := NewProvisioningRecordLogWithClient(wsmanMessageCreator, &client)

	t.Run("ips_ProvisioningRecordLog Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			body             string
			extraHeader      string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			{
				"should create a valid IPS_ProvisioningRecordLog Get wsman message",
				IPSProvisioningRecordLog,
				wsmantesting.Get,
				"",
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageGet

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetResponse: ProvisioningRecordLog{
						XMLName:                xml.Name{Space: fmt.Sprintf("%s%s", message.IPSSchema, IPSProvisioningRecordLog), Local: IPSProvisioningRecordLog},
						CurrentNumberOfRecords: 1,
						ElementName:            "Intel(r) AMT Provisioning Record Log",
						EnabledState:           2,
						HealthState:            0,
						InstanceID:             "Intel(r) AMT: RecordLog 1",
						LogState:               4,
						MaxNumberOfRecords:     1,
						Name:                   "Intel(r) AMT Provisioning Record Log",
						OverwritePolicy:        2,
						RequestedState:         12,
					},
				},
			},
			{
				"should create a valid IPS_ProvisioningRecordLog Enumerate wsman message",
				IPSProvisioningRecordLog,
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageEnumerate

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName:           xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{EnumerationContext: "17000000-0000-0000-0000-000000000000"},
				},
			},
			{
				"should create a valid IPS_ProvisioningRecordLog Pull wsman message",
				IPSProvisioningRecordLog,
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
						ProvisioningRecordLogItems: []ProvisioningRecordLog{{
							XMLName:                xml.Name{Space: fmt.Sprintf("%s%s", message.IPSSchema, IPSProvisioningRecordLog), Local: IPSProvisioningRecordLog},
							CurrentNumberOfRecords: 1,
							ElementName:            "Intel(r) AMT Provisioning Record Log",
							EnabledState:           2,
							HealthState:            0,
							InstanceID:             "Intel(r) AMT: RecordLog 1",
							LogState:               4,
							MaxNumberOfRecords:     1,
							Name:                   "Intel(r) AMT Provisioning Record Log",
							OverwritePolicy:        2,
							RequestedState:         12,
						}},
					},
				},
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
