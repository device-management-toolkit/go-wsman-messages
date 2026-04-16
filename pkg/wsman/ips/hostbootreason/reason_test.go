/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package hostbootreason

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
	expectedResult := "{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"GetResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"ElementName\":\"\",\"InstanceID\":\"\",\"PreviousSxState\":0,\"Reason\":0,\"ReasonDetails\":\"\"},\"PullResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"HostBootReasonItems\":null},\"EnumerateResponse\":{\"EnumerationContext\":\"\"}}"
	result := response.JSON()
	assert.Equal(t, expectedResult, result)
}

func TestYaml(t *testing.T) {
	response := Response{
		Body: Body{
			PullResponse: PullResponse{},
		},
	}
	expectedResult := "xmlname:\n    space: \"\"\n    local: \"\"\ngetresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    elementname: \"\"\n    instanceid: \"\"\n    previoussxstate: 0\n    reason: 0\n    reasondetails: \"\"\npullresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    hostbootreasonitems: []\nenumerateresponse:\n    enumerationcontext: \"\"\n"
	result := response.YAML()
	assert.Equal(t, expectedResult, result)
}

func TestPositiveIPS_HostBootReason(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.IPSResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "ips/hostbootreason",
	}
	elementUnderTest := NewHostBootReasonWithClient(wsmanMessageCreator, &client)

	t.Run("ips_HostBootReason Tests", func(t *testing.T) {
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
				"should create a valid IPS_HostBootReason Get wsman message",
				IPSHostBootReason,
				wsmantesting.Get,
				"",
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageGet

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetResponse: HostBootReasonResponse{
						XMLName:         xml.Name{Space: fmt.Sprintf("%s%s", message.IPSSchema, IPSHostBootReason), Local: IPSHostBootReason},
						ElementName:     "Intel(r) AMT:Host Boot Reason",
						InstanceID:      "Intel(r) AMT:Host Boot Reason",
						PreviousSxState: 0,
						Reason:          1,
						ReasonDetails:   "",
					},
				},
			},
			{
				"should create a valid IPS_HostBootReason Enumerate wsman message",
				IPSHostBootReason,
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageEnumerate

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName:           xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{EnumerationContext: "97000000-0000-0000-0000-000000000000"},
				},
			},
			{
				"should create a valid IPS_HostBootReason Pull wsman message",
				IPSHostBootReason,
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
						HostBootReasonItems: []HostBootReasonResponse{
							{
								XMLName:         xml.Name{Space: fmt.Sprintf("%s%s", message.IPSSchema, IPSHostBootReason), Local: IPSHostBootReason},
								ElementName:     "Intel(r) AMT:Host Boot Reason",
								InstanceID:      "Intel(r) AMT:Host Boot Reason",
								PreviousSxState: 0,
								Reason:          1,
								ReasonDetails:   "",
							},
						},
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
